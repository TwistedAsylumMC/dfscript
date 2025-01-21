package dfscript

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/enchantment"
	"github.com/dop251/goja"
)

func (r *Runtime) setupEnchantment() error {
	return newObject().
		Const("Rarity", map[string]item.EnchantmentRarity{
			"Common":   item.EnchantmentRarityCommon,
			"Uncommon": item.EnchantmentRarityUncommon,
			"Rare":     item.EnchantmentRarityRare,
			"VeryRare": item.EnchantmentRarityVeryRare,
		}).
		Const("aquaAffinity", enchantment.AquaAffinity).
		Const("blastProtection", enchantment.BlastProtection).
		Const("curseOfVanishing", enchantment.CurseOfVanishing).
		Const("depthStrider", enchantment.DepthStrider).
		Const("efficiency", enchantment.Efficiency).
		Const("featherFalling", enchantment.FeatherFalling).
		Const("fireAspect", enchantment.FireAspect).
		Const("fireProtection", enchantment.FireProtection).
		Const("flame", enchantment.Flame).
		Const("infinity", enchantment.Infinity).
		Const("knockback", enchantment.Knockback).
		Const("mending", enchantment.Mending).
		Const("power", enchantment.Power).
		Const("projectileProtection", enchantment.ProjectileProtection).
		Const("protection", enchantment.Protection).
		Const("punch", enchantment.Punch).
		Const("quickCharge", enchantment.QuickCharge).
		Const("respiration", enchantment.Respiration).
		Const("sharpness", enchantment.Sharpness).
		Const("silkTouch", enchantment.SilkTouch).
		Const("soulSpeed", enchantment.SoulSpeed).
		Const("swiftSneak", enchantment.SwiftSneak).
		Const("thorns", enchantment.Thorns).
		Const("unbreaking", enchantment.Unbreaking).
		Method("create", func(c goja.FunctionCall) goja.Value {
			t, ok := c.Argument(0).Export().(item.EnchantmentType)
			if !ok {
				panic(r.vm.NewTypeError("argument 0 is not an item.EnchantmentType"))
			}
			level, ok := c.Argument(1).Export().(int64)
			if !ok {
				panic(r.vm.NewTypeError("argument 1 is not an int64"))
			}
			return r.vm.ToValue(item.NewEnchantment(t, int(level)))
		}).
		Method("register", func(c goja.FunctionCall) goja.Value {
			id, ok := c.Argument(0).Export().(int64)
			if !ok {
				panic(r.vm.NewTypeError("argument 0 is not an int64"))
			}
			t, ok := c.Argument(1).Export().(item.EnchantmentType)
			if !ok {
				panic(r.vm.NewTypeError("argument 1 is not an item.EnchantmentType"))
			}
			item.RegisterEnchantment(int(id), t)
			return nil
		}).
		Apply(r.vm, "enchantment")
}
