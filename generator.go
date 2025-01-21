package dfscript

import (
	"github.com/df-mc/dragonfly/server/world"
	"github.com/df-mc/dragonfly/server/world/generator"
	"github.com/dop251/goja"
)

func (r *Runtime) setupGenerator() error {
	return newObject().
		Method("flat", func(c goja.FunctionCall) goja.Value {
			biome, ok := c.Argument(0).Export().(world.Biome)
			if !ok {
				panic(r.vm.NewTypeError("argument 0 must be a world.Biome"))
			}
			layers, ok := c.Argument(1).Export().([]world.Block)
			if !ok {
				panic(r.vm.NewTypeError("argument 1 must be a []world.Block"))
			}
			return r.vm.ToValue(generator.NewFlat(biome, layers))
		}).
		Apply(r.vm, "generator")
}
