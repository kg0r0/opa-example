package main

import (
	"context"
	"fmt"

	"github.com/open-policy-agent/opa/ast"
	"github.com/open-policy-agent/opa/rego"
)

// Ref: https://pkg.go.dev/github.com/open-policy-agent/opa/rego#example-Rego.Eval-Compiler
func main() {

	ctx := context.Background()

	// Define a simple policy.
	module := `
		package example

		import rego.v1

		default allow = false

		allow if {
			input.identity = "admin"
		}

		allow if {
			input.method = "GET"
		}
	`

	// Compile the module. The keys are used as identifiers in error messages.
	compiler, err := ast.CompileModules(map[string]string{
		"example.rego": module,
	})
	if err != nil {
		// If there is a syntax error in the policy, the compiler will return an error.
		panic(err)
	}

	// Create a new query that uses the compiled policy from above.
	rego := rego.New(
		rego.Query("data.example.allow"),
		rego.Compiler(compiler),
		rego.Input(
			map[string]interface{}{
				"identity": "bob",
				"method":   "GET",
			},
		),
	)

	// Run evaluation.
	rs, err := rego.Eval(ctx)
	if err != nil {
		panic(err)
	}

	// Inspect results.
	fmt.Println("len:", len(rs))
	fmt.Println("value:", rs[0].Expressions[0].Value)
	fmt.Println("allowed:", rs.Allowed()) // helper method

}
