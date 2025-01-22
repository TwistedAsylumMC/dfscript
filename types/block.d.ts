/// <reference path="cube.d.ts" />
/// <reference path="customblock.d.ts" />
/// <reference path="item.d.ts" />
/// <reference path="model.d.ts" />

declare namespace block {
    interface Block {
        encodeBlock(): [string, { [key: string]: any }];
        hash(): [number, number];
        model(): model.Model;
    }
    interface Liquid {
        liquidDepth(): number;
        spreadDecay(): number;
        withDepth(depth: number, falling: boolean): Liquid;
        liquidFalling(): boolean;
        blastResistance(): number;
        liquidType(): string;
        // TODO: harden
    }

    interface CustomBlock {
        properties(): customblock.Properties;
    }
    function nextHash(): number;

    function breakDuration(): boolean;
    function breaksInstantly(): boolean;

    interface Attachment {
        uint8(): number;
        faceUint8(): number;
        rotateLeft(): Attachment;
        rotateRight(): Attachment;
        rotation(): cube.Rotation;
    }
    function wallAttachment(facing: cube.Direction): Attachment;
    function standingAttachment(orientation: cube.Orientation): Attachment;

    class AnvilType {
        static readonly Undamaged: AnvilType;
        static readonly SlightlyDamaged: AnvilType;
        static readonly VeryDamaged: AnvilType;

        uint8(): number;
        string(): string;
    }

    class BannerPattern {
        static readonly Border: BannerPattern;
        static readonly Bricks: BannerPattern;
        static readonly Circle: BannerPattern;
        static readonly Creeper: BannerPattern;
        static readonly Cross: BannerPattern;
        static readonly CurlyBorder: BannerPattern;
        static readonly DiagonalLeft: BannerPattern;
        static readonly DiagonalRight: BannerPattern;
        static readonly DiagonalUpLeft: BannerPattern;
        static readonly DiagonalUpRight: BannerPattern;
        static readonly Flower: BannerPattern;
        static readonly Gradient: BannerPattern;
        static readonly GradientUp: BannerPattern;
        static readonly HalfHorizontal: BannerPattern;
        static readonly HalfHorizontalBottom: BannerPattern;
        static readonly HalfVertical: BannerPattern;
        static readonly HalfVerticalRight: BannerPattern;
        static readonly Mojang: BannerPattern;
        static readonly Rhombus: BannerPattern;
        static readonly Skull: BannerPattern;
        static readonly SmallStripes: BannerPattern;
        static readonly SquareBottomLeft: BannerPattern;
        static readonly SquareBottomRight: BannerPattern;
        static readonly SquareTopLeft: BannerPattern;
        static readonly SquareTopRight: BannerPattern;
        static readonly StraightCross: BannerPattern;
        static readonly StripeBottom: BannerPattern;
        static readonly StripeCenter: BannerPattern;
        static readonly StripeDownLeft: BannerPattern;
        static readonly StripeDownRight: BannerPattern;
        static readonly StripeLeft: BannerPattern;
        static readonly StripeMiddle: BannerPattern;
        static readonly StripeRight: BannerPattern;
        static readonly StripeTop: BannerPattern;
        static readonly TriangleBottom: BannerPattern;
        static readonly TriangleTop: BannerPattern;
        static readonly TrianglesBottom: BannerPattern;
        static readonly TrianglesTop: BannerPattern;
        static readonly Globe: BannerPattern;
        static readonly Piglin: BannerPattern;
        static readonly Flow: BannerPattern;
        static readonly Guster: BannerPattern;

        uint8(): number;
        string(): string;
        item(): [item.BannerPatternType, boolean]
    }

    class BlackstoneType {
        static readonly Normal: BlackstoneType;
        static readonly Gilded: BlackstoneType;
        static readonly Polished: BlackstoneType;
        static readonly ChiseledPolished: BlackstoneType;

        uint8(): number;
        name(): string;
        string(): string;
    }

    class CopperType {
        static readonly Normal: CopperType;
        static readonly Cut: CopperType;
        static readonly Chiseled: CopperType;

        uint8(): number;
        name(): string;
        string(): string;
    }

    class CoralType {
        static readonly Tube: CoralType;
        static readonly Brain: CoralType;
        static readonly Bubble: CoralType;
        static readonly Fire: CoralType;
        static readonly Horn: CoralType;

        uint8(): number;
        colour(): item.Colour;
        name(): string;
        string(): string;
    }

    class DeepslateType {
        static readonly Normal: DeepslateType;
        static readonly Cobbled: DeepslateType;
        static readonly Polished: DeepslateType;
        static readonly Chiseled: DeepslateType;

        uint8(): number;
        name(): string;
        string(): string;
    }

    class FireType {
        static readonly Normal: FireType;
        static readonly Soul: FireType;

        uint8(): number;
        lightLevel(): number;
        damage(): number;
        name(): string;
        string(): string;
    }

    class DoubleFlowerType {
        static readonly Sunflower: DoubleFlowerType;
        static readonly Lilac: DoubleFlowerType;
        static readonly RoseBush: DoubleFlowerType;
        static readonly Peony: DoubleFlowerType;

        uint8(): number;
        name(): string;
        string(): string;
    }

    class DoubleTallGrassType {
        static readonly Normal: DoubleTallGrassType;
        static readonly Fern: DoubleTallGrassType;

        uint8(): number;
        name(): string;
        string(): string;
    }

    class FlowerType {
        static readonly Dandelion: FlowerType;
        static readonly Poppy: FlowerType;
        static readonly BlueOrchid: FlowerType;
        static readonly Allium: FlowerType;
        static readonly AzureBluet: FlowerType;
        static readonly RedTulip: FlowerType;
        static readonly OrangeTulip: FlowerType;
        static readonly WhiteTulip: FlowerType;
        static readonly PinkTulip: FlowerType;
        static readonly OxeyeDaisy: FlowerType;
        static readonly Cornflower: FlowerType;
        static readonly LilyOfTheValley: FlowerType;
        static readonly WitherRose: FlowerType;

        uint8(): number;
        name(): string;
        string(): string;
    }

    class FroglightType {
        static readonly Pearlescent: FroglightType;
        static readonly Verdant: FroglightType;
        static readonly Ochre: FroglightType;

        uint8(): number;
        name(): string;
        string(): string;
    }

    class GrindstoneAttachment {
        static readonly Standing: GrindstoneAttachment;
        static readonly Hanging: GrindstoneAttachment;
        static readonly Wall: GrindstoneAttachment;

        uint8(): number;
        string(): string;
    }

    class NetherBricksType {
        static readonly Normal: NetherBricksType;
        static readonly Red: NetherBricksType;
        static readonly Cracked: NetherBricksType;
        static readonly Chiseled: NetherBricksType;

        uint8(): number;
        name(): string;
        string(): string;
    }

    class OreType {
        static readonly Stone: OreType;
        static readonly Deepslate: OreType;

        uint8(): number;
        name(): string;
        string(): string;
        prefix(): string;
        hardness(): number;
    }

    class Oxidation {
        static readonly Unoxidised: Oxidation;
        static readonly Exposed: Oxidation;
        static readonly Weathered: Oxidation;
        static readonly Oxidised: Oxidation;

        uint8(): number;
        name(): string;
        increase(): [Oxidation, boolean]
        decrease(): [Oxidation, boolean]
        string(): string;
    }

    class PrismarineType {
        static readonly Normal: PrismarineType;
        static readonly Dark: PrismarineType;
        static readonly Bricks: PrismarineType;

        uint8(): number;
        name(): string;
        string(): string;
    }

    class SandstoneType {
        static readonly Normal: SandstoneType;
        static readonly Cut: SandstoneType;
        static readonly Chiseled: SandstoneType;
        static readonly Smooth: SandstoneType;

        uint8(): number;
        name(): string;
        string(): string;
    }

    class SkullType {
        static readonly Skeleton: SkullType;
        static readonly WitherSkeleton: SkullType;
        static readonly Zombie: SkullType;
        static readonly Player: SkullType;
        static readonly Creeper: SkullType;
        static readonly Dragon: SkullType;
        static readonly Piglin: SkullType;

        uint8(): number;
        name(): string;
        string(): string;
    }

    class StoneBricksType {
        static readonly Normal: StoneBricksType;
        static readonly Mossy: StoneBricksType;
        static readonly Cracked: StoneBricksType;
        static readonly Chiseled: StoneBricksType;

        uint8(): number;
        name(): string;
        string(): string;
    }

    class WallConnection {
        static readonly None: WallConnection;
        static readonly Short: WallConnection;
        static readonly Tall: WallConnection;

        uint8(): number;
        string(): string;
        height(): number;
    }

    class WoodType {
        static readonly Oak: WoodType;
        static readonly Spruce: WoodType;
        static readonly Birch: WoodType;
        static readonly Jungle: WoodType;
        static readonly Acacia: WoodType;
        static readonly DarkOak: WoodType;
        static readonly Crimson: WoodType;
        static readonly Warped: WoodType;
        static readonly Mangrove: WoodType;
        static readonly Cherry: WoodType;
        static readonly PaleOak: WoodType;

        uint8(): number;
        name(): string;
        string(): string;
        flammable(): boolean;
    }

    interface BannerPatternLayer {
        type: BannerPattern;
        colour: item.Colour;
    }
    function bannerPatternLayer(props?: { type: BannerPattern, colour: item.Colour }): BannerPatternLayer;

    interface CampfireItem {
        item: item.Stack;
        time: Date;
    }
    function campfireItem(item: item.Stack, time: Date): CampfireItem;

    interface CloseAction {}
    function closeAction(): CloseAction;

    interface ContinueCrackAction {
        breakTime: Date;
    }
    function continueCrackAction(breakTime: Date): ContinueCrackAction;

    interface DamageSource {
        block: Block;
    }
    function damageSource(block: Block): DamageSource;

    interface DecoratedPotWobbleAction {
        decoratedPot: DecoratedPot;
        success: boolean;
    }
    function decoratedPotWobbleAction(decoratedPot: DecoratedPot, success?: boolean): DecoratedPotWobbleAction;

    interface FireDamageSource {}
    function fireDamageSource(): FireDamageSource;

    interface LavaDamageSource {}
    function lavaDamageSource(): LavaDamageSource;

    interface OpenAction {}
    function openAction(): OpenAction;

    interface SignText {
        text: string;
        owner: string;
        // TODO: baseColour
        glowing: boolean;
    }
    function signText(props?: { text?: string, owner?: string, /*baseColour,*/ waxed?: boolean }): SignText;

    interface StartCrackAction {}
    function startCrackAction(): StartCrackAction;

    interface Air extends Block, item.Item {}
    function air(): Air;
    interface Amethyst extends Block, item.Item {}
    function amethyst(): Amethyst;
    interface AncientDebris extends Block, item.Item {}
    function ancientDebris(): AncientDebris;
    interface Anvil extends Block, item.Item {
        type: AnvilType;
        facing: cube.Direction;
    }
    function anvil(props?: { type?: AnvilType, facing?: cube.Direction }): Anvil;
    interface Banner extends Block, item.Item {
        colour: item.Colour;
        attach: Attachment;
        patterns: BannerPatternLayer[];
        illager: boolean;
    }
    function banner(props?: { colour?: item.Colour, attach?: Attachment, patterns?: BannerPatternLayer[], illager: boolean }): Banner;
    interface Barrel extends Block, item.Item {
        facing: cube.Face;
        open: boolean;
        customName: string;
    }
    function barrel(props?: { facing?: cube.Face, open?: boolean, customName?: string }): Barrel;
    interface Barrier extends Block, item.Item {}
    function barrier(): Barrier;
    interface Basalt extends Block, item.Item {
        axis: cube.Axis;
        polished: boolean;
    }
    function basalt(props?: { axis?: cube.Axis, polished?: boolean }): Basalt;
    interface Beacon extends Block, item.Item {} // TODO
    function beacon(): Beacon;
    interface Bedrock extends Block, item.Item {
        infiniteBurning: boolean;
    }
    function bedrock(props?: { infiniteBurning?: boolean }): Bedrock;
    interface BeetrootSeeds extends Block, item.Item {
        growth: number;
    }
    function beetrootSeeds(props?: { growth?: number }): BeetrootSeeds;
    interface Blackstone extends Block, item.Item {
        type: BlackstoneType;
    }
    function blackstone(props?: { type?: BlackstoneType }): Blackstone;
    interface BlastFurnace extends Block, item.Item {
        facing: cube.Direction;
        lit: boolean;
    }
    function blastFurnace(props?: { facing?: cube.Direction, lit?: boolean }): BlastFurnace;
    interface BlueIce extends Block, item.Item {}
    function blueIce(): BlueIce;
    interface Bone extends Block, item.Item {
        axis: cube.Axis;
    }
    function bone(props?: { axis?: cube.Axis }): Bone;
    interface Bookshelf extends Block, item.Item {}
    function bookshelf(): Bookshelf;
    interface BrewingStand extends Block, item.Item {
        left: boolean;
        middle: boolean;
        right: boolean;
    }
    function brewingStand(props?: { left?: boolean, middle?: boolean, right?: boolean }): BrewingStand;
    interface Bricks extends Block, item.Item {}
    function bricks(): Bricks;
    interface Cactus extends Block, item.Item {
        age: number;
    }
    function cactus(props?: { age?: number }): Cactus;
    interface Cake extends Block, item.Item {
        bites: number;
    }
    function cake(props?: { bites?: number }): Cake;
    interface Calcite extends Block, item.Item {}
    function calcite(): Calcite;
    interface Campfire extends Block, item.Item {
        type: FireType;
        items: CampfireItem;
        facing: cube.Direction;
        extinguished: boolean;
    }
    function campfire(props?: { type?: FireType, items?: CampfireItem, facing?: cube.Direction, extinguished?: boolean }): Campfire;
    interface Carpet extends Block, item.Item {
        colour: item.Colour;
    }
    function carpet(props?: { colour?: item.Colour }): Carpet;
    interface Carrot extends Block, item.Item {
        growth: number;
    }
    function carrot(props?: { growth?: number }): Carrot;
    interface Chain extends Block, item.Item {
        axis: cube.Axis;
    }
    function chain(props?: { axis?: cube.Axis }): Chain;
    interface Chest extends Block, item.Item {
        facing: cube.Direction
        customName: string
    }
    function chest(props?: { facing?: cube.Direction, customName?: string }): Chest;
    interface ChiseledQuartz extends Block, item.Item {}
    function chiseledQuartz(): ChiseledQuartz;
    interface Clay extends Block, item.Item {}
    function clay(): Clay;
    interface Coal extends Block, item.Item {}
    function coal(): Coal;
    interface CoalOre extends Block, item.Item {
        type: OreType;
    }
    function coalOre(props?: { type?: OreType }): CoalOre;
    interface Cobblestone extends Block, item.Item {
        mossy: boolean;
    }
    function cobblestone(props?: { mossy?: boolean }): Cobblestone;
    interface CocoaBean extends Block, item.Item {
        facing: cube.Direction;
        age: number;
    }
    function cocoaBean(props?: { facing?: cube.Direction, age?: number }): CocoaBean;
    interface Composter extends Block, item.Item {
        level: number;
    }
    function composter(props?: { level?: number }): Composter;
    interface Concrete extends Block, item.Item {
        colour: item.Colour;
    }
    function concrete(props?: { colour?: item.Colour }): Concrete;
    interface ConcretePowder extends Block, item.Item {
        colour: item.Colour;
    }
    function concretePowder(props?: { colour?: item.Colour }): ConcretePowder;
    interface Copper extends Block, item.Item {
        type: CopperType;
        oxidation: Oxidation;
        waxed: boolean;
    }
    function copper(props?: { type?: CopperType, oxidation?: Oxidation, waxed?: boolean }): Copper;
    interface CopperDoor extends Block, item.Item {
        facing: cube.Direction;
        oxidation: Oxidation;
        waxed: boolean;
        open: boolean;
        top: boolean;
        right: boolean;
    }
    function copperDoor(props?: { facing?: cube.Direction, oxidation?: Oxidation, waxed?: boolean, open?: boolean, top?: boolean, right?: boolean }): CopperDoor;
    interface CopperGrate extends Block, item.Item {
        oxidation: Oxidation;
        waxed: boolean;
    }
    function copperGrate(props?: { oxidation?: Oxidation, waxed?: boolean }): CopperGrate;
    interface CopperOre extends Block, item.Item {
        type: OreType;
    }
    function copperOre(props?: { type?: OreType }): CopperOre;
    interface CopperTrapdoor extends Block, item.Item {
        facing: cube.Direction;
        oxidation: Oxidation;
        waxed: boolean;
        open: boolean;
        top: boolean;
    }
    function copperTrapdoor(props?: { facing?: cube.Direction, oxidation?: Oxidation, waxed?: boolean, open?: boolean, top?: boolean }): CopperTrapdoor;
    interface Coral extends Block, item.Item {
        type: CoralType;
        dead: boolean;
    }
    function coral(props?: { type?: CoralType, dead?: boolean }): Coral;
    interface CoralBlock extends Block, item.Item {
        type: CoralType;
        dead: boolean;
    }
    function coralBlock(props?: { type?: CoralType, dead?: boolean }): CoralBlock;
    interface CraftingTable extends Block, item.Item {}
    function craftingTable(): CraftingTable;
    interface DeadBush extends Block, item.Item {}
    function deadBush(): DeadBush;
    interface DecoratedPot extends Block, item.Item {
        item: item.Stack;
        facing: cube.Direction;
        decorations: item.Stack[];
    }
    function decoratedPot(props?: { item?: item.Stack, facing?: cube.Direction, decorations?: item.Stack[] }): DecoratedPot;
    interface Deepslate extends Block, item.Item {
        type: DeepslateType;
        axis: cube.Axis;
    }
    function deepslate(props?: { type?: DeepslateType, axis?: cube.Axis }): Deepslate;
    interface DeepslateBricks extends Block, item.Item {
        cracked: boolean;
    }
    function deepslateBricks(props?: { cracked?: boolean }): DeepslateBricks;
    interface DeepslateTiles extends Block, item.Item {
        cracked: boolean;
    }
    function deepslateTiles(props?: { cracked?: boolean }): DeepslateTiles;
    interface Diamond extends Block, item.Item {}
    function diamond(): Diamond;
    interface DiamondOre extends Block, item.Item {
        type: OreType;
    }
    function diamondOre(props?: { type?: OreType }): DiamondOre;
    interface Dirt extends Block, item.Item {
        coarse: boolean;
    }
    function dirt(props?: { coarse?: boolean }): Dirt;
    interface DirtPath extends Block, item.Item {}
    function dirtPath(): DirtPath;
    interface DoubleFlower extends Block, item.Item {
        type: DoubleFlowerType;
        upperPart: boolean;
    }
    function doubleFlower(props?: { type?: DoubleFlowerType, upperPart?: boolean }): DoubleFlower;
    interface DoubleTallGrass extends Block, item.Item {
        type: DoubleTallGrassType;
        upperPart: boolean;
    }
    function doubleTallGrass(props?: { type?: DoubleTallGrassType, upperPart?: boolean }): DoubleTallGrass;
    interface DragonEgg extends Block, item.Item {}
    function dragonEgg(): DragonEgg;
    interface DriedKelp extends Block, item.Item {}
    function driedKelp(): DriedKelp;
    interface Dripstone extends Block, item.Item {}
    function dripstone(): Dripstone;
    interface Emerald extends Block, item.Item {}
    function emerald(): Emerald;
    interface EmeraldOre extends Block, item.Item {
        type: OreType;
    }
    function emeraldOre(props?: { type?: OreType }): EmeraldOre;
    interface EnchantingTable extends Block, item.Item {}
    function enchantingTable(): EnchantingTable;
    interface EndBricks extends Block, item.Item {}
    function endBricks(): EndBricks;
    interface EndRod extends Block, item.Item {
        facing: cube.Face
    }
    function endRod(props?: { facing?: cube.Face }): EndRod;
    interface EndStone extends Block, item.Item {}
    function endStone(): EndStone;
    interface EnderChest extends Block, item.Item {
        facing: cube.Direction
    }
    function enderChest(props?: { facing?: cube.Direction }): EnderChest;
    interface Farmland extends Block, item.Item {
        hydration: number;
    }
    function farmland(props?: { hydration?: number }): Farmland;
    interface Fern extends Block, item.Item {}
    function fern(): Fern;
    interface Fire extends Block, item.Item {
        type: FireType;
        age: number;
    }
    function fire(props?: { type?: FireType, age?: number }): Fire;
    interface FletchingTable extends Block, item.Item {}
    function fletchingTable(): FletchingTable;
    interface Flower extends Block, item.Item {
        type: FlowerType;
    }
    function flower(props?: { type?: FlowerType }): Flower;
    interface Froglight extends Block, item.Item {
        type: FroglightType;
    }
    function froglight(props?: { type?: FroglightType }): Froglight;
    interface Furnace extends Block, item.Item {
        facing: cube.Direction;
        lit: boolean;
    }
    function furnace(props?: { facing?: cube.Direction, lit?: boolean }): Furnace;
    interface Glass extends Block, item.Item {}
    function glass(): Glass;
    interface GlassPane extends Block, item.Item {}
    function glassPane(): GlassPane;
    interface GlazedTerracotta extends Block, item.Item {
        colour: item.Colour;
    }
    function glazedTerracotta(props?: { colour?: item.Colour }): GlazedTerracotta;
    interface Glowstone extends Block, item.Item {}
    function glowstone(): Glowstone;
    interface Gold extends Block, item.Item {}
    function gold(): Gold;
    interface GoldOre extends Block, item.Item {
        type: OreType;
    }
    function goldOre(props?: { type?: OreType }): GoldOre;
    interface Grass extends Block, item.Item {}
    function grass(): Grass;
    interface Gravel extends Block, item.Item {}
    function gravel(): Gravel;
    interface Grindstone extends Block, item.Item {
        facing: cube.Direction;
        attach: GrindstoneAttachment;
    }
    function grindstone(props?: { facing?: cube.Direction, attach?: GrindstoneAttachment }): Grindstone;
    interface HayBale extends Block, item.Item {
        axis: cube.Axis;
    }
    function hayBale(props?: { axis?: cube.Axis }): HayBale;
    interface Honeycomb extends Block, item.Item {}
    function honeycomb(): Honeycomb;
    interface Hopper extends Block, item.Item {
        facing: cube.Face;
        powered: boolean;
        customName: string;
        lastTick: number;
        transferCooldown: number;
        collectCooldown: number;
    }
    function hopper(props?: { facing?: cube.Face, powered?: boolean, customName?: string, lastTick?: number, transferCooldown?: number, collectCooldown?: number }): Hopper;
    interface InvisibleBedrock extends Block, item.Item {}
    function invisibleBedrock(): InvisibleBedrock;
    interface Iron extends Block, item.Item {}
    function iron(): Iron;
    interface IronBars extends Block, item.Item {}
    function ironBars(): IronBars;
    interface IronOre extends Block, item.Item {
        type: OreType;
    }
    function ironOre(props?: { type?: OreType }): IronOre;
    interface ItemFrame extends Block, item.Item {
        facing: cube.Face;
        item: item.Stack;
        rotations: number;
        dropChance: number;
        glowing: boolean;
    }
    function itemFrame(props?: { facing?: cube.Face, item?: item.Stack, rotations?: number, dropChance?: number, glowing?: boolean }): ItemFrame;
    interface Jukebox extends Block, item.Item {
        item: item.Stack;
    }
    function jukebox(props?: { item?: item.Stack }): Jukebox;
    interface Kelp extends Block, item.Item {
        age: number;
    }
    function kelp(props?: { age?: number }): Kelp;
    interface Ladder extends Block, item.Item {
        facing: cube.Direction;
    }
    function ladder(props?: { facing?: cube.Direction }): Ladder;
    interface Lantern extends Block, item.Item {
        type: FireType;
        hanging: boolean;
    }
    function lantern(props?: { type?: FireType, hanging?: boolean }): Lantern;
    interface Lapis extends Block, item.Item {}
    function lapis(): Lapis;
    interface LapisOre extends Block, item.Item {
        type: OreType;
    }
    function lapisOre(props?: { type?: OreType }): LapisOre;
    interface Lava extends Block, Liquid {
        depth: number;
        still: boolean;
        falling: boolean;
    }
    function lava(props?: { depth?: number, still?: boolean, falling?: boolean }): Lava;
    interface Leaves extends Block, item.Item {
        wood: WoodType;
        persistent: boolean;
        shouldUpdate: boolean;
    }
    function leaves(props?: { wood?: WoodType, persistent?: boolean, shouldUpdate?: boolean }): Leaves;
    interface Lectern extends Block, item.Item {
        facing: cube.Direction;
        book: item.Stack;
        page: number;
    }
    function lectern(props?: { facing?: cube.Direction, book?: item.Stack, page?: number }): Lectern;
    interface Light extends Block, item.Item {
        level: number;
    }
    function light(props?: { level?: number }): Light;
    interface LitPumpkin extends Block, item.Item {
        facing: cube.Direction;
    }
    function litPumpkin(props?: { facing?: cube.Direction }): LitPumpkin;
    interface Log extends Block, item.Item {
        wood: WoodType;
        axis: cube.Axis;
        stripped: boolean;
    }
    function log(props?: { wood?: WoodType, axis?: cube.Axis, stripped?: boolean }): Log;
    interface Loom extends Block, item.Item {
        facing: cube.Direction;
    }
    function loom(props?: { facing?: cube.Direction }): Loom;
    interface Melon extends Block, item.Item {}
    function melon(): Melon;
    interface MelonSeeds extends Block, item.Item {
        direction: cube.Face;
        growth: number;
    }
    function melonSeeds(props?: { direction?: cube.Face, growth?: number }): MelonSeeds;
    interface MossCarpet extends Block, item.Item {}
    function mossCarpet(): MossCarpet;
    interface Mud extends Block, item.Item {}
    function mud(): Mud;
    interface MudBricks extends Block, item.Item {}
    function mudBricks(): MudBricks;
    interface MuddyMangroveRoots extends Block, item.Item {
        axis: cube.Axis;
    }
    function muddyMangroveRoots(props?: { axis?: cube.Axis }): MuddyMangroveRoots;
    interface NetherBrickFence extends Block, item.Item {}
    function netherBrickFence(): NetherBrickFence;
    interface NetherBricks extends Block, item.Item {
        type: NetherBricksType;
    }
    function netherBricks(props?: { type?: NetherBricksType }): NetherBricks;
    interface NetherGoldOre extends Block, item.Item {
        type: OreType;
    }
    function netherGoldOre(props?: { type?: OreType }): NetherGoldOre;
    interface NetherQuartzOre extends Block, item.Item {
        type: OreType;
    }
    function netherQuartzOre(props?: { type?: OreType }): NetherQuartzOre;
    interface NetherSprouts extends Block, item.Item {}
    function netherSprouts(): NetherSprouts;
    interface NetherWart extends Block, item.Item {
        age: number;
    }
    function netherWart(props?: { age?: number }): NetherWart;
    interface NetherWartBlock extends Block, item.Item {
        warped: boolean;
    }
    function netherWartBlock(props?: { warped?: boolean }): NetherWartBlock;
    interface Netherite extends Block, item.Item {}
    function netherite(): Netherite;
    interface Netherrack extends Block, item.Item {}
    function netherrack(): Netherrack;
    interface Note extends Block, item.Item {
        pitch: number;
    }
    function note(props?: { pitch?: number }): Note;
    interface Obsidian extends Block, item.Item {
        crying: boolean;
    }
    function obsidian(props?: { crying?: boolean }): Obsidian;
    interface PackedIce extends Block, item.Item {}
    function packedIce(): PackedIce;
    interface PackedMud extends Block, item.Item {}
    function packedMud(): PackedMud;
    interface PinkPetals extends Block, item.Item {
        facing: cube.Direction;
        additionalCount: number;
    }
    function pinkPetals(props?: { facing?: cube.Direction, additionalCount?: number }): PinkPetals;
    interface Planks extends Block, item.Item {
        wood: WoodType;
    }
    function planks(props?: { wood?: WoodType }): Planks;
    interface Podzol extends Block, item.Item {}
    function podzol(): Podzol;
    interface PolishedBlackstoneBrick extends Block, item.Item {
        cracked: boolean;
    }
    function polishedBlackstoneBrick(props?: { cracked?: boolean }): PolishedBlackstoneBrick;
    interface PolishedTuff extends Block, item.Item {}
    function polishedTuff(): PolishedTuff;
    interface Potato extends Block, item.Item {
        growth: number;
    }
    function potato(props?: { growth?: number }): Potato;
    interface Prismarine extends Block, item.Item {
        type: PrismarineType;
    }
    function prismarine(props?: { type?: PrismarineType }): Prismarine;
    interface Pumpkin extends Block, item.Item {
        facing: cube.Direction;
    }
    function pumpkin(props?: { facing?: cube.Direction }): Pumpkin;
    interface PumpkinSeeds extends Block, item.Item {
        direction: cube.Face;
        growth: number;
    }
    function pumpkinSeeds(props?: { direction?: cube.Face, growth?: number }): PumpkinSeeds;
    interface Purpur extends Block, item.Item {}
    function purpur(): Purpur;
    interface PurpurPillar extends Block, item.Item {
        axis: cube.Axis;
    }
    function purpurPillar(props?: { axis?: cube.Axis }): PurpurPillar;
    interface Quartz extends Block, item.Item {
        smooth: boolean;
    }
    function quartz(props?: { smooth?: boolean }): Quartz;
    interface QuartzBricks extends Block, item.Item {}
    function quartzBricks(): QuartzBricks;
    interface QuartzPillar extends Block, item.Item {
        axis: cube.Axis;
    }
    function quartzPillar(props?: { axis?: cube.Axis }): QuartzPillar;
    interface RawCopper extends Block, item.Item {}
    function rawCopper(): RawCopper;
    interface RawGold extends Block, item.Item {}
    function rawGold(): RawGold;
    interface RawIron extends Block, item.Item {}
    function rawIron(): RawIron;
    interface ReinforcedDeepslate extends Block, item.Item {}
    function reinforcedDeepslate(): ReinforcedDeepslate;
    interface Resin extends Block, item.Item {}
    function resin(): Resin;
    interface ResinBricks extends Block, item.Item {
        chiseled: boolean;
    }
    function resinBricks(props?: { chiseled?: boolean }): ResinBricks;
    interface Sand extends Block, item.Item {
        red: boolean;
    }
    function sand(props?: { red?: boolean }): Sand;
    interface Sandstone extends Block, item.Item {
        type: SandstoneType;
        red: boolean;
    }
    function sandstone(props?: { type?: SandstoneType, red?: boolean }): Sandstone;
    interface SeaLantern extends Block, item.Item {}
    function seaLantern(): SeaLantern;
    interface SeaPickle extends Block, item.Item {
        dead: boolean;
    }
    function seaPickle(props?: { dead?: boolean }): SeaPickle;
    interface ShortGrass extends Block, item.Item {
        double: boolean;
    }
    function shortGrass(props?: { double?: boolean }): ShortGrass;
    interface Shroomlight extends Block, item.Item {}
    function shroomlight(): Shroomlight;
    interface Sign extends Block, item.Item {
        wood: WoodType;
        attach: Attachment;
        front: SignText;
        back: SignText;
        waxed: boolean;
    }
    function sign(props?: { wood?: WoodType, attach?: Attachment, front?: SignText, back?: SignText, waxed?: boolean }): Sign;
    interface Skull extends Block, item.Item {
        type: SkullType;
        attach: Attachment;
    }
    function skull(props?: { type?: SkullType, attach?: Attachment }): Skull;
    interface Slab extends Block, item.Item {
        block: Block;
        double: boolean;
        top: boolean;
    }
    function slab(props?: { block?: Block, double?: boolean, top?: boolean }): Slab;
    interface SmithingTable extends Block, item.Item {}
    function smithingTable(): SmithingTable;
    interface Smoker extends Block, item.Item {
        facing: cube.Direction;
        lit: boolean;
    }
    function smoker(props?: { facing?: cube.Direction, lit?: boolean }): Smoker;
    interface Snow extends Block, item.Item {}
    function snow(): Snow;
    interface SoulSand extends Block, item.Item {}
    function soulSand(): SoulSand;
    interface SoulSoil extends Block, item.Item {}
    function soulSoil(): SoulSoil;
    interface Sponge extends Block, item.Item {
        wet: boolean;
    }
    function sponge(props?: { wet?: boolean }): Sponge;
    interface SporeBlossom extends Block, item.Item {}
    function sporeBlossom(): SporeBlossom;
    interface StainedGlass extends Block, item.Item {
        colour: item.Colour;
    }
    function stainedGlass(props?: { colour?: item.Colour }): StainedGlass;
    interface StainedGlassPane extends Block, item.Item {
        colour: item.Colour;
    }
    function stainedGlassPane(props?: { colour?: item.Colour }): StainedGlassPane;
    interface StainedTerracotta extends Block, item.Item {
        colour: item.Colour;
    }
    function stainedTerracotta(props?: { colour?: item.Colour }): StainedTerracotta;
    interface Stairs extends Block, item.Item {
        block: Block;
        facing: cube.Direction;
        upsideDown: boolean;
    }
    function stairs(props?: { block?: Block, facing?: cube.Direction, upsideDown?: boolean }): Stairs;
    interface Stone extends Block, item.Item {
        smooth: boolean;
    }
    function stone(props?: { smooth?: boolean }): Stone;
    interface StoneBricks extends Block, item.Item {
        type: StoneBricksType;
    }
    function stoneBricks(props?: { type?: StoneBricksType }): StoneBricks;
    interface Stonecutter extends Block, item.Item {
        facing: cube.Direction
    }
    function stonecutter(props?: { facing?: cube.Direction }): Stonecutter;
    interface StopCrackAction extends Block, item.Item {}
    function stopCrackAction(): StopCrackAction;
    interface SugarCane extends Block, item.Item {
        age: number;
    }
    function sugarCane(props?: { age?: number }): SugarCane;
    interface Tnt extends Block, item.Item {}
    function tnt(): Tnt;
    interface Terracotta extends Block, item.Item {}
    function terracotta(): Terracotta;
    interface Torch extends Block, item.Item {
        facing: cube.Face;
        type: FireType;
    }
    function torch(props?: { facing?: cube.Face, type?: FireType }): Torch;
    interface Tuff extends Block, item.Item {
        chiseled: boolean;
    }
    function tuff(props?: { chiseled?: boolean }): Tuff;
    interface TuffBricks extends Block, item.Item {
        chiseled: boolean;
    }
    function tuffBricks(props?: { chiseled?: boolean }): TuffBricks;
    interface Vines extends Block, item.Item {
        northDirection: boolean;
        eastDirection: boolean;
        southDirection: boolean;
        westDirection: boolean;
    }
    function vines(props?: { north?: boolean, east?: boolean, south?: boolean, west?: boolean }): Vines;
    interface Wall extends Block, item.Item {
        block: Block;
        northConnection: WallConnection;
        eastConnection: WallConnection;
        southConnection: WallConnection;
        westConnection: WallConnection;
        post: boolean;
    }
    function wall(props?: { block?: Block, north?: WallConnection, east?: WallConnection, south?: WallConnection, west?: WallConnection, post?: boolean }): Wall;
    interface Water extends Block, Liquid {
        depth: number;
        still: boolean;
        falling: boolean;
    }
    function water(props?: { depth?: number, still?: boolean, falling?: boolean }): Water;
    interface WheatSeeds extends Block, item.Item {
        growth: number;
    }
    function wheatSeeds(props?: { growth?: number }): WheatSeeds;
    interface Wood extends Block, item.Item {
        wood: WoodType;
        axis: cube.Axis;
        stripped: boolean;
    }
    function wood(props?: { wood?: WoodType, axis?: cube.Axis, stripped?: boolean }): Wood;
    interface WoodDoor extends Block, item.Item {
        wood: WoodType;
        facing: cube.Direction;
        open: boolean;
        top: boolean;
        right: boolean;
    }
    function woodDoor(props?: { wood?: WoodType, facing?: cube.Direction, open?: boolean, top?: boolean, right?: boolean }): WoodDoor;
    interface WoodFence extends Block, item.Item {
        wood: WoodType;
    }
    function woodFence(props?: { wood?: WoodType }): WoodFence;
    interface WoodFenceGate extends Block, item.Item {
        wood: WoodType;
        facing: cube.Direction;
        open: boolean;
        lowered: boolean;
    }
    function woodFenceGate(props?: { wood?: WoodType, facing?: cube.Direction, open?: boolean, lowered?: boolean }): WoodFenceGate;
    interface WoodTrapdoor extends Block, item.Item {
        wood: WoodType;
        facing: cube.Direction;
        open: boolean;
        top: boolean;
    }
    function woodTrapdoor(props?: { wood?: WoodType, facing?: cube.Direction, open?: boolean, top ?: boolean }): WoodTrapdoor;
    interface Wool extends Block, item.Item {
        colour: item.Colour;
    }
    function wool(props?: { colour?: item.Colour }): Wool;
}