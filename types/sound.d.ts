/// <reference path="block.d.ts" />
/// <reference path="item.d.ts" />
/// <reference path="mgl64.d.ts" />
/// <reference path="world.d.ts" />

declare namespace sound {
    interface Sound {
        play(world: world.World, pos: mgl64.Vec3): void;
    }

    class InstrumentType {
        static readonly Banjo: InstrumentType;
        static readonly Bass: InstrumentType;
        static readonly BassDrum: InstrumentType;
        static readonly Bell: InstrumentType;
        static readonly Bit: InstrumentType;
        static readonly Chimes: InstrumentType;
        static readonly ClicksAndSticks: InstrumentType;
        static readonly CowBell: InstrumentType;
        static readonly Didgeridoo: InstrumentType;
        static readonly Flute: InstrumentType;
        static readonly Guitar: InstrumentType;
        static readonly IronXylophone: InstrumentType;
        static readonly Piano: InstrumentType;
        static readonly Pling: InstrumentType;
        static readonly Snare: InstrumentType;
        static readonly Xylophone: InstrumentType;

        int32(): number;
    }

    class DiscType {
        static readonly No13: DiscType;
        static readonly Cat: DiscType;
        static readonly Blocks: DiscType;
        static readonly Chirp: DiscType;
        static readonly Far: DiscType;
        static readonly Mall: DiscType;
        static readonly Mellohi: DiscType;
        static readonly Stal: DiscType;
        static readonly Strad: DiscType;
        static readonly Ward: DiscType;
        static readonly No11: DiscType;
        static readonly Wait: DiscType;
        static readonly Otherside: DiscType;
        static readonly Pigstep: DiscType;
        static readonly No5: DiscType;
        static readonly Relic: DiscType;
        static readonly Creator: DiscType;
        static readonly CreatorMusicBox: DiscType;
        static readonly Precipice: DiscType;

        uint8(): number;
        string(): string;
        displayName(): string;
        author(): string
    }

    class HornType {
        static readonly Ponder: HornType;
        static readonly Sing: HornType;
        static readonly Seek: HornType;
        static readonly Feel: HornType;
        static readonly Admire: HornType;
        static readonly Call: HornType;
        static readonly Yearn: HornType;
        static readonly Dream: HornType;

        uint8(): number;
        name(): string;
    }

    class CrossbowLoading {
        static readonly Start: CrossbowLoading;
        static readonly Middle: CrossbowLoading;
        static readonly End: CrossbowLoading;
    }

    interface AnvilBreak extends Sound {}
    function anvilBreak(): AnvilBreak;
    interface AnvilLand extends Sound {}
    function anvilLand(): AnvilLand;
    interface AnvilUse extends Sound {}
    function anvilUse(): AnvilUse;
    interface ArrowHit extends Sound {}
    function arrowHit(): ArrowHit;
    interface Attack extends Sound {
        damage: boolean;
    }
    function attack(damage?: boolean): Attack;
    interface BarrelClose extends Sound {}
    function barrelClose(): BarrelClose;
    interface BarrelOpen extends Sound {}
    function barrelOpen(): BarrelOpen;
    interface BlastFurnaceCrackle extends Sound {}
    function blastFurnaceCrackle(): BlastFurnaceCrackle;
    interface BlockBreaking extends Sound {
        block: block.Block;
    }
    function blockBreaking(block: block.Block): BlockBreaking;
    interface BlockPlace extends Sound {
        block: block.Block;
    }
    function blockPlace(block: block.Block): BlockPlace;
    interface BowShoot extends Sound {}
    function bowShoot(): BowShoot;
    interface BucketEmpty extends Sound {
        liquid: block.Liquid;
    }
    function bucketEmpty(liquid: block.Liquid): BucketEmpty;
    interface BucketFill extends Sound {
        liquid: block.Liquid;
    }
    function bucketFill(liquid: block.Liquid): BucketFill;
    interface Burning extends Sound {}
    function burning(): Burning;
    interface Burp extends Sound {}
    function burp(): Burp;
    interface CampfireCrackle extends Sound {}
    function campfireCrackle(): CampfireCrackle;
    interface ChestClose extends Sound {}
    function chestClose(): ChestClose;
    interface ChestOpen extends Sound {}
    function chestOpen(): ChestOpen;
    interface Click extends Sound {}
    function click(): Click;
    interface ComposterEmpty extends Sound {}
    function composterEmpty(): ComposterEmpty;
    interface ComposterFill extends Sound {}
    function composterFill(): ComposterFill;
    interface ComposterFillLayer extends Sound {}
    function composterFillLayer(): ComposterFillLayer;
    interface ComposterReady extends Sound {}
    function composterReady(): ComposterReady;
    interface CopperScraped extends Sound {}
    function copperScraped(): CopperScraped;
    interface CrossbowLoad extends Sound {
        stage: CrossbowLoading;
    }
    function crossbowLoad(stage: CrossbowLoading): CrossbowLoad;
    interface CrossbowShoot extends Sound {}
    function crossbowShoot(): CrossbowShoot;
    interface DecoratedPotInsertFailed extends Sound {}
    function decoratedPotInsertFailed(): DecoratedPotInsertFailed;
    interface DecoratedPotInserted extends Sound {
        progress: number;
    }
    function decoratedPotInserted(progress: number): DecoratedPotInserted;
    interface Deny extends Sound {}
    function deny(): Deny;
    interface DoorClose extends Sound {
        block: block.Block;
    }
    function doorClose(block: block.Block): DoorClose;
    interface DoorCrash extends Sound {}
    function doorCrash(): DoorCrash;
    interface DoorOpen extends Sound {
        block: block.Block;
    }
    function doorOpen(block: block.Block): DoorOpen;
    interface Drowning extends Sound {}
    function drowning(): Drowning;
    interface EnderChestClose extends Sound {}
    function enderChestClose(): EnderChestClose;
    interface EnderChestOpen extends Sound {}
    function enderChestOpen(): EnderChestOpen;
    interface EquipItem extends Sound {
        item: item.Item;
    }
    function equipItem(item: item.Item): EquipItem;
    interface Experience extends Sound {}
    function experience(): Experience;
    interface Explosion extends Sound {}
    function explosion(): Explosion;
    interface Fall extends Sound {
        distance: number;
    }
    function fall(): Fall;
    interface FenceGateClose extends Sound {
        block: block.Block;
    }
    function fenceGateClose(block: block.Block): FenceGateClose;
    interface FenceGateOpen extends Sound {
        block: block.Block;
    }
    function fenceGateOpen(block: block.Block): FenceGateOpen;
    interface FireCharge extends Sound {}
    function fireCharge(): FireCharge;
    interface FireExtinguish extends Sound {}
    function fireExtinguish(): FireExtinguish;
    interface FireworkBlast extends Sound {}
    function fireworkBlast(): FireworkBlast;
    interface FireworkHugeBlast extends Sound {}
    function fireworkHugeBlast(): FireworkHugeBlast;
    interface FireworkLaunch extends Sound {}
    function fireworkLaunch(): FireworkLaunch;
    interface FireworkTwinkle extends Sound {}
    function fireworkTwinkle(): FireworkTwinkle;
    interface Fizz extends Sound {}
    function fizz(): Fizz;
    interface FurnaceCrackle extends Sound {}
    function furnaceCrackle(): FurnaceCrackle;
    interface GhastShoot extends Sound {}
    function ghastShoot(): GhastShoot;
    interface GhastWarning extends Sound {}
    function ghastWarning(): GhastWarning;
    interface GlassBreak extends Sound {}
    function glassBreak(): GlassBreak;
    interface GoatHorn extends Sound {
        horn: HornType;
    }
    function goatHorn(horn: HornType): GoatHorn;
    interface Ignite extends Sound {}
    function ignite(): Ignite;
    interface Instrument extends Sound {}
    function instrument(): Instrument;
    interface ItemAdd extends Sound {}
    function itemAdd(): ItemAdd;
    interface ItemBreak extends Sound {}
    function itemBreak(): ItemBreak;
    interface ItemFrameRemove extends Sound {}
    function itemFrameRemove(): ItemFrameRemove;
    interface ItemFrameRotate extends Sound {}
    function itemFrameRotate(): ItemFrameRotate;
    interface ItemThrow extends Sound {}
    function itemThrow(): ItemThrow;
    interface ItemUseOn extends Sound {
        block: block.Block;
    }
    function itemUseOn(block: block.Block): ItemUseOn;
    interface LecternBookPlace extends Sound {}
    function lecternBookPlace(): LecternBookPlace;
    interface LevelUp extends Sound {}
    function levelUp(): LevelUp;
    interface MusicDiscEnd extends Sound {}
    function musicDiscEnd(): MusicDiscEnd;
    interface MusicDiscPlay extends Sound {
        discType: DiscType;
    }
    function musicDiscPlay(discType: DiscType): MusicDiscPlay;
    interface Note extends Sound {
        instrument: InstrumentType;
        pitch: number;
    }
    function note(instrument: Instrument, pitch: number): Note;
    interface Pop extends Sound {}
    function pop(): Pop;
    interface PotionBrewed extends Sound {}
    function potionBrewed(): PotionBrewed;
    interface SignWaxed extends Sound {}
    function signWaxed(): SignWaxed;
    interface SmokerCrackle extends Sound {}
    function smokerCrackle(): SmokerCrackle;
    interface StopUsingSpyglass extends Sound {}
    function stopUsingSpyglass(): StopUsingSpyglass;
    interface Tnt extends Sound {}
    function tnt(): Tnt;
    interface Teleport extends Sound {}
    function teleport(): Teleport;
    interface Thunder extends Sound {}
    function thunder(): Thunder;
    interface Totem extends Sound {}
    function totem(): Totem;
    interface TrapdoorClose extends Sound {
        block: block.Block;
    }
    function trapdoorClose(block: block.Block): TrapdoorClose;
    interface TrapdoorOpen extends Sound {
        block: block.Block;
    }
    function trapdoorOpen(block: block.Block): TrapdoorOpen;
    interface UseSpyglass extends Sound {}
    function useSpyglass(): UseSpyglass;
    interface WaxRemoved extends Sound {}
    function waxRemoved(): WaxRemoved;
    interface WaxedSignFailedInteraction extends Sound {}
    function waxedSignFailedInteraction(): WaxedSignFailedInteraction;
}