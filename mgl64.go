package dfscript

import (
	"github.com/dop251/goja"
	"github.com/go-gl/mathgl/mgl64"
)

func (r *Runtime) setupMgl64() error {
	return newObject().
		Method("vec3", func(c goja.FunctionCall) goja.Value {
			if len(c.Arguments) < 3 {
				panic(r.vm.NewTypeError("vec3 expects 3 arguments"))
			}
			return r.vm.ToValue(mgl64.Vec3{
				c.Argument(0).ToFloat(),
				c.Argument(1).ToFloat(),
				c.Argument(2).ToFloat(),
			})
		}).
		Apply(r.vm, "mgl64")
}
