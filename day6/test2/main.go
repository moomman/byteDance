package main

import (
	"gen/cmd/generate"
	"gorm.io/gen"
)

func main() {
	generate.Execute(generators)
}

func generators(g *gen.Generator) {
	g.ApplyBasic(
		g.GenerateModel("users"),
	)

	// apply diy interfaces on structs or table models

	// execute the action of code generation
	g.Execute()
}
