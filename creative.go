package dfscript

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/creative"
	"github.com/dop251/goja"
)

func (r *Runtime) setupCreative() error {
	return newObject().
		Method("items", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(creative.Items())
		}).Method("registerItem", func(c goja.FunctionCall) goja.Value {
		i, ok := c.Argument(0).Export().(item.Stack)
		if !ok {
			panic(r.vm.NewTypeError("argument 0 must be an item.Stack"))
		}
		creative.RegisterItem(i)
		return nil
	}).
		Apply(r.vm, "creative")
}
