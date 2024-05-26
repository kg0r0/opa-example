package main

import (
	"context"
	"fmt"

	"github.com/open-policy-agent/opa/ast"
	"github.com/open-policy-agent/opa/rego"
	"github.com/open-policy-agent/opa/types"
)

// Ref: https://www.openpolicyagent.org/docs/latest/extensions/
func main() {
	ctx := context.Background()
	r := rego.New(
		rego.Query(`x = hello("bob")`),
		rego.Function1(
			&rego.Function{
				Name: "hello",
				Decl: types.NewFunction(types.Args(types.S), types.S),
			},
			func(_ rego.BuiltinContext, a *ast.Term) (*ast.Term, error) {
				if str, ok := a.Value.(ast.String); ok {
					return ast.StringTerm("hello, " + string(str)), nil
				}
				return nil, nil
			}),
	)

	query, err := r.PrepareForEval(ctx)
	if err != nil {
		// handle error.
	}

	rs, err := query.Eval(ctx)
	if err != nil {
		// handle error.
	}

	// Do something with result.
	fmt.Println(rs[0].Bindings["x"])
}
