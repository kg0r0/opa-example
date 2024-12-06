package main

import (
	"context"
	"fmt"

	"github.com/open-policy-agent/opa/ast"
	"github.com/open-policy-agent/opa/loader"
	"github.com/open-policy-agent/opa/rego"
)

func main() {
	ctx := context.TODO()
	path := "./policy/example.rego"
	regoFile, err := loader.RegoWithOpts(path, ast.ParserOptions{})
	if err != nil {
		panic(err)
	}
	rego := rego.New(
		rego.Query("data.example.allow"),
		rego.Module("example.rego", regoFile.Parsed.String()),
		rego.Input(
			map[string]interface{}{
				"user":   "bob",
				"path":   []string{"accounts", "bob"},
				"method": "GET",
			}),
	)
	rs, err := rego.Eval(ctx)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v\n", rs)
	println(rs.Allowed())
}
