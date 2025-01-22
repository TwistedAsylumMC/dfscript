/// <reference path="block.d.ts" />
/// <reference path="category.d.ts" />
/// <reference path="effect.d.ts" />
/// <reference path="enchantment.d.ts" />
/// <reference path="potion.d.ts" />
/// <reference path="sound.d.ts" />

declare namespace item {
    interface Item {
        encodeItem(): [string, number];
    }
    interface CustomItem {
        name(): string;
        // TODO: texture
        category(): category.Category;
    }

    interface Stack {
        count(): number;
        maxCount(): number;
        grow(n: number): Stack;
        durability(): number;
        maxDurability(): number;
        damage(n: number): Stack;
        withDurability(durability: number): Stack;
        empty(): boolean;
        item(): Item;
        attackDamage(): number;
        withCustomName(...args: any[]): Stack;
        customName(): string;
        withLore(...lines: string[]): Stack;
        lore(): string[];
        withValue(key: string, value: any): Stack;
        value(key: string): [any, boolean];
        withEnchantments(...enchants: enchantment.Enchantment[]): Stack;
        withoutEnchantments(...enchants: enchantment.Enchantment[]): Stack;
        enchantment(type: enchantment.EnchantmentType): [enchantment.Enchantment, boolean];
        enchantments(): enchantment.Enchantment[];
        anvilCost(): number;
        withAnvilCost(cost: number): Stack;
        withItem(item: Item): Stack;
        addStack(other: Stack): Stack;
        equal(other: Stack): boolean;
        comparable(other: Stack): boolean;
        string(): string;
        values(): { [key: string]: any };
    }
    function stack(item: item.Item, count?: number): Stack;

    class BannerPatternType {
        static readonly Creeper: BannerPatternType;
        static readonly Skull: BannerPatternType;
        static readonly Flower: BannerPatternType;
        static readonly Mojang: BannerPatternType;
        static readonly FieldMasoned: BannerPatternType;
        static readonly BordureIndented: BannerPatternType;
        static readonly Piglin: BannerPatternType;
        static readonly Globe: BannerPatternType;
        static readonly Flow: BannerPatternType;
        static readonly Guster: BannerPatternType;

        uint8(): number;
        string(): string;
    }

    class BucketContents {
        static readonly Lava: BucketContents;
        static readonly Water: BucketContents;
        static readonly Milk: BucketContents;

        string(): string;
        liquidType(): string;
        liquid(): [block.Liquid, boolean];
    }

    class Colour {
        static readonly White: Colour;
        static readonly Orange: Colour;
        static readonly Magenta: Colour;
        static readonly LightBlue: Colour;
        static readonly Yellow: Colour;
        static readonly Lime: Colour;
        static readonly Pink: Colour;
        static readonly Grey: Colour;
        static readonly LightGrey: Colour;
        static readonly Cyan: Colour;
        static readonly Purple: Colour;
        static readonly Blue: Colour;
        static readonly Brown: Colour;
        static readonly Green: Colour;
        static readonly Red: Colour;
        static readonly Black: Colour;

        // TODO: rgba
        // TODO: signRGBA
        uint8(): number;
        string(): string;
        silverString(): string;
    }

    class FireworkShape {
        static readonly SmallSphere: FireworkShape;
        static readonly HugeSphere: FireworkShape;
        static readonly Star: FireworkShape;
        static readonly CreeperHead: FireworkShape;
        static readonly Burst: FireworkShape;

        uint8(): number;
        name(): string;
        string(): string;
    }

    class SherdType {
        static readonly Angler: SherdType;
        static readonly Archer: SherdType;
        static readonly ArmsUp: SherdType;
        static readonly Blade: SherdType;
        static readonly Brewer: SherdType;
        static readonly Burn: SherdType;
        static readonly Danger: SherdType;
        static readonly Explorer: SherdType;
        static readonly Friend: SherdType;
        static readonly Heart: SherdType;
        static readonly Heartbreak: SherdType;
        static readonly Howl: SherdType;
        static readonly Miner: SherdType;
        static readonly Mourner: SherdType;
        static readonly Plenty: SherdType;
        static readonly Prize: SherdType;
        static readonly Sheaf: SherdType;
        static readonly Shelter: SherdType;
        static readonly Skull: SherdType;
        static readonly Snort: SherdType;
        static readonly Flow: SherdType;
        static readonly Guster: SherdType;
        static readonly Scrape: SherdType;

        uint8(): number;
        string(): string;
    }

    class StewType {
        static readonly NightVisionPoppy: StewType;
        static readonly JumpBoost: StewType;
        static readonly Weakness: StewType;
        static readonly BlindnessBluet: StewType;
        static readonly Poison: StewType;
        static readonly SaturationDandelion: StewType;
        static readonly SaturationOrchid: StewType;
        static readonly FireResistance: StewType;
        static readonly Regeneration: StewType;
        static readonly Wither: StewType;
        static readonly NightVisionTorchflower: StewType;
        static readonly BlindnessEyeBlossom: StewType;
        static readonly NauseaBlossom: StewType;

        uint8(): number;
        effects: effect.Effect[];
    }

    class SmithingTemplateType {
        static readonly NetheriteUpgrade: SmithingTemplateType;
        static readonly Sentry: SmithingTemplateType;
        static readonly Vex: SmithingTemplateType;
        static readonly Wild: SmithingTemplateType;
        static readonly Coast: SmithingTemplateType;
        static readonly Dune: SmithingTemplateType;
        static readonly WayFinder: SmithingTemplateType;
        static readonly Raiser: SmithingTemplateType;
        static readonly Shaper: SmithingTemplateType;
        static readonly Host: SmithingTemplateType;
        static readonly Ward: SmithingTemplateType;
        static readonly Silence: SmithingTemplateType;
        static readonly Tide: SmithingTemplateType;
        static readonly Snout: SmithingTemplateType;
        static readonly Rib: SmithingTemplateType;
        static readonly Eye: SmithingTemplateType;
        static readonly Spire: SmithingTemplateType;
        static readonly Flow: SmithingTemplateType;
        static readonly Bolt: SmithingTemplateType;

        string(): string;
    }

    class ToolType {
        static readonly None: ToolType;
        static readonly Pickaxe: ToolType;
        static readonly Axe: ToolType;
        static readonly Hoe: ToolType;
        static readonly Shovel: ToolType;
        static readonly Shears: ToolType;
    }

    class ToolTier {
        static readonly Wood: ToolTier;
        static readonly Stone: ToolTier;
        static readonly Iron: ToolTier;
        static readonly Gold: ToolTier;
        static readonly Diamond: ToolTier;
        static readonly Netherite: ToolTier;

        harvestLevel: number;
        baseMiningEfficiency: number;
        baseAttackDamag: number;
        enchantmentValue: number;
        durability: number;
        name: string;
    }

    class WrittenBookGeneration {
        static readonly Original: WrittenBookGeneration;
        static readonly Copy: WrittenBookGeneration;
        static readonly CopyOfCopy: WrittenBookGeneration;

        uint8(): number;
        string(): string;
    }

    interface Armour {
        defencePoints(): number;
        toughness(): number;
        knockBackResistance(): number;
    }
    interface Helmet extends Armour {
        helmet(): boolean;
    }
    interface Chestplate extends Armour {
        chestplate(): boolean;
    }
    interface Leggings extends Armour {
        leggings(): boolean;
    }
    interface Boots extends Armour {
        boots(): boolean;
    }

    interface ArmourTier {
        baseDurability(): number;
        toughness(): number;
        knockBackResistance(): number;
        enchantmentValue(): number;
        name(): string;
    }
    interface ArmourTierChain extends ArmourTier {}
    function armourTierChain(): ArmourTierChain;
    interface ArmourTierDiamond extends ArmourTier {}
    function armourTierDiamond(): ArmourTierDiamond;
    interface ArmourTierGold extends ArmourTier {}
    function armourTierGold(): ArmourTierGold;
    interface ArmourTierIron extends ArmourTier {}
    function armourTierIron(): ArmourTierIron;
    interface ArmourTierLeather extends ArmourTier {
        // TODO: colour
    }
    function armourTierLeather(): ArmourTierLeather;
    interface ArmourTierNetherite extends ArmourTier {}
    function armourTierNetherite(): ArmourTierNetherite;

    interface ArmourTrimMaterial {
        trimMaterial(): string;
        materialColour(): string;
    }

    interface FireworkExplosion {
        shape: FireworkShape;
        colour: Colour;
        fade: Colour;
        fades: boolean;
        twinkle: boolean;
        trail: boolean;
    }
    function fireworkExplosion(props?: { shape?: FireworkShape, colour?: Colour, fade?: Colour, fades?: boolean, twinkle?: boolean, trail?: boolean }): FireworkExplosion;

    interface Tool {
        toolType(): ToolType;
        harvestLevel(): number;
        baseMiningEfficiency(block: block.Block): number;
    }
    interface ToolNone extends Tool {}
    function toolNone(): ToolNone;

    interface AmethystShard extends Item, ArmourTrimMaterial {}
    function amethystShard(): AmethystShard;
    interface Apple extends Item {}
    function apple(): Apple;
    interface ArmourTrim extends Item {
        template: SmithingTemplateType;
        material: ArmourTrimMaterial;
    }
    function armourTrim(props?: { template?: SmithingTemplateType, material?: ArmourTrimMaterial }): ArmourTrim;
    interface Arrow extends Item {
        Tip: potion.Potion;
    }
    function arrow(props?: { tip?: potion.Potion }): Arrow;
    interface Axe extends Item, Tool {
        tier: ToolTier;
    }
    function axe(props?: { tier?: ToolTier }): Axe;
    interface BakedPotato extends Item {}
    function bakedPotato(): BakedPotato;
    interface BannerPattern extends Item {
        type: BannerPatternType;
    }
    function bannerPattern(props?: { type?: BannerPatternType }): BannerPattern;
    interface Beef extends Item {
        cooked: boolean;
    }
    function beef(props?: { cooked?: boolean }): Beef;
    interface Beetroot extends Item {}
    function beetroot(): Beetroot;
    interface BeetrootSoup extends Item {}
    function beetrootSoup(): BeetrootSoup;
    interface BlazePowder extends Item {}
    function blazePowder(): BlazePowder;
    interface BlazeRod extends Item {}
    function blazeRod(): BlazeRod;
    interface Bone extends Item {}
    function bone(): Bone;
    interface BoneMeal extends Item {}
    function boneMeal(): BoneMeal;
    interface Book extends Item {}
    function book(): Book;
    interface BookAndQuill extends Item {
        pages: string[];
    }
    function bookAndQuill(props?: { pages?: string[] }): BookAndQuill;
    interface Boots extends Item {
        tier: ArmourTier;
        trim: ArmourTrim;
    }
    function boots(props?: { tier?: ArmourTier, trim?: ArmourTrim }): Boots;
    interface BottleOfEnchanting extends Item {}
    function bottleOfEnchanting(): BottleOfEnchanting;
    interface Bow extends Item {}
    function bow(): Bow;
    interface Bowl extends Item {}
    function bowl(): Bowl;
    interface Bread extends Item {}
    function bread(): Bread;
    interface Brick extends Item {}
    function brick(): Brick;
    interface Bucket extends Item {
        content: BucketContents;
    }
    function bucket(props?: { content?: BucketContents }): Bucket;
    interface CarrotOnAStick extends Item {}
    function carrotOnAStick(): CarrotOnAStick;
    interface Charcoal extends Item {}
    function charcoal(): Charcoal;
    interface Chestplate extends Item {
        tier: ArmourTier;
        trim: ArmourTrim;
    }
    function chestplate(props?: { tier?: ArmourTier, trim?: ArmourTrim }): Chestplate;
    interface Chicken extends Item {
        cooked: boolean;
    }
    function chicken(props?: { cooked?: boolean }): Chicken;
    interface ClayBall extends Item {}
    function clayBall(): ClayBall;
    interface Clock extends Item {}
    function clock(): Clock;
    interface Coal extends Item {}
    function coal(): Coal;
    interface Cod extends Item {
        cooked: boolean;
    }
    function cod(props?: { cooked?: boolean }): Cod;
    interface Compass extends Item {}
    function compass(): Compass;
    interface Cookie extends Item {}
    function cookie(): Cookie;
    interface CopperIngot extends Item, ArmourTrimMaterial {}
    function copperIngot(): CopperIngot;
    interface Crossbow extends Item {}
    function crossbow(): Crossbow;
    interface Diamond extends Item, ArmourTrimMaterial {}
    function diamond(): Diamond;
    interface DiscFragment extends Item {}
    function discFragment(): DiscFragment;
    interface DragonBreath extends Item {}
    function dragonBreath(): DragonBreath;
    interface DriedKelp extends Item {}
    function driedKelp(): DriedKelp;
    interface Dye extends Item {
        colour: Colour;
    }
    function dye(props?: { colour?: Colour }): Dye;
    interface EchoShard extends Item {}
    function echoShard(): EchoShard;
    interface Egg extends Item {}
    function egg(): Egg;
    interface Elytra extends Item {}
    function elytra(): Elytra;
    interface Emerald extends Item, ArmourTrimMaterial {}
    function emerald(): Emerald;
    interface EnchantedApple extends Item {}
    function enchantedApple(): EnchantedApple;
    interface EnchantedBook extends Item {}
    function enchantedBook(): EnchantedBook;
    interface EnderPearl extends Item {}
    function enderPearl(): EnderPearl;
    interface Feather extends Item {}
    function feather(): Feather;
    interface FermentedSpiderEye extends Item {}
    function fermentedSpiderEye(): FermentedSpiderEye;
    interface FireCharge extends Item {}
    function fireCharge(): FireCharge;
    interface Firework extends Item {
        duration: Date;
        explosions: FireworkExplosion[];
    }
    function firework(props?: { duration?: Date, explosions?: FireworkExplosion[] }): Firework;
    interface FireworkStar extends Item {
        fireworkExplosion: FireworkExplosion;
    }
    function fireworkStar(props?: { explosion?: FireworkExplosion }): FireworkStar;
    interface Flint extends Item {}
    function flint(): Flint;
    interface FlintAndSteel extends Item {}
    function flintAndSteel(): FlintAndSteel;
    interface GhastTear extends Item {}
    function ghastTear(): GhastTear;
    interface GlassBottle extends Item {}
    function glassBottle(): GlassBottle;
    interface GlisteringMelonSlice extends Item {}
    function glisteringMelonSlice(): GlisteringMelonSlice;
    interface GlowstoneDust extends Item {}
    function glowstoneDust(): GlowstoneDust;
    interface GoatHorn extends Item {
        type: sound.HornType;
    }
    function goatHorn(props?: { type?: sound.HornType }): GoatHorn;
    interface GoldIngot extends Item, ArmourTrimMaterial {}
    function goldIngot(): GoldIngot;
    interface GoldNugget extends Item {}
    function goldNugget(): GoldNugget;
    interface GoldenApple extends Item {}
    function goldenApple(): GoldenApple;
    interface GoldenCarrot extends Item {}
    function goldenCarrot(): GoldenCarrot;
    interface Gunpowder extends Item {}
    function gunpowder(): Gunpowder;
    interface HeartOfTheSea extends Item {}
    function heartOfTheSea(): HeartOfTheSea;
    interface Helmet extends Item {
        tier: ArmourTier;
        trim: ArmourTrim;
    }
    function helmet(props?: { tier?: ArmourTier, trim?: ArmourTrim }): Helmet;
    interface Hoe extends Item, Tool {
        tier: ToolTier;
    }
    function hoe(props?: { tier?: ToolTier }): Hoe;
    interface Honeycomb extends Item {}
    function honeycomb(): Honeycomb;
    interface InkSac extends Item {}
    function inkSac(): InkSac;
    interface IronIngot extends Item, ArmourTrimMaterial {}
    function ironIngot(): IronIngot;
    interface IronNugget extends Item {}
    function ironNugget(): IronNugget;
    interface LapisLazuli extends Item, ArmourTrimMaterial {}
    function lapisLazuli(): LapisLazuli;
    interface Leather extends Item {}
    function leather(): Leather;
    interface Leggings extends Item {
        tier: ArmourTier;
        trim: ArmourTrim;
    }
    function leggings(props?: { tier?: ArmourTier, trim?: ArmourTrim }): Leggings;
    interface LingeringPotion extends Item {
        type: potion.Potion;
    }
    function lingeringPotion(props?: { type?: potion.Potion }): LingeringPotion;
    interface MagmaCream extends Item {}
    function magmaCream(): MagmaCream;
    interface MelonSlice extends Item {}
    function melonSlice(): MelonSlice;
    interface MushroomStew extends Item {}
    function mushroomStew(): MushroomStew;
    interface MusicDisc extends Item {
        type: sound.DiscType;
    }
    function musicDisc(props?: { type?: sound.DiscType }): MusicDisc;
    interface Mutton extends Item {
        cooked: boolean;
    }
    function mutton(props?: { cooked?: boolean }): Mutton;
    interface NautilusShell extends Item {}
    function nautilusShell(): NautilusShell;
    interface NetherBrick extends Item {}
    function netherBrick(): NetherBrick;
    interface NetherQuartz extends Item, ArmourTrimMaterial {}
    function netherQuartz(): NetherQuartz;
    interface NetherStar extends Item {}
    function netherStar(): NetherStar;
    interface NetheriteIngot extends Item, ArmourTrimMaterial {}
    function netheriteIngot(): NetheriteIngot;
    interface NetheriteScrap extends Item {}
    function netheriteScrap(): NetheriteScrap;
    interface Paper extends Item {}
    function paper(): Paper;
    interface PhantomMembrane extends Item {}
    function phantomMembrane(): PhantomMembrane;
    interface Pickaxe extends Item, Tool {
        tier: ToolTier;
    }
    function pickaxe(props?: { tier?: ToolTier }): Pickaxe;
    interface PoisonousPotato extends Item {}
    function poisonousPotato(): PoisonousPotato;
    interface PoppedChorusFruit extends Item {}
    function poppedChorusFruit(): PoppedChorusFruit;
    interface Porkchop extends Item {
        cooked: boolean;
    }
    function porkchop(props?: { cooked?: boolean }): Porkchop;
    interface Potion extends Item {
        type: potion.Potion;
    }
    function potion(props?: { type?: potion.Potion }): Potion;
    interface PotterySherd extends Item {
        type: SherdType;
    }
    function potterySherd(props?: { type?: SherdType }): PotterySherd;
    interface PrismarineCrystals extends Item {}
    function prismarineCrystals(): PrismarineCrystals;
    interface PrismarineShard extends Item {}
    function prismarineShard(): PrismarineShard;
    interface Pufferfish extends Item {}
    function pufferfish(): Pufferfish;
    interface PumpkinPie extends Item {}
    function pumpkinPie(): PumpkinPie;
    interface Rabbit extends Item {
        cooked: boolean;
    }
    function rabbit(props?: { cooked?: boolean }): Rabbit;
    interface RabbitFoot extends Item {}
    function rabbitFoot(): RabbitFoot;
    interface RabbitHide extends Item {}
    function rabbitHide(): RabbitHide;
    interface RabbitStew extends Item {}
    function rabbitStew(): RabbitStew;
    interface RawCopper extends Item {}
    function rawCopper(): RawCopper;
    interface RawGold extends Item {}
    function rawGold(): RawGold;
    interface RawIron extends Item {}
    function rawIron(): RawIron;
    interface RecoveryCompass extends Item {}
    function recoveryCompass(): RecoveryCompass;
    interface ResinBrick extends Item, ArmourTrimMaterial {}
    function resinBrick(): ResinBrick;
    interface RottenFlesh extends Item {}
    function rottenFlesh(): RottenFlesh;
    interface Salmon extends Item {
        cooked: boolean;
    }
    function salmon(props?: { cooked?: boolean }): Salmon;
    interface Scute extends Item {}
    function scute(): Scute;
    interface Shears extends Item, Tool {}
    function shears(): Shears;
    interface Shovel extends Item, Tool {
        tier: ToolTier;
    }
    function shovel(props?: { tier?: ToolTier }): Shovel;
    interface ShulkerShell extends Item {}
    function shulkerShell(): ShulkerShell;
    interface Slimeball extends Item {}
    function slimeball(): Slimeball;
    interface SmithingTemplate extends Item {
        template: SmithingTemplateType;
    }
    function smithingTemplate(props?: { template?: SmithingTemplateType }): SmithingTemplate;
    interface Snowball extends Item {}
    function snowball(): Snowball;
    interface SpiderEye extends Item {}
    function spiderEye(): SpiderEye;
    interface SplashPotion extends Item {
        type: potion.Potion;
    }
    function splashPotion(props?: { type?: potion.Potion }): SplashPotion;
    interface Spyglass extends Item {}
    function spyglass(): Spyglass;
    interface Stick extends Item {}
    function stick(): Stick;
    interface Sugar extends Item {}
    function sugar(): Sugar;
    interface SuspiciousStew extends Item {
        type: StewType;
    }
    function suspiciousStew(props?: { type?: StewType }): SuspiciousStew;
    interface Sword extends Item, Tool {
        tier: ToolTier;
    }
    function sword(props?: { tier?: ToolTier }): Sword;
    interface Totem extends Item {}
    function totem(): Totem;
    interface TropicalFish extends Item {}
    function tropicalFish(): TropicalFish;
    interface TurtleShell extends Item {}
    function turtleShell(): TurtleShell;
    interface WarpedFungusOnAStick extends Item {}
    function warpedFungusOnAStick(): WarpedFungusOnAStick;
    interface Wheat extends Item {}
    function wheat(): Wheat;
    interface WrittenBook extends Item {
        title: string;
        author: string;
        generation: WrittenBookGeneration;
        pages: string[];
    }
    function writtenBook(props?: { title?: string, author?: string, generation?: WrittenBookGeneration, pages?: string[] }): WrittenBook;
}