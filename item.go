package dfscript

import (
	"github.com/df-mc/dragonfly/server/block"
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/potion"
	"github.com/df-mc/dragonfly/server/world"
	"github.com/df-mc/dragonfly/server/world/sound"
	"github.com/dop251/goja"
	"image/color"
	"time"
)

func (r *Runtime) setupItem() error {
	return newObject().
		Const("toolType", map[string]item.ToolType{
			"none":    item.TypeNone,
			"pickaxe": item.TypePickaxe,
			"axe":     item.TypeAxe,
			"hoe":     item.TypeHoe,
			"shovel":  item.TypeShovel,
			"shears":  item.TypeShears,
		}).
		Const("toolTier", map[string]item.ToolTier{
			"wood":      item.ToolTierWood,
			"gold":      item.ToolTierGold,
			"stone":     item.ToolTierStone,
			"iron":      item.ToolTierIron,
			"diamond":   item.ToolTierDiamond,
			"netherite": item.ToolTierNetherite,
		}).
		Const("writtenBookGeneration", map[string]item.WrittenBookGeneration{
			"original":   item.OriginalGeneration(),
			"copy":       item.CopyGeneration(),
			"copyOfCopy": item.CopyOfCopyGeneration(),
		}).
		Const("stewType", map[string]item.StewType{
			"nightVisionPoppy":       item.NightVisionPoppyStew(),
			"jumpBoost":              item.JumpBoostStew(),
			"weakness":               item.WeaknessStew(),
			"blindnessBluet":         item.BlindnessBluetStew(),
			"poison":                 item.PoisonStew(),
			"saturationDandelion":    item.SaturationDandelionStew(),
			"saturationOrchid":       item.SaturationOrchidStew(),
			"fireResistance":         item.FireResistanceStew(),
			"regeneration":           item.RegenerationStew(),
			"wither":                 item.WitherStew(),
			"nightVisionTorchflower": item.NightVisionTorchflowerStew(),
			"blindnessEyeBlossom":    item.BlindnessEyeblossomStew(),
			"nauseaBlossom":          item.NauseaStew(),
		}).
		Const("smithingTemplateType", map[string]item.SmithingTemplateType{
			"netheriteUpgrade": item.TemplateNetheriteUpgrade(),
			"sentry":           item.TemplateSentry(),
			"vex":              item.TemplateVex(),
			"wild":             item.TemplateWild(),
			"coast":            item.TemplateCoast(),
			"dune":             item.TemplateDune(),
			"wayFinder":        item.TemplateWayFinder(),
			"raiser":           item.TemplateRaiser(),
			"shaper":           item.TemplateShaper(),
			"host":             item.TemplateHost(),
			"ward":             item.TemplateWard(),
			"silence":          item.TemplateSilence(),
			"tide":             item.TemplateTide(),
			"snout":            item.TemplateSnout(),
			"rib":              item.TemplateRib(),
			"eye":              item.TemplateEye(),
			"spire":            item.TemplateSpire(),
			"flow":             item.TemplateFlow(),
			"bolt":             item.TemplateBolt(),
		}).
		Const("sherdType", map[string]item.SherdType{
			"angler":     item.SherdTypeAngler(),
			"archer":     item.SherdTypeArcher(),
			"armsUp":     item.SherdTypeArmsUp(),
			"blade":      item.SherdTypeBlade(),
			"brewer":     item.SherdTypeBrewer(),
			"burn":       item.SherdTypeBurn(),
			"danger":     item.SherdTypeDanger(),
			"explorer":   item.SherdTypeExplorer(),
			"friend":     item.SherdTypeFriend(),
			"heart":      item.SherdTypeHeart(),
			"heartbreak": item.SherdTypeHeartbreak(),
			"howl":       item.SherdTypeHowl(),
			"miner":      item.SherdTypeMiner(),
			"mourner":    item.SherdTypeMourner(),
			"plenty":     item.SherdTypePlenty(),
			"prize":      item.SherdTypePrize(),
			"sheaf":      item.SherdTypeSheaf(),
			"shelter":    item.SherdTypeShelter(),
			"skull":      item.SherdTypeSkull(),
			"snort":      item.SherdTypeSnort(),
			"flow":       item.SherdTypeFlow(),
			"guster":     item.SherdTypeGuster(),
			"scrape":     item.SherdTypeScrape(),
		}).
		Const("colour", map[string]item.Colour{
			"white":     item.ColourWhite(),
			"orange":    item.ColourOrange(),
			"magenta":   item.ColourMagenta(),
			"lightBlue": item.ColourLightBlue(),
			"yellow":    item.ColourYellow(),
			"lime":      item.ColourLime(),
			"pink":      item.ColourPink(),
			"grey":      item.ColourGrey(),
			"lightGrey": item.ColourLightGrey(),
			"cyan":      item.ColourCyan(),
			"purple":    item.ColourPurple(),
			"blue":      item.ColourBlue(),
			"brown":     item.ColourBrown(),
			"green":     item.ColourGreen(),
			"red":       item.ColourRed(),
			"black":     item.ColourBlack(),
		}).
		Const("bucketContents", map[string]item.BucketContent{
			"milk":  item.MilkBucketContent(),
			"water": item.LiquidBucketContent(block.Water{}),
			"lava":  item.LiquidBucketContent(block.Lava{}),
		}).
		Const("bannerPatternType", map[string]item.BannerPatternType{
			"creeper":         item.CreeperBannerPattern(),
			"skull":           item.SkullBannerPattern(),
			"flower":          item.FlowerBannerPattern(),
			"mojang":          item.MojangBannerPattern(),
			"fieldMasoned":    item.FieldMasonedBannerPattern(),
			"bordureIndented": item.BordureIndentedBannerPattern(),
			"piglin":          item.PiglinBannerPattern(),
			"globe":           item.GlobeBannerPattern(),
			"flow":            item.FlowBannerPattern(),
			"guster":          item.GusterBannerPattern(),
		}).
		Const("fireworkShape", map[string]item.FireworkShape{
			"smallSphere": item.FireworkShapeSmallSphere(),
			"hugeSphere":  item.FireworkShapeHugeSphere(),
			"star":        item.FireworkShapeStar(),
			"creeperHead": item.FireworkShapeCreeperHead(),
			"burst":       item.FireworkShapeBurst(),
		}).
		Method("amethystShard", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(item.AmethystShard{})
		}).
		Method("apple", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(item.Apple{})
		}).
		Method("armourTierChain", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(item.ArmourTierChain{})
		}).
		Method("armourTierDiamond", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(item.ArmourTierDiamond{})
		}).
		Method("armourTierGold", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(item.ArmourTierGold{})
		}).
		Method("armourTierIron", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(item.ArmourTierIron{})
		}).
		Method("armourTierLeather", func(c goja.FunctionCall) goja.Value {
			colour, ok := c.Argument(0).Export().(color.RGBA)
			if !ok {
				panic(r.vm.NewTypeError("argument 0 is not a color.RGBA, got %T", c.Argument(0).Export()))
			}
			return r.vm.ToValue(item.ArmourTierLeather{Colour: colour})
		}).
		Method("armourTierNetherite", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(item.ArmourTierNetherite{})
		}).
		Method("armourTrim", func(c goja.FunctionCall) goja.Value {
			t, ok := c.Argument(0).Export().(item.SmithingTemplateType)
			if !ok {
				panic(r.vm.NewTypeError("argument 0 is not an item.SmithingTemplateType, got %T", c.Argument(0).Export()))
			}
			m, ok := c.Argument(1).Export().(item.ArmourTrimMaterial)
			if !ok {
				panic(r.vm.NewTypeError("argument 1 is not an item.ArmourTrimMaterial, got %T", c.Argument(1).Export()))
			}
			return r.vm.ToValue(item.ArmourTrim{Template: t, Material: m})
		}).
		Method("arrow", func(c goja.FunctionCall) goja.Value {
			tip, _ := c.Argument(0).Export().(potion.Potion)
			return r.vm.ToValue(item.Arrow{Tip: tip})
		}).
		Method("axe", func(c goja.FunctionCall) goja.Value {
			tier, ok := c.Argument(0).Export().(item.ToolTier)
			if !ok {
				panic(r.vm.NewTypeError("argument 0 is not a ToolTier, got %T", c.Argument(0).Export()))
			}
			return r.vm.ToValue(item.Axe{Tier: tier})
		}).
		Method("bakedPotato", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(item.BakedPotato{})
		}).
		Method("bannerPattern", func(c goja.FunctionCall) goja.Value {
			t, ok := c.Argument(0).Export().(item.BannerPatternType)
			if !ok {
				panic(r.vm.NewTypeError("argument 0 is not an item.BannerPatternType, got %T", c.Argument(0).Export()))
			}
			return r.vm.ToValue(item.BannerPattern{Type: t})
		}).
		Method("beef", func(c goja.FunctionCall) goja.Value {
			cooked, _ := c.Argument(0).Export().(bool)
			return r.vm.ToValue(item.Beef{Cooked: cooked})
		}).
		Method("beetroot", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(item.Beetroot{})
		}).
		Method("beetrootSoup", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(item.BeetrootSoup{})
		}).
		Method("blazePowder", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(item.BlazePowder{})
		}).
		Method("blazeRod", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(item.BlazeRod{})
		}).
		Method("bone", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(item.Bone{})
		}).
		Method("boneMeal", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(item.BoneMeal{})
		}).
		Method("book", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(item.Book{})
		}).
		Method("bookAndQuill", func(c goja.FunctionCall) goja.Value {
			pages, ok := c.Argument(0).Export().([]string)
			if !ok {
				panic(r.vm.NewTypeError("argument 0 is not a []string, got %T", c.Argument(0).Export()))
			}
			return r.vm.ToValue(item.BookAndQuill{Pages: pages})
		}).
		Method("boots", func(c goja.FunctionCall) goja.Value {
			tier, ok := c.Argument(0).Export().(item.ArmourTier)
			if !ok {
				panic(r.vm.NewTypeError("argument 0 is not an ArmourTier, got %T", c.Argument(0).Export()))
			}
			trim, _ := c.Argument(1).Export().(item.ArmourTrim)
			return r.vm.ToValue(item.Boots{Tier: tier, Trim: trim})
		}).
		Method("bottleOfEnchanting", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(item.BottleOfEnchanting{})
		}).
		Method("bow", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(item.Bow{})
		}).
		Method("bowl", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(item.Bowl{})
		}).
		Method("bread", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(item.Bread{})
		}).
		Method("brick", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(item.Brick{})
		}).
		Method("bucket", func(c goja.FunctionCall) goja.Value {
			cont, ok := c.Argument(0).Export().(item.BucketContent)
			if !ok {
				panic(r.vm.NewTypeError("argument 0 is not an item.BucketContent, got %T", c.Argument(0).Export()))
			}
			return r.vm.ToValue(item.Bucket{Content: cont})
		}).
		Method("carrotOnAStick", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(item.CarrotOnAStick{})
		}).
		Method("charcoal", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(item.Charcoal{})
		}).
		Method("chestplate", func(c goja.FunctionCall) goja.Value {
			tier, ok := c.Argument(0).Export().(item.ArmourTier)
			if !ok {
				panic(r.vm.NewTypeError("argument 0 is not an ArmourTier, got %T", c.Argument(0).Export()))
			}
			trim, _ := c.Argument(1).Export().(item.ArmourTrim)
			return r.vm.ToValue(item.Chestplate{Tier: tier, Trim: trim})
		}).
		Method("chicken", func(c goja.FunctionCall) goja.Value {
			cooked, _ := c.Argument(0).Export().(bool)
			return r.vm.ToValue(item.Chicken{Cooked: cooked})
		}).
		Method("clayBall", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(item.ClayBall{})
		}).
		Method("clock", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(item.Clock{})
		}).
		Method("coal", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(item.Coal{})
		}).
		Method("cod", func(c goja.FunctionCall) goja.Value {
			cooked, _ := c.Argument(0).Export().(bool)
			return r.vm.ToValue(item.Cod{Cooked: cooked})
		}).
		Method("compass", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(item.Compass{})
		}).
		Method("cookie", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(item.Cookie{})
		}).
		Method("copperIngot", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(item.CopperIngot{})
		}).
		Method("crossbow", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(item.Crossbow{})
		}).
		Method("diamond", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(item.Diamond{})
		}).
		Method("discFragment", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(item.DiscFragment{})
		}).
		Method("dragonBreath", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(item.DragonBreath{})
		}).
		Method("driedKelp", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(item.DriedKelp{})
		}).
		Method("dye", func(c goja.FunctionCall) goja.Value {
			col, ok := c.Argument(0).Export().(item.Colour)
			if !ok {
				panic(r.vm.NewTypeError("argument 0 is not an item.Colour, got %T", c.Argument(0).Export()))
			}
			return r.vm.ToValue(item.Dye{Colour: col})
		}).
		Method("echoShard", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(item.EchoShard{})
		}).
		Method("egg", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(item.Egg{})
		}).
		Method("elytra", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(item.Elytra{})
		}).
		Method("emerald", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(item.Emerald{})
		}).
		Method("enchantedApple", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(item.EnchantedApple{})
		}).
		Method("enchantedBook", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(item.EnchantedBook{})
		}).
		Method("enderPearl", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(item.EnderPearl{})
		}).
		Method("feather", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(item.Feather{})
		}).
		Method("fermentedSpiderEye", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(item.FermentedSpiderEye{})
		}).
		Method("fireCharge", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(item.FireCharge{})
		}).
		Method("firework", func(c goja.FunctionCall) goja.Value {
			d, ok := c.Argument(0).Export().(time.Duration)
			if !ok {
				panic(r.vm.NewTypeError("argument 0 is not a time.Duration, got %T", c.Argument(0).Export()))
			}
			explosions, ok := c.Argument(1).Export().([]item.FireworkExplosion)
			if !ok {
				panic(r.vm.NewTypeError("argument 1 is not a []item.FireworkExplosion, got %T", c.Argument(1).Export()))
			}
			return r.vm.ToValue(item.Firework{Duration: d, Explosions: explosions})
		}).
		Method("fireworkExplosion", func(c goja.FunctionCall) goja.Value {
			shape, ok := c.Argument(0).Export().(item.FireworkShape)
			if !ok {
				panic(r.vm.NewTypeError("argument 0 is not a FireworkShape, got %T", c.Argument(0).Export()))
			}
			colour, ok := c.Argument(1).Export().(item.Colour)
			if !ok {
				panic(r.vm.NewTypeError("argument 1 is not a Colour, got %T", c.Argument(1).Export()))
			}
			fade, ok := c.Argument(2).Export().(item.Colour)
			if !ok {
				panic(r.vm.NewTypeError("argument 2 is not a Colour, got %T", c.Argument(2).Export()))
			}
			fades, _ := c.Argument(3).Export().(bool)
			twinkle, _ := c.Argument(4).Export().(bool)
			trail, _ := c.Argument(5).Export().(bool)
			return r.vm.ToValue(item.FireworkExplosion{
				Shape:   shape,
				Colour:  colour,
				Fade:    fade,
				Fades:   fades,
				Twinkle: twinkle,
				Trail:   trail,
			})
		}).
		Method("fireworkStar", func(c goja.FunctionCall) goja.Value {
			explosion, ok := c.Argument(0).Export().(item.FireworkExplosion)
			if !ok {
				panic(r.vm.NewTypeError("argument 0 is not a FireworkExplosion, got %T", c.Argument(0).Export()))
			}
			return r.vm.ToValue(item.FireworkStar{FireworkExplosion: explosion})
		}).
		Method("flint", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(item.Flint{})
		}).
		Method("flintAndSteel", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(item.FlintAndSteel{})
		}).
		Method("ghastTear", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(item.GhastTear{})
		}).
		Method("glassBottle", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(item.GlassBottle{})
		}).
		Method("glisteringMelonSlice", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(item.GlisteringMelonSlice{})
		}).
		Method("glowstoneDust", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(item.GlowstoneDust{})
		}).
		Method("goatHorn", func(c goja.FunctionCall) goja.Value {
			t, ok := c.Argument(0).Export().(sound.Horn)
			if !ok {
				panic(r.vm.NewTypeError("argument 0 is not a sound.Horn, got %T", c.Argument(0).Export()))
			}
			return r.vm.ToValue(item.GoatHorn{Type: t})
		}).
		Method("goldIngot", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(item.GoldIngot{})
		}).
		Method("goldNugget", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(item.GoldNugget{})
		}).
		Method("goldenApple", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(item.GoldenApple{})
		}).
		Method("goldenCarrot", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(item.GoldenCarrot{})
		}).
		Method("gunpowder", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(item.Gunpowder{})
		}).
		Method("heartOfTheSea", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(item.HeartOfTheSea{})
		}).
		Method("helmet", func(c goja.FunctionCall) goja.Value {
			tier, ok := c.Argument(0).Export().(item.ArmourTier)
			if !ok {
				panic(r.vm.NewTypeError("argument 0 is not an ArmourTier, got %T", c.Argument(0).Export()))
			}
			trim, _ := c.Argument(1).Export().(item.ArmourTrim)
			return r.vm.ToValue(item.Helmet{Tier: tier, Trim: trim})
		}).
		Method("hoe", func(c goja.FunctionCall) goja.Value {
			tier, ok := c.Argument(0).Export().(item.ToolTier)
			if !ok {
				panic(r.vm.NewTypeError("argument 0 is not a ToolTier, got %T", c.Argument(0).Export()))
			}
			return r.vm.ToValue(item.Hoe{Tier: tier})
		}).
		Method("honeycomb", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(item.Honeycomb{})
		}).
		Method("inkSac", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(item.InkSac{})
		}).
		Method("ironIngot", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(item.IronIngot{})
		}).
		Method("ironNugget", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(item.IronNugget{})
		}).
		Method("lapisLazuli", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(item.LapisLazuli{})
		}).
		Method("leather", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(item.Leather{})
		}).
		Method("leggings", func(c goja.FunctionCall) goja.Value {
			tier, ok := c.Argument(0).Export().(item.ArmourTier)
			if !ok {
				panic(r.vm.NewTypeError("argument 0 is not an ArmourTier, got %T", c.Argument(0).Export()))
			}
			trim, _ := c.Argument(1).Export().(item.ArmourTrim)
			return r.vm.ToValue(item.Leggings{Tier: tier, Trim: trim})
		}).
		Method("lingeringPotion", func(c goja.FunctionCall) goja.Value {
			t, ok := c.Argument(0).Export().(potion.Potion)
			if !ok {
				panic(r.vm.NewTypeError("argument 0 is not a Potion, got %T", c.Argument(0).Export()))
			}
			return r.vm.ToValue(item.LingeringPotion{Type: t})
		}).
		Method("magmaCream", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(item.MagmaCream{})
		}).
		Method("melonSlice", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(item.MelonSlice{})
		}).
		Method("mushroomStew", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(item.MushroomStew{})
		}).
		Method("musicDisc", func(c goja.FunctionCall) goja.Value {
			t, ok := c.Argument(0).Export().(sound.DiscType)
			if !ok {
				panic(r.vm.NewTypeError("argument 0 is not a DiscType, got %T", c.Argument(0).Export()))
			}
			return r.vm.ToValue(item.MusicDisc{DiscType: t})
		}).
		Method("mutton", func(c goja.FunctionCall) goja.Value {
			cooked, _ := c.Argument(0).Export().(bool)
			return r.vm.ToValue(item.Mutton{Cooked: cooked})
		}).
		Method("nautilusShell", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(item.NautilusShell{})
		}).
		Method("netherBrick", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(item.NetherBrick{})
		}).
		Method("netherQuartz", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(item.NetherQuartz{})
		}).
		Method("netherStar", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(item.NetherStar{})
		}).
		Method("netheriteIngot", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(item.NetheriteIngot{})
		}).
		Method("netheriteScrap", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(item.NetheriteScrap{})
		}).
		Method("paper", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(item.Paper{})
		}).
		Method("phantomMembrane", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(item.PhantomMembrane{})
		}).
		Method("pickaxe", func(c goja.FunctionCall) goja.Value {
			tier, ok := c.Argument(0).Export().(item.ToolTier)
			if !ok {
				panic(r.vm.NewTypeError("argument 0 is not a ToolTier, got %T", c.Argument(0).Export()))
			}
			return r.vm.ToValue(item.Pickaxe{Tier: tier})
		}).
		Method("poisonousPotato", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(item.PoisonousPotato{})
		}).
		Method("poppedChorusFruit", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(item.PoppedChorusFruit{})
		}).
		Method("porkchop", func(c goja.FunctionCall) goja.Value {
			cooked, _ := c.Argument(0).Export().(bool)
			return r.vm.ToValue(item.Porkchop{Cooked: cooked})
		}).
		Method("potion", func(c goja.FunctionCall) goja.Value {
			t, ok := c.Argument(0).Export().(potion.Potion)
			if !ok {
				panic(r.vm.NewTypeError("argument 0 is not a Potion, got %T", c.Argument(0).Export()))
			}
			return r.vm.ToValue(item.Potion{Type: t})
		}).
		Method("potterySherd", func(c goja.FunctionCall) goja.Value {
			t, ok := c.Argument(0).Export().(item.SherdType)
			if !ok {
				panic(r.vm.NewTypeError("argument 0 is not a SherdType, got %T", c.Argument(0).Export()))
			}
			return r.vm.ToValue(item.PotterySherd{Type: t})
		}).
		Method("prismarineCrystals", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(item.PrismarineCrystals{})
		}).
		Method("prismarineShard", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(item.PrismarineShard{})
		}).
		Method("pufferfish", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(item.Pufferfish{})
		}).
		Method("pumpkinPie", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(item.PumpkinPie{})
		}).
		Method("rabbit", func(c goja.FunctionCall) goja.Value {
			cooked, _ := c.Argument(0).Export().(bool)
			return r.vm.ToValue(item.Rabbit{Cooked: cooked})
		}).
		Method("rabbitFoot", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(item.RabbitFoot{})
		}).
		Method("rabbitHide", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(item.RabbitHide{})
		}).
		Method("rabbitStew", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(item.RabbitStew{})
		}).
		Method("rawCopper", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(item.RawCopper{})
		}).
		Method("rawGold", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(item.RawGold{})
		}).
		Method("rawIron", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(item.RawIron{})
		}).
		Method("recoveryCompass", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(item.RecoveryCompass{})
		}).
		Method("resinBrick", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(item.ResinBrick{})
		}).
		Method("rottenFlesh", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(item.RottenFlesh{})
		}).
		Method("salmon", func(c goja.FunctionCall) goja.Value {
			cooked, _ := c.Argument(0).Export().(bool)
			return r.vm.ToValue(item.Salmon{Cooked: cooked})
		}).
		Method("scute", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(item.Scute{})
		}).
		Method("shears", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(item.Shears{})
		}).
		Method("shovel", func(c goja.FunctionCall) goja.Value {
			tier, ok := c.Argument(0).Export().(item.ToolTier)
			if !ok {
				panic(r.vm.NewTypeError("argument 0 is not a ToolTier, got %T", c.Argument(0).Export()))
			}
			return r.vm.ToValue(item.Shovel{Tier: tier})
		}).
		Method("shulkerShell", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(item.ShulkerShell{})
		}).
		Method("slimeball", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(item.Slimeball{})
		}).
		Method("smithingTemplate", func(c goja.FunctionCall) goja.Value {
			t, ok := c.Argument(0).Export().(item.SmithingTemplateType)
			if !ok {
				panic(r.vm.NewTypeError("argument 0 is not a SmithingTemplateType, got %T", c.Argument(0).Export()))
			}
			return r.vm.ToValue(item.SmithingTemplate{Template: t})
		}).
		Method("snowball", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(item.Snowball{})
		}).
		Method("spiderEye", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(item.SpiderEye{})
		}).
		Method("splashPotion", func(c goja.FunctionCall) goja.Value {
			t, ok := c.Argument(0).Export().(potion.Potion)
			if !ok {
				panic(r.vm.NewTypeError("argument 0 is not a Potion, got %T", c.Argument(0).Export()))
			}
			return r.vm.ToValue(item.SplashPotion{Type: t})
		}).
		Method("spyglass", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(item.Spyglass{})
		}).
		Method("stick", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(item.Stick{})
		}).
		Method("sugar", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(item.Sugar{})
		}).
		Method("suspiciousStew", func(c goja.FunctionCall) goja.Value {
			t, ok := c.Argument(0).Export().(item.StewType)
			if !ok {
				panic(r.vm.NewTypeError("argument 0 is not a StewType, got %T", c.Argument(0).Export()))
			}
			return r.vm.ToValue(item.SuspiciousStew{Type: t})
		}).
		Method("sword", func(c goja.FunctionCall) goja.Value {
			tier, ok := c.Argument(0).Export().(item.ToolTier)
			if !ok {
				panic(r.vm.NewTypeError("argument 0 is not a ToolTier, got %T", c.Argument(0).Export()))
			}
			return r.vm.ToValue(item.Sword{Tier: tier})
		}).
		Method("toolNone", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(item.ToolNone{})
		}).
		Method("totem", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(item.Totem{})
		}).
		Method("tropicalFish", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(item.TropicalFish{})
		}).
		Method("turtleShell", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(item.TurtleShell{})
		}).
		Method("warpedFungusOnAStick", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(item.WarpedFungusOnAStick{})
		}).
		Method("wheat", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(item.Wheat{})
		}).
		Method("writtenBook", func(c goja.FunctionCall) goja.Value {
			title, ok := c.Argument(0).Export().(string)
			if !ok {
				panic(r.vm.NewTypeError("argument 0 is not a string, got %T", c.Argument(0).Export()))
			}
			author, ok := c.Argument(1).Export().(string)
			if !ok {
				panic(r.vm.NewTypeError("argument 1 is not a string, got %T", c.Argument(1).Export()))
			}
			gen, ok := c.Argument(2).Export().(item.WrittenBookGeneration)
			if !ok {
				panic(r.vm.NewTypeError("argument 2 is not a WrittenBookGeneration, got %T", c.Argument(0).Export()))
			}
			pages, ok := c.Argument(3).Export().([]string)
			if !ok {
				panic(r.vm.NewTypeError("argument 3 is not a []string, got %T", c.Argument(3).Export()))
			}
			return r.vm.ToValue(item.WrittenBook{Title: title, Author: author, Generation: gen, Pages: pages})
		}).
		Method("stack", func(c goja.FunctionCall) goja.Value {
			it, ok := c.Argument(0).Export().(world.Item)
			if !ok {
				panic(r.vm.NewTypeError("argument 0 is not an item"))
			}
			count, _ := c.Argument(1).Export().(int64)
			if count == 0 {
				count = 1
			}
			return r.vm.ToValue(item.NewStack(it, int(count)))
		}).
		Apply(r.vm, "item")
}
