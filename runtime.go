package dfscript

import (
	"github.com/df-mc/dragonfly/server"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/dop251/goja"
	"github.com/dop251/goja_nodejs/buffer"
	"github.com/dop251/goja_nodejs/console"
	"github.com/dop251/goja_nodejs/process"
	"github.com/dop251/goja_nodejs/require"
	"github.com/dop251/goja_nodejs/url"
	"github.com/google/uuid"
	"go.uber.org/multierr"
)

type callable struct {
	f    goja.Callable
	this goja.Value
}

type Runtime struct {
	s *server.Server

	vm   *goja.Runtime
	reg  *require.Registry
	loop *EventLoop

	worlds       map[string]*World
	players      map[uuid.UUID]*Player
	onPlayerJoin []callable
}

func NewRuntime(s *server.Server) (*Runtime, error) {
	r := &Runtime{
		s: s,

		vm:  goja.New(),
		reg: new(require.Registry),

		worlds:  make(map[string]*World),
		players: make(map[uuid.UUID]*Player),
	}
	r.vm.SetFieldNameMapper(goja.UncapFieldNameMapper())
	r.reg.Enable(r.vm)
	r.loop = NewEventLoop(r.vm, EnableConsole(true), WithRegistry(r.reg))
	r.loop.Start()

	buffer.Enable(r.vm)
	console.Enable(r.vm)
	process.Enable(r.vm)
	url.Enable(r.vm)

	r.worlds[s.World().Name()] = NewWorld(r, s.World())

	var err error
	multierr.AppendInto(&err, r.setupBiome())
	multierr.AppendInto(&err, r.setupBlock())
	multierr.AppendInto(&err, r.setupBossBar())
	multierr.AppendInto(&err, r.setupCategory())
	multierr.AppendInto(&err, r.setupChat())
	multierr.AppendInto(&err, r.setupChunk())
	multierr.AppendInto(&err, r.setupCreative())
	multierr.AppendInto(&err, r.setupCube())
	multierr.AppendInto(&err, r.setupEffect())
	multierr.AppendInto(&err, r.setupEnchantment())
	// TODO: entity
	multierr.AppendInto(&err, r.setupForm())
	multierr.AppendInto(&err, r.setupGenerator())
	multierr.AppendInto(&err, r.setupItem())
	// TODO: item/inventory
	// TODO: world/mcdb/leveldat
	// TODO: world/mcdb
	multierr.AppendInto(&err, r.setupModel())
	multierr.AppendInto(&err, r.setupParticle())
	multierr.AppendInto(&err, r.setupPotion())
	multierr.AppendInto(&err, r.setupRecipe())
	multierr.AppendInto(&err, r.setupScoreboard())
	multierr.AppendInto(&err, r.setupServer())
	multierr.AppendInto(&err, r.setupSkin())
	multierr.AppendInto(&err, r.setupSound())
	multierr.AppendInto(&err, r.setupText())
	multierr.AppendInto(&err, r.setupTitle())
	multierr.AppendInto(&err, r.setupWorld())

	multierr.AppendInto(&err, r.setupMgl64())

	if err != nil {
		return nil, err
	}
	return r, nil
}

func (r *Runtime) Run(script string) bool {
	return r.loop.RunOnLoop(func(vm *goja.Runtime) {
		_, err := vm.RunString(script)
		if err != nil {
			panic(err)
		}
	})
}

func (r *Runtime) World(name string) *World {
	if w, ok := r.worlds[name]; ok {
		return w
	}
	return nil
}

func (r *Runtime) WorldFromValue(v goja.Value) *World {
	if v == nil {
		return nil
	}
	m, ok := v.Export().(map[string]any)
	if !ok {
		return nil
	}
	n, ok := m["name"].(string)
	if !ok {
		return nil
	}
	return r.World(n)
}

func (r *Runtime) Player(uuid uuid.UUID) *Player {
	if p, ok := r.players[uuid]; ok {
		return p
	}
	return nil
}

func (r *Runtime) PlayerFromValue(v goja.Value) *Player {
	if v == nil {
		return nil
	}
	m, ok := v.Export().(map[string]any)
	if !ok {
		return nil
	}
	u, ok := m["uuid"].(uuid.UUID)
	if !ok {
		return nil
	}
	return r.Player(u)
}

func (r *Runtime) PlayerJoin(p *player.Player) player.Handler {
	pl := NewPlayer(r, p)
	r.players[p.UUID()] = pl
	for _, c := range r.onPlayerJoin {
		_, err := c.f(c.this, r.vm.ToValue(r.Player(p.UUID()).obj))
		if err != nil {
			panic(err)
		}
	}
	pl.p = nil
	return pl
}

func (r *Runtime) exportUUID(c goja.FunctionCall, idx int) uuid.UUID {
	switch arg := c.Argument(idx).Export().(type) {
	case string:
		return uuid.MustParse(arg)
	case uuid.UUID:
		return arg
	default:
		panic(r.vm.NewTypeError("argument %d must be a string or uuid.UUID, got %T", idx, c.Argument(idx).Export()))
	}
}
