package dfscript

import (
	"github.com/df-mc/dragonfly/server/block/customblock"
	"github.com/dop251/goja"
)

func (r *Runtime) setupCustomBlock() error {
	return newObject().
		Const("RenderMethod", map[string]customblock.Method{
			"Opaque":      customblock.OpaqueRenderMethod(),
			"AlphaTest":   customblock.AlphaTestRenderMethod(),
			"Blend":       customblock.BlendRenderMethod(),
			"DoubleSided": customblock.DoubleSidedRenderMethod(),
		}).
		Method("material", func(c goja.FunctionCall) goja.Value {
			texture, ok := c.Argument(0).Export().(string)
			if !ok {
				panic(r.vm.NewTypeError("argument 0 is not a string"))
			}
			method, ok := c.Argument(1).Export().(customblock.Method)
			if !ok {
				panic(r.vm.NewTypeError("argument 1 is not a customblock.Method"))
			}
			return r.vm.ToValue(customblock.NewMaterial(texture, method))
		}).
		Apply(r.vm, "customblock")
}
