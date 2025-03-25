package dfscript

import (
	"github.com/df-mc/dragonfly/server/item/creative"
	"github.com/dop251/goja"
)

func (r *Runtime) setupCreative() error {
	// TODO: Allow registering items and groups, waiting for https://github.com/df-mc/dragonfly/pull/1014 to be merged.
	return newObject().
		Method("items", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(creative.Items())
		}).
		Apply(r.vm, "creative")
}
