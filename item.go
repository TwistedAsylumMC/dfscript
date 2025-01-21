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
		Const("ToolType", map[string]item.ToolType{
			"None":    item.TypeNone,
			"Pickaxe": item.TypePickaxe,
			"Axe":     item.TypeAxe,
			"Hoe":     item.TypeHoe,
			"Shovel":  item.TypeShovel,
			"Shears":  item.TypeShears,
		}).
		Const("ToolTier", map[string]item.ToolTier{
			"Wood":      item.ToolTierWood,
			"Gold":      item.ToolTierGold,
			"Stone":     item.ToolTierStone,
			"Iron":      item.ToolTierIron,
			"Diamond":   item.ToolTierDiamond,
			"Netherite": item.ToolTierNetherite,
		}).
		Const("WrittenBookGeneration", map[string]item.WrittenBookGeneration{
			"Original":   item.OriginalGeneration(),
			"Copy":       item.CopyGeneration(),
			"CopyOfCopy": item.CopyOfCopyGeneration(),
		}).
		Const("StewType", map[string]item.StewType{
			"NightVisionPoppy":       item.NightVisionPoppyStew(),
			"JumpBoost":              item.JumpBoostStew(),
			"Weakness":               item.WeaknessStew(),
			"BlindnessBluet":         item.BlindnessBluetStew(),
			"Poison":                 item.PoisonStew(),
			"SaturationDandelion":    item.SaturationDandelionStew(),
			"SaturationOrchid":       item.SaturationOrchidStew(),
			"FireResistance":         item.FireResistanceStew(),
			"Regeneration":           item.RegenerationStew(),
			"Wither":                 item.WitherStew(),
			"NightVisionTorchflower": item.NightVisionTorchflowerStew(),
			"BlindnessEyeBlossom":    item.BlindnessEyeblossomStew(),
			"NauseaBlossom":          item.NauseaStew(),
		}).
		Const("SmithingTemplateType", map[string]item.SmithingTemplateType{
			"NetheriteUpgrade": item.TemplateNetheriteUpgrade(),
			"Sentry":           item.TemplateSentry(),
			"Vex":              item.TemplateVex(),
			"Wild":             item.TemplateWild(),
			"Coast":            item.TemplateCoast(),
			"Dune":             item.TemplateDune(),
			"WayFinder":        item.TemplateWayFinder(),
			"Raiser":           item.TemplateRaiser(),
			"Shaper":           item.TemplateShaper(),
			"Host":             item.TemplateHost(),
			"Ward":             item.TemplateWard(),
			"Silence":          item.TemplateSilence(),
			"Tide":             item.TemplateTide(),
			"Snout":            item.TemplateSnout(),
			"Rib":              item.TemplateRib(),
			"Eye":              item.TemplateEye(),
			"Spire":            item.TemplateSpire(),
			"Flow":             item.TemplateFlow(),
			"Bolt":             item.TemplateBolt(),
		}).
		Const("SherdType", map[string]item.SherdType{
			"Angler":     item.SherdTypeAngler(),
			"Archer":     item.SherdTypeArcher(),
			"ArmsUp":     item.SherdTypeArmsUp(),
			"Blade":      item.SherdTypeBlade(),
			"Brewer":     item.SherdTypeBrewer(),
			"Burn":       item.SherdTypeBurn(),
			"Danger":     item.SherdTypeDanger(),
			"Explorer":   item.SherdTypeExplorer(),
			"Friend":     item.SherdTypeFriend(),
			"Heart":      item.SherdTypeHeart(),
			"Heartbreak": item.SherdTypeHeartbreak(),
			"Howl":       item.SherdTypeHowl(),
			"Miner":      item.SherdTypeMiner(),
			"Mourner":    item.SherdTypeMourner(),
			"Plenty":     item.SherdTypePlenty(),
			"Prize":      item.SherdTypePrize(),
			"Sheaf":      item.SherdTypeSheaf(),
			"Shelter":    item.SherdTypeShelter(),
			"Skull":      item.SherdTypeSkull(),
			"Snort":      item.SherdTypeSnort(),
			"Flow":       item.SherdTypeFlow(),
			"Guster":     item.SherdTypeGuster(),
			"Scrape":     item.SherdTypeScrape(),
		}).
		Const("Colour", map[string]item.Colour{
			"White":     item.ColourWhite(),
			"Orange":    item.ColourOrange(),
			"Magenta":   item.ColourMagenta(),
			"LightBlue": item.ColourLightBlue(),
			"Yellow":    item.ColourYellow(),
			"Lime":      item.ColourLime(),
			"Pink":      item.ColourPink(),
			"Grey":      item.ColourGrey(),
			"LightGrey": item.ColourLightGrey(),
			"Cyan":      item.ColourCyan(),
			"Purple":    item.ColourPurple(),
			"Blue":      item.ColourBlue(),
			"Brown":     item.ColourBrown(),
			"Green":     item.ColourGreen(),
			"Red":       item.ColourRed(),
			"Black":     item.ColourBlack(),
		}).
		Const("BucketContents", map[string]item.BucketContent{
			"Milk":  item.MilkBucketContent(),
			"Water": item.LiquidBucketContent(block.Water{}),
			"Lava":  item.LiquidBucketContent(block.Lava{}),
		}).
		Const("BannerPatternType", map[string]item.BannerPatternType{
			"Creeper":         item.CreeperBannerPattern(),
			"Skull":           item.SkullBannerPattern(),
			"Flower":          item.FlowerBannerPattern(),
			"Mojang":          item.MojangBannerPattern(),
			"FieldMasoned":    item.FieldMasonedBannerPattern(),
			"BordureIndented": item.BordureIndentedBannerPattern(),
			"Piglin":          item.PiglinBannerPattern(),
			"Globe":           item.GlobeBannerPattern(),
			"Flow":            item.FlowBannerPattern(),
			"Guster":          item.GusterBannerPattern(),
		}).
		Const("FireworkShape", map[string]item.FireworkShape{
			"SmallSphere": item.FireworkShapeSmallSphere(),
			"HugeSphere":  item.FireworkShapeHugeSphere(),
			"Star":        item.FireworkShapeStar(),
			"CreeperHead": item.FireworkShapeCreeperHead(),
			"Burst":       item.FireworkShapeBurst(),
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
			m := c.Argument(0).Export().(map[string]any)
			t, _ := m["template"].(item.SmithingTemplateType)
			mat, _ := m["material"].(item.ArmourTrimMaterial)
			return r.vm.ToValue(item.ArmourTrim{Template: t, Material: mat})
		}).
		Method("arrow", func(c goja.FunctionCall) goja.Value {
			m := c.Argument(0).Export().(map[string]any)
			tip, _ := m["tip"].(potion.Potion)
			return r.vm.ToValue(item.Arrow{Tip: tip})
		}).
		Method("axe", func(c goja.FunctionCall) goja.Value {
			m := c.Argument(0).Export().(map[string]any)
			tier, _ := m["tier"].(item.ToolTier)
			return r.vm.ToValue(item.Axe{Tier: tier})
		}).
		Method("bakedPotato", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(item.BakedPotato{})
		}).
		Method("bannerPattern", func(c goja.FunctionCall) goja.Value {
			m := c.Argument(0).Export().(map[string]any)
			t, _ := m["type"].(item.BannerPatternType)
			return r.vm.ToValue(item.BannerPattern{Type: t})
		}).
		Method("beef", func(c goja.FunctionCall) goja.Value {
			m := c.Argument(0).Export().(map[string]any)
			cooked, _ := m["cooked"].(bool)
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
			m := c.Argument(0).Export().(map[string]any)
			pages, _ := m["pages"].([]string)
			return r.vm.ToValue(item.BookAndQuill{Pages: pages})
		}).
		Method("boots", func(c goja.FunctionCall) goja.Value {
			m := c.Argument(0).Export().(map[string]any)
			tier, _ := m["tier"].(item.ArmourTier)
			trim, _ := m["trim"].(item.ArmourTrim)
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
			m := c.Argument(0).Export().(map[string]any)
			cont, _ := m["content"].(item.BucketContent)
			return r.vm.ToValue(item.Bucket{Content: cont})
		}).
		Method("carrotOnAStick", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(item.CarrotOnAStick{})
		}).
		Method("charcoal", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(item.Charcoal{})
		}).
		Method("chestplate", func(c goja.FunctionCall) goja.Value {
			m := c.Argument(0).Export().(map[string]any)
			tier, _ := m["tier"].(item.ArmourTier)
			trim, _ := m["trim"].(item.ArmourTrim)
			return r.vm.ToValue(item.Chestplate{Tier: tier, Trim: trim})
		}).
		Method("chicken", func(c goja.FunctionCall) goja.Value {
			m := c.Argument(0).Export().(map[string]any)
			cooked, _ := m["cooked"].(bool)
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
			m := c.Argument(0).Export().(map[string]any)
			cooked, _ := m["cooked"].(bool)
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
			m := c.Argument(0).Export().(map[string]any)
			col, _ := m["colour"].(item.Colour)
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
			m := c.Argument(0).Export().(map[string]any)
			d, _ := m["duration"].(time.Duration)
			explosions, _ := m["explosions"].([]item.FireworkExplosion)
			return r.vm.ToValue(item.Firework{Duration: d, Explosions: explosions})
		}).
		Method("fireworkExplosion", func(c goja.FunctionCall) goja.Value {
			m := c.Argument(0).Export().(map[string]any)
			shape, _ := m["shape"].(item.FireworkShape)
			colour, _ := m["colour"].(item.Colour)
			fade, _ := m["fade"].(item.Colour)
			fades, _ := m["fades"].(bool)
			twinkle, _ := m["twinkle"].(bool)
			trail, _ := m["trail"].(bool)
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
			m := c.Argument(0).Export().(map[string]any)
			explosion, _ := m["explosion"].(item.FireworkExplosion)
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
			m := c.Argument(0).Export().(map[string]any)
			t, _ := m["type"].(sound.Horn)
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
			m := c.Argument(0).Export().(map[string]any)
			tier, _ := m["tier"].(item.ArmourTier)
			trim, _ := m["trim"].(item.ArmourTrim)
			return r.vm.ToValue(item.Helmet{Tier: tier, Trim: trim})
		}).
		Method("hoe", func(c goja.FunctionCall) goja.Value {
			m := c.Argument(0).Export().(map[string]any)
			tier, _ := m["tier"].(item.ToolTier)
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
			m := c.Argument(0).Export().(map[string]any)
			tier, _ := m["tier"].(item.ArmourTier)
			trim, _ := m["trim"].(item.ArmourTrim)
			return r.vm.ToValue(item.Leggings{Tier: tier, Trim: trim})
		}).
		Method("lingeringPotion", func(c goja.FunctionCall) goja.Value {
			m := c.Argument(0).Export().(map[string]any)
			t, _ := m["type"].(potion.Potion)
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
			m := c.Argument(0).Export().(map[string]any)
			t, _ := m["type"].(sound.DiscType)
			return r.vm.ToValue(item.MusicDisc{DiscType: t})
		}).
		Method("mutton", func(c goja.FunctionCall) goja.Value {
			m := c.Argument(0).Export().(map[string]any)
			cooked, _ := m["cooked"].(bool)
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
			m := c.Argument(0).Export().(map[string]any)
			tier, _ := m["tier"].(item.ToolTier)
			return r.vm.ToValue(item.Pickaxe{Tier: tier})
		}).
		Method("poisonousPotato", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(item.PoisonousPotato{})
		}).
		Method("poppedChorusFruit", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(item.PoppedChorusFruit{})
		}).
		Method("porkchop", func(c goja.FunctionCall) goja.Value {
			m := c.Argument(0).Export().(map[string]any)
			cooked, _ := m["cooked"].(bool)
			return r.vm.ToValue(item.Porkchop{Cooked: cooked})
		}).
		Method("potion", func(c goja.FunctionCall) goja.Value {
			m := c.Argument(0).Export().(map[string]any)
			t, _ := m["type"].(potion.Potion)
			return r.vm.ToValue(item.Potion{Type: t})
		}).
		Method("potterySherd", func(c goja.FunctionCall) goja.Value {
			m := c.Argument(0).Export().(map[string]any)
			t, _ := m["type"].(item.SherdType)
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
			m := c.Argument(0).Export().(map[string]any)
			cooked, _ := m["cooked"].(bool)
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
			m := c.Argument(0).Export().(map[string]any)
			cooked, _ := m["cooked"].(bool)
			return r.vm.ToValue(item.Salmon{Cooked: cooked})
		}).
		Method("scute", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(item.Scute{})
		}).
		Method("shears", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(item.Shears{})
		}).
		Method("shovel", func(c goja.FunctionCall) goja.Value {
			m := c.Argument(0).Export().(map[string]any)
			tier, _ := m["tier"].(item.ToolTier)
			return r.vm.ToValue(item.Shovel{Tier: tier})
		}).
		Method("shulkerShell", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(item.ShulkerShell{})
		}).
		Method("slimeball", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(item.Slimeball{})
		}).
		Method("smithingTemplate", func(c goja.FunctionCall) goja.Value {
			m := c.Argument(0).Export().(map[string]any)
			t, _ := m["template"].(item.SmithingTemplateType)
			return r.vm.ToValue(item.SmithingTemplate{Template: t})
		}).
		Method("snowball", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(item.Snowball{})
		}).
		Method("spiderEye", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(item.SpiderEye{})
		}).
		Method("splashPotion", func(c goja.FunctionCall) goja.Value {
			m := c.Argument(0).Export().(map[string]any)
			t, _ := m["type"].(potion.Potion)
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
			m := c.Argument(0).Export().(map[string]any)
			t, _ := m["type"].(item.StewType)
			return r.vm.ToValue(item.SuspiciousStew{Type: t})
		}).
		Method("sword", func(c goja.FunctionCall) goja.Value {
			m := c.Argument(0).Export().(map[string]any)
			tier, _ := m["tier"].(item.ToolTier)
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
			m := c.Argument(0).Export().(map[string]any)
			title, _ := m["title"].(string)
			author, _ := m["author"].(string)
			gen, _ := m["generation"].(item.WrittenBookGeneration)
			pages, _ := m["pages"].([]string)
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
