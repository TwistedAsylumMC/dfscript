package dfscript

import (
	"github.com/df-mc/dragonfly/server/entity/effect"
	"github.com/dop251/goja"
	"time"
)

func (r *Runtime) setupEffect() error {
	return newObject().
		Const("absorption", effect.Absorption).
		Const("blindness", effect.Blindness).
		Const("conduitPower", effect.ConduitPower).
		Const("darkness", effect.Darkness).
		Const("fatalPoison", effect.FatalPoison).
		Const("fireResistance", effect.FireResistance).
		Const("haste", effect.Haste).
		Const("healthBoost", effect.HealthBoost).
		Const("hunger", effect.Hunger).
		Const("instantDamage", effect.InstantDamage).
		Const("instantHealth", effect.InstantHealth).
		Const("invisibility", effect.Invisibility).
		Const("jumpBoost", effect.JumpBoost).
		Const("levitation", effect.Levitation).
		Const("miningFatigue", effect.MiningFatigue).
		Const("nausea", effect.Nausea).
		Const("nightVision", effect.NightVision).
		Const("poison", effect.Poison).
		Const("regeneration", effect.Regeneration).
		Const("resistance", effect.Resistance).
		Const("saturation", effect.Saturation).
		Const("slowFalling", effect.SlowFalling).
		Const("slowness", effect.Slowness).
		Const("speed", effect.Speed).
		Const("strength", effect.Strength).
		Const("waterBreathing", effect.WaterBreathing).
		Const("weakness", effect.Weakness).
		Const("wither", effect.Wither).
		Method("byId", func(c goja.FunctionCall) goja.Value {
			id, ok := c.Argument(0).Export().(int64)
			if !ok {
				return nil
			}
			t, ok := effect.ByID(int(id))
			if !ok {
				return nil
			}
			return r.vm.ToValue(t)
		}).
		Method("id", func(c goja.FunctionCall) goja.Value {
			t, ok := c.Argument(0).Export().(effect.Type)
			if !ok {
				return nil
			}
			id, ok := effect.ID(t)
			if !ok {
				return nil
			}
			return r.vm.ToValue(id)
		}).
		Method("create", func(c goja.FunctionCall) goja.Value {
			t, ok := c.Argument(0).Export().(effect.LastingType)
			if !ok {
				return nil
			}
			lvl, ok := c.Argument(1).Export().(int64)
			if !ok {
				return nil
			}
			d, ok := c.Argument(2).Export().(time.Duration)
			if !ok {
				return nil
			}
			return r.vm.ToValue(effect.New(t, int(lvl), d))
		}).
		Method("createAmbient", func(c goja.FunctionCall) goja.Value {
			t, ok := c.Argument(0).Export().(effect.LastingType)
			if !ok {
				return nil
			}
			lvl, ok := c.Argument(1).Export().(int64)
			if !ok {
				return nil
			}
			d, ok := c.Argument(2).Export().(time.Duration)
			if !ok {
				return nil
			}
			return r.vm.ToValue(effect.NewAmbient(t, int(lvl), d))
		}).
		Method("createInstant", func(c goja.FunctionCall) goja.Value {
			t, ok := c.Argument(0).Export().(effect.Type)
			if !ok {
				return nil
			}
			lvl, ok := c.Argument(1).Export().(int64)
			if !ok {
				return nil
			}
			potency, ok := c.Argument(2).Export().(float64)
			if !ok {
				potency = 1.0
			}
			return r.vm.ToValue(effect.NewInstantWithPotency(t, int(lvl), potency))
		}).
		Apply(r.vm, "effect")
}
