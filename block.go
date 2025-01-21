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
		Const("anvilType", map[string]block.AnvilType{
			"undamaged":       block.UndamagedAnvil(),
			"slightlyDamaged": block.SlightlyDamagedAnvil(),
			"veryDamaged":     block.VeryDamagedAnvil(),
		}).
		Const("bannerPattern", map[string]block.BannerPatternType{
			"border":               block.BorderBannerPattern(),
			"bricks":               block.BricksBannerPattern(),
			"circle":               block.CircleBannerPattern(),
			"creeper":              block.CreeperBannerPattern(),
			"cross":                block.CrossBannerPattern(),
			"curlyBorder":          block.CurlyBorderBannerPattern(),
			"diagonalLeft":         block.DiagonalLeftBannerPattern(),
			"diagonalRight":        block.DiagonalRightBannerPattern(),
			"diagonalUpLeft":       block.DiagonalUpLeftBannerPattern(),
			"diagonalUpRight":      block.DiagonalUpRightBannerPattern(),
			"flower":               block.FlowerBannerPattern(),
			"gradient":             block.GradientBannerPattern(),
			"gradientUp":           block.GradientUpBannerPattern(),
			"halfHorizontal":       block.HalfHorizontalBannerPattern(),
			"halfHorizontalBottom": block.HalfHorizontalBottomBannerPattern(),
			"halfVertical":         block.HalfVerticalBannerPattern(),
			"halfVerticalRight":    block.HalfVerticalRightBannerPattern(),
			"mojang":               block.MojangBannerPattern(),
			"rhombus":              block.RhombusBannerPattern(),
			"skull":                block.SkullBannerPattern(),
			"smallStripes":         block.SmallStripesBannerPattern(),
			"squareBottomLeft":     block.SquareBottomLeftBannerPattern(),
			"squareBottomRight":    block.SquareBottomRightBannerPattern(),
			"squareTopLeft":        block.SquareTopLeftBannerPattern(),
			"squareTopRight":       block.SquareTopRightBannerPattern(),
			"straightCross":        block.StraightCrossBannerPattern(),
			"stripeBottom":         block.StripeBottomBannerPattern(),
			"stripeCenter":         block.StripeCenterBannerPattern(),
			"stripeDownLeft":       block.StripeDownLeftBannerPattern(),
			"stripeDownRight":      block.StripeDownRightBannerPattern(),
			"stripeLeft":           block.StripeLeftBannerPattern(),
			"stripeMiddle":         block.StripeMiddleBannerPattern(),
			"stripeRight":          block.StripeRightBannerPattern(),
			"stripeTop":            block.StripeTopBannerPattern(),
			"triangleBottom":       block.TriangleBottomBannerPattern(),
			"triangleTop":          block.TriangleTopBannerPattern(),
			"trianglesBottom":      block.TrianglesBottomBannerPattern(),
			"trianglesTop":         block.TrianglesTopBannerPattern(),
			"globe":                block.GlobeBannerPattern(),
			"piglin":               block.PiglinBannerPattern(),
			"flow":                 block.FlowBannerPattern(),
			"guster":               block.GusterBannerPattern(),
		}).
		Const("blackstoneType", map[string]block.BlackstoneType{
			"normal":           block.NormalBlackstone(),
			"gilded":           block.GildedBlackstone(),
			"polished":         block.PolishedBlackstone(),
			"chiseledPolished": block.ChiseledPolishedBlackstone(),
		}).
		Const("copperType", map[string]block.CopperType{
			"normal":   block.NormalCopper(),
			"cut":      block.CutCopper(),
			"chiseled": block.ChiseledCopper(),
		}).
		Const("coralType", map[string]block.CoralType{
			"tube":   block.TubeCoral(),
			"brain":  block.BrainCoral(),
			"bubble": block.BubbleCoral(),
			"fire":   block.FireCoral(),
			"horn":   block.HornCoral(),
		}).
		Const("deepslateType", map[string]block.DeepslateType{
			"normal":   block.NormalDeepslate(),
			"cobbled":  block.CobbledDeepslate(),
			"polished": block.PolishedDeepslate(),
			"chiseled": block.ChiseledDeepslate(),
		}).
		Const("fireType", map[string]block.FireType{
			"normal": block.NormalFire(),
			"soul":   block.SoulFire(),
		}).
		Const("doubleFlowerType", map[string]block.DoubleFlowerType{
			"sunflower": block.Sunflower(),
			"lilac":     block.Lilac(),
			"roseBush":  block.RoseBush(),
			"peony":     block.Peony(),
		}).
		Const("doubleTallGrassType", map[string]block.DoubleTallGrassType{
			"normal": block.NormalDoubleTallGrass(),
			"fern":   block.FernDoubleTallGrass(),
		}).
		Const("flowerType", map[string]block.FlowerType{
			"dandelion":       block.Dandelion(),
			"poppy":           block.Poppy(),
			"blueOrchid":      block.BlueOrchid(),
			"allium":          block.Allium(),
			"azureBluet":      block.AzureBluet(),
			"redTulip":        block.RedTulip(),
			"orangeTulip":     block.OrangeTulip(),
			"whiteTulip":      block.WhiteTulip(),
			"pinkTulip":       block.PinkTulip(),
			"oxeyeDaisy":      block.OxeyeDaisy(),
			"cornflower":      block.Cornflower(),
			"lilyOfTheValley": block.LilyOfTheValley(),
			"witherRose":      block.WitherRose(),
		}).
		Const("froglightType", map[string]block.FroglightType{
			"pearlescent": block.Pearlescent(),
			"verdant":     block.Verdant(),
			"ochre":       block.Ochre(),
		}).
		Const("netherBricksType", map[string]block.NetherBricksType{
			"normal":   block.NormalNetherBricks(),
			"red":      block.RedNetherBricks(),
			"cracked":  block.CrackedNetherBricks(),
			"chiseled": block.ChiseledNetherBricks(),
		}).
		Const("oreType", map[string]block.OreType{
			"stone":     block.StoneOre(),
			"deepslate": block.DeepslateOre(),
		}).
		Const("oxidation", map[string]block.OxidationType{
			"unoxidised": block.UnoxidisedOxidation(),
			"exposed":    block.ExposedOxidation(),
			"weathered":  block.WeatheredOxidation(),
			"oxidised":   block.OxidisedOxidation(),
		}).
		Const("prismarineType", map[string]block.PrismarineType{
			"normal": block.NormalPrismarine(),
			"dark":   block.DarkPrismarine(),
			"bricks": block.BrickPrismarine(),
		}).
		Const("sandstoneType", map[string]block.SandstoneType{
			"normal":   block.NormalSandstone(),
			"cut":      block.CutSandstone(),
			"chiseled": block.ChiseledSandstone(),
			"smooth":   block.SmoothSandstone(),
		}).
		Const("skullType", map[string]block.SkullType{
			"skeleton":       block.SkeletonSkull(),
			"witherSkeleton": block.WitherSkeletonSkull(),
			"zombie":         block.ZombieHead(),
			"player":         block.PlayerHead(),
			"creeper":        block.CreeperHead(),
			"dragon":         block.DragonHead(),
			"piglin":         block.PiglinHead(),
		}).
		Const("stoneBricksType", map[string]block.StoneBricksType{
			"normal":   block.NormalStoneBricks(),
			"mossy":    block.MossyStoneBricks(),
			"cracked":  block.CrackedStoneBricks(),
			"chiseled": block.ChiseledStoneBricks(),
		}).
		Const("wallConnection", map[string]block.WallConnectionType{
			"none":  block.NoWallConnection(),
			"short": block.ShortWallConnection(),
			"tall":  block.TallWallConnection(),
		}).
		Const("woodType", map[string]block.WoodType{
			"oak":      block.OakWood(),
			"spruce":   block.SpruceWood(),
			"birch":    block.BirchWood(),
			"jungle":   block.JungleWood(),
			"acacia":   block.AcaciaWood(),
			"darkOak":  block.DarkOakWood(),
			"crimson":  block.CrimsonWood(),
			"warped":   block.WarpedWood(),
			"mangrove": block.MangroveWood(),
			"cherry":   block.CherryWood(),
			"paleOak":  block.PaleOakWood(),
		}).
		Const("grindstoneAttachment", map[string]block.GrindstoneAttachment{
			"standing": block.StandingGrindstoneAttachment(),
			"hanging":  block.HangingGrindstoneAttachment(),
			"wall":     block.WallGrindstoneAttachment(),
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
		Method("anvil", func(c goja.FunctionCall) goja.Value {
			t, ok := c.Argument(0).Export().(block.AnvilType)
			if !ok {
				panic(r.vm.NewTypeError("argument 0 is not a block.AnvilType"))
			}
			facing, ok := c.Argument(1).Export().(cube.Direction)
			if !ok {
				panic(r.vm.NewTypeError("argument 1 is not a cube.Direction"))
			}
			return r.vm.ToValue(block.Anvil{Type: t, Facing: facing})
		}).
		Method("banner", func(c goja.FunctionCall) goja.Value {
			colour, ok := c.Argument(0).Export().(item.Colour)
			if !ok {
				panic(r.vm.NewTypeError("argument 0 is not an item.Colour"))
			}
			attachment, ok := c.Argument(1).Export().(block.Attachment)
			if !ok {
				panic(r.vm.NewTypeError("argument 1 is not a block.Attachment"))
			}
			patterns, ok := c.Argument(2).Export().([]block.BannerPatternLayer)
			if !ok {
				panic(r.vm.NewTypeError("argument 2 is not a []block.BannerPatternLayer"))
			}
			illager, _ := c.Argument(3).Export().(bool)
			return r.vm.ToValue(block.Banner{Colour: colour, Attach: attachment, Patterns: patterns, Illager: illager})
		}).
		Method("bannerPatternLayer", func(c goja.FunctionCall) goja.Value {
			t, ok := c.Argument(0).Export().(block.BannerPatternType)
			if !ok {
				panic(r.vm.NewTypeError("argument 0 is not a block.BannerPatternType"))
			}
			colour, ok := c.Argument(1).Export().(item.Colour)
			if !ok {
				panic(r.vm.NewTypeError("argument 1 is not an item.Colour"))
			}
			return r.vm.ToValue(block.BannerPatternLayer{Type: t, Colour: colour})
		}).
		Method("barrel", func(c goja.FunctionCall) goja.Value {
			facing, ok := c.Argument(0).Export().(cube.Face)
			if !ok {
				panic(r.vm.NewTypeError("argument 0 is not a cube.Face"))
			}
			open, _ := c.Argument(1).Export().(bool)
			customName, _ := c.Argument(2).Export().(string)
			return r.vm.ToValue(block.Barrel{Facing: facing, Open: open, CustomName: customName})
		}).
		Method("barrier", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(block.Barrier{})
		}).
		Method("basalt", func(c goja.FunctionCall) goja.Value {
			axis, ok := c.Argument(0).Export().(cube.Axis)
			if !ok {
				panic(r.vm.NewTypeError("argument 0 is not a cube.Axis"))
			}
			polished, _ := c.Argument(1).Export().(bool)
			return r.vm.ToValue(block.Basalt{Axis: axis, Polished: polished})
		}).
		Method("beacon", func(c goja.FunctionCall) goja.Value {
			primary, ok := c.Argument(0).Export().(effect.LastingType)
			if !ok {
				panic(r.vm.NewTypeError("argument 0 is not an effect.LastingType"))
			}
			secondary, ok := c.Argument(1).Export().(effect.LastingType)
			if !ok {
				panic(r.vm.NewTypeError("argument 1 is not an effect.LastingType"))
			}
			return r.vm.ToValue(block.Beacon{Primary: primary, Secondary: secondary})
		}).
		Method("bedrock", func(c goja.FunctionCall) goja.Value {
			infinite, _ := c.Argument(0).Export().(bool)
			return r.vm.ToValue(block.Bedrock{InfiniteBurning: infinite})
		}).
		Method("beetrootSeeds", func(c goja.FunctionCall) goja.Value {
			growth, ok := c.Argument(0).Export().(int64)
			if !ok {
				panic(r.vm.NewTypeError("argument 0 is not an int64"))
			}
			b := block.BeetrootSeeds{}
			b.Growth = int(growth)
			return r.vm.ToValue(b)
		}).
		Method("blackstone", func(c goja.FunctionCall) goja.Value {
			t, ok := c.Argument(0).Export().(block.BlackstoneType)
			if !ok {
				panic(r.vm.NewTypeError("argument 0 is not a block.BlackstoneType"))
			}
			return r.vm.ToValue(block.Blackstone{Type: t})
		}).
		Method("blastFurnace", func(c goja.FunctionCall) goja.Value {
			facing, ok := c.Argument(0).Export().(cube.Direction)
			if !ok {
				panic(r.vm.NewTypeError("argument 0 is not a cube.Direction"))
			}
			lit, _ := c.Argument(1).Export().(bool)
			return r.vm.ToValue(block.BlastFurnace{Facing: facing, Lit: lit})
		}).
		Method("blueIce", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(block.BlueIce{})
		}).
		Method("bone", func(c goja.FunctionCall) goja.Value {
			axis, ok := c.Argument(0).Export().(cube.Axis)
			if !ok {
				panic(r.vm.NewTypeError("argument 0 is not a cube.Axis"))
			}
			return r.vm.ToValue(block.Bone{Axis: axis})
		}).
		Method("bookshelf", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(block.Bookshelf{})
		}).
		Method("brewingStand", func(c goja.FunctionCall) goja.Value {
			left, _ := c.Argument(0).Export().(bool)
			middle, _ := c.Argument(1).Export().(bool)
			right, _ := c.Argument(2).Export().(bool)
			return r.vm.ToValue(block.BrewingStand{LeftSlot: left, MiddleSlot: middle, RightSlot: right})
		}).
		Method("bricks", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(block.Bricks{})
		}).
		Method("cactus", func(c goja.FunctionCall) goja.Value {
			age, ok := c.Argument(0).Export().(int64)
			if !ok {
				panic(r.vm.NewTypeError("argument 0 is not an int64"))
			}
			return r.vm.ToValue(block.Cactus{Age: int(age)})
		}).
		Method("cake", func(c goja.FunctionCall) goja.Value {
			bites, ok := c.Argument(0).Export().(int64)
			if !ok {
				panic(r.vm.NewTypeError("argument 0 is not an int64"))
			}
			return r.vm.ToValue(block.Cake{Bites: int(bites)})
		}).
		Method("calcite", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(block.Calcite{})
		}).
		Method("campfire", func(c goja.FunctionCall) goja.Value {
			t, ok := c.Argument(0).Export().(block.FireType)
			if !ok {
				panic(r.vm.NewTypeError("argument 0 is not a block.FireType"))
			}
			rawItems, ok := c.Argument(1).Export().([]block.CampfireItem)
			if !ok {
				panic(r.vm.NewTypeError("argument 1 is not a []block.CampfireItem"))
			}
			facing, ok := c.Argument(2).Export().(cube.Direction)
			if !ok {
				panic(r.vm.NewTypeError("argument 2 is not a cube.Direction"))
			}
			extinguished, _ := c.Argument(3).Export().(bool)
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
			colour, ok := c.Argument(0).Export().(item.Colour)
			if !ok {
				panic(r.vm.NewTypeError("argument 0 is not an item.Colour"))
			}
			return r.vm.ToValue(block.Carpet{Colour: colour})
		}).
		Method("carrot", func(c goja.FunctionCall) goja.Value {
			growth, ok := c.Argument(0).Export().(int64)
			if !ok {
				panic(r.vm.NewTypeError("argument 0 is not an int64"))
			}
			b := block.Carrot{}
			b.Growth = int(growth)
			return r.vm.ToValue(b)
		}).
		Method("chain", func(c goja.FunctionCall) goja.Value {
			axis, ok := c.Argument(0).Export().(cube.Axis)
			if !ok {
				panic(r.vm.NewTypeError("argument 0 is not a cube.Axis"))
			}
			return r.vm.ToValue(block.Chain{Axis: axis})
		}).
		Method("chest", func(c goja.FunctionCall) goja.Value {
			facing, ok := c.Argument(0).Export().(cube.Direction)
			if !ok {
				panic(r.vm.NewTypeError("argument 0 is not a cube.Direction"))
			}
			customName, _ := c.Argument(1).Export().(string)
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
			t, ok := c.Argument(0).Export().(block.OreType)
			if !ok {
				panic(r.vm.NewTypeError("argument 0 is not a block.OreType"))
			}
			return r.vm.ToValue(block.CoalOre{Type: t})
		}).
		Method("cobblestone", func(c goja.FunctionCall) goja.Value {
			mossy, _ := c.Argument(0).Export().(bool)
			return r.vm.ToValue(block.Cobblestone{Mossy: mossy})
		}).
		Method("cocoaBean", func(c goja.FunctionCall) goja.Value {
			facing, ok := c.Argument(0).Export().(cube.Direction)
			if !ok {
				panic(r.vm.NewTypeError("argument 0 is not a cube.Direction"))
			}
			age, ok := c.Argument(1).Export().(int64)
			if !ok {
				panic(r.vm.NewTypeError("argument 1 is not an int64"))
			}
			return r.vm.ToValue(block.CocoaBean{Facing: facing, Age: int(age)})
		}).
		Method("composter", func(c goja.FunctionCall) goja.Value {
			level, ok := c.Argument(0).Export().(int64)
			if !ok {
				panic(r.vm.NewTypeError("argument 0 is not an int64"))
			}
			return r.vm.ToValue(block.Composter{Level: int(level)})
		}).
		Method("concrete", func(c goja.FunctionCall) goja.Value {
			colour, ok := c.Argument(0).Export().(item.Colour)
			if !ok {
				panic(r.vm.NewTypeError("argument 0 is not an item.Colour"))
			}
			return r.vm.ToValue(block.Concrete{Colour: colour})
		}).
		Method("concretePowder", func(c goja.FunctionCall) goja.Value {
			colour, ok := c.Argument(0).Export().(item.Colour)
			if !ok {
				panic(r.vm.NewTypeError("argument 0 is not an item.Colour"))
			}
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
			t, ok := c.Argument(0).Export().(block.CopperType)
			if !ok {
				panic(r.vm.NewTypeError("argument 0 is not a block.CopperType"))
			}
			o, ok := c.Argument(1).Export().(block.OxidationType)
			if !ok {
				panic(r.vm.NewTypeError("argument 1 is not a block.OxidationType"))
			}
			waxed, _ := c.Argument(2).Export().(bool)
			return r.vm.ToValue(block.Copper{Type: t, Oxidation: o, Waxed: waxed})
		}).
		Method("copperDoor", func(c goja.FunctionCall) goja.Value {
			facing, ok := c.Argument(0).Export().(cube.Direction)
			if !ok {
				panic(r.vm.NewTypeError("argument 0 is not a cube.Direction"))
			}
			o, ok := c.Argument(1).Export().(block.OxidationType)
			if !ok {
				panic(r.vm.NewTypeError("argument 1 is not a block.OxidationType"))
			}
			waxed, _ := c.Argument(2).Export().(bool)
			open, _ := c.Argument(3).Export().(bool)
			top, _ := c.Argument(4).Export().(bool)
			right, _ := c.Argument(5).Export().(bool)
			return r.vm.ToValue(block.CopperDoor{Facing: facing, Oxidation: o, Waxed: waxed, Open: open, Top: top, Right: right})
		}).
		Method("copperGrate", func(c goja.FunctionCall) goja.Value {
			o, ok := c.Argument(0).Export().(block.OxidationType)
			if !ok {
				panic(r.vm.NewTypeError("argument 0 is not a block.OxidationType"))
			}
			waxed, _ := c.Argument(1).Export().(bool)
			return r.vm.ToValue(block.CopperGrate{Oxidation: o, Waxed: waxed})
		}).
		Method("copperOre", func(c goja.FunctionCall) goja.Value {
			t, ok := c.Argument(0).Export().(block.OreType)
			if !ok {
				panic(r.vm.NewTypeError("argument 0 is not a block.OreType"))
			}
			return r.vm.ToValue(block.CopperOre{Type: t})
		}).
		Method("copperTrapdoor", func(c goja.FunctionCall) goja.Value {
			facing, ok := c.Argument(0).Export().(cube.Direction)
			if !ok {
				panic(r.vm.NewTypeError("argument 0 is not a cube.Direction"))
			}
			o, ok := c.Argument(1).Export().(block.OxidationType)
			if !ok {
				panic(r.vm.NewTypeError("argument 1 is not a block.OxidationType"))
			}
			waxed, _ := c.Argument(2).Export().(bool)
			open, _ := c.Argument(3).Export().(bool)
			top, _ := c.Argument(4).Export().(bool)
			return r.vm.ToValue(block.CopperTrapdoor{Facing: facing, Oxidation: o, Waxed: waxed, Open: open, Top: top})
		}).
		Method("coral", func(c goja.FunctionCall) goja.Value {
			t, ok := c.Argument(0).Export().(block.CoralType)
			if !ok {
				panic(r.vm.NewTypeError("argument 0 is not a block.CoralType"))
			}
			dead, _ := c.Argument(1).Export().(bool)
			return r.vm.ToValue(block.Coral{Type: t, Dead: dead})
		}).
		Method("coralBlock", func(c goja.FunctionCall) goja.Value {
			t, ok := c.Argument(0).Export().(block.CoralType)
			if !ok {
				panic(r.vm.NewTypeError("argument 0 is not a block.CoralType"))
			}
			dead, _ := c.Argument(1).Export().(bool)
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
			i, ok := c.Argument(0).Export().(item.Stack)
			if !ok {
				panic(r.vm.NewTypeError("argument 0 is not an item.Stack"))
			}
			facing, ok := c.Argument(1).Export().(cube.Direction)
			if !ok {
				panic(r.vm.NewTypeError("argument 1 is not a cube.Direction"))
			}
			rawDeco, ok := c.Argument(2).Export().([]block.PotDecoration)
			if !ok {
				panic(r.vm.NewTypeError("argument 2 is not a []block.PotDecoration"))
			}
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
			t, ok := c.Argument(0).Export().(block.DeepslateType)
			if !ok {
				panic(r.vm.NewTypeError("argument 0 is not a block.DeepslateType"))
			}
			axis, ok := c.Argument(1).Export().(cube.Axis)
			if !ok {
				panic(r.vm.NewTypeError("argument 1 is not a cube.Axis"))
			}
			return r.vm.ToValue(block.Deepslate{Type: t, Axis: axis})
		}).
		Method("deepslateBricks", func(c goja.FunctionCall) goja.Value {
			cracked, _ := c.Argument(0).Export().(bool)
			return r.vm.ToValue(block.DeepslateBricks{Cracked: cracked})
		}).
		Method("deepslateTiles", func(c goja.FunctionCall) goja.Value {
			cracked, _ := c.Argument(0).Export().(bool)
			return r.vm.ToValue(block.DeepslateTiles{Cracked: cracked})
		}).
		Method("diamond", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(block.Diamond{})
		}).
		Method("diamondOre", func(c goja.FunctionCall) goja.Value {
			t, ok := c.Argument(0).Export().(block.OreType)
			if !ok {
				panic(r.vm.NewTypeError("argument 0 is not a block.OreType"))
			}
			return r.vm.ToValue(block.DiamondOre{Type: t})
		}).
		Method("dirt", func(c goja.FunctionCall) goja.Value {
			coarse, _ := c.Argument(0).Export().(bool)
			return r.vm.ToValue(block.Dirt{Coarse: coarse})
		}).
		Method("dirtPath", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(block.DirtPath{})
		}).
		Method("doubleFlower", func(c goja.FunctionCall) goja.Value {
			t, ok := c.Argument(0).Export().(block.DoubleFlowerType)
			if !ok {
				panic(r.vm.NewTypeError("argument 0 is not a block.DoubleFlowerType"))
			}
			top, _ := c.Argument(1).Export().(bool)
			return r.vm.ToValue(block.DoubleFlower{Type: t, UpperPart: top})
		}).
		Method("doubleTallGrass", func(c goja.FunctionCall) goja.Value {
			t, ok := c.Argument(0).Export().(block.DoubleTallGrassType)
			if !ok {
				panic(r.vm.NewTypeError("argument 0 is not a block.DoubleTallGrassType"))
			}
			top, _ := c.Argument(1).Export().(bool)
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
			t, ok := c.Argument(0).Export().(block.OreType)
			if !ok {
				panic(r.vm.NewTypeError("argument 0 is not a block.OreType"))
			}
			return r.vm.ToValue(block.EmeraldOre{Type: t})
		}).
		Method("enchantingTable", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(block.EnchantingTable{})
		}).
		Method("endBricks", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(block.EndBricks{})
		}).
		Method("endRod", func(c goja.FunctionCall) goja.Value {
			facing, ok := c.Argument(0).Export().(cube.Face)
			if !ok {
				panic(r.vm.NewTypeError("argument 0 is not a cube.Face"))
			}
			return r.vm.ToValue(block.EndRod{Facing: facing})
		}).
		Method("endStone", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(block.EndStone{})
		}).
		Method("enderChest", func(c goja.FunctionCall) goja.Value {
			facing, ok := c.Argument(0).Export().(cube.Direction)
			if !ok {
				panic(r.vm.NewTypeError("argument 0 is not a cube.Direction"))
			}
			return r.vm.ToValue(block.EnderChest{Facing: facing})
		}).
		Method("farmland", func(c goja.FunctionCall) goja.Value {
			hydration, ok := c.Argument(0).Export().(int64)
			if !ok {
				panic(r.vm.NewTypeError("argument 0 is not an int64"))
			}
			return r.vm.ToValue(block.Farmland{Hydration: int(hydration)})
		}).
		Method("fern", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(block.Fern{})
		}).
		Method("fire", func(c goja.FunctionCall) goja.Value {
			t, ok := c.Argument(0).Export().(block.FireType)
			if !ok {
				panic(r.vm.NewTypeError("argument 0 is not a block.FireType"))
			}
			age, ok := c.Argument(1).Export().(int64)
			if !ok {
				panic(r.vm.NewTypeError("argument 1 is not an int64"))
			}
			return r.vm.ToValue(block.Fire{Type: t, Age: int(age)})
		}).
		Method("fireDamageSource", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(block.FireDamageSource{})
		}).
		Method("fletchingTable", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(block.FletchingTable{})
		}).
		Method("flower", func(c goja.FunctionCall) goja.Value {
			t, ok := c.Argument(0).Export().(block.FlowerType)
			if !ok {
				panic(r.vm.NewTypeError("argument 0 is not a block.FlowerType"))
			}
			return r.vm.ToValue(block.Flower{Type: t})
		}).
		Method("froglight", func(c goja.FunctionCall) goja.Value {
			t, ok := c.Argument(0).Export().(block.FroglightType)
			if !ok {
				panic(r.vm.NewTypeError("argument 0 is not a block.FroglightType"))
			}
			axis, ok := c.Argument(1).Export().(cube.Axis)
			if !ok {
				panic(r.vm.NewTypeError("argument 1 is not a cube.Axis"))
			}
			return r.vm.ToValue(block.Froglight{Type: t, Axis: axis})
		}).
		Method("furnace", func(c goja.FunctionCall) goja.Value {
			facing, ok := c.Argument(0).Export().(cube.Direction)
			if !ok {
				panic(r.vm.NewTypeError("argument 0 is not a cube.Direction"))
			}
			lit, _ := c.Argument(1).Export().(bool)
			return r.vm.ToValue(block.Furnace{Facing: facing, Lit: lit})
		}).
		Method("glass", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(block.Glass{})
		}).
		Method("glassPane", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(block.GlassPane{})
		}).
		Method("glazedTerracotta", func(c goja.FunctionCall) goja.Value {
			colour, ok := c.Argument(0).Export().(item.Colour)
			if !ok {
				panic(r.vm.NewTypeError("argument 0 is not an item.Colour"))
			}
			facing, ok := c.Argument(1).Export().(cube.Direction)
			if !ok {
				panic(r.vm.NewTypeError("argument 1 is not a cube.Direction"))
			}
			return r.vm.ToValue(block.GlazedTerracotta{Colour: colour, Facing: facing})
		}).
		Method("glowstone", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(block.Glowstone{})
		}).
		Method("gold", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(block.Gold{})
		}).
		Method("goldOre", func(c goja.FunctionCall) goja.Value {
			t, ok := c.Argument(0).Export().(block.OreType)
			if !ok {
				panic(r.vm.NewTypeError("argument 0 is not a block.OreType"))
			}
			return r.vm.ToValue(block.GoldOre{Type: t})
		}).
		Method("grass", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(block.Grass{})
		}).
		Method("gravel", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(block.Gravel{})
		}).
		Method("grindstone", func(c goja.FunctionCall) goja.Value {
			facing, ok := c.Argument(0).Export().(cube.Direction)
			if !ok {
				panic(r.vm.NewTypeError("argument 0 is not a cube.Direction"))
			}
			a, ok := c.Argument(1).Export().(block.GrindstoneAttachment)
			if !ok {
				panic(r.vm.NewTypeError("argument 1 is not a block.GrindstoneAttachment"))
			}
			return r.vm.ToValue(block.Grindstone{Facing: facing, Attach: a})
		}).
		Method("hayBale", func(c goja.FunctionCall) goja.Value {
			axis, ok := c.Argument(0).Export().(cube.Axis)
			if !ok {
				panic(r.vm.NewTypeError("argument 0 is not a cube.Axis"))
			}
			return r.vm.ToValue(block.HayBale{Axis: axis})
		}).
		Method("honeycomb", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(block.Honeycomb{})
		}).
		Method("hopper", func(c goja.FunctionCall) goja.Value {
			facing, ok := c.Argument(0).Export().(cube.Face)
			if !ok {
				panic(r.vm.NewTypeError("argument 0 is not a cube.Face"))
			}
			powered, _ := c.Argument(1).Export().(bool)
			customName, _ := c.Argument(2).Export().(string)
			lastTick, _ := c.Argument(3).Export().(int64)
			transferCooldown, _ := c.Argument(4).Export().(int64)
			collectCooldown, _ := c.Argument(5).Export().(int64)
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
			t, ok := c.Argument(0).Export().(block.OreType)
			if !ok {
				panic(r.vm.NewTypeError("argument 0 is not a block.OreType"))
			}
			return r.vm.ToValue(block.IronOre{Type: t})
		}).
		Method("itemFrame", func(c goja.FunctionCall) goja.Value {
			facing, ok := c.Argument(0).Export().(cube.Face)
			if !ok {
				panic(r.vm.NewTypeError("argument 0 is not a cube.Face"))
			}
			i, ok := c.Argument(1).Export().(item.Stack)
			if !ok {
				panic(r.vm.NewTypeError("argument 1 is not an item.Stack"))
			}
			rotations, ok := c.Argument(2).Export().(int64)
			if !ok {
				panic(r.vm.NewTypeError("argument 2 is not an int64"))
			}
			dropChance, ok := c.Argument(3).Export().(float64)
			if !ok {
				panic(r.vm.NewTypeError("argument 3 is not a float64"))
			}
			glowing, _ := c.Argument(4).Export().(bool)
			return r.vm.ToValue(block.ItemFrame{Facing: facing, Item: i, Rotations: int(rotations), DropChance: dropChance, Glowing: glowing})
		}).
		Method("jukebox", func(c goja.FunctionCall) goja.Value {
			i, ok := c.Argument(0).Export().(item.Stack)
			if !ok {
				panic(r.vm.NewTypeError("argument 0 is not an item.Stack"))
			}
			return r.vm.ToValue(block.Jukebox{Item: i})
		}).
		Method("kelp", func(c goja.FunctionCall) goja.Value {
			age, ok := c.Argument(0).Export().(int64)
			if !ok {
				panic(r.vm.NewTypeError("argument 0 is not an int64"))
			}
			return r.vm.ToValue(block.Kelp{Age: int(age)})
		}).
		Method("ladder", func(c goja.FunctionCall) goja.Value {
			facing, ok := c.Argument(0).Export().(cube.Direction)
			if !ok {
				panic(r.vm.NewTypeError("argument 0 is not a cube.Direction"))
			}
			return r.vm.ToValue(block.Ladder{Facing: facing})
		}).
		Method("lantern", func(c goja.FunctionCall) goja.Value {
			t, ok := c.Argument(0).Export().(block.FireType)
			if !ok {
				panic(r.vm.NewTypeError("argument 0 is not a block.LanternType"))
			}
			hanging, _ := c.Argument(1).Export().(bool)
			return r.vm.ToValue(block.Lantern{Type: t, Hanging: hanging})
		}).
		Method("lapis", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(block.Lapis{})
		}).
		Method("lapisOre", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(block.LapisOre{})
		}).
		Method("lava", func(c goja.FunctionCall) goja.Value {
			depth, ok := c.Argument(0).Export().(int64)
			if !ok {
				panic(r.vm.NewTypeError("argument 0 is not an int64"))
			}
			still, _ := c.Argument(1).Export().(bool)
			falling, _ := c.Argument(2).Export().(bool)
			return r.vm.ToValue(block.Lava{Depth: int(depth), Still: still, Falling: falling})
		}).
		Method("lavaDamageSource", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(block.LavaDamageSource{})
		}).
		Method("leaves", func(c goja.FunctionCall) goja.Value {
			wood, ok := c.Argument(0).Export().(block.WoodType)
			if !ok {
				panic(r.vm.NewTypeError("argument 0 is not a block.WoodType"))
			}
			persistent, _ := c.Argument(1).Export().(bool)
			shouldUpdate, _ := c.Argument(2).Export().(bool)
			return r.vm.ToValue(block.Leaves{Wood: wood, Persistent: persistent, ShouldUpdate: shouldUpdate})
		}).
		Method("lectern", func(c goja.FunctionCall) goja.Value {
			facing, ok := c.Argument(0).Export().(cube.Direction)
			if !ok {
				panic(r.vm.NewTypeError("argument 0 is not a cube.Direction"))
			}
			i, ok := c.Argument(1).Export().(item.Stack)
			if !ok {
				panic(r.vm.NewTypeError("argument 1 is not an item.Stack"))
			}
			page, ok := c.Argument(2).Export().(int64)
			if !ok {
				panic(r.vm.NewTypeError("argument 2 is not an int64"))
			}
			return r.vm.ToValue(block.Lectern{Facing: facing, Book: i, Page: int(page)})
		}).
		Method("light", func(c goja.FunctionCall) goja.Value {
			level, ok := c.Argument(0).Export().(int64)
			if !ok {
				panic(r.vm.NewTypeError("argument 0 is not an int64"))
			}
			return r.vm.ToValue(block.Light{Level: int(level)})
		}).
		Method("litPumpkin", func(c goja.FunctionCall) goja.Value {
			facing, ok := c.Argument(0).Export().(cube.Direction)
			if !ok {
				panic(r.vm.NewTypeError("argument 0 is not a cube.Direction"))
			}
			return r.vm.ToValue(block.LitPumpkin{Facing: facing})
		}).
		Method("log", func(c goja.FunctionCall) goja.Value {
			wood, ok := c.Argument(0).Export().(block.WoodType)
			if !ok {
				panic(r.vm.NewTypeError("argument 0 is not a block.WoodType"))
			}
			axis, ok := c.Argument(1).Export().(cube.Axis)
			if !ok {
				panic(r.vm.NewTypeError("argument 1 is not a cube.Axis"))
			}
			stripped, _ := c.Argument(2).Export().(bool)
			return r.vm.ToValue(block.Log{Wood: wood, Axis: axis, Stripped: stripped})
		}).
		Method("loom", func(c goja.FunctionCall) goja.Value {
			facing, ok := c.Argument(0).Export().(cube.Direction)
			if !ok {
				panic(r.vm.NewTypeError("argument 0 is not a cube.Direction"))
			}
			return r.vm.ToValue(block.Loom{Facing: facing})
		}).
		Method("melon", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(block.Melon{})
		}).
		Method("melonSeeds", func(c goja.FunctionCall) goja.Value {
			direction, ok := c.Argument(0).Export().(cube.Face)
			if !ok {
				panic(r.vm.NewTypeError("argument 0 is not a cube.Face"))
			}
			growth, ok := c.Argument(0).Export().(int64)
			if !ok {
				panic(r.vm.NewTypeError("argument 0 is not an int64"))
			}
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
			axis, ok := c.Argument(0).Export().(cube.Axis)
			if !ok {
				panic(r.vm.NewTypeError("argument 0 is not a cube.Axis"))
			}
			return r.vm.ToValue(block.MuddyMangroveRoots{Axis: axis})
		}).
		Method("netherBrickFence", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(block.NetherBrickFence{})
		}).
		Method("netherBricks", func(c goja.FunctionCall) goja.Value {
			t, ok := c.Argument(0).Export().(block.NetherBricksType)
			if !ok {
				panic(r.vm.NewTypeError("argument 0 is not a block.NetherBricksType"))
			}
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
			age, ok := c.Argument(0).Export().(int64)
			if !ok {
				panic(r.vm.NewTypeError("argument 0 is not an int64"))
			}
			return r.vm.ToValue(block.NetherWart{Age: int(age)})
		}).
		Method("netherWartBlock", func(c goja.FunctionCall) goja.Value {
			warped, _ := c.Argument(0).Export().(bool)
			return r.vm.ToValue(block.NetherWartBlock{Warped: warped})
		}).
		Method("netherite", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(block.Netherite{})
		}).
		Method("netherrack", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(block.Netherrack{})
		}).
		Method("note", func(c goja.FunctionCall) goja.Value {
			pitch, ok := c.Argument(0).Export().(int64)
			if !ok {
				panic(r.vm.NewTypeError("argument 0 is not an int64"))
			}
			return r.vm.ToValue(block.Note{Pitch: int(pitch)})
		}).
		Method("obsidian", func(c goja.FunctionCall) goja.Value {
			crying, _ := c.Argument(0).Export().(bool)
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
			facing, ok := c.Argument(0).Export().(cube.Direction)
			if !ok {
				panic(r.vm.NewTypeError("argument 0 is not a cube.Direction"))
			}
			additional, ok := c.Argument(1).Export().(int64)
			if !ok {
				panic(r.vm.NewTypeError("argument 1 is not an int64"))
			}
			return r.vm.ToValue(block.PinkPetals{Facing: facing, AdditionalCount: int(additional)})
		}).
		Method("planks", func(c goja.FunctionCall) goja.Value {
			wood, ok := c.Argument(0).Export().(block.WoodType)
			if !ok {
				panic(r.vm.NewTypeError("argument 0 is not a block.WoodType"))
			}
			return r.vm.ToValue(block.Planks{Wood: wood})
		}).
		Method("podzol", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(block.Podzol{})
		}).
		Method("polishedBlackstoneBrick", func(c goja.FunctionCall) goja.Value {
			cracked, _ := c.Argument(0).Export().(bool)
			return r.vm.ToValue(block.PolishedBlackstoneBrick{Cracked: cracked})
		}).
		Method("polishedTuff", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(block.PolishedTuff{})
		}).
		Method("potato", func(c goja.FunctionCall) goja.Value {
			growth, ok := c.Argument(0).Export().(int64)
			if !ok {
				panic(r.vm.NewTypeError("argument 0 is not an int64"))
			}
			b := block.Potato{}
			b.Growth = int(growth)
			return r.vm.ToValue(b)
		}).
		Method("prismarine", func(c goja.FunctionCall) goja.Value {
			t, ok := c.Argument(0).Export().(block.PrismarineType)
			if !ok {
				panic(r.vm.NewTypeError("argument 0 is not a block.PrismarineType"))
			}
			return r.vm.ToValue(block.Prismarine{Type: t})
		}).
		Method("pumpkin", func(c goja.FunctionCall) goja.Value {
			facing, ok := c.Argument(0).Export().(cube.Direction)
			if !ok {
				panic(r.vm.NewTypeError("argument 0 is not a cube.Direction"))
			}
			return r.vm.ToValue(block.Pumpkin{Facing: facing})
		}).
		Method("pumpkinSeeds", func(c goja.FunctionCall) goja.Value {
			direction, ok := c.Argument(0).Export().(cube.Face)
			if !ok {
				panic(r.vm.NewTypeError("argument 0 is not a cube.Face"))
			}
			growth, ok := c.Argument(1).Export().(int64)
			if !ok {
				panic(r.vm.NewTypeError("argument 1 not an int64"))
			}
			b := block.PumpkinSeeds{Direction: direction}
			b.Growth = int(growth)
			return r.vm.ToValue(b)
		}).
		Method("purpur", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(block.Purpur{})
		}).
		Method("purpurPillar", func(c goja.FunctionCall) goja.Value {
			axis, ok := c.Argument(0).Export().(cube.Axis)
			if !ok {
				panic(r.vm.NewTypeError("argument 0 is not a cube.Axis"))
			}
			return r.vm.ToValue(block.PurpurPillar{Axis: axis})
		}).
		Method("quartz", func(c goja.FunctionCall) goja.Value {
			smooth, _ := c.Argument(0).Export().(bool)
			return r.vm.ToValue(block.Quartz{Smooth: smooth})
		}).
		Method("quartzBricks", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(block.QuartzBricks{})
		}).
		Method("quartzPillar", func(c goja.FunctionCall) goja.Value {
			axis, ok := c.Argument(0).Export().(cube.Axis)
			if !ok {
				panic(r.vm.NewTypeError("argument 0 is not a cube.Axis"))
			}
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
			chiseled, _ := c.Argument(0).Export().(bool)
			return r.vm.ToValue(block.ResinBricks{Chiseled: chiseled})
		}).
		Method("sand", func(c goja.FunctionCall) goja.Value {
			red, _ := c.Argument(0).Export().(bool)
			return r.vm.ToValue(block.Sand{Red: red})
		}).
		Method("sandstone", func(c goja.FunctionCall) goja.Value {
			t, ok := c.Argument(0).Export().(block.SandstoneType)
			if !ok {
				panic(r.vm.NewTypeError("argument 0 is not a block.SandstoneType"))
			}
			red, _ := c.Argument(1).Export().(bool)
			return r.vm.ToValue(block.Sandstone{Type: t, Red: red})
		}).
		Method("seaLantern", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(block.SeaLantern{})
		}).
		Method("seaPickle", func(c goja.FunctionCall) goja.Value {
			dead, _ := c.Argument(0).Export().(bool)
			return r.vm.ToValue(block.SeaPickle{Dead: dead})
		}).
		Method("shortGrass", func(c goja.FunctionCall) goja.Value {
			double, _ := c.Argument(0).Export().(bool)
			return r.vm.ToValue(block.ShortGrass{Double: double})
		}).
		Method("shroomlight", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(block.Shroomlight{})
		}).
		Method("sign", func(c goja.FunctionCall) goja.Value {
			wood, ok := c.Argument(0).Export().(block.WoodType)
			if !ok {
				panic(r.vm.NewTypeError("argument 0 is not a block.WoodType"))
			}
			a, ok := c.Argument(1).Export().(block.Attachment)
			if !ok {
				panic(r.vm.NewTypeError("argument 1 is not a block.Attachment"))
			}
			front, ok := c.Argument(2).Export().(block.SignText)
			if !ok {
				panic(r.vm.NewTypeError("argument 2 is not a block.SignText"))
			}
			back, ok := c.Argument(3).Export().(block.SignText)
			if !ok {
				panic(r.vm.NewTypeError("argument 3 is not a block.SignText"))
			}
			waxed, _ := c.Argument(4).Export().(bool)
			return r.vm.ToValue(block.Sign{Wood: wood, Attach: a, Front: front, Back: back, Waxed: waxed})
		}).
		Method("signText", func(c goja.FunctionCall) goja.Value {
			text, ok := c.Argument(0).Export().(string)
			if !ok {
				panic(r.vm.NewTypeError("argument 0 is not a string"))
			}
			owner, ok := c.Argument(1).Export().(string)
			if !ok {
				panic(r.vm.NewTypeError("argument 1 is not a string"))
			}
			colour, ok := c.Argument(2).Export().(color.RGBA)
			if !ok {
				panic(r.vm.NewTypeError("argument 2 is not a color.RGBA"))
			}
			glowing, _ := c.Argument(3).Export().(bool)
			return r.vm.ToValue(block.SignText{Text: text, Owner: owner, BaseColour: colour, Glowing: glowing})
		}).
		Method("skull", func(c goja.FunctionCall) goja.Value {
			t, ok := c.Argument(0).Export().(block.SkullType)
			if !ok {
				panic(r.vm.NewTypeError("argument 0 is not a block.SkullType"))
			}
			a, ok := c.Argument(1).Export().(block.Attachment)
			if !ok {
				panic(r.vm.NewTypeError("argument 1 is not a block.Attachment"))
			}
			return r.vm.ToValue(block.Skull{Type: t, Attach: a})
		}).
		Method("slab", func(c goja.FunctionCall) goja.Value {
			b, ok := c.Argument(0).Export().(world.Block)
			if !ok {
				panic(r.vm.NewTypeError("argument 0 is not a world.Block"))
			}
			double, _ := c.Argument(1).Export().(bool)
			top, _ := c.Argument(2).Export().(bool)
			return r.vm.ToValue(block.Slab{Block: b, Double: double, Top: top})
		}).
		Method("smithingTable", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(block.SmithingTable{})
		}).
		Method("smoker", func(c goja.FunctionCall) goja.Value {
			facing, ok := c.Argument(0).Export().(cube.Direction)
			if !ok {
				panic(r.vm.NewTypeError("argument 0 is not a cube.Direction"))
			}
			lit, _ := c.Argument(1).Export().(bool)
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
			wet, _ := c.Argument(0).Export().(bool)
			return r.vm.ToValue(block.Sponge{Wet: wet})
		}).
		Method("sporeBlossom", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(block.SporeBlossom{})
		}).
		Method("stainedGlass", func(c goja.FunctionCall) goja.Value {
			colour, ok := c.Argument(0).Export().(item.Colour)
			if !ok {
				panic(r.vm.NewTypeError("argument 0 is not an item.Colour"))
			}
			return r.vm.ToValue(block.StainedGlass{Colour: colour})
		}).
		Method("stainedGlassPane", func(c goja.FunctionCall) goja.Value {
			colour, ok := c.Argument(0).Export().(item.Colour)
			if !ok {
				panic(r.vm.NewTypeError("argument 0 is not an item.Colour"))
			}
			return r.vm.ToValue(block.StainedGlassPane{Colour: colour})
		}).
		Method("stainedTerracotta", func(c goja.FunctionCall) goja.Value {
			colour, ok := c.Argument(0).Export().(item.Colour)
			if !ok {
				panic(r.vm.NewTypeError("argument 0 is not an item.Colour"))
			}
			return r.vm.ToValue(block.StainedTerracotta{Colour: colour})
		}).
		Method("stairs", func(c goja.FunctionCall) goja.Value {
			b, ok := c.Argument(0).Export().(world.Block)
			if !ok {
				panic(r.vm.NewTypeError("argument 0 is not a world.Block"))
			}
			facing, ok := c.Argument(1).Export().(cube.Direction)
			if !ok {
				panic(r.vm.NewTypeError("argument 1 is not a cube.Direction"))
			}
			upsideDown, _ := c.Argument(2).Export().(bool)
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
			smooth, _ := c.Argument(0).Export().(bool)
			return r.vm.ToValue(block.Stone{Smooth: smooth})
		}).
		Method("stoneBricks", func(c goja.FunctionCall) goja.Value {
			t, ok := c.Argument(0).Export().(block.StoneBricksType)
			if !ok {
				panic(r.vm.NewTypeError("argument 0 is not a block.StoneBricksType"))
			}
			return r.vm.ToValue(block.StoneBricks{Type: t})
		}).
		Method("stonecutter", func(c goja.FunctionCall) goja.Value {
			facing, ok := c.Argument(0).Export().(cube.Direction)
			if !ok {
				panic(r.vm.NewTypeError("argument 0 is not a cube.Direction"))
			}
			return r.vm.ToValue(block.Stonecutter{Facing: facing})
		}).
		Method("stopCrackAction", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(block.StopCrackAction{})
		}).
		Method("sugarCane", func(c goja.FunctionCall) goja.Value {
			age, ok := c.Argument(0).Export().(int64)
			if !ok {
				panic(r.vm.NewTypeError("argument 0 is not an int64"))
			}
			return r.vm.ToValue(block.SugarCane{Age: int(age)})
		}).
		Method("tnt", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(block.TNT{})
		}).
		Method("terracotta", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(block.Terracotta{})
		}).
		Method("torch", func(c goja.FunctionCall) goja.Value {
			facing, ok := c.Argument(0).Export().(cube.Face)
			if !ok {
				panic(r.vm.NewTypeError("argument 0 is not a cube.Face"))
			}
			t, ok := c.Argument(1).Export().(block.FireType)
			if !ok {
				panic(r.vm.NewTypeError("argument 1 is not a block.FireType"))
			}
			return r.vm.ToValue(block.Torch{Facing: facing, Type: t})
		}).
		Method("tuff", func(c goja.FunctionCall) goja.Value {
			chiseled, _ := c.Argument(0).Export().(bool)
			return r.vm.ToValue(block.Tuff{Chiseled: chiseled})
		}).
		Method("tuffBricks", func(c goja.FunctionCall) goja.Value {
			chiseled, _ := c.Argument(0).Export().(bool)
			return r.vm.ToValue(block.TuffBricks{Chiseled: chiseled})
		}).
		Method("vines", func(c goja.FunctionCall) goja.Value {
			north, _ := c.Argument(0).Export().(bool)
			east, _ := c.Argument(1).Export().(bool)
			south, _ := c.Argument(2).Export().(bool)
			west, _ := c.Argument(3).Export().(bool)
			return r.vm.ToValue(block.Vines{NorthDirection: north, EastDirection: east, SouthDirection: south, WestDirection: west})
		}).
		Method("wall", func(c goja.FunctionCall) goja.Value {
			b, ok := c.Argument(0).Export().(world.Block)
			if !ok {
				panic(r.vm.NewTypeError("argument 0 is not a world.Block"))
			}
			north, ok := c.Argument(1).Export().(block.WallConnectionType)
			if !ok {
				panic(r.vm.NewTypeError("argument 1 is not a block.WallConnectionType"))
			}
			east, ok := c.Argument(2).Export().(block.WallConnectionType)
			if !ok {
				panic(r.vm.NewTypeError("argument 2 is not a block.WallConnectionType"))
			}
			south, ok := c.Argument(3).Export().(block.WallConnectionType)
			if !ok {
				panic(r.vm.NewTypeError("argument 3 is not a block.WallConnectionType"))
			}
			west, ok := c.Argument(4).Export().(block.WallConnectionType)
			if !ok {
				panic(r.vm.NewTypeError("argument 4 is not a block.WallConnectionType"))
			}
			post, _ := c.Argument(5).Export().(bool)
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
			depth, ok := c.Argument(0).Export().(int64)
			if !ok {
				panic(r.vm.NewTypeError("argument 0 is not an int64"))
			}
			still, _ := c.Argument(1).Export().(bool)
			falling, _ := c.Argument(2).Export().(bool)
			return r.vm.ToValue(block.Water{Depth: int(depth), Still: still, Falling: falling})
		}).
		Method("wheatSeeds", func(c goja.FunctionCall) goja.Value {
			growth, ok := c.Argument(0).Export().(int64)
			if !ok {
				panic(r.vm.NewTypeError("argument 0 is not an int64"))
			}
			b := block.WheatSeeds{}
			b.Growth = int(growth)
			return r.vm.ToValue(b)
		}).
		Method("wood", func(c goja.FunctionCall) goja.Value {
			wood, ok := c.Argument(0).Export().(block.WoodType)
			if !ok {
				panic(r.vm.NewTypeError("argument 0 is not a block.WoodType"))
			}
			axis, ok := c.Argument(1).Export().(cube.Axis)
			if !ok {
				panic(r.vm.NewTypeError("argument 1 is not a cube.Axis"))
			}
			stripped, _ := c.Argument(2).Export().(bool)
			return r.vm.ToValue(block.Wood{Wood: wood, Axis: axis, Stripped: stripped})
		}).
		Method("woodDoor", func(c goja.FunctionCall) goja.Value {
			wood, ok := c.Argument(0).Export().(block.WoodType)
			if !ok {
				panic(r.vm.NewTypeError("argument 0 is not a block.WoodType"))
			}
			facing, ok := c.Argument(1).Export().(cube.Direction)
			if !ok {
				panic(r.vm.NewTypeError("argument 1 is not a cube.Direction"))
			}
			open, _ := c.Argument(2).Export().(bool)
			top, _ := c.Argument(3).Export().(bool)
			right, _ := c.Argument(4).Export().(bool)
			return r.vm.ToValue(block.WoodDoor{Wood: wood, Facing: facing, Open: open, Top: top, Right: right})
		}).
		Method("woodFence", func(c goja.FunctionCall) goja.Value {
			wood, ok := c.Argument(0).Export().(block.WoodType)
			if !ok {
				panic(r.vm.NewTypeError("argument 0 is not a block.WoodType"))
			}
			return r.vm.ToValue(block.WoodFence{Wood: wood})
		}).
		Method("woodFenceGate", func(c goja.FunctionCall) goja.Value {
			wood, ok := c.Argument(0).Export().(block.WoodType)
			if !ok {
				panic(r.vm.NewTypeError("argument 0 is not a block.WoodType"))
			}
			facing, ok := c.Argument(1).Export().(cube.Direction)
			if !ok {
				panic(r.vm.NewTypeError("argument 1 is not a cube.Direction"))
			}
			open, _ := c.Argument(2).Export().(bool)
			lowered, _ := c.Argument(3).Export().(bool)
			return r.vm.ToValue(block.WoodFenceGate{Wood: wood, Facing: facing, Open: open, Lowered: lowered})
		}).
		Method("woodTrapdoor", func(c goja.FunctionCall) goja.Value {
			wood, ok := c.Argument(0).Export().(block.WoodType)
			if !ok {
				panic(r.vm.NewTypeError("argument 0 is not a block.WoodType"))
			}
			facing, ok := c.Argument(1).Export().(cube.Direction)
			if !ok {
				panic(r.vm.NewTypeError("argument 1 is not a cube.Direction"))
			}
			open, _ := c.Argument(2).Export().(bool)
			top, _ := c.Argument(3).Export().(bool)
			return r.vm.ToValue(block.WoodTrapdoor{Wood: wood, Facing: facing, Open: open, Top: top})
		}).
		Method("wool", func(c goja.FunctionCall) goja.Value {
			colour, ok := c.Argument(0).Export().(item.Colour)
			if !ok {
				panic(r.vm.NewTypeError("argument 0 is not an item.Colour"))
			}
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
