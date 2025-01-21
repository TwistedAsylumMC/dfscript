package dfscript

import (
	"github.com/df-mc/dragonfly/server/item/potion"
	"github.com/dop251/goja"
)

func (r *Runtime) setupPotion() error {
	return newObject().
		Const("all", potion.All()).
		Const("awkward", potion.Awkward()).
		Const("fireResistance", potion.FireResistance()).
		Const("harming", potion.Harming()).
		Const("healing", potion.Healing()).
		Const("invisibility", potion.Invisibility()).
		Const("leaping", potion.Leaping()).
		Const("longFireResistance", potion.LongFireResistance()).
		Const("longInvisibility", potion.LongInvisibility()).
		Const("longLeaping", potion.LongLeaping()).
		Const("longMundane", potion.LongMundane()).
		Const("longNightVision", potion.LongNightVision()).
		Const("longPoison", potion.LongPoison()).
		Const("longRegeneration", potion.LongRegeneration()).
		Const("longSlowFalling", potion.LongSlowFalling()).
		Const("longSlowness", potion.LongSlowness()).
		Const("longStrength", potion.LongStrength()).
		Const("longSwiftness", potion.LongSwiftness()).
		Const("longTurtleMaster", potion.LongTurtleMaster()).
		Const("longWaterBreathing", potion.LongWaterBreathing()).
		Const("longWeakness", potion.LongWeakness()).
		Const("mundane", potion.Mundane()).
		Const("nightVision", potion.NightVision()).
		Const("poison", potion.Poison()).
		Const("regeneration", potion.Regeneration()).
		Const("slowFalling", potion.SlowFalling()).
		Const("slowness", potion.Slowness()).
		Const("strength", potion.Strength()).
		Const("strongHarming", potion.StrongHarming()).
		Const("strongHealing", potion.StrongHealing()).
		Const("strongLeaping", potion.StrongLeaping()).
		Const("strongPoison", potion.StrongPoison()).
		Const("strongRegeneration", potion.StrongRegeneration()).
		Const("strongSlowness", potion.StrongSlowness()).
		Const("strongStrength", potion.StrongStrength()).
		Const("strongSwiftness", potion.StrongSwiftness()).
		Const("strongTurtleMaster", potion.StrongTurtleMaster()).
		Const("swiftness", potion.Swiftness()).
		Const("thick", potion.Thick()).
		Const("turtleMaster", potion.TurtleMaster()).
		Const("water", potion.Water()).
		Const("waterBreathing", potion.WaterBreathing()).
		Const("weakness", potion.Weakness()).
		Const("wither", potion.Wither()).
		Method("from", func(c goja.FunctionCall) goja.Value {
			id, ok := c.Argument(0).Export().(int64)
			if !ok {
				panic(r.vm.NewTypeError("argument 0 must be an int64"))
			}
			return r.vm.ToValue(potion.From(int32(id)))
		}).
		Apply(r.vm, "potion")
}
