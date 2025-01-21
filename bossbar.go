package dfscript

import (
	"github.com/df-mc/dragonfly/server/player/bossbar"
	"github.com/dop251/goja"
)

func (r *Runtime) setupBossBar() error {
	return newObject().
		Const("colour", map[string]bossbar.Colour{
			"blue":   bossbar.Blue(),
			"green":  bossbar.Green(),
			"grey":   bossbar.Grey(),
			"purple": bossbar.Purple(),
			"red":    bossbar.Red(),
			"white":  bossbar.White(),
			"yellow": bossbar.Yellow(),
		}).
		Method("create", func(c goja.FunctionCall) goja.Value {
			var args []any
			for _, arg := range c.Arguments {
				args = append(args, arg.Export())
			}
			return r.vm.ToValue(bossbar.New(args...))
		}).
		Apply(r.vm, "bossbar")
}
