/// <reference path="block.d.ts" />
/// <reference path="cube.d.ts" />
/// <reference path="mgl64.d.ts" />
/// <reference path="sound.d.ts" />
/// <reference path="world.d.ts" />

declare namespace particle {
    interface Particle {
        spawn(world: world.World, pos: mgl64.Vec3): void;
    }

    interface BlockBreak extends Particle {
        block: block.Block;
    }
    function blockBreak(block: block.Block): BlockBreak;
    interface BlockForceField extends Particle {}
    function blockForceField(): BlockForceField;
    interface BoneMeal extends Particle {}
    function boneMeal(): BoneMeal;
    interface DragonEggTeleport extends Particle {
        diff: cube.Pos;
    }
    function dragonEggTeleport(diff: cube.Pos): DragonEggTeleport;
    interface Dust extends Particle {
        // TODO: colour
    }
    function dust(): Dust;
    interface DustPlume extends Particle {}
    function dustPlume(): DustPlume;
    interface Effect extends Particle {
        // TODO: colour
    }
    function effect(): Effect;
    interface EggSmash extends Particle {}
    function eggSmash(): EggSmash;
    interface EndermanTeleport extends Particle {}
    function endermanTeleport(): EndermanTeleport;
    interface EntityFlame extends Particle {}
    function entityFlame(): EntityFlame;
    interface Evaporate extends Particle {}
    function evaporate(): Evaporate;
    interface Flame extends Particle {
        // TODO: colour
    }
    function flame(): Flame;
    interface HugeExplosion extends Particle {}
    function hugeExplosion(): HugeExplosion;
    interface Lava extends Particle {}
    function lava(): Lava;
    interface LavaDrip extends Particle {}
    function lavaDrip(): LavaDrip;
    interface Note extends Particle {
        instrument: sound.InstrumentType;
        pitch: number;
    }
    function note(instrument: sound.InstrumentType, pitch: number): Note;
    interface PunchBlock extends Particle {
        block: block.Block;
        face: cube.Face;
    }
    function punchBlock(block: block.Block, face: cube.Face): PunchBlock;
    interface SnowballPoof extends Particle {}
    function snowballPoof(): SnowballPoof;
    interface Splash extends Particle {
        // TODO: colour
    }
    function splash(): Splash;
    interface WaterDrip extends Particle {}
    function waterDrip(): WaterDrip;
}