package dfscript

import (
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
			var players []*goja.Object
			for _, p := range r.players {
				players = append(players, p.obj)
			}
			return r.vm.ToValue(players)
		}).
		Method("player", func(c goja.FunctionCall) goja.Value {
			id := r.exportUUID(c, 0)
			p, ok := r.players[id]
			if !ok {
				return nil
			}
			return r.vm.ToValue(p.obj)
		}).
		Method("playerByName", func(c goja.FunctionCall) goja.Value {
			name := c.Argument(0).String()
			h, ok := r.s.PlayerByName(name)
			if !ok {
				return nil
			}
			return r.vm.ToValue(r.players[h.UUID()].obj)
		}).
		Method("playerByXUID", func(c goja.FunctionCall) goja.Value {
			xuid := c.Argument(0).String()
			h, ok := r.s.PlayerByXUID(xuid)
			if !ok {
				return nil
			}
			return r.vm.ToValue(r.players[h.UUID()].obj)
		}).
		Method("onPlayerJoin", func(c goja.FunctionCall) goja.Value {
			f, ok := goja.AssertFunction(c.Argument(0))
			if !ok {
				panic(r.vm.NewTypeError("onPlayerJoin expects a function as its first argument"))
			}
			r.onPlayerJoin = append(r.onPlayerJoin, callable{f, c.This})
			return nil
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
			var worlds []*goja.Object
			for _, w := range r.worlds {
				worlds = append(worlds, w.obj)
			}
			return r.vm.ToValue(worlds)
		}).
		Apply(r.vm, "server")
}
