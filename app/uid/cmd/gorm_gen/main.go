package main

import (
	"rpc102/app/uid/biz/dal/model"

	"gorm.io/gen"
)

func main() {

	g := gen.NewGenerator(gen.Config{
		OutPath: "biz/dal/query",
		Mode:    gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface,
	})

	g.ApplyInterface(func(model.Querier) {}, model.AllModels...)
	g.Execute()
}
