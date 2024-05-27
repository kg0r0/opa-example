package main

import (
	"context"
	"fmt"

	"github.com/open-policy-agent/opa/rego"
)

func main() {
	ctx := context.Background()

	module := `package example
	p = ["Hello", "World"] { true }`

	rego := rego.New(
		rego.Query("data.example.p"),
		rego.Module("example.rego", module),
	)

	rs, err := rego.Eval(ctx)
	if err != nil {
		panic(err)
	}
	fmt.Println(rs[0].Expressions[0].Value)
}
