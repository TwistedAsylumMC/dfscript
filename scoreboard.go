package dfscript

import (
	"github.com/df-mc/dragonfly/server/player/scoreboard"
	"github.com/dop251/goja"
)

func (r *Runtime) setupScoreboard() error {
	return newObject().
		Method("create", func(c goja.FunctionCall) goja.Value {
			var args []any
			for _, arg := range c.Arguments {
				args = append(args, arg.Export())
			}
			return r.vm.ToValue(scoreboard.New(args...))
		}).
		Apply(r.vm, "scoreboard")
}
