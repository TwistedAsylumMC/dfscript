package dfscript

import (
	"github.com/df-mc/dragonfly/server/world"
	"github.com/df-mc/dragonfly/server/world/sound"
	"github.com/dop251/goja"
)

func (r *Runtime) setupSound() error {
	return newObject().
		Const("instrumentType", map[string]sound.Instrument{
			"banjo":           sound.Banjo(),
			"bass":            sound.Bass(),
			"bassDrum":        sound.BassDrum(),
			"bell":            sound.Bell(),
			"bit":             sound.Bit(),
			"chimes":          sound.Chimes(),
			"clicksAndSticks": sound.ClicksAndSticks(),
			"cowBell":         sound.CowBell(),
			"didgeridoo":      sound.Didgeridoo(),
			"flute":           sound.Flute(),
			"guitar":          sound.Guitar(),
			"ironXylophone":   sound.IronXylophone(),
			"piano":           sound.Piano(),
			"pling":           sound.Pling(),
			"snare":           sound.Snare(),
			"xylophone":       sound.Xylophone(),
		}).
		Const("discType", map[string]sound.DiscType{
			"13":              sound.Disc13(),
			"cat":             sound.DiscCat(),
			"blocks":          sound.DiscBlocks(),
			"chirp":           sound.DiscChirp(),
			"far":             sound.DiscFar(),
			"mall":            sound.DiscMall(),
			"mellohi":         sound.DiscMellohi(),
			"stal":            sound.DiscStal(),
			"strad":           sound.DiscStrad(),
			"ward":            sound.DiscWard(),
			"11":              sound.Disc11(),
			"wait":            sound.DiscWait(),
			"otherside":       sound.DiscOtherside(),
			"pigstep":         sound.DiscPigstep(),
			"5":               sound.Disc5(),
			"relic":           sound.DiscRelic(),
			"creator":         sound.DiscCreator(),
			"creatorMusicBox": sound.DiscCreatorMusicBox(),
			"precipice":       sound.DiscPrecipice(),
		}).
		Const("hornType", map[string]sound.Horn{
			"ponder": sound.Ponder(),
			"sing":   sound.Sing(),
			"seek":   sound.Seek(),
			"feel":   sound.Feel(),
			"admire": sound.Admire(),
			"call":   sound.Call(),
			"yearn":  sound.Yearn(),
			"dream":  sound.Dream(),
		}).
		Const("crossbowLoading", map[string]int{
			"start":  sound.CrossbowLoadingStart,
			"middle": sound.CrossbowLoadingMiddle,
			"end":    sound.CrossbowLoadingEnd,
		}).
		Method("anvilBreak", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(sound.AnvilBreak{})
		}).
		Method("anvilLand", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(sound.AnvilLand{})
		}).
		Method("anvilUse", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(sound.AnvilUse{})
		}).
		Method("arrowHit", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(sound.ArrowHit{})
		}).
		Method("attack", func(c goja.FunctionCall) goja.Value {
			damage, _ := c.Argument(0).Export().(bool)
			return r.vm.ToValue(sound.Attack{Damage: damage})
		}).
		Method("barrelClose", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(sound.BarrelClose{})
		}).
		Method("barrelOpen", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(sound.BarrelOpen{})
		}).
		Method("blastFurnaceCrackle", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(sound.BlastFurnaceCrackle{})
		}).
		Method("blockBreaking", func(c goja.FunctionCall) goja.Value {
			b, ok := c.Argument(0).Export().(world.Block)
			if !ok {
				panic(r.vm.NewTypeError("argument 0 must be a world.Block"))
			}
			return r.vm.ToValue(sound.BlockBreaking{Block: b})
		}).
		Method("blockPlace", func(c goja.FunctionCall) goja.Value {
			b, ok := c.Argument(0).Export().(world.Block)
			if !ok {
				panic(r.vm.NewTypeError("argument 0 must be a world.Block"))
			}
			return r.vm.ToValue(sound.BlockPlace{Block: b})
		}).
		Method("bowShoot", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(sound.BowShoot{})
		}).
		Method("bucketEmpty", func(c goja.FunctionCall) goja.Value {
			l, ok := c.Argument(0).Export().(world.Liquid)
			if !ok {
				panic(r.vm.NewTypeError("argument 0 must be a world.Liquid"))
			}
			return r.vm.ToValue(sound.BucketEmpty{Liquid: l})
		}).
		Method("bucketFill", func(c goja.FunctionCall) goja.Value {
			l, ok := c.Argument(0).Export().(world.Liquid)
			if !ok {
				panic(r.vm.NewTypeError("argument 0 must be a world.Liquid"))
			}
			return r.vm.ToValue(sound.BucketFill{Liquid: l})
		}).
		Method("burning", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(sound.Burning{})
		}).
		Method("burp", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(sound.Burp{})
		}).
		Method("campfireCrackle", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(sound.CampfireCrackle{})
		}).
		Method("chestClose", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(sound.ChestClose{})
		}).
		Method("chestOpen", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(sound.ChestOpen{})
		}).
		Method("click", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(sound.Click{})
		}).
		Method("composterEmpty", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(sound.ComposterEmpty{})
		}).
		Method("composterFill", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(sound.ComposterFill{})
		}).
		Method("composterFillLayer", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(sound.ComposterFillLayer{})
		}).
		Method("composterReady", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(sound.ComposterReady{})
		}).
		Method("copperScraped", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(sound.CopperScraped{})
		}).
		Method("crossbowLoad", func(c goja.FunctionCall) goja.Value {
			stage, ok := c.Argument(0).Export().(int64)
			if !ok {
				panic(r.vm.NewTypeError("argument 0 must be an int64"))
			}
			quickCharge, _ := c.Argument(1).Export().(bool)
			return r.vm.ToValue(sound.CrossbowLoad{Stage: int(stage), QuickCharge: quickCharge})
		}).
		Method("crossbowShoot", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(sound.CrossbowShoot{})
		}).
		Method("decoratedPotInsertFailed", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(sound.DecoratedPotInsertFailed{})
		}).
		Method("decoratedPotInserted", func(c goja.FunctionCall) goja.Value {
			progress, ok := c.Argument(0).Export().(float64)
			if !ok {
				panic(r.vm.NewTypeError("argument 0 must be a float64"))
			}
			return r.vm.ToValue(sound.DecoratedPotInserted{Progress: progress})
		}).
		Method("deny", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(sound.Deny{})
		}).
		Method("doorClose", func(c goja.FunctionCall) goja.Value {
			b, ok := c.Argument(0).Export().(world.Block)
			if !ok {
				panic(r.vm.NewTypeError("argument 0 must be a world.Block"))
			}
			return r.vm.ToValue(sound.DoorClose{Block: b})
		}).
		Method("doorCrash", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(sound.DoorCrash{})
		}).
		Method("doorOpen", func(c goja.FunctionCall) goja.Value {
			b, ok := c.Argument(0).Export().(world.Block)
			if !ok {
				panic(r.vm.NewTypeError("argument 0 must be a world.Block"))
			}
			return r.vm.ToValue(sound.DoorOpen{Block: b})
		}).
		Method("drowning", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(sound.Drowning{})
		}).
		Method("enderChestClose", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(sound.EnderChestClose{})
		}).
		Method("enderChestOpen", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(sound.EnderChestOpen{})
		}).
		Method("equipItem", func(c goja.FunctionCall) goja.Value {
			it, ok := c.Argument(0).Export().(world.Item)
			if !ok {
				panic(r.vm.NewTypeError("argument 0 must be a world.Item"))
			}
			return r.vm.ToValue(sound.EquipItem{Item: it})
		}).
		Method("experience", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(sound.Experience{})
		}).
		Method("explosion", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(sound.Explosion{})
		}).
		Method("fall", func(c goja.FunctionCall) goja.Value {
			distance, ok := c.Argument(0).Export().(float64)
			if !ok {
				panic(r.vm.NewTypeError("argument 0 must be a float64"))
			}
			return r.vm.ToValue(sound.Fall{Distance: distance})
		}).
		Method("fenceGateClose", func(c goja.FunctionCall) goja.Value {
			b, ok := c.Argument(0).Export().(world.Block)
			if !ok {
				panic(r.vm.NewTypeError("argument 0 must be a world.Block"))
			}
			return r.vm.ToValue(sound.FenceGateClose{Block: b})
		}).
		Method("fenceGateOpen", func(c goja.FunctionCall) goja.Value {
			b, ok := c.Argument(0).Export().(world.Block)
			if !ok {
				panic(r.vm.NewTypeError("argument 0 must be a world.Block"))
			}
			return r.vm.ToValue(sound.FenceGateOpen{Block: b})
		}).
		Method("fireCharge", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(sound.FireCharge{})
		}).
		Method("fireExtinguish", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(sound.FireExtinguish{})
		}).
		Method("fireworkBlast", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(sound.FireworkBlast{})
		}).
		Method("fireworkHugeBlast", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(sound.FireworkHugeBlast{})
		}).
		Method("fireworkLaunch", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(sound.FireworkLaunch{})
		}).
		Method("fireworkTwinkle", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(sound.FireworkTwinkle{})
		}).
		Method("fizz", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(sound.Fizz{})
		}).
		Method("furnaceCrackle", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(sound.FurnaceCrackle{})
		}).
		Method("ghastShoot", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(sound.GhastShoot{})
		}).
		Method("ghastWarning", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(sound.GhastWarning{})
		}).
		Method("glassBreak", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(sound.GlassBreak{})
		}).
		Method("goatHorn", func(c goja.FunctionCall) goja.Value {
			horn, ok := c.Argument(0).Export().(sound.Horn)
			if !ok {
				panic(r.vm.NewTypeError("argument 0 must be a sound.Horn"))
			}
			return r.vm.ToValue(sound.GoatHorn{Horn: horn})
		}).
		Method("ignite", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(sound.Ignite{})
		}).
		Method("instrument", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(sound.Instrument{})
		}).
		Method("itemAdd", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(sound.ItemAdd{})
		}).
		Method("itemBreak", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(sound.ItemBreak{})
		}).
		Method("itemFrameRemove", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(sound.ItemFrameRemove{})
		}).
		Method("itemFrameRotate", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(sound.ItemFrameRotate{})
		}).
		Method("itemThrow", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(sound.ItemThrow{})
		}).
		Method("itemUseOn", func(c goja.FunctionCall) goja.Value {
			b, ok := c.Argument(0).Export().(world.Block)
			if !ok {
				panic(r.vm.NewTypeError("argument 0 must be a world.Block"))
			}
			return r.vm.ToValue(sound.ItemUseOn{Block: b})
		}).
		Method("lecternBookPlace", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(sound.LecternBookPlace{})
		}).
		Method("levelUp", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(sound.LevelUp{})
		}).
		Method("musicDiscEnd", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(sound.MusicDiscEnd{})
		}).
		Method("musicDiscPlay", func(c goja.FunctionCall) goja.Value {
			disc, ok := c.Argument(0).Export().(sound.DiscType)
			if !ok {
				panic(r.vm.NewTypeError("argument 0 must be a sound.DiscType"))
			}
			return r.vm.ToValue(sound.MusicDiscPlay{DiscType: disc})
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
			return r.vm.ToValue(sound.Note{Instrument: instrument, Pitch: int(pitch)})
		}).
		Method("pop", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(sound.Pop{})
		}).
		Method("potionBrewed", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(sound.PotionBrewed{})
		}).
		Method("signWaxed", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(sound.SignWaxed{})
		}).
		Method("smokerCrackle", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(sound.SmokerCrackle{})
		}).
		Method("stopUsingSpyglass", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(sound.StopUsingSpyglass{})
		}).
		Method("tnt", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(sound.TNT{})
		}).
		Method("teleport", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(sound.Teleport{})
		}).
		Method("thunder", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(sound.Thunder{})
		}).
		Method("totem", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(sound.Totem{})
		}).
		Method("trapdoorClose", func(c goja.FunctionCall) goja.Value {
			b, ok := c.Argument(0).Export().(world.Block)
			if !ok {
				panic(r.vm.NewTypeError("argument 0 must be a world.Block"))
			}
			return r.vm.ToValue(sound.TrapdoorClose{Block: b})
		}).
		Method("trapdoorOpen", func(c goja.FunctionCall) goja.Value {
			b, ok := c.Argument(0).Export().(world.Block)
			if !ok {
				panic(r.vm.NewTypeError("argument 0 must be a world.Block"))
			}
			return r.vm.ToValue(sound.TrapdoorOpen{Block: b})
		}).
		Method("useSpyglass", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(sound.UseSpyglass{})
		}).
		Method("waxRemoved", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(sound.WaxRemoved{})
		}).
		Method("waxedSignFailedInteraction", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(sound.WaxedSignFailedInteraction{})
		}).
		Apply(r.vm, "sound")
}
