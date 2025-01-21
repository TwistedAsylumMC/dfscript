/// <reference path="biome.d.ts" />
/// <reference path="block.d.ts" />
/// <reference path="cube.d.ts" />
/// <reference path="generator.d.ts" />
/// <reference path="item.d.ts" />

declare namespace world {
    interface Config {
        // TODO: log
        dim: Dimension;
        portalDestination: (dim: Dimension) => World;
        // TODO: provider
        generator: generator.Generator;
        readOnly: boolean;
        saveInterval: Date;
        randomTickSpeed: number;
        // TODO: randSource
        // TODO: entities
    }
    function config(): Config;

    interface BlockSource {
        block(pos: cube.Pos): block.Block;
    }
    interface World extends BlockSource {}
    function create(config?: Config): World;

    interface Viewer {}

    interface Difficulty {
        foodRegenerates(): boolean;
        starvationHealthLimit(): number;
        fireSpreadIncrease(): number;
    }
    const difficulty: {
        peaceful: Difficulty;
        easy: Difficulty;
        normal: Difficulty;
        hard: Difficulty;
    }
    function difficultyId(difficulty: Difficulty): number | null;
    function difficultyById(id: number): Difficulty | null;

    interface GameMode {
        allowsEditing(): boolean;
        allowsTakingDamage(): boolean;
        creativeInventory(): boolean;
        hasCollision(): boolean;
        allowsFlying(): boolean;
        allowsInteraction(): boolean;
        visible(): boolean;
    }
    const gameMode: {
        survival: GameMode;
        creative: GameMode;
        adventure: GameMode;
        spectator: GameMode;
    }
    function gameModeId(gameMode: GameMode): number | null;
    function gameModeById(id: number): GameMode | null;

    interface Dimension {
        range(): cube.Range;
        waterEvaporates(): boolean;
        lavaSpreadDuration(): Date;
        weatherCycle(): number;
        timeCycle(): number;
    }
    const dimension: {
        overworld: Dimension;
        nether: Dimension;
        end: Dimension;
    }
    function dimensionId(dimension: Dimension): number | null;
    function dimensionById(id: number): Dimension | null;

    interface ChunkPos {
        x(): number;
        z(): number;
        string(): string;
    }
    function chunkPos(x: number, z: number): ChunkPos;

    interface SubChunkPos {
        x(): number;
        y(): number;
        z(): number;
        string(): string;
    }
    function subChunkPos(x: number, y: number, z: number): SubChunkPos;

    function biomeById(id: number): biome.Biome | null;
    function biomeByName(name: string): biome.Biome | null;
    function biomes(): biome.Biome[];
    function registerBiome(biome: biome.Biome): void;

    function blockByName(name: string): block.Block | null;
    function blockByRuntimeId(id: number): block.Block | null;
    function blockHash(block: block.Block): number;
    function blockRuntimeId(block: block.Block): number;
    function customBlocks(): { [key: string]: block.CustomBlock };
    function registerBlock(block: block.Block): void;

    function itemByName(name: string): item.Stack | null;
    function itemByRuntimeId(id: number): item.Stack | null;
    function itemRuntimeId(item: item.Stack): { runtimeId: number, meta: number };
    function items(): item.Stack[];
    function customItems(): { [key: string]: item.CustomItem };
    function registerItem(item: item.Stack): void;

    interface EntityAnimation {
        name(): string;
        controller(): string;
        withController(controller: string): EntityAnimation;
        nextState(): string;
        withNextState(nextState: string): EntityAnimation;
        stopCondition(): string;
        withStopCondition(stopCondition: string): EntityAnimation;
    }
    function entityAnimation(name: string): EntityAnimation;

    interface Loader {
        world(): World;
        // TODO: changeWorld
        // TODO: changeRadius
        // TODO: move
        // TODO: load
        // TODO: chunk
        // TODO: close
    }
    function loader(radius: number, world: World, viewer: Viewer): Loader;

    interface SetOpts {
        disableBlockUpdates: boolean;
        disableLiquidDisplacement: boolean;
    }
    function setOpts(disableBlockUpdates: boolean, disableLiquidDisplacement: boolean): SetOpts;
}
