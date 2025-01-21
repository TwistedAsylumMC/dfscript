package dfscript

import (
	"github.com/df-mc/dragonfly/server/player/skin"
	"github.com/dop251/goja"
)

func (r *Runtime) setupSkin() error {
	return newObject().
		Const("AnimationType", map[string]skin.AnimationType{
			"Head":        skin.AnimationHead,
			"Body32x32":   skin.AnimationBody32x32,
			"Body128x128": skin.AnimationBody128x128,
		}).
		Method("animation", func(c goja.FunctionCall) goja.Value {
			width, ok := c.Argument(0).Export().(int64)
			if !ok {
				panic(r.vm.NewTypeError("argument 0 must be an int64"))
			}
			height, ok := c.Argument(1).Export().(int64)
			if !ok {
				panic(r.vm.NewTypeError("argument 1 must be an int64"))
			}
			expression, ok := c.Argument(2).Export().(int64)
			if !ok {
				panic(r.vm.NewTypeError("argument 2 must be an int64"))
			}
			t, ok := c.Argument(3).Export().(skin.AnimationType)
			if !ok {
				panic(r.vm.NewTypeError("argument 3 must be a skin.AnimationType"))
			}
			return r.vm.ToValue(skin.NewAnimation(int(width), int(height), int(expression), t))
		}).
		Method("cape", func(c goja.FunctionCall) goja.Value {
			width, ok := c.Argument(0).Export().(int64)
			if !ok {
				panic(r.vm.NewTypeError("argument 0 must be an int64"))
			}
			height, ok := c.Argument(1).Export().(int64)
			if !ok {
				panic(r.vm.NewTypeError("argument 1 must be an int64"))
			}
			return r.vm.ToValue(skin.NewCape(int(width), int(height)))
		}).
		Method("modelConfig", func(c goja.FunctionCall) goja.Value {
			def, ok := c.Argument(0).Export().(string)
			if !ok {
				panic(r.vm.NewTypeError("argument 0 must be a string"))
			}
			face, ok := c.Argument(1).Export().(string)
			if !ok {
				panic(r.vm.NewTypeError("argument 1 must be a string"))
			}
			return r.vm.ToValue(skin.ModelConfig{Default: def, AnimatedFace: face})
		}).
		Method("decodeModelConfig", func(c goja.FunctionCall) goja.Value {
			b, ok := c.Argument(0).Export().([]byte)
			if !ok {
				panic(r.vm.NewTypeError("argument 0 must be a []byte"))
			}
			conf, err := skin.DecodeModelConfig(b)
			if err != nil {
				panic(r.vm.ToValue(err))
			}
			return r.vm.ToValue(conf)
		}).
		Method("create", func(c goja.FunctionCall) goja.Value {
			width, ok := c.Argument(0).Export().(int64)
			if !ok {
				panic(r.vm.NewTypeError("argument 0 must be an int64"))
			}
			height, ok := c.Argument(1).Export().(int64)
			if !ok {
				panic(r.vm.NewTypeError("argument 1 must be an int64"))
			}
			return r.vm.ToValue(skin.New(int(width), int(height)))
		}).
		Apply(r.vm, "skin")
}
