package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/open-policy-agent/opa/ast"
	"github.com/open-policy-agent/opa/rego"
	"github.com/open-policy-agent/opa/types"
)

func main() {

	r := rego.New(
		rego.Query(`github.repo("open-policy-agent", "opa")`),
		rego.Function2(
			&rego.Function{
				Name:             "github.repo",
				Decl:             types.NewFunction(types.Args(types.S, types.S), types.A),
				Memoize:          true,
				Nondeterministic: true,
			},
			func(bctx rego.BuiltinContext, a, b *ast.Term) (*ast.Term, error) {

				var org, repo string

				if err := ast.As(a.Value, &org); err != nil {
					return nil, err
				} else if err := ast.As(b.Value, &repo); err != nil {
					return nil, err
				}

				req, err := http.NewRequest("GET", fmt.Sprintf("https://api.github.com/repos/%v/%v", org, repo), nil)
				if err != nil {
					return nil, err
				}

				resp, err := http.DefaultClient.Do(req.WithContext(bctx.Context))
				if err != nil {
					return nil, err
				}

				defer resp.Body.Close()

				if resp.StatusCode != http.StatusOK {
					return nil, fmt.Errorf(resp.Status)
				}

				v, err := ast.ValueFromReader(resp.Body)
				if err != nil {
					return nil, err
				}

				return ast.NewTerm(v), nil
			},
		),
	)

	rs, err := r.Eval(context.Background())
	if err != nil {
		log.Fatal(err)
	} else if len(rs) == 0 {
		fmt.Println("undefined")
	} else {
		bs, _ := json.MarshalIndent(rs[0].Expressions[0].Value, "", "  ")
		fmt.Println(string(bs))
	}
}
