/// <reference path="cube.d.ts" />
/// <reference path="world.d.ts" />

declare namespace model {
    interface Model {
        bbox(pos: cube.Pos, source: world.BlockSource): cube.BBox[];
        faceSolid(pos: cube.Pos, face: cube.Face, source: world.BlockSource): boolean;
    }

    interface Anvil extends Model {
        facing: cube.Direction;
    }
    function anvil(facing: cube.Direction): Anvil;
    interface BrewingStand extends Model {}
    function brewingStand(): BrewingStand;
    interface Cactus extends Model {}
    function cactus(): Cactus;
    interface Cake extends Model {
        bites: number;
    }
    function cake(bites: number): Cake;
    interface Campfire extends Model {}
    function campfire(): Campfire;
    interface Carpet extends Model {}
    function carpet(): Carpet;
    interface Chain extends Model {
        axis: cube.Axis;
    }
    function chain(axis: cube.Axis): Chain;
    interface Chest extends Model {}
    function chest(): Chest;
    interface CocoaBean extends Model {
        facing: cube.Direction;
        age: number;
    }
    function cocoaBean(facing: cube.Direction, age: number): CocoaBean;
    interface Composter extends Model {
        level: number;
    }
    function composter(level: number): Composter;
    interface DecoratedPot extends Model {}
    function decoratedPot(): DecoratedPot;
    interface Door extends Model {
        facing: cube.Direction;
        open: boolean;
        right: boolean;
    }
    function door(facing: cube.Direction, open?: boolean, right?: boolean): Door;
    interface Empty extends Model {}
    function empty(): Empty;
    interface EnchantingTable extends Model {}
    function enchantingTable(): EnchantingTable;
    interface EndRod extends Model {
        axis: cube.Axis;
    }
    function endRod(axis: cube.Axis): EndRod;
    interface Fence extends Model {
        wood: boolean;
    }
    function fence(wood?: boolean): Fence;
    interface FenceGate extends Model {
        facing: cube.Direction;
        open: boolean;
    }
    function fenceGate(facing: cube.Direction, open?: boolean): FenceGate;
    interface Grindstone extends Model {
        axis: cube.Axis;
    }
    function grindstone(axis: cube.Axis): Grindstone;
    interface Hopper extends Model {}
    function hopper(): Hopper;
    interface Ladder extends Model {
        facing: cube.Direction;
    }
    function ladder(facing: cube.Direction): Ladder;
    interface Lantern extends Model {
        hanging: boolean;
    }
    function lantern(hanging?: boolean): Lantern;
    interface Leaves extends Model {}
    function leaves(): Leaves
    interface Lectern extends Model {}
    function lectern(): Lectern;
    interface Skull extends Model {
        direction: cube.Direction;
        hanging: boolean;
    }
    function skull(direction: cube.Direction, hanging?: boolean): Skull;
    interface Solid extends Model {}
    function solid(): Solid;
    interface Stair extends Model {
        facing: cube.Direction;
        upsideDown: boolean;
    }
    function stair(facing: cube.Direction, upsideDown?: boolean): Stair;
    interface Stonecutter extends Model {}
    function stonecutter(): Stonecutter;
    interface Thin extends Model {}
    function thin(): Thin;
    interface TilledGrass extends Model {}
    function tilledGrass(): TilledGrass;
    interface Trapdoor extends Model {
        facing: cube.Direction;
        open: boolean;
        top: boolean;
    }
    function trapdoor(facing: cube.Direction, open?: boolean, top?: boolean): Trapdoor;
    interface Wall extends Model {
        northConnection: number;
        eastConnection: number;
        southConnection: number;
        westConnection: number;
        post: boolean;
    }
    function wall(northConnection: number, eastConnection: number, southConnection: number, westConnection: number, post?: boolean): Wall;
}