package generate

import (
	"gorm.io/gen"
	"gorm.io/gen/examples/dal"
)

var gener *gen.Generator

func CreateGen() {
	generator := gen.NewGenerator(gen.Config{
		OutPath: "./dal/query",
	})

	generator.UseDB(dai.Db)

}

func Execute(execs ...func(*gen.Generator)) {
	for _, exec := range execs {
		exec(gener)
	}
	gener.Execute()
}
