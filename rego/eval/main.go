package main

import (
	"context"
	"fmt"

	"github.com/open-policy-agent/opa/rego"
)

func main() {
	ctx := context.Background()

	rego := rego.New(rego.Query("x = 1"))

	rs, err := rego.Eval(ctx)
	if err != nil {
		panic(err)
	}

	fmt.Println(rs[0].Bindings["x"])
}
