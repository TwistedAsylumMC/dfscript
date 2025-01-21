/// <reference path="cube.d.ts" />
/// <reference path="customblock.d.ts" />
/// <reference path="item.d.ts" />
/// <reference path="model.d.ts" />

declare namespace block {
    interface Block {
        encodeBlock(): [string, object];
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

    interface AnvilType {
        uint8(): number;
        string(): string;
    }
    const anvilType: {
        undamaged: AnvilType;
        slightlyDamaged: AnvilType;
        veryDamaged: AnvilType;
    }

    interface BannerPattern {
        uint8(): number;
        string(): string;
        item(): [item.BannerPatternType, boolean]
    }
    const bannerPattern: {
        border: BannerPattern;
        bricks: BannerPattern;
        circle: BannerPattern;
        creeper: BannerPattern;
        cross: BannerPattern;
        curlyBorder: BannerPattern;
        diagonalLeft: BannerPattern;
        diagonalRight: BannerPattern;
        diagonalUpLeft: BannerPattern;
        diagonalUpRight: BannerPattern;
        flower: BannerPattern;
        gradient: BannerPattern;
        gradientUp: BannerPattern;
        halfHorizontal: BannerPattern;
        halfHorizontalBottom: BannerPattern;
        halfVertical: BannerPattern;
        halfVerticalRight: BannerPattern;
        mojang: BannerPattern;
        rhombus: BannerPattern;
        skull: BannerPattern;
        smallStripes: BannerPattern;
        squareBottomLeft: BannerPattern;
        squareBottomRight: BannerPattern;
        squareTopLeft: BannerPattern;
        squareTopRight: BannerPattern;
        straightCross: BannerPattern;
        stripeBottom: BannerPattern;
        stripeCenter: BannerPattern;
        stripeDownLeft: BannerPattern;
        stripeDownRight: BannerPattern;
        stripeLeft: BannerPattern;
        stripeMiddle: BannerPattern;
        stripeRight: BannerPattern;
        stripeTop: BannerPattern;
        triangleBottom: BannerPattern;
        triangleTop: BannerPattern;
        trianglesBottom: BannerPattern;
        trianglesTop: BannerPattern;
        globe: BannerPattern;
        piglin: BannerPattern;
        flow: BannerPattern;
        guster: BannerPattern;
    }

    interface BlackstoneType {
        uint8(): number;
        name(): string;
        string(): string;
    }
    const blackstoneType: {
        normal: BlackstoneType;
        gilded: BlackstoneType;
        polished: BlackstoneType;
        chiseledPolished: BlackstoneType;
    }

    interface CopperType {
        uint8(): number;
        name(): string;
        string(): string;
    }
    const copperType: {
        normal: CopperType;
        cut: CopperType;
        chiseled: CopperType;
    }

    interface CoralType {
        uint8(): number;
        colour(): item.Colour;
        name(): string;
        string(): string;
    }
    const coralType: {
        tube: CoralType;
        brain: CoralType;
        bubble: CoralType;
        fire: CoralType;
        horn: CoralType;
    }

    interface DeepslateType {
        uint8(): number;
        name(): string;
        string(): string;
    }
    const deepslateType: {
        normal: DeepslateType;
        cobbled: DeepslateType;
        polished: DeepslateType;
        chiseled: DeepslateType;
    }

    interface FireType {
        uint8(): number;
        lightLevel(): number;
        damage(): number;
        name(): string;
        string(): string;
    }
    const fireType: {
        normal: FireType;
        soul: FireType;
    }

    interface DoubleFlowerType {
        uint8(): number;
        name(): string;
        string(): string;
    }
    const doubleFlowerType: {
        sunflower: DoubleFlowerType;
        lilac: DoubleFlowerType;
        roseBush: DoubleFlowerType;
        peony: DoubleFlowerType;
    }

    interface DoubleTallGrassType {
        uint8(): number;
        name(): string;
        string(): string;
    }
    const doubleTallGrassType: {
        normal: DoubleTallGrassType;
        fern: DoubleTallGrassType;
    }

    interface FlowerType {
        uint8(): number;
        name(): string;
        string(): string;
    }
    const flowerType: {
        dandelion: FlowerType;
        poppy: FlowerType;
        blueOrchid: FlowerType;
        allium: FlowerType;
        azureBluet: FlowerType;
        redTulip: FlowerType;
        orangeTulip: FlowerType;
        whiteTulip: FlowerType;
        pinkTulip: FlowerType;
        oxeyeDaisy: FlowerType;
        cornflower: FlowerType;
        lilyOfTheValley: FlowerType;
        witherRose: FlowerType;
    }

    interface FroglightType {
        uint8(): number;
        name(): string;
        string(): string;
    }
    const froglightType: {
        pearlescent: FroglightType;
        verdant: FroglightType;
        ochre: FroglightType;
    }

    interface GrindstoneAttachment {
        uint8(): number;
        string(): string;
    }
    const grindstoneAttachment: {
        standing: GrindstoneAttachment;
        hanging: GrindstoneAttachment;
        wall: GrindstoneAttachment;
    }

    interface NetherBricksType {
        uint8(): number;
        name(): string;
        string(): string;
    }
    const netherBricksType: {
        normal: NetherBricksType;
        red: NetherBricksType;
        cracked: NetherBricksType;
        chiseled: NetherBricksType;
    }

    interface OreType {
        uint8(): number;
        name(): string;
        string(): string;
        prefix(): string;
        hardness(): number;
    }
    const oreType: {
        stone: OreType;
        deepslate: OreType;
    }

    interface Oxidation {
        uint8(): number;
        name(): string;
        increase(): [Oxidation, boolean]
        decrease(): [Oxidation, boolean]
        string(): string;
    }
    const oxidation: {
        unoxidised: Oxidation;
        exposed: Oxidation;
        weathered: Oxidation;
        oxidised: Oxidation;
    }

    interface PrismarineType {
        uint8(): number;
        name(): string;
        string(): string;
    }
    const prismarineType: {
        normal: PrismarineType;
        dark: PrismarineType;
        bricks: PrismarineType;
    }

    interface SandstoneType {
        uint8(): number;
        name(): string;
        string(): string;
    }
    const sandstoneType: {
        normal: SandstoneType;
        cut: SandstoneType;
        chiseled: SandstoneType;
        smooth: SandstoneType;
    }

    interface SkullType {
        uint8(): number;
        name(): string;
        string(): string;
    }
    const skullType: {
        skeleton: SkullType;
        witherSkeleton: SkullType;
        zombie: SkullType;
        player: SkullType;
        creeper: SkullType;
        dragon: SkullType;
        piglin: SkullType;
    }

    interface StoneBricksType {
        uint8(): number;
        name(): string;
        string(): string;
    }
    const stoneBricksType: {
        normal: StoneBricksType;
        mossy: StoneBricksType;
        cracked: StoneBricksType;
        chiseled: StoneBricksType;
    }

    interface WallConnection {
        uint8(): number;
        string(): string;
        height(): number;
    }
    const wallConnection: {
        none: WallConnection;
        short: WallConnection;
        tall: WallConnection;
    }

    interface WoodType {
        uint8(): number;
        name(): string;
        string(): string;
        flammable(): boolean;
    }
    const woodType: {
        oak: WoodType;
        spruce: WoodType;
        birch: WoodType;
        jungle: WoodType;
        acacia: WoodType;
        darkOak: WoodType;
        crimson: WoodType;
        warped: WoodType;
        mangrove: WoodType;
        cherry: WoodType;
        paleOak: WoodType;
    }

    interface BannerPatternLayer {
        type: BannerPattern;
        colour: item.Colour;
    }
    function bannerPatternLayer(type: BannerPattern, colour: item.Colour): BannerPatternLayer;

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
    function signText(text: string, owner: string, /*baseColour,*/ waxed?: boolean): SignText;

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
    function anvil(type: AnvilType, facing: cube.Direction): Anvil;
    interface Banner extends Block, item.Item {
        colour: item.Colour;
        attach: Attachment;
        patterns: BannerPatternLayer[];
        illager: boolean;
    }
    function banner(colour: item.Colour, attach: Attachment, patterns: BannerPatternLayer[], illager?: boolean): Banner;
    interface Barrel extends Block, item.Item {
        facing: cube.Face;
        open: boolean;
        customName: string;
    }
    function barrel(facing: cube.Face, open?: boolean, customName?: string): Barrel;
    interface Barrier extends Block, item.Item {}
    function barrier(): Barrier;
    interface Basalt extends Block, item.Item {
        axis: cube.Axis;
        polished: boolean;
    }
    function basalt(axis: cube.Axis, polished?: boolean): Basalt;
    interface Beacon extends Block, item.Item {} // TODO
    function beacon(): Beacon;
    interface Bedrock extends Block, item.Item {
        infiniteBurning: boolean;
    }
    function bedrock(infiniteBurning?: boolean): Bedrock;
    interface BeetrootSeeds extends Block, item.Item {
        growth: number;
    }
    function beetrootSeeds(growth: number): BeetrootSeeds;
    interface Blackstone extends Block, item.Item {
        type: BlackstoneType;
    }
    function blackstone(type: BlackstoneType): Blackstone;
    interface BlastFurnace extends Block, item.Item {
        facing: cube.Direction;
        lit: boolean;
    }
    function blastFurnace(facing: cube.Direction, lit?: boolean): BlastFurnace;
    interface BlueIce extends Block, item.Item {}
    function blueIce(): BlueIce;
    interface Bone extends Block, item.Item {
        axis: cube.Axis;
    }
    function bone(axis: cube.Axis): Bone;
    interface Bookshelf extends Block, item.Item {}
    function bookshelf(): Bookshelf;
    interface BrewingStand extends Block, item.Item {
        left: boolean;
        middle: boolean;
        right: boolean;
    }
    function brewingStand(left?: boolean, middle?: boolean, right?: boolean): BrewingStand;
    interface Bricks extends Block, item.Item {}
    function bricks(): Bricks;
    interface Cactus extends Block, item.Item {
        age: number;
    }
    function cactus(age: number): Cactus;
    interface Cake extends Block, item.Item {
        bites: number;
    }
    function cake(bites: number): Cake;
    interface Calcite extends Block, item.Item {}
    function calcite(): Calcite;
    interface Campfire extends Block, item.Item {
        type: FireType;
        items: CampfireItem;
        facing: cube.Direction;
        extinguished: boolean;
    }
    function campfire(type: FireType, items: CampfireItem, facing: cube.Direction, extinguished?: boolean): Campfire;
    interface Carpet extends Block, item.Item {
        colour: item.Colour;
    }
    function carpet(colour: item.Colour): Carpet;
    interface Carrot extends Block, item.Item {
        growth: number;
    }
    function carrot(growth: number): Carrot;
    interface Chain extends Block, item.Item {
        axis: cube.Axis;
    }
    function chain(axis: cube.Axis): Chain;
    interface Chest extends Block, item.Item {
        facing: cube.Direction
        customName: string
    }
    function chest(facing: cube.Direction, customName?: string): Chest;
    interface ChiseledQuartz extends Block, item.Item {}
    function chiseledQuartz(): ChiseledQuartz;
    interface Clay extends Block, item.Item {}
    function clay(): Clay;
    interface Coal extends Block, item.Item {}
    function coal(): Coal;
    interface CoalOre extends Block, item.Item {
        type: OreType;
    }
    function coalOre(type: OreType): CoalOre;
    interface Cobblestone extends Block, item.Item {
        mossy: boolean;
    }
    function cobblestone(mossy?: boolean): Cobblestone;
    interface CocoaBean extends Block, item.Item {
        facing: cube.Direction;
        age: number;
    }
    function cocoaBean(facing: cube.Direction, age: number): CocoaBean;
    interface Composter extends Block, item.Item {
        level: number;
    }
    function composter(level: number): Composter;
    interface Concrete extends Block, item.Item {
        colour: item.Colour;
    }
    function concrete(colour: item.Colour): Concrete;
    interface ConcretePowder extends Block, item.Item {
        colour: item.Colour;
    }
    function concretePowder(colour: item.Colour): ConcretePowder;
    interface Copper extends Block, item.Item {
        type: CopperType;
        oxidation: Oxidation;
        waxed: boolean;
    }
    function copper(type: CopperType, oxidation: Oxidation, waxed?: boolean): Copper;
    interface CopperDoor extends Block, item.Item {
        facing: cube.Direction;
        oxidation: Oxidation;
        waxed: boolean;
        open: boolean;
        top: boolean;
        right: boolean;
    }
    function copperDoor(facing: cube.Direction, oxidation: Oxidation, waxed?: boolean, open?: boolean, top?: boolean, right?: boolean): CopperDoor;
    interface CopperGrate extends Block, item.Item {
        oxidation: Oxidation;
        waxed: boolean;
    }
    function copperGrate(oxidation: Oxidation, waxed?: boolean): CopperGrate;
    interface CopperOre extends Block, item.Item {
        type: OreType;
    }
    function copperOre(type: OreType): CopperOre;
    interface CopperTrapdoor extends Block, item.Item {
        facing: cube.Direction;
        oxidation: Oxidation;
        waxed: boolean;
        open: boolean;
        top: boolean;
    }
    function copperTrapdoor(facing: cube.Direction, oxidation: Oxidation, waxed?: boolean, open?: boolean, top?: boolean): CopperTrapdoor;
    interface Coral extends Block, item.Item {
        type: CoralType;
        dead: boolean;
    }
    function coral(type: CoralType, dead?: boolean): Coral;
    interface CoralBlock extends Block, item.Item {
        type: CoralType;
        dead: boolean;
    }
    function coralBlock(type: CoralType, dead?: boolean): CoralBlock;
    interface CraftingTable extends Block, item.Item {}
    function craftingTable(): CraftingTable;
    interface DeadBush extends Block, item.Item {}
    function deadBush(): DeadBush;
    interface DecoratedPot extends Block, item.Item {
        item: item.Stack;
        facing: cube.Direction;
        decorations: item.Stack[];
    }
    function decoratedPot(item: item.Stack, facing: cube.Direction, decorations: item.Stack[]): DecoratedPot;
    interface Deepslate extends Block, item.Item {
        type: DeepslateType;
        axis: cube.Axis;
    }
    function deepslate(type: DeepslateType, axis: cube.Axis): Deepslate;
    interface DeepslateBricks extends Block, item.Item {
        cracked: boolean;
    }
    function deepslateBricks(cracked?: boolean): DeepslateBricks;
    interface DeepslateTiles extends Block, item.Item {
        cracked: boolean;
    }
    function deepslateTiles(cracked?: boolean): DeepslateTiles;
    interface Diamond extends Block, item.Item {}
    function diamond(): Diamond;
    interface DiamondOre extends Block, item.Item {
        type: OreType;
    }
    function diamondOre(type: OreType): DiamondOre;
    interface Dirt extends Block, item.Item {
        coarse: boolean;
    }
    function dirt(coarse?: boolean): Dirt;
    interface DirtPath extends Block, item.Item {}
    function dirtPath(): DirtPath;
    interface DoubleFlower extends Block, item.Item {
        type: DoubleFlowerType;
        upperPart: boolean;
    }
    function doubleFlower(type: DoubleFlowerType, upperPart?: boolean): DoubleFlower;
    interface DoubleTallGrass extends Block, item.Item {
        type: DoubleTallGrassType;
        upperPart: boolean;
    }
    function doubleTallGrass(type: DoubleTallGrassType, upperPart?: boolean): DoubleTallGrass;
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
    function emeraldOre(type: OreType): EmeraldOre;
    interface EnchantingTable extends Block, item.Item {}
    function enchantingTable(): EnchantingTable;
    interface EndBricks extends Block, item.Item {}
    function endBricks(): EndBricks;
    interface EndRod extends Block, item.Item {
        facing: cube.Face
    }
    function endRod(facing: cube.Face): EndRod;
    interface EndStone extends Block, item.Item {}
    function endStone(): EndStone;
    interface EnderChest extends Block, item.Item {
        facing: cube.Direction
    }
    function enderChest(facing: cube.Direction): EnderChest;
    interface Farmland extends Block, item.Item {
        hydration: number;
    }
    function farmland(hydration: number): Farmland;
    interface Fern extends Block, item.Item {}
    function fern(): Fern;
    interface Fire extends Block, item.Item {
        type: FireType;
        age: number;
    }
    function fire(type: FireType, age: number): Fire;
    interface FletchingTable extends Block, item.Item {}
    function fletchingTable(): FletchingTable;
    interface Flower extends Block, item.Item {
        type: FlowerType;
    }
    function flower(type: FlowerType): Flower;
    interface Froglight extends Block, item.Item {
        type: FroglightType;
    }
    function froglight(type: FroglightType): Froglight;
    interface Furnace extends Block, item.Item {
        facing: cube.Direction;
        lit: boolean;
    }
    function furnace(facing: cube.Direction, lit?: boolean): Furnace;
    interface Glass extends Block, item.Item {}
    function glass(): Glass;
    interface GlassPane extends Block, item.Item {}
    function glassPane(): GlassPane;
    interface GlazedTerracotta extends Block, item.Item {
        colour: item.Colour;
    }
    function glazedTerracotta(colour: item.Colour): GlazedTerracotta;
    interface Glowstone extends Block, item.Item {}
    function glowstone(): Glowstone;
    interface Gold extends Block, item.Item {}
    function gold(): Gold;
    interface GoldOre extends Block, item.Item {
        type: OreType;
    }
    function goldOre(type: OreType): GoldOre;
    interface Grass extends Block, item.Item {}
    function grass(): Grass;
    interface Gravel extends Block, item.Item {}
    function gravel(): Gravel;
    interface Grindstone extends Block, item.Item {
        facing: cube.Direction;
        attach: GrindstoneAttachment;
    }
    function grindstone(facing: cube.Direction, attach: GrindstoneAttachment): Grindstone;
    interface HayBale extends Block, item.Item {
        axis: cube.Axis;
    }
    function hayBale(axis: cube.Axis): HayBale;
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
    function hopper(facing: cube.Face, powered?: boolean, customName?: string, lastTick?: number, transferCooldown?: number, collectCooldown?: number): Hopper;
    interface InvisibleBedrock extends Block, item.Item {}
    function invisibleBedrock(): InvisibleBedrock;
    interface Iron extends Block, item.Item {}
    function iron(): Iron;
    interface IronBars extends Block, item.Item {}
    function ironBars(): IronBars;
    interface IronOre extends Block, item.Item {
        type: OreType;
    }
    function ironOre(type: OreType): IronOre;
    interface ItemFrame extends Block, item.Item {
        facing: cube.Face;
        item: item.Stack;
        rotations: number;
        dropChance: number;
        glowing: boolean;
    }
    function itemFrame(facing: cube.Face, item: item.Stack, rotations: number, dropChance: number, glowing?: boolean): ItemFrame;
    interface Jukebox extends Block, item.Item {
        item: item.Stack;
    }
    function jukebox(item: item.Stack): Jukebox;
    interface Kelp extends Block, item.Item {
        age: number;
    }
    function kelp(age: number): Kelp;
    interface Ladder extends Block, item.Item {
        facing: cube.Direction;
    }
    function ladder(facing: cube.Direction): Ladder;
    interface Lantern extends Block, item.Item {
        type: FireType;
        hanging: boolean;
    }
    function lantern(type: FireType, hanging?: boolean): Lantern;
    interface Lapis extends Block, item.Item {}
    function lapis(): Lapis;
    interface LapisOre extends Block, item.Item {
        type: OreType;
    }
    function lapisOre(type: OreType): LapisOre;
    interface Lava extends Block, Liquid {
        depth: number;
        still: boolean;
        falling: boolean;
    }
    function lava(depth: number, still?: boolean, falling?: boolean): Lava;
    interface Leaves extends Block, item.Item {
        wood: WoodType;
        persistent: boolean;
        shouldUpdate: boolean;
    }
    function leaves(wood: WoodType, persistent?: boolean, shouldUpdate?: boolean): Leaves;
    interface Lectern extends Block, item.Item {
        facing: cube.Direction;
        book: item.Stack;
        page: number;
    }
    function lectern(facing: cube.Direction, book: item.Stack, page: number): Lectern;
    interface Light extends Block, item.Item {
        level: number;
    }
    function light(level: number): Light;
    interface LitPumpkin extends Block, item.Item {
        facing: cube.Direction;
    }
    function litPumpkin(facing: cube.Direction): LitPumpkin;
    interface Log extends Block, item.Item {
        wood: WoodType;
        axis: cube.Axis;
        stripped: boolean;
    }
    function log(wood: WoodType, axis: cube.Axis, stripped?: boolean): Log;
    interface Loom extends Block, item.Item {
        facing: cube.Direction;
    }
    function loom(facing: cube.Direction): Loom;
    interface Melon extends Block, item.Item {}
    function melon(): Melon;
    interface MelonSeeds extends Block, item.Item {
        direction: cube.Face;
        growth: number;
    }
    function melonSeeds(direction: cube.Face, growth: number): MelonSeeds;
    interface MossCarpet extends Block, item.Item {}
    function mossCarpet(): MossCarpet;
    interface Mud extends Block, item.Item {}
    function mud(): Mud;
    interface MudBricks extends Block, item.Item {}
    function mudBricks(): MudBricks;
    interface MuddyMangroveRoots extends Block, item.Item {
        axis: cube.Axis;
    }
    function muddyMangroveRoots(axis: cube.Axis): MuddyMangroveRoots;
    interface NetherBrickFence extends Block, item.Item {}
    function netherBrickFence(): NetherBrickFence;
    interface NetherBricks extends Block, item.Item {
        type: NetherBricksType;
    }
    function netherBricks(type: NetherBricksType): NetherBricks;
    interface NetherGoldOre extends Block, item.Item {
        type: OreType;
    }
    function netherGoldOre(type: OreType): NetherGoldOre;
    interface NetherQuartzOre extends Block, item.Item {
        type: OreType;
    }
    function netherQuartzOre(type: OreType): NetherQuartzOre;
    interface NetherSprouts extends Block, item.Item {}
    function netherSprouts(): NetherSprouts;
    interface NetherWart extends Block, item.Item {
        age: number;
    }
    function netherWart(age: number): NetherWart;
    interface NetherWartBlock extends Block, item.Item {
        warped: boolean;
    }
    function netherWartBlock(warped?: boolean): NetherWartBlock;
    interface Netherite extends Block, item.Item {}
    function netherite(): Netherite;
    interface Netherrack extends Block, item.Item {}
    function netherrack(): Netherrack;
    interface Note extends Block, item.Item {
        pitch: number;
    }
    function note(pitch: number): Note;
    interface Obsidian extends Block, item.Item {
        crying: boolean;
    }
    function obsidian(crying?: boolean): Obsidian;
    interface PackedIce extends Block, item.Item {}
    function packedIce(): PackedIce;
    interface PackedMud extends Block, item.Item {}
    function packedMud(): PackedMud;
    interface PinkPetals extends Block, item.Item {
        facing: cube.Direction;
        additionalCount: number;
    }
    function pinkPetals(facing: cube.Direction, additionalCount: number): PinkPetals;
    interface Planks extends Block, item.Item {
        wood: WoodType;
    }
    function planks(wood: WoodType): Planks;
    interface Podzol extends Block, item.Item {}
    function podzol(): Podzol;
    interface PolishedBlackstoneBrick extends Block, item.Item {
        cracked: boolean;
    }
    function polishedBlackstoneBrick(cracked?: boolean): PolishedBlackstoneBrick;
    interface PolishedTuff extends Block, item.Item {}
    function polishedTuff(): PolishedTuff;
    interface Potato extends Block, item.Item {
        growth: number;
    }
    function potato(growth: number): Potato;
    interface Prismarine extends Block, item.Item {
        type: PrismarineType;
    }
    function prismarine(type: PrismarineType): Prismarine;
    interface Pumpkin extends Block, item.Item {
        facing: cube.Direction;
    }
    function pumpkin(facing: cube.Direction): Pumpkin;
    interface PumpkinSeeds extends Block, item.Item {
        direction: cube.Face;
        growth: number;
    }
    function pumpkinSeeds(direction: cube.Face, growth: number): PumpkinSeeds;
    interface Purpur extends Block, item.Item {}
    function purpur(): Purpur;
    interface PurpurPillar extends Block, item.Item {
        axis: cube.Axis;
    }
    function purpurPillar(axis: cube.Axis): PurpurPillar;
    interface Quartz extends Block, item.Item {
        smooth: boolean;
    }
    function quartz(smooth?: boolean): Quartz;
    interface QuartzBricks extends Block, item.Item {}
    function quartzBricks(): QuartzBricks;
    interface QuartzPillar extends Block, item.Item {
        axis: cube.Axis;
    }
    function quartzPillar(axis: cube.Axis): QuartzPillar;
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
    function resinBricks(chiseled?: boolean): ResinBricks;
    interface Sand extends Block, item.Item {
        red: boolean;
    }
    function sand(red?: boolean): Sand;
    interface Sandstone extends Block, item.Item {
        type: SandstoneType;
        red: boolean;
    }
    function sandstone(type: SandstoneType, red?: boolean): Sandstone;
    interface SeaLantern extends Block, item.Item {}
    function seaLantern(): SeaLantern;
    interface SeaPickle extends Block, item.Item {
        dead: boolean;
    }
    function seaPickle(dead?: boolean): SeaPickle;
    interface ShortGrass extends Block, item.Item {
        double: boolean;
    }
    function shortGrass(double?: boolean): ShortGrass;
    interface Shroomlight extends Block, item.Item {}
    function shroomlight(): Shroomlight;
    interface Sign extends Block, item.Item {
        wood: WoodType;
        attach: Attachment;
        front: SignText;
        back: SignText;
        waxed: boolean;
    }
    function sign(wood: WoodType, attach: Attachment, front: SignText, back: SignText, waxed?: boolean): Sign;
    interface Skull extends Block, item.Item {
        type: SkullType;
        attach: Attachment;
    }
    function skull(type: SkullType, attach: Attachment): Skull;
    interface Slab extends Block, item.Item {
        block: Block;
        double: boolean;
        top: boolean;
    }
    function slab(block: Block, double?: boolean, top?: boolean): Slab;
    interface SmithingTable extends Block, item.Item {}
    function smithingTable(): SmithingTable;
    interface Smoker extends Block, item.Item {
        facing: cube.Direction;
        lit: boolean;
    }
    function smoker(facing: cube.Direction, lit?: boolean): Smoker;
    interface Snow extends Block, item.Item {}
    function snow(): Snow;
    interface SoulSand extends Block, item.Item {}
    function soulSand(): SoulSand;
    interface SoulSoil extends Block, item.Item {}
    function soulSoil(): SoulSoil;
    interface Sponge extends Block, item.Item {
        wet: boolean;
    }
    function sponge(wet?: boolean): Sponge;
    interface SporeBlossom extends Block, item.Item {}
    function sporeBlossom(): SporeBlossom;
    interface StainedGlass extends Block, item.Item {
        colour: item.Colour;
    }
    function stainedGlass(colour: item.Colour): StainedGlass;
    interface StainedGlassPane extends Block, item.Item {
        colour: item.Colour;
    }
    function stainedGlassPane(colour: item.Colour): StainedGlassPane;
    interface StainedTerracotta extends Block, item.Item {
        colour: item.Colour;
    }
    function stainedTerracotta(colour: item.Colour): StainedTerracotta;
    interface Stairs extends Block, item.Item {
        block: Block;
        facing: cube.Direction;
        upsideDown: boolean;
    }
    function stairs(block: Block, facing: cube.Direction, upsideDown?: boolean): Stairs;
    interface Stone extends Block, item.Item {
        smooth: boolean;
    }
    function stone(smooth?: boolean): Stone;
    interface StoneBricks extends Block, item.Item {
        type: StoneBricksType;
    }
    function stoneBricks(type: StoneBricksType): StoneBricks;
    interface Stonecutter extends Block, item.Item {
        facing: cube.Direction
    }
    function stonecutter(facing: cube.Direction): Stonecutter;
    interface StopCrackAction extends Block, item.Item {}
    function stopCrackAction(): StopCrackAction;
    interface SugarCane extends Block, item.Item {
        age: number;
    }
    function sugarCane(age: number): SugarCane;
    interface Tnt extends Block, item.Item {}
    function tnt(): Tnt;
    interface Terracotta extends Block, item.Item {}
    function terracotta(): Terracotta;
    interface Torch extends Block, item.Item {
        facing: cube.Face;
        type: FireType;
    }
    function torch(facing: cube.Face, type: FireType): Torch;
    interface Tuff extends Block, item.Item {
        chiseled: boolean;
    }
    function tuff(chiseled?: boolean): Tuff;
    interface TuffBricks extends Block, item.Item {
        chiseled: boolean;
    }
    function tuffBricks(chiseled?: boolean): TuffBricks;
    interface Vines extends Block, item.Item {
        northDirection: boolean;
        eastDirection: boolean;
        southDirection: boolean;
        westDirection: boolean;
    }
    function vines(north?: boolean, east?: boolean, south?: boolean, west?: boolean): Vines;
    interface Wall extends Block, item.Item {
        block: Block;
        northConnection: WallConnection;
        eastConnection: WallConnection;
        southConnection: WallConnection;
        westConnection: WallConnection;
        post: boolean;
    }
    function wall(block: Block, north: WallConnection, east: WallConnection, south: WallConnection, west: WallConnection, post?: boolean): Wall;
    interface Water extends Block, Liquid {
        depth: number;
        still: boolean;
        falling: boolean;
    }
    function water(depth: number, still?: boolean, falling?: boolean): Water;
    interface WheatSeeds extends Block, item.Item {
        growth: number;
    }
    function wheatSeeds(growth: number): WheatSeeds;
    interface Wood extends Block, item.Item {
        wood: WoodType;
        axis: cube.Axis;
        stripped: boolean;
    }
    function wood(wood: WoodType, axis: cube.Axis, stripped?: boolean): Wood;
    interface WoodDoor extends Block, item.Item {
        wood: WoodType;
        facing: cube.Direction;
        open: boolean;
        top: boolean;
        right: boolean;
    }
    function woodDoor(wood: WoodType, facing: cube.Direction, open?: boolean, top?: boolean, right?: boolean): WoodDoor;
    interface WoodFence extends Block, item.Item {
        wood: WoodType;
    }
    function woodFence(wood: WoodType): WoodFence;
    interface WoodFenceGate extends Block, item.Item {
        wood: WoodType;
        facing: cube.Direction;
        open: boolean;
        lowered: boolean;
    }
    function woodFenceGate(wood: WoodType, facing: cube.Direction, open?: boolean, lowered?: boolean): WoodFenceGate;
    interface WoodTrapdoor extends Block, item.Item {
        wood: WoodType;
        facing: cube.Direction;
        open: boolean;
        top: boolean;
    }
    function woodTrapdoor(wood: WoodType, facing: cube.Direction, open?: boolean, top ?: boolean): WoodTrapdoor;
    interface Wool extends Block, item.Item {
        colour: item.Colour;
    }
    function wool(colour: item.Colour): Wool;
}