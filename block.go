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
		Method("banner", func(c goja.FunctionCall) goja.Value {
			m := c.Argument(0).Export().(map[string]any)
			colour, _ := m["colour"].(item.Colour)
			attachment, _ := m["attach"].(block.Attachment)
			patterns, _ := m["patterns"].([]block.BannerPatternLayer)
			illager, _ := m["illager"].(bool)
			return r.vm.ToValue(block.Banner{Colour: colour, Attach: attachment, Patterns: patterns, Illager: illager})
		}).
		Method("bannerPatternLayer", func(c goja.FunctionCall) goja.Value {
			m := c.Argument(0).Export().(map[string]any)
			t, _ := m["type"].(block.BannerPatternType)
			colour, _ := m["colour"].(item.Colour)
			return r.vm.ToValue(block.BannerPatternLayer{Type: t, Colour: colour})
		}).
		Method("barrel", func(c goja.FunctionCall) goja.Value {
			m := c.Argument(0).Export().(map[string]any)
			facing, _ := m["facing"].(cube.Face)
			open, _ := m["open"].(bool)
			customName, _ := m["customName"].(string)
			return r.vm.ToValue(block.Barrel{Facing: facing, Open: open, CustomName: customName})
		}).
		Method("barrier", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(block.Barrier{})
		}).
		Method("basalt", func(c goja.FunctionCall) goja.Value {
			m := c.Argument(0).Export().(map[string]any)
			axis, _ := m["axis"].(cube.Axis)
			polished, _ := m["polished"].(bool)
			return r.vm.ToValue(block.Basalt{Axis: axis, Polished: polished})
		}).
		Method("beacon", func(c goja.FunctionCall) goja.Value {
			m := c.Argument(0).Export().(map[string]any)
			primary, _ := m["primary"].(effect.LastingType)
			secondary, _ := m["secondary"].(effect.LastingType)
			return r.vm.ToValue(block.Beacon{Primary: primary, Secondary: secondary})
		}).
		Method("bedrock", func(c goja.FunctionCall) goja.Value {
			m := c.Argument(0).Export().(map[string]any)
			infinite, _ := m["infiniteBurning"].(bool)
			return r.vm.ToValue(block.Bedrock{InfiniteBurning: infinite})
		}).
		Method("beetrootSeeds", func(c goja.FunctionCall) goja.Value {
			m := c.Argument(0).Export().(map[string]any)
			growth, _ := m["growth"].(int64)
			b := block.BeetrootSeeds{}
			b.Growth = int(growth)
			return r.vm.ToValue(b)
		}).
		Method("blackstone", func(c goja.FunctionCall) goja.Value {
			m := c.Argument(0).Export().(map[string]any)
			t, _ := m["type"].(block.BlackstoneType)
			return r.vm.ToValue(block.Blackstone{Type: t})
		}).
		Method("blastFurnace", func(c goja.FunctionCall) goja.Value {
			m := c.Argument(0).Export().(map[string]any)
			facing, _ := m["facing"].(cube.Direction)
			lit, _ := m["lit"].(bool)
			return r.vm.ToValue(block.BlastFurnace{Facing: facing, Lit: lit})
		}).
		Method("blueIce", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(block.BlueIce{})
		}).
		Method("bone", func(c goja.FunctionCall) goja.Value {
			m := c.Argument(0).Export().(map[string]any)
			axis, _ := m["axis"].(cube.Axis)
			return r.vm.ToValue(block.Bone{Axis: axis})
		}).
		Method("bookshelf", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(block.Bookshelf{})
		}).
		Method("brewingStand", func(c goja.FunctionCall) goja.Value {
			m := c.Argument(0).Export().(map[string]any)
			left, _ := m["left"].(bool)
			middle, _ := m["middle"].(bool)
			right, _ := m["right"].(bool)
			return r.vm.ToValue(block.BrewingStand{LeftSlot: left, MiddleSlot: middle, RightSlot: right})
		}).
		Method("bricks", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(block.Bricks{})
		}).
		Method("cactus", func(c goja.FunctionCall) goja.Value {
			m := c.Argument(0).Export().(map[string]any)
			age, _ := m["age"].(int64)
			return r.vm.ToValue(block.Cactus{Age: int(age)})
		}).
		Method("cake", func(c goja.FunctionCall) goja.Value {
			m := c.Argument(0).Export().(map[string]any)
			bites, _ := m["bites"].(int64)
			return r.vm.ToValue(block.Cake{Bites: int(bites)})
		}).
		Method("calcite", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(block.Calcite{})
		}).
		Method("campfire", func(c goja.FunctionCall) goja.Value {
			m := c.Argument(0).Export().(map[string]any)
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
		Method("carpet", func(c goja.FunctionCall) goja.Value {
			m := c.Argument(0).Export().(map[string]any)
			colour, _ := m["colour"].(item.Colour)
			return r.vm.ToValue(block.Carpet{Colour: colour})
		}).
		Method("carrot", func(c goja.FunctionCall) goja.Value {
			m := c.Argument(0).Export().(map[string]any)
			growth, _ := m["growth"].(int64)
			b := block.Carrot{}
			b.Growth = int(growth)
			return r.vm.ToValue(b)
		}).
		Method("chain", func(c goja.FunctionCall) goja.Value {
			m := c.Argument(0).Export().(map[string]any)
			axis, _ := m["axis"].(cube.Axis)
			return r.vm.ToValue(block.Chain{Axis: axis})
		}).
		Method("chest", func(c goja.FunctionCall) goja.Value {
			m := c.Argument(0).Export().(map[string]any)
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
		Method("coalOre", func(c goja.FunctionCall) goja.Value {
			m := c.Argument(0).Export().(map[string]any)
			t, _ := m["type"].(block.OreType)
			return r.vm.ToValue(block.CoalOre{Type: t})
		}).
		Method("cobblestone", func(c goja.FunctionCall) goja.Value {
			m := c.Argument(0).Export().(map[string]any)
			mossy, _ := m["mossy"].(bool)
			return r.vm.ToValue(block.Cobblestone{Mossy: mossy})
		}).
		Method("cocoaBean", func(c goja.FunctionCall) goja.Value {
			m := c.Argument(0).Export().(map[string]any)
			facing, _ := m["facing"].(cube.Direction)
			age, _ := m["age"].(int64)
			return r.vm.ToValue(block.CocoaBean{Facing: facing, Age: int(age)})
		}).
		Method("composter", func(c goja.FunctionCall) goja.Value {
			m := c.Argument(0).Export().(map[string]any)
			level, _ := m["level"].(int64)
			return r.vm.ToValue(block.Composter{Level: int(level)})
		}).
		Method("concrete", func(c goja.FunctionCall) goja.Value {
			m := c.Argument(0).Export().(map[string]any)
			colour, _ := m["colour"].(item.Colour)
			return r.vm.ToValue(block.Concrete{Colour: colour})
		}).
		Method("concretePowder", func(c goja.FunctionCall) goja.Value {
			m := c.Argument(0).Export().(map[string]any)
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
		Method("copper", func(c goja.FunctionCall) goja.Value {
			m := c.Argument(0).Export().(map[string]any)
			t, _ := m["type"].(block.CopperType)
			o, _ := m["oxidation"].(block.OxidationType)
			waxed, _ := m["waxed"].(bool)
			return r.vm.ToValue(block.Copper{Type: t, Oxidation: o, Waxed: waxed})
		}).
		Method("copperDoor", func(c goja.FunctionCall) goja.Value {
			m := c.Argument(0).Export().(map[string]any)
			facing, _ := m["facing"].(cube.Direction)
			o, _ := m["oxidation"].(block.OxidationType)
			waxed, _ := m["waxed"].(bool)
			open, _ := m["open"].(bool)
			top, _ := m["top"].(bool)
			right, _ := m["right"].(bool)
			return r.vm.ToValue(block.CopperDoor{Facing: facing, Oxidation: o, Waxed: waxed, Open: open, Top: top, Right: right})
		}).
		Method("copperGrate", func(c goja.FunctionCall) goja.Value {
			m := c.Argument(0).Export().(map[string]any)
			o, _ := m["oxidation"].(block.OxidationType)
			waxed, _ := m["waxed"].(bool)
			return r.vm.ToValue(block.CopperGrate{Oxidation: o, Waxed: waxed})
		}).
		Method("copperOre", func(c goja.FunctionCall) goja.Value {
			m := c.Argument(0).Export().(map[string]any)
			t, _ := m["type"].(block.OreType)
			return r.vm.ToValue(block.CopperOre{Type: t})
		}).
		Method("copperTrapdoor", func(c goja.FunctionCall) goja.Value {
			m := c.Argument(0).Export().(map[string]any)
			facing, _ := m["facing"].(cube.Direction)
			o, _ := m["oxidation"].(block.OxidationType)
			waxed, _ := m["waxed"].(bool)
			open, _ := m["open"].(bool)
			top, _ := m["top"].(bool)
			return r.vm.ToValue(block.CopperTrapdoor{Facing: facing, Oxidation: o, Waxed: waxed, Open: open, Top: top})
		}).
		Method("coral", func(c goja.FunctionCall) goja.Value {
			m := c.Argument(0).Export().(map[string]any)
			t, _ := m["type"].(block.CoralType)
			dead, _ := m["dead"].(bool)
			return r.vm.ToValue(block.Coral{Type: t, Dead: dead})
		}).
		Method("coralBlock", func(c goja.FunctionCall) goja.Value {
			m := c.Argument(0).Export().(map[string]any)
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
		Method("decoratedPot", func(c goja.FunctionCall) goja.Value {
			m := c.Argument(0).Export().(map[string]any)
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
		Method("deepslate", func(c goja.FunctionCall) goja.Value {
			m := c.Argument(0).Export().(map[string]any)
			t, _ := m["type"].(block.DeepslateType)
			axis, _ := m["axis"].(cube.Axis)
			return r.vm.ToValue(block.Deepslate{Type: t, Axis: axis})
		}).
		Method("deepslateBricks", func(c goja.FunctionCall) goja.Value {
			m := c.Argument(0).Export().(map[string]any)
			cracked, _ := m["cracked"].(bool)
			return r.vm.ToValue(block.DeepslateBricks{Cracked: cracked})
		}).
		Method("deepslateTiles", func(c goja.FunctionCall) goja.Value {
			m := c.Argument(0).Export().(map[string]any)
			cracked, _ := m["cracked"].(bool)
			return r.vm.ToValue(block.DeepslateTiles{Cracked: cracked})
		}).
		Method("diamond", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(block.Diamond{})
		}).
		Method("diamondOre", func(c goja.FunctionCall) goja.Value {
			m := c.Argument(0).Export().(map[string]any)
			t, _ := m["type"].(block.OreType)
			return r.vm.ToValue(block.DiamondOre{Type: t})
		}).
		Method("dirt", func(c goja.FunctionCall) goja.Value {
			m := c.Argument(0).Export().(map[string]any)
			coarse, _ := m["coarse"].(bool)
			return r.vm.ToValue(block.Dirt{Coarse: coarse})
		}).
		Method("dirtPath", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(block.DirtPath{})
		}).
		Method("doubleFlower", func(c goja.FunctionCall) goja.Value {
			m := c.Argument(0).Export().(map[string]any)
			t, _ := m["type"].(block.DoubleFlowerType)
			top, _ := m["upperPart"].(bool)
			return r.vm.ToValue(block.DoubleFlower{Type: t, UpperPart: top})
		}).
		Method("doubleTallGrass", func(c goja.FunctionCall) goja.Value {
			m := c.Argument(0).Export().(map[string]any)
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
		Method("emeraldOre", func(c goja.FunctionCall) goja.Value {
			m := c.Argument(0).Export().(map[string]any)
			t, _ := m["type"].(block.OreType)
			return r.vm.ToValue(block.EmeraldOre{Type: t})
		}).
		Method("enchantingTable", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(block.EnchantingTable{})
		}).
		Method("endBricks", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(block.EndBricks{})
		}).
		Method("endRod", func(c goja.FunctionCall) goja.Value {
			m := c.Argument(0).Export().(map[string]any)
			facing, _ := m["facing"].(cube.Face)
			return r.vm.ToValue(block.EndRod{Facing: facing})
		}).
		Method("endStone", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(block.EndStone{})
		}).
		Method("enderChest", func(c goja.FunctionCall) goja.Value {
			m := c.Argument(0).Export().(map[string]any)
			facing, _ := m["facing"].(cube.Direction)
			return r.vm.ToValue(block.EnderChest{Facing: facing})
		}).
		Method("farmland", func(c goja.FunctionCall) goja.Value {
			m := c.Argument(0).Export().(map[string]any)
			hydration, _ := m["hydration"].(int64)
			return r.vm.ToValue(block.Farmland{Hydration: int(hydration)})
		}).
		Method("fern", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(block.Fern{})
		}).
		Method("fire", func(c goja.FunctionCall) goja.Value {
			m := c.Argument(0).Export().(map[string]any)
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
		Method("flower", func(c goja.FunctionCall) goja.Value {
			m := c.Argument(0).Export().(map[string]any)
			t, _ := m["type"].(block.FlowerType)
			return r.vm.ToValue(block.Flower{Type: t})
		}).
		Method("froglight", func(c goja.FunctionCall) goja.Value {
			m := c.Argument(0).Export().(map[string]any)
			t, _ := m["type"].(block.FroglightType)
			axis, _ := m["axis"].(cube.Axis)
			return r.vm.ToValue(block.Froglight{Type: t, Axis: axis})
		}).
		Method("furnace", func(c goja.FunctionCall) goja.Value {
			m := c.Argument(0).Export().(map[string]any)
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
		Method("glazedTerracotta", func(c goja.FunctionCall) goja.Value {
			m := c.Argument(0).Export().(map[string]any)
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
		Method("goldOre", func(c goja.FunctionCall) goja.Value {
			m := c.Argument(0).Export().(map[string]any)
			t, _ := m["type"].(block.OreType)
			return r.vm.ToValue(block.GoldOre{Type: t})
		}).
		Method("grass", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(block.Grass{})
		}).
		Method("gravel", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(block.Gravel{})
		}).
		Method("grindstone", func(c goja.FunctionCall) goja.Value {
			m := c.Argument(0).Export().(map[string]any)
			facing, _ := m["facing"].(cube.Direction)
			a, _ := m["attach"].(block.GrindstoneAttachment)
			return r.vm.ToValue(block.Grindstone{Facing: facing, Attach: a})
		}).
		Method("hayBale", func(c goja.FunctionCall) goja.Value {
			m := c.Argument(0).Export().(map[string]any)
			axis, _ := m["axis"].(cube.Axis)
			return r.vm.ToValue(block.HayBale{Axis: axis})
		}).
		Method("honeycomb", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(block.Honeycomb{})
		}).
		Method("hopper", func(c goja.FunctionCall) goja.Value {
			m := c.Argument(0).Export().(map[string]any)
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
		Method("ironOre", func(c goja.FunctionCall) goja.Value {
			m := c.Argument(0).Export().(map[string]any)
			t, _ := m["type"].(block.OreType)
			return r.vm.ToValue(block.IronOre{Type: t})
		}).
		Method("itemFrame", func(c goja.FunctionCall) goja.Value {
			m := c.Argument(0).Export().(map[string]any)
			i, _ := m["item"].(item.Stack)
			facing, _ := m["facing"].(cube.Face)
			rotations, _ := m["rotations"].(int64)
			dropChance, _ := m["dropChance"].(float64)
			glowing, _ := m["glowing"].(bool)
			return r.vm.ToValue(block.ItemFrame{Facing: facing, Item: i, Rotations: int(rotations), DropChance: dropChance, Glowing: glowing})
		}).
		Method("jukebox", func(c goja.FunctionCall) goja.Value {
			m := c.Argument(0).Export().(map[string]any)
			i, _ := m["item"].(item.Stack)
			return r.vm.ToValue(block.Jukebox{Item: i})
		}).
		Method("kelp", func(c goja.FunctionCall) goja.Value {
			m := c.Argument(0).Export().(map[string]any)
			age, _ := m["age"].(int64)
			return r.vm.ToValue(block.Kelp{Age: int(age)})
		}).
		Method("ladder", func(c goja.FunctionCall) goja.Value {
			m := c.Argument(0).Export().(map[string]any)
			facing, _ := m["facing"].(cube.Direction)
			return r.vm.ToValue(block.Ladder{Facing: facing})
		}).
		Method("lantern", func(c goja.FunctionCall) goja.Value {
			m := c.Argument(0).Export().(map[string]any)
			t, _ := m["type"].(block.FireType)
			hanging, _ := m["hanging"].(bool)
			return r.vm.ToValue(block.Lantern{Type: t, Hanging: hanging})
		}).
		Method("lapis", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(block.Lapis{})
		}).
		Method("lapisOre", func(c goja.FunctionCall) goja.Value {
			m := c.Argument(0).Export().(map[string]any)
			t, _ := m["type"].(block.OreType)
			return r.vm.ToValue(block.LapisOre{Type: t})
		}).
		Method("lava", func(c goja.FunctionCall) goja.Value {
			m := c.Argument(0).Export().(map[string]any)
			depth, _ := m["depth"].(int64)
			still, _ := m["still"].(bool)
			falling, _ := m["falling"].(bool)
			return r.vm.ToValue(block.Lava{Depth: int(depth), Still: still, Falling: falling})
		}).
		Method("lavaDamageSource", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(block.LavaDamageSource{})
		}).
		Method("leaves", func(c goja.FunctionCall) goja.Value {
			m := c.Argument(0).Export().(map[string]any)
			wood, _ := m["wood"].(block.WoodType)
			persistent, _ := m["persistent"].(bool)
			shouldUpdate, _ := m["shouldUpdate"].(bool)
			return r.vm.ToValue(block.Leaves{Wood: wood, Persistent: persistent, ShouldUpdate: shouldUpdate})
		}).
		Method("lectern", func(c goja.FunctionCall) goja.Value {
			m := c.Argument(0).Export().(map[string]any)
			facing, _ := m["facing"].(cube.Direction)
			i, _ := m["item"].(item.Stack)
			page, _ := m["page"].(int64)
			return r.vm.ToValue(block.Lectern{Facing: facing, Book: i, Page: int(page)})
		}).
		Method("light", func(c goja.FunctionCall) goja.Value {
			m := c.Argument(0).Export().(map[string]any)
			level, _ := m["level"].(int64)
			return r.vm.ToValue(block.Light{Level: int(level)})
		}).
		Method("litPumpkin", func(c goja.FunctionCall) goja.Value {
			m := c.Argument(0).Export().(map[string]any)
			facing, _ := m["facing"].(cube.Direction)
			return r.vm.ToValue(block.LitPumpkin{Facing: facing})
		}).
		Method("log", func(c goja.FunctionCall) goja.Value {
			m := c.Argument(0).Export().(map[string]any)
			wood, _ := m["wood"].(block.WoodType)
			axis, _ := m["axis"].(cube.Axis)
			stripped, _ := m["stripped"].(bool)
			return r.vm.ToValue(block.Log{Wood: wood, Axis: axis, Stripped: stripped})
		}).
		Method("loom", func(c goja.FunctionCall) goja.Value {
			m := c.Argument(0).Export().(map[string]any)
			facing, _ := m["facing"].(cube.Direction)
			return r.vm.ToValue(block.Loom{Facing: facing})
		}).
		Method("melon", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(block.Melon{})
		}).
		Method("melonSeeds", func(c goja.FunctionCall) goja.Value {
			m := c.Argument(0).Export().(map[string]any)
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
		Method("muddyMangroveRoots", func(c goja.FunctionCall) goja.Value {
			m := c.Argument(0).Export().(map[string]any)
			axis, _ := m["axis"].(cube.Axis)
			return r.vm.ToValue(block.MuddyMangroveRoots{Axis: axis})
		}).
		Method("netherBrickFence", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(block.NetherBrickFence{})
		}).
		Method("netherBricks", func(c goja.FunctionCall) goja.Value {
			m := c.Argument(0).Export().(map[string]any)
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
		Method("netherWart", func(c goja.FunctionCall) goja.Value {
			m := c.Argument(0).Export().(map[string]any)
			age, _ := m["age"].(int64)
			return r.vm.ToValue(block.NetherWart{Age: int(age)})
		}).
		Method("netherWartBlock", func(c goja.FunctionCall) goja.Value {
			m := c.Argument(0).Export().(map[string]any)
			warped, _ := m["warped"].(bool)
			return r.vm.ToValue(block.NetherWartBlock{Warped: warped})
		}).
		Method("netherite", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(block.Netherite{})
		}).
		Method("netherrack", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(block.Netherrack{})
		}).
		Method("note", func(c goja.FunctionCall) goja.Value {
			m := c.Argument(0).Export().(map[string]any)
			pitch, _ := m["pitch"].(int64)
			return r.vm.ToValue(block.Note{Pitch: int(pitch)})
		}).
		Method("obsidian", func(c goja.FunctionCall) goja.Value {
			m := c.Argument(0).Export().(map[string]any)
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
		Method("pinkPetals", func(c goja.FunctionCall) goja.Value {
			m := c.Argument(0).Export().(map[string]any)
			facing, _ := m["facing"].(cube.Direction)
			additional, _ := m["additionalCount"].(int64)
			return r.vm.ToValue(block.PinkPetals{Facing: facing, AdditionalCount: int(additional)})
		}).
		Method("planks", func(c goja.FunctionCall) goja.Value {
			m := c.Argument(0).Export().(map[string]any)
			wood, _ := m["wood"].(block.WoodType)
			return r.vm.ToValue(block.Planks{Wood: wood})
		}).
		Method("podzol", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(block.Podzol{})
		}).
		Method("polishedBlackstoneBrick", func(c goja.FunctionCall) goja.Value {
			m := c.Argument(0).Export().(map[string]any)
			cracked, _ := m["cracked"].(bool)
			return r.vm.ToValue(block.PolishedBlackstoneBrick{Cracked: cracked})
		}).
		Method("polishedTuff", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(block.PolishedTuff{})
		}).
		Method("potato", func(c goja.FunctionCall) goja.Value {
			m := c.Argument(0).Export().(map[string]any)
			growth, _ := m["growth"].(int64)
			b := block.Potato{}
			b.Growth = int(growth)
			return r.vm.ToValue(b)
		}).
		Method("prismarine", func(c goja.FunctionCall) goja.Value {
			m := c.Argument(0).Export().(map[string]any)
			t, _ := m["type"].(block.PrismarineType)
			return r.vm.ToValue(block.Prismarine{Type: t})
		}).
		Method("pumpkin", func(c goja.FunctionCall) goja.Value {
			m := c.Argument(0).Export().(map[string]any)
			facing, _ := m["facing"].(cube.Direction)
			return r.vm.ToValue(block.Pumpkin{Facing: facing})
		}).
		Method("pumpkinSeeds", func(c goja.FunctionCall) goja.Value {
			m := c.Argument(0).Export().(map[string]any)
			direction, _ := m["direction"].(cube.Face)
			growth, _ := m["growth"].(int64)
			b := block.PumpkinSeeds{Direction: direction}
			b.Growth = int(growth)
			return r.vm.ToValue(b)
		}).
		Method("purpur", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(block.Purpur{})
		}).
		Method("purpurPillar", func(c goja.FunctionCall) goja.Value {
			m := c.Argument(0).Export().(map[string]any)
			axis, _ := m["axis"].(cube.Axis)
			return r.vm.ToValue(block.PurpurPillar{Axis: axis})
		}).
		Method("quartz", func(c goja.FunctionCall) goja.Value {
			m := c.Argument(0).Export().(map[string]any)
			smooth, _ := m["smooth"].(bool)
			return r.vm.ToValue(block.Quartz{Smooth: smooth})
		}).
		Method("quartzBricks", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(block.QuartzBricks{})
		}).
		Method("quartzPillar", func(c goja.FunctionCall) goja.Value {
			m := c.Argument(0).Export().(map[string]any)
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
		Method("resinBricks", func(c goja.FunctionCall) goja.Value {
			m := c.Argument(0).Export().(map[string]any)
			chiseled, _ := m["chiseled"].(bool)
			return r.vm.ToValue(block.ResinBricks{Chiseled: chiseled})
		}).
		Method("sand", func(c goja.FunctionCall) goja.Value {
			m := c.Argument(0).Export().(map[string]any)
			red, _ := m["red"].(bool)
			return r.vm.ToValue(block.Sand{Red: red})
		}).
		Method("sandstone", func(c goja.FunctionCall) goja.Value {
			m := c.Argument(0).Export().(map[string]any)
			t, _ := m["type"].(block.SandstoneType)
			red, _ := m["red"].(bool)
			return r.vm.ToValue(block.Sandstone{Type: t, Red: red})
		}).
		Method("seaLantern", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(block.SeaLantern{})
		}).
		Method("seaPickle", func(c goja.FunctionCall) goja.Value {
			m := c.Argument(0).Export().(map[string]any)
			dead, _ := m["dead"].(bool)
			return r.vm.ToValue(block.SeaPickle{Dead: dead})
		}).
		Method("shortGrass", func(c goja.FunctionCall) goja.Value {
			m := c.Argument(0).Export().(map[string]any)
			double, _ := m["double"].(bool)
			return r.vm.ToValue(block.ShortGrass{Double: double})
		}).
		Method("shroomlight", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(block.Shroomlight{})
		}).
		Method("sign", func(c goja.FunctionCall) goja.Value {
			m := c.Argument(0).Export().(map[string]any)
			wood, _ := m["wood"].(block.WoodType)
			a, _ := m["attach"].(block.Attachment)
			front, _ := m["front"].(block.SignText)
			back, _ := m["back"].(block.SignText)
			waxed, _ := m["waxed"].(bool)
			return r.vm.ToValue(block.Sign{Wood: wood, Attach: a, Front: front, Back: back, Waxed: waxed})
		}).
		Method("signText", func(c goja.FunctionCall) goja.Value {
			m := c.Argument(0).Export().(map[string]any)
			text, _ := m["text"].(string)
			owner, _ := m["owner"].(string)
			colour, _ := m["baseColour"].(color.RGBA)
			glowing, _ := m["glowing"].(bool)
			return r.vm.ToValue(block.SignText{Text: text, Owner: owner, BaseColour: colour, Glowing: glowing})
		}).
		Method("skull", func(c goja.FunctionCall) goja.Value {
			m := c.Argument(0).Export().(map[string]any)
			t, _ := m["type"].(block.SkullType)
			a, _ := m["attach"].(block.Attachment)
			return r.vm.ToValue(block.Skull{Type: t, Attach: a})
		}).
		Method("slab", func(c goja.FunctionCall) goja.Value {
			m := c.Argument(0).Export().(map[string]any)
			b, _ := m["block"].(world.Block)
			double, _ := m["double"].(bool)
			top, _ := m["top"].(bool)
			return r.vm.ToValue(block.Slab{Block: b, Double: double, Top: top})
		}).
		Method("smithingTable", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(block.SmithingTable{})
		}).
		Method("smoker", func(c goja.FunctionCall) goja.Value {
			m := c.Argument(0).Export().(map[string]any)
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
		Method("sponge", func(c goja.FunctionCall) goja.Value {
			m := c.Argument(0).Export().(map[string]any)
			wet, _ := m["wet"].(bool)
			return r.vm.ToValue(block.Sponge{Wet: wet})
		}).
		Method("sporeBlossom", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(block.SporeBlossom{})
		}).
		Method("stainedGlass", func(c goja.FunctionCall) goja.Value {
			m := c.Argument(0).Export().(map[string]any)
			colour, _ := m["colour"].(item.Colour)
			return r.vm.ToValue(block.StainedGlass{Colour: colour})
		}).
		Method("stainedGlassPane", func(c goja.FunctionCall) goja.Value {
			m := c.Argument(0).Export().(map[string]any)
			colour, _ := m["colour"].(item.Colour)
			return r.vm.ToValue(block.StainedGlassPane{Colour: colour})
		}).
		Method("stainedTerracotta", func(c goja.FunctionCall) goja.Value {
			m := c.Argument(0).Export().(map[string]any)
			colour, _ := m["colour"].(item.Colour)
			return r.vm.ToValue(block.StainedTerracotta{Colour: colour})
		}).
		Method("stairs", func(c goja.FunctionCall) goja.Value {
			m := c.Argument(0).Export().(map[string]any)
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
		Method("stone", func(c goja.FunctionCall) goja.Value {
			m := c.Argument(0).Export().(map[string]any)
			smooth, _ := m["smooth"].(bool)
			return r.vm.ToValue(block.Stone{Smooth: smooth})
		}).
		Method("stoneBricks", func(c goja.FunctionCall) goja.Value {
			m := c.Argument(0).Export().(map[string]any)
			t, _ := m["type"].(block.StoneBricksType)
			return r.vm.ToValue(block.StoneBricks{Type: t})
		}).
		Method("stonecutter", func(c goja.FunctionCall) goja.Value {
			m := c.Argument(0).Export().(map[string]any)
			facing, _ := m["facing"].(cube.Direction)
			return r.vm.ToValue(block.Stonecutter{Facing: facing})
		}).
		Method("stopCrackAction", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(block.StopCrackAction{})
		}).
		Method("sugarCane", func(c goja.FunctionCall) goja.Value {
			m := c.Argument(0).Export().(map[string]any)
			age, _ := m["age"].(int64)
			return r.vm.ToValue(block.SugarCane{Age: int(age)})
		}).
		Method("tnt", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(block.TNT{})
		}).
		Method("terracotta", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(block.Terracotta{})
		}).
		Method("torch", func(c goja.FunctionCall) goja.Value {
			m := c.Argument(0).Export().(map[string]any)
			facing, _ := m["facing"].(cube.Face)
			t, _ := m["type"].(block.FireType)
			return r.vm.ToValue(block.Torch{Facing: facing, Type: t})
		}).
		Method("tuff", func(c goja.FunctionCall) goja.Value {
			m := c.Argument(0).Export().(map[string]any)
			chiseled, _ := m["chiseled"].(bool)
			return r.vm.ToValue(block.Tuff{Chiseled: chiseled})
		}).
		Method("tuffBricks", func(c goja.FunctionCall) goja.Value {
			m := c.Argument(0).Export().(map[string]any)
			chiseled, _ := m["chiseled"].(bool)
			return r.vm.ToValue(block.TuffBricks{Chiseled: chiseled})
		}).
		Method("vines", func(c goja.FunctionCall) goja.Value {
			m := c.Argument(0).Export().(map[string]any)
			north, _ := m["north"].(bool)
			east, _ := m["east"].(bool)
			south, _ := m["south"].(bool)
			west, _ := m["west"].(bool)
			return r.vm.ToValue(block.Vines{NorthDirection: north, EastDirection: east, SouthDirection: south, WestDirection: west})
		}).
		Method("wall", func(c goja.FunctionCall) goja.Value {
			m := c.Argument(0).Export().(map[string]any)
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
		Method("water", func(c goja.FunctionCall) goja.Value {
			m := c.Argument(0).Export().(map[string]any)
			depth, _ := m["depth"].(int64)
			still, _ := m["still"].(bool)
			falling, _ := m["falling"].(bool)
			return r.vm.ToValue(block.Water{Depth: int(depth), Still: still, Falling: falling})
		}).
		Method("wheatSeeds", func(c goja.FunctionCall) goja.Value {
			m := c.Argument(0).Export().(map[string]any)
			growth, _ := m["growth"].(int64)
			b := block.WheatSeeds{}
			b.Growth = int(growth)
			return r.vm.ToValue(b)
		}).
		Method("wood", func(c goja.FunctionCall) goja.Value {
			m := c.Argument(0).Export().(map[string]any)
			wood, _ := m["wood"].(block.WoodType)
			axis, _ := m["axis"].(cube.Axis)
			stripped, _ := m["stripped"].(bool)
			return r.vm.ToValue(block.Wood{Wood: wood, Axis: axis, Stripped: stripped})
		}).
		Method("woodDoor", func(c goja.FunctionCall) goja.Value {
			m := c.Argument(0).Export().(map[string]any)
			wood, _ := m["wood"].(block.WoodType)
			facing, _ := m["facing"].(cube.Direction)
			open, _ := m["open"].(bool)
			top, _ := m["top"].(bool)
			right, _ := m["right"].(bool)
			return r.vm.ToValue(block.WoodDoor{Wood: wood, Facing: facing, Open: open, Top: top, Right: right})
		}).
		Method("woodFence", func(c goja.FunctionCall) goja.Value {
			m := c.Argument(0).Export().(map[string]any)
			wood, _ := m["wood"].(block.WoodType)
			return r.vm.ToValue(block.WoodFence{Wood: wood})
		}).
		Method("woodFenceGate", func(c goja.FunctionCall) goja.Value {
			m := c.Argument(0).Export().(map[string]any)
			wood, _ := m["wood"].(block.WoodType)
			facing, _ := m["facing"].(cube.Direction)
			open, _ := m["open"].(bool)
			lowered, _ := m["lowered"].(bool)
			return r.vm.ToValue(block.WoodFenceGate{Wood: wood, Facing: facing, Open: open, Lowered: lowered})
		}).
		Method("woodTrapdoor", func(c goja.FunctionCall) goja.Value {
			m := c.Argument(0).Export().(map[string]any)
			wood, _ := m["wood"].(block.WoodType)
			facing, _ := m["facing"].(cube.Direction)
			open, _ := m["open"].(bool)
			top, _ := m["top"].(bool)
			return r.vm.ToValue(block.WoodTrapdoor{Wood: wood, Facing: facing, Open: open, Top: top})
		}).
		Method("wool", func(c goja.FunctionCall) goja.Value {
			m := c.Argument(0).Export().(map[string]any)
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
