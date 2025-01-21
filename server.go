package dfscript

import (
	"github.com/dop251/goja"
	"github.com/google/uuid"
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
			var id uuid.UUID
			switch arg := c.Argument(0).Export().(type) {
			case string:
				id, _ = uuid.Parse(arg)
			case uuid.UUID:
				id = arg
			default:
				panic(r.vm.NewTypeError("argument 0 must be a string or uuid.UUID"))
			}
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
		Apply(r.vm, "server")
}
