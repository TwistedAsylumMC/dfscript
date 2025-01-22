package dfscript

import (
	"github.com/df-mc/dragonfly/server/block"
	"github.com/df-mc/dragonfly/server/block/cube"
	"github.com/df-mc/dragonfly/server/entity/effect"
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/world"
	"github.com/dop251/goja"
	"image/color"
	"time"
)

func (r *Runtime) setupBlock() error {
	return newObject().
		Const("AnvilType", map[string]block.AnvilType{
			"Undamaged":       block.UndamagedAnvil(),
			"SlightlyDamaged": block.SlightlyDamagedAnvil(),
			"VeryDamaged":     block.VeryDamagedAnvil(),
		}).
		Const("BannerPattern", map[string]block.BannerPatternType{
			"Border":               block.BorderBannerPattern(),
			"Bricks":               block.BricksBannerPattern(),
			"Circle":               block.CircleBannerPattern(),
			"Creeper":              block.CreeperBannerPattern(),
			"Cross":                block.CrossBannerPattern(),
			"CurlyBorder":          block.CurlyBorderBannerPattern(),
			"DiagonalLeft":         block.DiagonalLeftBannerPattern(),
			"DiagonalRight":        block.DiagonalRightBannerPattern(),
			"DiagonalUpLeft":       block.DiagonalUpLeftBannerPattern(),
			"DiagonalUpRight":      block.DiagonalUpRightBannerPattern(),
			"Flower":               block.FlowerBannerPattern(),
			"Gradient":             block.GradientBannerPattern(),
			"GradientUp":           block.GradientUpBannerPattern(),
			"HalfHorizontal":       block.HalfHorizontalBannerPattern(),
			"HalfHorizontalBottom": block.HalfHorizontalBottomBannerPattern(),
			"HalfVertical":         block.HalfVerticalBannerPattern(),
			"HalfVerticalRight":    block.HalfVerticalRightBannerPattern(),
			"Mojang":               block.MojangBannerPattern(),
			"Rhombus":              block.RhombusBannerPattern(),
			"Skull":                block.SkullBannerPattern(),
			"SmallStripes":         block.SmallStripesBannerPattern(),
			"SquareBottomLeft":     block.SquareBottomLeftBannerPattern(),
			"SquareBottomRight":    block.SquareBottomRightBannerPattern(),
			"SquareTopLeft":        block.SquareTopLeftBannerPattern(),
			"SquareTopRight":       block.SquareTopRightBannerPattern(),
			"StraightCross":        block.StraightCrossBannerPattern(),
			"StripeBottom":         block.StripeBottomBannerPattern(),
			"StripeCenter":         block.StripeCenterBannerPattern(),
			"StripeDownLeft":       block.StripeDownLeftBannerPattern(),
			"StripeDownRight":      block.StripeDownRightBannerPattern(),
			"StripeLeft":           block.StripeLeftBannerPattern(),
			"StripeMiddle":         block.StripeMiddleBannerPattern(),
			"StripeRight":          block.StripeRightBannerPattern(),
			"StripeTop":            block.StripeTopBannerPattern(),
			"TriangleBottom":       block.TriangleBottomBannerPattern(),
			"TriangleTop":          block.TriangleTopBannerPattern(),
			"TrianglesBottom":      block.TrianglesBottomBannerPattern(),
			"TrianglesTop":         block.TrianglesTopBannerPattern(),
			"Globe":                block.GlobeBannerPattern(),
			"Piglin":               block.PiglinBannerPattern(),
			"Flow":                 block.FlowBannerPattern(),
			"Guster":               block.GusterBannerPattern(),
		}).
		Const("BlackstoneType", map[string]block.BlackstoneType{
			"Normal":           block.NormalBlackstone(),
			"Gilded":           block.GildedBlackstone(),
			"Polished":         block.PolishedBlackstone(),
			"ChiseledPolished": block.ChiseledPolishedBlackstone(),
		}).
		Const("CopperType", map[string]block.CopperType{
			"Normal":   block.NormalCopper(),
			"Cut":      block.CutCopper(),
			"Chiseled": block.ChiseledCopper(),
		}).
		Const("CoralType", map[string]block.CoralType{
			"Tube":   block.TubeCoral(),
			"Brain":  block.BrainCoral(),
			"Bubble": block.BubbleCoral(),
			"Fire":   block.FireCoral(),
			"Horn":   block.HornCoral(),
		}).
		Const("DeepslateType", map[string]block.DeepslateType{
			"Normal":   block.NormalDeepslate(),
			"Cobbled":  block.CobbledDeepslate(),
			"Polished": block.PolishedDeepslate(),
			"Chiseled": block.ChiseledDeepslate(),
		}).
		Const("FireType", map[string]block.FireType{
			"Normal": block.NormalFire(),
			"Soul":   block.SoulFire(),
		}).
		Const("DoubleFlowerType", map[string]block.DoubleFlowerType{
			"Sunflower": block.Sunflower(),
			"Lilac":     block.Lilac(),
			"RoseBush":  block.RoseBush(),
			"Peony":     block.Peony(),
		}).
		Const("DoubleTallGrassType", map[string]block.DoubleTallGrassType{
			"Normal": block.NormalDoubleTallGrass(),
			"Fern":   block.FernDoubleTallGrass(),
		}).
		Const("FlowerType", map[string]block.FlowerType{
			"Dandelion":       block.Dandelion(),
			"Poppy":           block.Poppy(),
			"BlueOrchid":      block.BlueOrchid(),
			"Allium":          block.Allium(),
			"AzureBluet":      block.AzureBluet(),
			"RedTulip":        block.RedTulip(),
			"OrangeTulip":     block.OrangeTulip(),
			"WhiteTulip":      block.WhiteTulip(),
			"PinkTulip":       block.PinkTulip(),
			"OxeyeDaisy":      block.OxeyeDaisy(),
			"Cornflower":      block.Cornflower(),
			"LilyOfTheValley": block.LilyOfTheValley(),
			"WitherRose":      block.WitherRose(),
		}).
		Const("FroglightType", map[string]block.FroglightType{
			"Pearlescent": block.Pearlescent(),
			"Verdant":     block.Verdant(),
			"Ochre":       block.Ochre(),
		}).
		Const("NetherBricksType", map[string]block.NetherBricksType{
			"Normal":   block.NormalNetherBricks(),
			"Red":      block.RedNetherBricks(),
			"Cracked":  block.CrackedNetherBricks(),
			"Chiseled": block.ChiseledNetherBricks(),
		}).
		Const("OreType", map[string]block.OreType{
			"Stone":     block.StoneOre(),
			"Deepslate": block.DeepslateOre(),
		}).
		Const("Oxidation", map[string]block.OxidationType{
			"Unoxidised": block.UnoxidisedOxidation(),
			"Exposed":    block.ExposedOxidation(),
			"Weathered":  block.WeatheredOxidation(),
			"Oxidised":   block.OxidisedOxidation(),
		}).
		Const("PrismarineType", map[string]block.PrismarineType{
			"Normal": block.NormalPrismarine(),
			"Dark":   block.DarkPrismarine(),
			"Bricks": block.BrickPrismarine(),
		}).
		Const("SandstoneType", map[string]block.SandstoneType{
			"Normal":   block.NormalSandstone(),
			"Cut":      block.CutSandstone(),
			"Chiseled": block.ChiseledSandstone(),
			"Smooth":   block.SmoothSandstone(),
		}).
		Const("SkullType", map[string]block.SkullType{
			"Skeleton":       block.SkeletonSkull(),
			"WitherSkeleton": block.WitherSkeletonSkull(),
			"Zombie":         block.ZombieHead(),
			"Player":         block.PlayerHead(),
			"Creeper":        block.CreeperHead(),
			"Dragon":         block.DragonHead(),
			"Piglin":         block.PiglinHead(),
		}).
		Const("StoneBricksType", map[string]block.StoneBricksType{
			"Normal":   block.NormalStoneBricks(),
			"Mossy":    block.MossyStoneBricks(),
			"Cracked":  block.CrackedStoneBricks(),
			"Chiseled": block.ChiseledStoneBricks(),
		}).
		Const("WallConnection", map[string]block.WallConnectionType{
			"None":  block.NoWallConnection(),
			"Short": block.ShortWallConnection(),
			"Tall":  block.TallWallConnection(),
		}).
		Const("WoodType", map[string]block.WoodType{
			"Oak":      block.OakWood(),
			"Spruce":   block.SpruceWood(),
			"Birch":    block.BirchWood(),
			"Jungle":   block.JungleWood(),
			"Acacia":   block.AcaciaWood(),
			"DarkOak":  block.DarkOakWood(),
			"Crimson":  block.CrimsonWood(),
			"Warped":   block.WarpedWood(),
			"Mangrove": block.MangroveWood(),
			"Cherry":   block.CherryWood(),
			"PaleOak":  block.PaleOakWood(),
		}).
		Const("GrindstoneAttachment", map[string]block.GrindstoneAttachment{
			"Standing": block.StandingGrindstoneAttachment(),
			"Hanging":  block.HangingGrindstoneAttachment(),
			"Wall":     block.WallGrindstoneAttachment(),
		}).
		Method("air", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(block.Air{})
		}).
		Method("amethyst", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(block.Amethyst{})
		}).
		Method("ancientDebris", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(block.AncientDebris{})
		}).
		Constructor("anvil", func(c goja.ConstructorCall) any {
			m := c.Argument(0).Export().(map[string]any)
			t, _ := m["type"].(block.AnvilType)
			facing, _ := m["facing"].(cube.Direction)
			return block.Anvil{Type: t, Facing: facing}
		}).
		PropsMethod("banner", func(m map[string]any) goja.Value {
			colour, _ := m["colour"].(item.Colour)
			attachment, _ := m["attach"].(block.Attachment)
			patterns, _ := m["patterns"].([]block.BannerPatternLayer)
			illager, _ := m["illager"].(bool)
			return r.vm.ToValue(block.Banner{Colour: colour, Attach: attachment, Patterns: patterns, Illager: illager})
		}).
		PropsMethod("bannerPatternLayer", func(m map[string]any) goja.Value {
			t, _ := m["type"].(block.BannerPatternType)
			colour, _ := m["colour"].(item.Colour)
			return r.vm.ToValue(block.BannerPatternLayer{Type: t, Colour: colour})
		}).
		PropsMethod("barrel", func(m map[string]any) goja.Value {
			facing, _ := m["facing"].(cube.Face)
			open, _ := m["open"].(bool)
			customName, _ := m["customName"].(string)
			return r.vm.ToValue(block.Barrel{Facing: facing, Open: open, CustomName: customName})
		}).
		Method("barrier", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(block.Barrier{})
		}).
		PropsMethod("basalt", func(m map[string]any) goja.Value {
			axis, _ := m["axis"].(cube.Axis)
			polished, _ := m["polished"].(bool)
			return r.vm.ToValue(block.Basalt{Axis: axis, Polished: polished})
		}).
		PropsMethod("beacon", func(m map[string]any) goja.Value {
			primary, _ := m["primary"].(effect.LastingType)
			secondary, _ := m["secondary"].(effect.LastingType)
			return r.vm.ToValue(block.Beacon{Primary: primary, Secondary: secondary})
		}).
		PropsMethod("bedrock", func(m map[string]any) goja.Value {
			infinite, _ := m["infiniteBurning"].(bool)
			return r.vm.ToValue(block.Bedrock{InfiniteBurning: infinite})
		}).
		PropsMethod("beetrootSeeds", func(m map[string]any) goja.Value {
			growth, _ := m["growth"].(int64)
			b := block.BeetrootSeeds{}
			b.Growth = int(growth)
			return r.vm.ToValue(b)
		}).
		PropsMethod("blackstone", func(m map[string]any) goja.Value {
			t, _ := m["type"].(block.BlackstoneType)
			return r.vm.ToValue(block.Blackstone{Type: t})
		}).
		PropsMethod("blastFurnace", func(m map[string]any) goja.Value {
			facing, _ := m["facing"].(cube.Direction)
			lit, _ := m["lit"].(bool)
			return r.vm.ToValue(block.BlastFurnace{Facing: facing, Lit: lit})
		}).
		Method("blueIce", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(block.BlueIce{})
		}).
		PropsMethod("bone", func(m map[string]any) goja.Value {
			axis, _ := m["axis"].(cube.Axis)
			return r.vm.ToValue(block.Bone{Axis: axis})
		}).
		Method("bookshelf", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(block.Bookshelf{})
		}).
		PropsMethod("brewingStand", func(m map[string]any) goja.Value {
			left, _ := m["left"].(bool)
			middle, _ := m["middle"].(bool)
			right, _ := m["right"].(bool)
			return r.vm.ToValue(block.BrewingStand{LeftSlot: left, MiddleSlot: middle, RightSlot: right})
		}).
		Method("bricks", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(block.Bricks{})
		}).
		PropsMethod("cactus", func(m map[string]any) goja.Value {
			age, _ := m["age"].(int64)
			return r.vm.ToValue(block.Cactus{Age: int(age)})
		}).
		PropsMethod("cake", func(m map[string]any) goja.Value {
			bites, _ := m["bites"].(int64)
			return r.vm.ToValue(block.Cake{Bites: int(bites)})
		}).
		Method("calcite", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(block.Calcite{})
		}).
		PropsMethod("campfire", func(m map[string]any) goja.Value {
			t, _ := m["type"].(block.FireType)
			rawItems, _ := m["items"].([]block.CampfireItem)
			facing, _ := m["facing"].(cube.Direction)
			extinguished, _ := m["extinguished"].(bool)
			var items [4]block.CampfireItem
			for i, it := range rawItems {
				if i >= 4 {
					break
				}
				items[i] = it
			}
			return r.vm.ToValue(block.Campfire{Type: t, Items: items, Facing: facing, Extinguished: extinguished})
		}).
		Method("campfireItem", func(c goja.FunctionCall) goja.Value {
			i, ok := c.Argument(0).Export().(item.Stack)
			if !ok {
				panic(r.vm.NewTypeError("argument 0 is not an item.Stack"))
			}
			t, ok := c.Argument(1).Export().(time.Duration)
			if !ok {
				panic(r.vm.NewTypeError("argument 1 is not a time.Duration"))
			}
			return r.vm.ToValue(block.CampfireItem{Item: i, Time: t})
		}).
		PropsMethod("carpet", func(m map[string]any) goja.Value {
			colour, _ := m["colour"].(item.Colour)
			return r.vm.ToValue(block.Carpet{Colour: colour})
		}).
		PropsMethod("carrot", func(m map[string]any) goja.Value {
			growth, _ := m["growth"].(int64)
			b := block.Carrot{}
			b.Growth = int(growth)
			return r.vm.ToValue(b)
		}).
		PropsMethod("chain", func(m map[string]any) goja.Value {
			axis, _ := m["axis"].(cube.Axis)
			return r.vm.ToValue(block.Chain{Axis: axis})
		}).
		PropsMethod("chest", func(m map[string]any) goja.Value {
			facing, _ := m["facing"].(cube.Direction)
			customName, _ := m["customName"].(string)
			return r.vm.ToValue(block.Chest{Facing: facing, CustomName: customName})
		}).
		Method("chiseledQuartz", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(block.ChiseledQuartz{})
		}).
		Method("clay", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(block.Clay{})
		}).
		Method("closeAction", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(block.CloseAction{})
		}).
		Method("coal", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(block.Coal{})
		}).
		PropsMethod("coalOre", func(m map[string]any) goja.Value {
			t, _ := m["type"].(block.OreType)
			return r.vm.ToValue(block.CoalOre{Type: t})
		}).
		PropsMethod("cobblestone", func(m map[string]any) goja.Value {
			mossy, _ := m["mossy"].(bool)
			return r.vm.ToValue(block.Cobblestone{Mossy: mossy})
		}).
		PropsMethod("cocoaBean", func(m map[string]any) goja.Value {
			facing, _ := m["facing"].(cube.Direction)
			age, _ := m["age"].(int64)
			return r.vm.ToValue(block.CocoaBean{Facing: facing, Age: int(age)})
		}).
		PropsMethod("composter", func(m map[string]any) goja.Value {
			level, _ := m["level"].(int64)
			return r.vm.ToValue(block.Composter{Level: int(level)})
		}).
		PropsMethod("concrete", func(m map[string]any) goja.Value {
			colour, _ := m["colour"].(item.Colour)
			return r.vm.ToValue(block.Concrete{Colour: colour})
		}).
		PropsMethod("concretePowder", func(m map[string]any) goja.Value {
			colour, _ := m["colour"].(item.Colour)
			return r.vm.ToValue(block.ConcretePowder{Colour: colour})
		}).
		Method("continueCrackAction", func(c goja.FunctionCall) goja.Value {
			t, ok := c.Argument(0).Export().(time.Duration)
			if !ok {
				panic(r.vm.NewTypeError("argument 0 is not a time.Duration"))
			}
			return r.vm.ToValue(block.ContinueCrackAction{BreakTime: t})
		}).
		PropsMethod("copper", func(m map[string]any) goja.Value {
			t, _ := m["type"].(block.CopperType)
			o, _ := m["oxidation"].(block.OxidationType)
			waxed, _ := m["waxed"].(bool)
			return r.vm.ToValue(block.Copper{Type: t, Oxidation: o, Waxed: waxed})
		}).
		PropsMethod("copperDoor", func(m map[string]any) goja.Value {
			facing, _ := m["facing"].(cube.Direction)
			o, _ := m["oxidation"].(block.OxidationType)
			waxed, _ := m["waxed"].(bool)
			open, _ := m["open"].(bool)
			top, _ := m["top"].(bool)
			right, _ := m["right"].(bool)
			return r.vm.ToValue(block.CopperDoor{Facing: facing, Oxidation: o, Waxed: waxed, Open: open, Top: top, Right: right})
		}).
		PropsMethod("copperGrate", func(m map[string]any) goja.Value {
			o, _ := m["oxidation"].(block.OxidationType)
			waxed, _ := m["waxed"].(bool)
			return r.vm.ToValue(block.CopperGrate{Oxidation: o, Waxed: waxed})
		}).
		PropsMethod("copperOre", func(m map[string]any) goja.Value {
			t, _ := m["type"].(block.OreType)
			return r.vm.ToValue(block.CopperOre{Type: t})
		}).
		PropsMethod("copperTrapdoor", func(m map[string]any) goja.Value {
			facing, _ := m["facing"].(cube.Direction)
			o, _ := m["oxidation"].(block.OxidationType)
			waxed, _ := m["waxed"].(bool)
			open, _ := m["open"].(bool)
			top, _ := m["top"].(bool)
			return r.vm.ToValue(block.CopperTrapdoor{Facing: facing, Oxidation: o, Waxed: waxed, Open: open, Top: top})
		}).
		PropsMethod("coral", func(m map[string]any) goja.Value {
			t, _ := m["type"].(block.CoralType)
			dead, _ := m["dead"].(bool)
			return r.vm.ToValue(block.Coral{Type: t, Dead: dead})
		}).
		PropsMethod("coralBlock", func(m map[string]any) goja.Value {
			t, _ := m["type"].(block.CoralType)
			dead, _ := m["dead"].(bool)
			return r.vm.ToValue(block.CoralBlock{Type: t, Dead: dead})
		}).
		Method("craftingTable", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(block.CraftingTable{})
		}).
		Method("damageSource", func(c goja.FunctionCall) goja.Value {
			b, ok := c.Argument(0).Export().(world.Block)
			if !ok {
				panic(r.vm.NewTypeError("argument 0 is not a world.Block"))
			}
			return r.vm.ToValue(block.DamageSource{Block: b})
		}).
		Method("deadBush", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(block.DeadBush{})
		}).
		PropsMethod("decoratedPot", func(m map[string]any) goja.Value {
			i, _ := m["item"].(item.Stack)
			facing, _ := m["facing"].(cube.Direction)
			rawDeco, _ := m["decorations"].([]block.PotDecoration)
			var deco [4]block.PotDecoration
			for i, d := range rawDeco {
				if i >= 4 {
					break
				}
				deco[i] = d
			}
			return r.vm.ToValue(block.DecoratedPot{Item: i, Facing: facing, Decorations: deco})
		}).
		Method("decoratedPotWobbleAction", func(c goja.FunctionCall) goja.Value {
			b, ok := c.Argument(0).Export().(block.DecoratedPot)
			if !ok {
				panic(r.vm.NewTypeError("argument 0 is not a block.DecoratedPot"))
			}
			success, _ := c.Argument(1).Export().(bool)
			return r.vm.ToValue(block.DecoratedPotWobbleAction{DecoratedPot: b, Success: success})
		}).
		PropsMethod("deepslate", func(m map[string]any) goja.Value {
			t, _ := m["type"].(block.DeepslateType)
			axis, _ := m["axis"].(cube.Axis)
			return r.vm.ToValue(block.Deepslate{Type: t, Axis: axis})
		}).
		PropsMethod("deepslateBricks", func(m map[string]any) goja.Value {
			cracked, _ := m["cracked"].(bool)
			return r.vm.ToValue(block.DeepslateBricks{Cracked: cracked})
		}).
		PropsMethod("deepslateTiles", func(m map[string]any) goja.Value {
			cracked, _ := m["cracked"].(bool)
			return r.vm.ToValue(block.DeepslateTiles{Cracked: cracked})
		}).
		Method("diamond", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(block.Diamond{})
		}).
		PropsMethod("diamondOre", func(m map[string]any) goja.Value {
			t, _ := m["type"].(block.OreType)
			return r.vm.ToValue(block.DiamondOre{Type: t})
		}).
		PropsMethod("dirt", func(m map[string]any) goja.Value {
			coarse, _ := m["coarse"].(bool)
			return r.vm.ToValue(block.Dirt{Coarse: coarse})
		}).
		Method("dirtPath", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(block.DirtPath{})
		}).
		PropsMethod("doubleFlower", func(m map[string]any) goja.Value {
			t, _ := m["type"].(block.DoubleFlowerType)
			top, _ := m["upperPart"].(bool)
			return r.vm.ToValue(block.DoubleFlower{Type: t, UpperPart: top})
		}).
		PropsMethod("doubleTallGrass", func(m map[string]any) goja.Value {
			t, _ := m["type"].(block.DoubleTallGrassType)
			top, _ := m["upperPart"].(bool)
			return r.vm.ToValue(block.DoubleTallGrass{Type: t, UpperPart: top})
		}).
		Method("dragonEgg", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(block.DragonEgg{})
		}).
		Method("driedKelp", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(block.DriedKelp{})
		}).
		Method("dripstone", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(block.Dripstone{})
		}).
		Method("emerald", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(block.Emerald{})
		}).
		PropsMethod("emeraldOre", func(m map[string]any) goja.Value {
			t, _ := m["type"].(block.OreType)
			return r.vm.ToValue(block.EmeraldOre{Type: t})
		}).
		Method("enchantingTable", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(block.EnchantingTable{})
		}).
		Method("endBricks", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(block.EndBricks{})
		}).
		PropsMethod("endRod", func(m map[string]any) goja.Value {
			facing, _ := m["facing"].(cube.Face)
			return r.vm.ToValue(block.EndRod{Facing: facing})
		}).
		Method("endStone", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(block.EndStone{})
		}).
		PropsMethod("enderChest", func(m map[string]any) goja.Value {
			facing, _ := m["facing"].(cube.Direction)
			return r.vm.ToValue(block.EnderChest{Facing: facing})
		}).
		PropsMethod("farmland", func(m map[string]any) goja.Value {
			hydration, _ := m["hydration"].(int64)
			return r.vm.ToValue(block.Farmland{Hydration: int(hydration)})
		}).
		Method("fern", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(block.Fern{})
		}).
		PropsMethod("fire", func(m map[string]any) goja.Value {
			t, _ := m["type"].(block.FireType)
			age, _ := m["age"].(int64)
			return r.vm.ToValue(block.Fire{Type: t, Age: int(age)})
		}).
		Method("fireDamageSource", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(block.FireDamageSource{})
		}).
		Method("fletchingTable", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(block.FletchingTable{})
		}).
		PropsMethod("flower", func(m map[string]any) goja.Value {
			t, _ := m["type"].(block.FlowerType)
			return r.vm.ToValue(block.Flower{Type: t})
		}).
		PropsMethod("froglight", func(m map[string]any) goja.Value {
			t, _ := m["type"].(block.FroglightType)
			axis, _ := m["axis"].(cube.Axis)
			return r.vm.ToValue(block.Froglight{Type: t, Axis: axis})
		}).
		PropsMethod("furnace", func(m map[string]any) goja.Value {
			facing, _ := m["facing"].(cube.Direction)
			lit, _ := m["lit"].(bool)
			return r.vm.ToValue(block.Furnace{Facing: facing, Lit: lit})
		}).
		Method("glass", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(block.Glass{})
		}).
		Method("glassPane", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(block.GlassPane{})
		}).
		PropsMethod("glazedTerracotta", func(m map[string]any) goja.Value {
			colour, _ := m["colour"].(item.Colour)
			facing, _ := m["facing"].(cube.Direction)
			return r.vm.ToValue(block.GlazedTerracotta{Colour: colour, Facing: facing})
		}).
		Method("glowstone", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(block.Glowstone{})
		}).
		Method("gold", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(block.Gold{})
		}).
		PropsMethod("goldOre", func(m map[string]any) goja.Value {
			t, _ := m["type"].(block.OreType)
			return r.vm.ToValue(block.GoldOre{Type: t})
		}).
		Method("grass", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(block.Grass{})
		}).
		Method("gravel", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(block.Gravel{})
		}).
		PropsMethod("grindstone", func(m map[string]any) goja.Value {
			facing, _ := m["facing"].(cube.Direction)
			a, _ := m["attach"].(block.GrindstoneAttachment)
			return r.vm.ToValue(block.Grindstone{Facing: facing, Attach: a})
		}).
		PropsMethod("hayBale", func(m map[string]any) goja.Value {
			axis, _ := m["axis"].(cube.Axis)
			return r.vm.ToValue(block.HayBale{Axis: axis})
		}).
		Method("honeycomb", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(block.Honeycomb{})
		}).
		PropsMethod("hopper", func(m map[string]any) goja.Value {
			facing, _ := m["facing"].(cube.Face)
			powered, _ := m["powered"].(bool)
			customName, _ := m["customName"].(string)
			lastTick, _ := m["lastTick"].(int64)
			transferCooldown, _ := m["transferCooldown"].(int64)
			collectCooldown, _ := m["collectCooldown"].(int64)
			return r.vm.ToValue(block.Hopper{Facing: facing, Powered: powered, CustomName: customName, LastTick: lastTick, TransferCooldown: transferCooldown, CollectCooldown: collectCooldown})
		}).
		Method("invisibleBedrock", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(block.InvisibleBedrock{})
		}).
		Method("iron", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(block.Iron{})
		}).
		Method("ironBars", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(block.IronBars{})
		}).
		PropsMethod("ironOre", func(m map[string]any) goja.Value {
			t, _ := m["type"].(block.OreType)
			return r.vm.ToValue(block.IronOre{Type: t})
		}).
		PropsMethod("itemFrame", func(m map[string]any) goja.Value {
			i, _ := m["item"].(item.Stack)
			facing, _ := m["facing"].(cube.Face)
			rotations, _ := m["rotations"].(int64)
			dropChance, _ := m["dropChance"].(float64)
			glowing, _ := m["glowing"].(bool)
			return r.vm.ToValue(block.ItemFrame{Facing: facing, Item: i, Rotations: int(rotations), DropChance: dropChance, Glowing: glowing})
		}).
		PropsMethod("jukebox", func(m map[string]any) goja.Value {
			i, _ := m["item"].(item.Stack)
			return r.vm.ToValue(block.Jukebox{Item: i})
		}).
		PropsMethod("kelp", func(m map[string]any) goja.Value {
			age, _ := m["age"].(int64)
			return r.vm.ToValue(block.Kelp{Age: int(age)})
		}).
		PropsMethod("ladder", func(m map[string]any) goja.Value {
			facing, _ := m["facing"].(cube.Direction)
			return r.vm.ToValue(block.Ladder{Facing: facing})
		}).
		PropsMethod("lantern", func(m map[string]any) goja.Value {
			t, _ := m["type"].(block.FireType)
			hanging, _ := m["hanging"].(bool)
			return r.vm.ToValue(block.Lantern{Type: t, Hanging: hanging})
		}).
		Method("lapis", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(block.Lapis{})
		}).
		PropsMethod("lapisOre", func(m map[string]any) goja.Value {
			t, _ := m["type"].(block.OreType)
			return r.vm.ToValue(block.LapisOre{Type: t})
		}).
		PropsMethod("lava", func(m map[string]any) goja.Value {
			depth, _ := m["depth"].(int64)
			still, _ := m["still"].(bool)
			falling, _ := m["falling"].(bool)
			return r.vm.ToValue(block.Lava{Depth: int(depth), Still: still, Falling: falling})
		}).
		Method("lavaDamageSource", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(block.LavaDamageSource{})
		}).
		PropsMethod("leaves", func(m map[string]any) goja.Value {
			wood, _ := m["wood"].(block.WoodType)
			persistent, _ := m["persistent"].(bool)
			shouldUpdate, _ := m["shouldUpdate"].(bool)
			return r.vm.ToValue(block.Leaves{Wood: wood, Persistent: persistent, ShouldUpdate: shouldUpdate})
		}).
		PropsMethod("lectern", func(m map[string]any) goja.Value {
			facing, _ := m["facing"].(cube.Direction)
			i, _ := m["item"].(item.Stack)
			page, _ := m["page"].(int64)
			return r.vm.ToValue(block.Lectern{Facing: facing, Book: i, Page: int(page)})
		}).
		PropsMethod("light", func(m map[string]any) goja.Value {
			level, _ := m["level"].(int64)
			return r.vm.ToValue(block.Light{Level: int(level)})
		}).
		PropsMethod("litPumpkin", func(m map[string]any) goja.Value {
			facing, _ := m["facing"].(cube.Direction)
			return r.vm.ToValue(block.LitPumpkin{Facing: facing})
		}).
		PropsMethod("log", func(m map[string]any) goja.Value {
			wood, _ := m["wood"].(block.WoodType)
			axis, _ := m["axis"].(cube.Axis)
			stripped, _ := m["stripped"].(bool)
			return r.vm.ToValue(block.Log{Wood: wood, Axis: axis, Stripped: stripped})
		}).
		PropsMethod("loom", func(m map[string]any) goja.Value {
			facing, _ := m["facing"].(cube.Direction)
			return r.vm.ToValue(block.Loom{Facing: facing})
		}).
		Method("melon", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(block.Melon{})
		}).
		PropsMethod("melonSeeds", func(m map[string]any) goja.Value {
			direction, _ := m["direction"].(cube.Face)
			growth, _ := m["growth"].(int64)
			b := block.MelonSeeds{Direction: direction}
			b.Growth = int(growth)
			return r.vm.ToValue(b)
		}).
		Method("mossCarpet", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(block.MossCarpet{})
		}).
		Method("mud", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(block.Mud{})
		}).
		Method("mudBricks", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(block.MudBricks{})
		}).
		PropsMethod("muddyMangroveRoots", func(m map[string]any) goja.Value {
			axis, _ := m["axis"].(cube.Axis)
			return r.vm.ToValue(block.MuddyMangroveRoots{Axis: axis})
		}).
		Method("netherBrickFence", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(block.NetherBrickFence{})
		}).
		PropsMethod("netherBricks", func(m map[string]any) goja.Value {
			t, _ := m["type"].(block.NetherBricksType)
			return r.vm.ToValue(block.NetherBricks{Type: t})
		}).
		Method("netherGoldOre", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(block.NetherGoldOre{})
		}).
		Method("netherQuartzOre", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(block.NetherQuartzOre{})
		}).
		Method("netherSprouts", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(block.NetherSprouts{})
		}).
		PropsMethod("netherWart", func(m map[string]any) goja.Value {
			age, _ := m["age"].(int64)
			return r.vm.ToValue(block.NetherWart{Age: int(age)})
		}).
		PropsMethod("netherWartBlock", func(m map[string]any) goja.Value {
			warped, _ := m["warped"].(bool)
			return r.vm.ToValue(block.NetherWartBlock{Warped: warped})
		}).
		Method("netherite", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(block.Netherite{})
		}).
		Method("netherrack", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(block.Netherrack{})
		}).
		PropsMethod("note", func(m map[string]any) goja.Value {
			pitch, _ := m["pitch"].(int64)
			return r.vm.ToValue(block.Note{Pitch: int(pitch)})
		}).
		PropsMethod("obsidian", func(m map[string]any) goja.Value {
			crying, _ := m["crying"].(bool)
			return r.vm.ToValue(block.Obsidian{Crying: crying})
		}).
		Method("openAction", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(block.OpenAction{})
		}).
		Method("packedIce", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(block.PackedIce{})
		}).
		Method("packedMud", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(block.PackedMud{})
		}).
		PropsMethod("pinkPetals", func(m map[string]any) goja.Value {
			facing, _ := m["facing"].(cube.Direction)
			additional, _ := m["additionalCount"].(int64)
			return r.vm.ToValue(block.PinkPetals{Facing: facing, AdditionalCount: int(additional)})
		}).
		PropsMethod("planks", func(m map[string]any) goja.Value {
			wood, _ := m["wood"].(block.WoodType)
			return r.vm.ToValue(block.Planks{Wood: wood})
		}).
		Method("podzol", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(block.Podzol{})
		}).
		PropsMethod("polishedBlackstoneBrick", func(m map[string]any) goja.Value {
			cracked, _ := m["cracked"].(bool)
			return r.vm.ToValue(block.PolishedBlackstoneBrick{Cracked: cracked})
		}).
		Method("polishedTuff", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(block.PolishedTuff{})
		}).
		PropsMethod("potato", func(m map[string]any) goja.Value {
			growth, _ := m["growth"].(int64)
			b := block.Potato{}
			b.Growth = int(growth)
			return r.vm.ToValue(b)
		}).
		PropsMethod("prismarine", func(m map[string]any) goja.Value {
			t, _ := m["type"].(block.PrismarineType)
			return r.vm.ToValue(block.Prismarine{Type: t})
		}).
		PropsMethod("pumpkin", func(m map[string]any) goja.Value {
			facing, _ := m["facing"].(cube.Direction)
			return r.vm.ToValue(block.Pumpkin{Facing: facing})
		}).
		PropsMethod("pumpkinSeeds", func(m map[string]any) goja.Value {
			direction, _ := m["direction"].(cube.Face)
			growth, _ := m["growth"].(int64)
			b := block.PumpkinSeeds{Direction: direction}
			b.Growth = int(growth)
			return r.vm.ToValue(b)
		}).
		Method("purpur", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(block.Purpur{})
		}).
		PropsMethod("purpurPillar", func(m map[string]any) goja.Value {
			axis, _ := m["axis"].(cube.Axis)
			return r.vm.ToValue(block.PurpurPillar{Axis: axis})
		}).
		PropsMethod("quartz", func(m map[string]any) goja.Value {
			smooth, _ := m["smooth"].(bool)
			return r.vm.ToValue(block.Quartz{Smooth: smooth})
		}).
		Method("quartzBricks", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(block.QuartzBricks{})
		}).
		PropsMethod("quartzPillar", func(m map[string]any) goja.Value {
			axis, _ := m["axis"].(cube.Axis)
			return r.vm.ToValue(block.QuartzPillar{Axis: axis})
		}).
		Method("rawCopper", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(block.RawCopper{})
		}).
		Method("rawGold", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(block.RawGold{})
		}).
		Method("rawIron", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(block.RawIron{})
		}).
		Method("reinforcedDeepslate", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(block.ReinforcedDeepslate{})
		}).
		Method("resin", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(block.Resin{})
		}).
		PropsMethod("resinBricks", func(m map[string]any) goja.Value {
			chiseled, _ := m["chiseled"].(bool)
			return r.vm.ToValue(block.ResinBricks{Chiseled: chiseled})
		}).
		PropsMethod("sand", func(m map[string]any) goja.Value {
			red, _ := m["red"].(bool)
			return r.vm.ToValue(block.Sand{Red: red})
		}).
		PropsMethod("sandstone", func(m map[string]any) goja.Value {
			t, _ := m["type"].(block.SandstoneType)
			red, _ := m["red"].(bool)
			return r.vm.ToValue(block.Sandstone{Type: t, Red: red})
		}).
		Method("seaLantern", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(block.SeaLantern{})
		}).
		PropsMethod("seaPickle", func(m map[string]any) goja.Value {
			dead, _ := m["dead"].(bool)
			return r.vm.ToValue(block.SeaPickle{Dead: dead})
		}).
		PropsMethod("shortGrass", func(m map[string]any) goja.Value {
			double, _ := m["double"].(bool)
			return r.vm.ToValue(block.ShortGrass{Double: double})
		}).
		Method("shroomlight", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(block.Shroomlight{})
		}).
		PropsMethod("sign", func(m map[string]any) goja.Value {
			wood, _ := m["wood"].(block.WoodType)
			a, _ := m["attach"].(block.Attachment)
			front, _ := m["front"].(block.SignText)
			back, _ := m["back"].(block.SignText)
			waxed, _ := m["waxed"].(bool)
			return r.vm.ToValue(block.Sign{Wood: wood, Attach: a, Front: front, Back: back, Waxed: waxed})
		}).
		PropsMethod("signText", func(m map[string]any) goja.Value {
			text, _ := m["text"].(string)
			owner, _ := m["owner"].(string)
			colour, _ := m["baseColour"].(color.RGBA)
			glowing, _ := m["glowing"].(bool)
			return r.vm.ToValue(block.SignText{Text: text, Owner: owner, BaseColour: colour, Glowing: glowing})
		}).
		PropsMethod("skull", func(m map[string]any) goja.Value {
			t, _ := m["type"].(block.SkullType)
			a, _ := m["attach"].(block.Attachment)
			return r.vm.ToValue(block.Skull{Type: t, Attach: a})
		}).
		PropsMethod("slab", func(m map[string]any) goja.Value {
			b, _ := m["block"].(world.Block)
			double, _ := m["double"].(bool)
			top, _ := m["top"].(bool)
			return r.vm.ToValue(block.Slab{Block: b, Double: double, Top: top})
		}).
		Method("smithingTable", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(block.SmithingTable{})
		}).
		PropsMethod("smoker", func(m map[string]any) goja.Value {
			facing, _ := m["facing"].(cube.Direction)
			lit, _ := m["lit"].(bool)
			return r.vm.ToValue(block.Smoker{Facing: facing, Lit: lit})
		}).
		Method("snow", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(block.Snow{})
		}).
		Method("soulSand", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(block.SoulSand{})
		}).
		Method("soulSoil", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(block.SoulSoil{})
		}).
		PropsMethod("sponge", func(m map[string]any) goja.Value {
			wet, _ := m["wet"].(bool)
			return r.vm.ToValue(block.Sponge{Wet: wet})
		}).
		Method("sporeBlossom", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(block.SporeBlossom{})
		}).
		PropsMethod("stainedGlass", func(m map[string]any) goja.Value {
			colour, _ := m["colour"].(item.Colour)
			return r.vm.ToValue(block.StainedGlass{Colour: colour})
		}).
		PropsMethod("stainedGlassPane", func(m map[string]any) goja.Value {
			colour, _ := m["colour"].(item.Colour)
			return r.vm.ToValue(block.StainedGlassPane{Colour: colour})
		}).
		PropsMethod("stainedTerracotta", func(m map[string]any) goja.Value {
			colour, _ := m["colour"].(item.Colour)
			return r.vm.ToValue(block.StainedTerracotta{Colour: colour})
		}).
		PropsMethod("stairs", func(m map[string]any) goja.Value {
			b, _ := m["block"].(world.Block)
			facing, _ := m["facing"].(cube.Direction)
			upsideDown, _ := m["upsideDown"].(bool)
			return r.vm.ToValue(block.Stairs{Block: b, Facing: facing, UpsideDown: upsideDown})
		}).
		Method("startCrackAction", func(c goja.FunctionCall) goja.Value {
			t, ok := c.Argument(0).Export().(time.Duration)
			if !ok {
				panic(r.vm.NewTypeError("argument 0 is not a time.Duration"))
			}
			return r.vm.ToValue(block.StartCrackAction{BreakTime: t})
		}).
		PropsMethod("stone", func(m map[string]any) goja.Value {
			smooth, _ := m["smooth"].(bool)
			return r.vm.ToValue(block.Stone{Smooth: smooth})
		}).
		PropsMethod("stoneBricks", func(m map[string]any) goja.Value {
			t, _ := m["type"].(block.StoneBricksType)
			return r.vm.ToValue(block.StoneBricks{Type: t})
		}).
		PropsMethod("stonecutter", func(m map[string]any) goja.Value {
			facing, _ := m["facing"].(cube.Direction)
			return r.vm.ToValue(block.Stonecutter{Facing: facing})
		}).
		Method("stopCrackAction", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(block.StopCrackAction{})
		}).
		PropsMethod("sugarCane", func(m map[string]any) goja.Value {
			age, _ := m["age"].(int64)
			return r.vm.ToValue(block.SugarCane{Age: int(age)})
		}).
		Method("tnt", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(block.TNT{})
		}).
		Method("terracotta", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(block.Terracotta{})
		}).
		PropsMethod("torch", func(m map[string]any) goja.Value {
			facing, _ := m["facing"].(cube.Face)
			t, _ := m["type"].(block.FireType)
			return r.vm.ToValue(block.Torch{Facing: facing, Type: t})
		}).
		PropsMethod("tuff", func(m map[string]any) goja.Value {
			chiseled, _ := m["chiseled"].(bool)
			return r.vm.ToValue(block.Tuff{Chiseled: chiseled})
		}).
		PropsMethod("tuffBricks", func(m map[string]any) goja.Value {
			chiseled, _ := m["chiseled"].(bool)
			return r.vm.ToValue(block.TuffBricks{Chiseled: chiseled})
		}).
		PropsMethod("vines", func(m map[string]any) goja.Value {
			north, _ := m["north"].(bool)
			east, _ := m["east"].(bool)
			south, _ := m["south"].(bool)
			west, _ := m["west"].(bool)
			return r.vm.ToValue(block.Vines{NorthDirection: north, EastDirection: east, SouthDirection: south, WestDirection: west})
		}).
		PropsMethod("wall", func(m map[string]any) goja.Value {
			b, _ := m["block"].(world.Block)
			north, _ := m["north"].(block.WallConnectionType)
			east, _ := m["east"].(block.WallConnectionType)
			south, _ := m["south"].(block.WallConnectionType)
			west, _ := m["west"].(block.WallConnectionType)
			post, _ := m["post"].(bool)
			return r.vm.ToValue(block.Wall{
				Block:           b,
				NorthConnection: north,
				EastConnection:  east,
				SouthConnection: south,
				WestConnection:  west,
				Post:            post,
			})
		}).
		PropsMethod("water", func(m map[string]any) goja.Value {
			depth, _ := m["depth"].(int64)
			still, _ := m["still"].(bool)
			falling, _ := m["falling"].(bool)
			return r.vm.ToValue(block.Water{Depth: int(depth), Still: still, Falling: falling})
		}).
		PropsMethod("wheatSeeds", func(m map[string]any) goja.Value {
			growth, _ := m["growth"].(int64)
			b := block.WheatSeeds{}
			b.Growth = int(growth)
			return r.vm.ToValue(b)
		}).
		PropsMethod("wood", func(m map[string]any) goja.Value {
			wood, _ := m["wood"].(block.WoodType)
			axis, _ := m["axis"].(cube.Axis)
			stripped, _ := m["stripped"].(bool)
			return r.vm.ToValue(block.Wood{Wood: wood, Axis: axis, Stripped: stripped})
		}).
		PropsMethod("woodDoor", func(m map[string]any) goja.Value {
			wood, _ := m["wood"].(block.WoodType)
			facing, _ := m["facing"].(cube.Direction)
			open, _ := m["open"].(bool)
			top, _ := m["top"].(bool)
			right, _ := m["right"].(bool)
			return r.vm.ToValue(block.WoodDoor{Wood: wood, Facing: facing, Open: open, Top: top, Right: right})
		}).
		PropsMethod("woodFence", func(m map[string]any) goja.Value {
			wood, _ := m["wood"].(block.WoodType)
			return r.vm.ToValue(block.WoodFence{Wood: wood})
		}).
		PropsMethod("woodFenceGate", func(m map[string]any) goja.Value {
			wood, _ := m["wood"].(block.WoodType)
			facing, _ := m["facing"].(cube.Direction)
			open, _ := m["open"].(bool)
			lowered, _ := m["lowered"].(bool)
			return r.vm.ToValue(block.WoodFenceGate{Wood: wood, Facing: facing, Open: open, Lowered: lowered})
		}).
		PropsMethod("woodTrapdoor", func(m map[string]any) goja.Value {
			wood, _ := m["wood"].(block.WoodType)
			facing, _ := m["facing"].(cube.Direction)
			open, _ := m["open"].(bool)
			top, _ := m["top"].(bool)
			return r.vm.ToValue(block.WoodTrapdoor{Wood: wood, Facing: facing, Open: open, Top: top})
		}).
		PropsMethod("wool", func(m map[string]any) goja.Value {
			colour, _ := m["colour"].(item.Colour)
			return r.vm.ToValue(block.Wool{Colour: colour})
		}).
		Method("breakDuration", func(c goja.FunctionCall) goja.Value {
			b, ok := c.Argument(0).Export().(world.Block)
			if !ok {
				panic(r.vm.NewTypeError("argument 0 is not a world.Block"))
			}
			i, ok := c.Argument(1).Export().(item.Stack)
			if !ok {
				panic(r.vm.NewTypeError("argument 1 is not an item.Stack"))
			}
			return r.vm.ToValue(block.BreakDuration(b, i))
		}).
		Method("breaksInstantly", func(c goja.FunctionCall) goja.Value {
			b, ok := c.Argument(0).Export().(world.Block)
			if !ok {
				panic(r.vm.NewTypeError("argument 0 is not a world.Block"))
			}
			i, ok := c.Argument(1).Export().(item.Stack)
			if !ok {
				panic(r.vm.NewTypeError("argument 1 is not an item.Stack"))
			}
			return r.vm.ToValue(block.BreaksInstantly(b, i))
		}).
		Method("nextHash", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(block.NextHash())
		}).
		Method("wallAttachment", func(c goja.FunctionCall) goja.Value {
			facing, ok := c.Argument(0).Export().(cube.Direction)
			if !ok {
				panic(r.vm.NewTypeError("argument 0 is not a cube.Direction"))
			}
			return r.vm.ToValue(block.WallAttachment(facing))
		}).
		Method("standingAttachment", func(c goja.FunctionCall) goja.Value {
			o, ok := c.Argument(0).Export().(cube.Orientation)
			if !ok {
				panic(r.vm.NewTypeError("argument 0 is not a cube.Orientation"))
			}
			return r.vm.ToValue(block.StandingAttachment(o))
		}).
		Apply(r.vm, "block")
}
