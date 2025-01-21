package dfscript

import (
	"github.com/df-mc/dragonfly/server/block/cube"
	"github.com/df-mc/dragonfly/server/world"
	"github.com/df-mc/dragonfly/server/world/particle"
	"github.com/df-mc/dragonfly/server/world/sound"
	"github.com/dop251/goja"
	"image/color"
)

func (r *Runtime) setupParticle() error {
	return newObject().
		Method("blockBreak", func(c goja.FunctionCall) goja.Value {
			b, ok := c.Argument(0).Export().(world.Block)
			if !ok {
				panic(r.vm.NewTypeError("argument 0 must be a world.Block"))
			}
			return r.vm.ToValue(particle.BlockBreak{Block: b})
		}).
		Method("blockForceField", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(particle.BlockForceField{})
		}).
		Method("boneMeal", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(particle.BoneMeal{})
		}).
		Method("dragonEggTeleport", func(c goja.FunctionCall) goja.Value {
			diff, ok := c.Argument(0).Export().(cube.Pos)
			if !ok {
				panic(r.vm.NewTypeError("argument 0 must be a cube.Pos"))
			}
			return r.vm.ToValue(particle.DragonEggTeleport{Diff: diff})
		}).
		Method("dust", func(c goja.FunctionCall) goja.Value {
			colour, ok := c.Argument(0).Export().(color.RGBA)
			if !ok {
				panic(r.vm.NewTypeError("argument 0 must be a color.RGBA"))
			}
			return r.vm.ToValue(particle.Dust{Colour: colour})
		}).
		Method("dustPlume", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(particle.DustPlume{})
		}).
		Method("effect", func(c goja.FunctionCall) goja.Value {
			colour, ok := c.Argument(0).Export().(color.RGBA)
			if !ok {
				panic(r.vm.NewTypeError("argument 0 must be a color.RGBA"))
			}
			return r.vm.ToValue(particle.Effect{Colour: colour})
		}).
		Method("eggSmash", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(particle.EggSmash{})
		}).
		Method("endermanTeleport", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(particle.EndermanTeleport{})
		}).
		Method("entityFlame", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(particle.EntityFlame{})
		}).
		Method("evaporate", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(particle.Evaporate{})
		}).
		Method("flame", func(c goja.FunctionCall) goja.Value {
			colour, ok := c.Argument(0).Export().(color.RGBA)
			if !ok {
				panic(r.vm.NewTypeError("argument 0 must be a color.RGBA"))
			}
			return r.vm.ToValue(particle.Flame{Colour: colour})
		}).
		Method("hugeExplosion", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(particle.HugeExplosion{})
		}).
		Method("lava", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(particle.Lava{})
		}).
		Method("lavaDrip", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(particle.LavaDrip{})
		}).
		Method("note", func(c goja.FunctionCall) goja.Value {
			instrument, ok := c.Argument(0).Export().(sound.Instrument)
			if !ok {
				panic(r.vm.NewTypeError("argument 0 must be a sound.Instrument"))
			}
			pitch, ok := c.Argument(1).Export().(int64)
			if !ok {
				panic(r.vm.NewTypeError("argument 1 must be an int64"))
			}
			return r.vm.ToValue(particle.Note{Instrument: instrument, Pitch: int(pitch)})
		}).
		Method("punchBlock", func(c goja.FunctionCall) goja.Value {
			b, ok := c.Argument(0).Export().(world.Block)
			if !ok {
				panic(r.vm.NewTypeError("argument 0 must be a world.Block"))
			}
			face, ok := c.Argument(1).Export().(cube.Face)
			if !ok {
				panic(r.vm.NewTypeError("argument 1 must be a cube.Face"))
			}
			return r.vm.ToValue(particle.PunchBlock{Block: b, Face: face})
		}).
		Method("snowballPoof", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(particle.SnowballPoof{})
		}).
		Method("splash", func(c goja.FunctionCall) goja.Value {
			colour, ok := c.Argument(0).Export().(color.RGBA)
			if !ok {
				panic(r.vm.NewTypeError("argument 0 must be a color.RGBA"))
			}
			return r.vm.ToValue(particle.Splash{Colour: colour})
		}).
		Method("waterDrip", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(particle.WaterDrip{})
		}).
		Apply(r.vm, "particle")
}
