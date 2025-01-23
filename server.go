package dfscript

import (
	"github.com/df-mc/dragonfly/server/world"
	"github.com/dop251/goja"
)

func (r *Runtime) setupServer() error {
	return newObject().
		Method("playerCount", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(r.s.PlayerCount())
		}).
		Method("maxPlayerCount", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(r.s.MaxPlayerCount())
		}).
		Method("players", func(c goja.FunctionCall) goja.Value {
			tx, _ := c.Argument(0).Export().(*world.Tx)
			return r.vm.ToValue(r.s.Players(tx))
		}).
		Method("player", func(c goja.FunctionCall) goja.Value {
			id := r.exportUUID(c, 0)
			h, ok := r.s.Player(id)
			if !ok {
				return nil
			}
			return r.vm.ToValue(h)
		}).
		Method("playerByName", func(c goja.FunctionCall) goja.Value {
			name := c.Argument(0).String()
			h, ok := r.s.PlayerByName(name)
			if !ok {
				return nil
			}
			return r.vm.ToValue(h)
		}).
		Method("playerByXUID", func(c goja.FunctionCall) goja.Value {
			xuid := c.Argument(0).String()
			h, ok := r.s.PlayerByXUID(xuid)
			if !ok {
				return nil
			}
			return r.vm.ToValue(h)
		}).
		Method("world", func(c goja.FunctionCall) goja.Value {
			name, ok := c.Argument(0).Export().(string)
			if !ok {
				panic(r.vm.NewTypeError("world expects a string as its first argument"))
			}
			w, ok := r.worlds[name]
			if !ok {
				return nil
			}
			return r.vm.ToValue(w)
		}).
		Method("worlds", func(c goja.FunctionCall) goja.Value {
			var worlds []*world.World
			for _, w := range r.worlds {
				worlds = append(worlds, w)
			}
			return r.vm.ToValue(worlds)
		}).
		Apply(r.vm, "server")
}
