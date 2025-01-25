package dfscript

import (
	"fmt"
	"github.com/df-mc/dragonfly/server"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/df-mc/dragonfly/server/world"
	"github.com/dop251/goja"
	"github.com/dop251/goja_nodejs/buffer"
	"github.com/dop251/goja_nodejs/console"
	"github.com/dop251/goja_nodejs/process"
	"github.com/dop251/goja_nodejs/require"
	"github.com/dop251/goja_nodejs/url"
	"github.com/google/uuid"
	"go.uber.org/multierr"
	"reflect"
	"strings"
	"unicode"
)

type callable struct {
	id   int
	f    goja.Callable
	this goja.Value
}

type Runtime struct {
	s *server.Server

	vm   *goja.Runtime
	reg  *require.Registry
	loop *EventLoop

	worlds       map[string]*world.World
	eventHandles map[string][]callable
	nextEventID  int
}

type uncapFieldNameMapper struct{}

func (u uncapFieldNameMapper) FieldName(_ reflect.Type, f reflect.StructField) string {
	return uncap(f.Name)
}

func (u uncapFieldNameMapper) MethodName(_ reflect.Type, m reflect.Method) string {
	return uncap(m.Name)
}

func uncap(s string) string {
	b := strings.Builder{}
	for i, k := range s {
		if unicode.IsUpper(k) {
			b.WriteRune(unicode.ToLower(k))
			continue
		}
		b.WriteString(s[i:])
		break
	}
	return b.String()
}

func NewRuntime(s *server.Server) (*Runtime, error) {
	r := &Runtime{
		s: s,

		vm:  goja.New(),
		reg: new(require.Registry),

		worlds:       make(map[string]*world.World),
		eventHandles: make(map[string][]callable),
	}
	r.vm.SetFieldNameMapper(uncapFieldNameMapper{})
	r.reg.Enable(r.vm)
	r.loop = NewEventLoop(r.vm, EnableConsole(true), WithRegistry(r.reg))
	r.loop.Start()

	buffer.Enable(r.vm)
	console.Enable(r.vm)
	process.Enable(r.vm)
	url.Enable(r.vm)

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

	multierr.AppendInto(&err, r.setupEvent())

	multierr.AppendInto(&err, r.setupMgl64())
	multierr.AppendInto(&err, r.setupStd())

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

func (r *Runtime) VM() *goja.Runtime {
	return r.vm
}

func (r *Runtime) Loop() *EventLoop {
	return r.loop
}

func (r *Runtime) AddWorld(w *world.World) world.Handler {
	r.worlds[w.Name()] = w
	return &WorldHandler{r: r}
}

func (r *Runtime) World(name string) *world.World {
	if w, ok := r.worlds[name]; ok {
		return w
	}
	return nil
}

func (r *Runtime) RemoveWorld(name string) {
	delete(r.worlds, name)
}

func (r *Runtime) PlayerJoin(p *player.Player) player.Handler {
	calls, ok := r.eventHandles["onPlayerJoin"]
	if ok {
		for _, c := range calls {
			_, err := c.f(c.this, r.vm.ToValue(p))
			if err != nil {
				panic(err)
			}
		}
	}
	return &PlayerHandler{r: r}
}

func (r *Runtime) Close() {
	r.loop.Stop()
	r.vm.Interrupt(fmt.Errorf("runtime closed"))

	r.s = nil
	r.vm = nil
	r.reg = nil
	r.loop = nil
	r.worlds = nil
	r.eventHandles = nil
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
