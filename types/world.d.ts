/// <reference path="biome.d.ts" />
/// <reference path="block.d.ts" />
/// <reference path="cube.d.ts" />
/// <reference path="entity.d.ts" />
/// <reference path="generator.d.ts" />
/// <reference path="item.d.ts" />
/// <reference path="particle.d.ts" />
/// <reference path="player.d.ts" />
/// <reference path="sound.d.ts" />
/// <reference path="uuid.d.ts" />

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
    interface World extends BlockSource {
        name(): string;
        dimension(): Dimension;
        range(): cube.Range;
        exec(tx: (tx: Tx) => void): void;
        time(): number;
        setTime(time: number): void;
        stopTime(): void;
        startTime(): void;
        spawn(): cube.Pos;
        setSpawn(pos: cube.Pos): void;
        playerSpawn(uuid: uuid.UUID): cube.Pos;
        setPlayerSpawn(uuid: uuid.UUID, pos: cube.Pos): void;
        defaultGameMode(): GameMode;
        setDefaultGameMode(gameMode: GameMode): void;
        setTickRange(range: cube.Range): void;
        difficulty(): Difficulty;
        setDifficulty(difficulty: Difficulty): void;
        save(): void;
        close(): void;
    }
    function create(config?: Config): World;

    interface Tx {
        block(pos: cube.Pos): block.Block;
        setBlock(pos: cube.Pos, block: block.Block, opts?: SetOpts): void;
        liquid(pos: cube.Pos): block.Liquid;
        setLiquid(pos: cube.Pos, liquid: block.Liquid): void;
        // TODO: buildStructure
        scheduleBlockUpdate(pos: cube.Pos, block: block.Block, delay: Date): void;
        highestLightBlocker(x: number, z: number): number;
        highestBlock(x: number, z: number): number;
        light(pos: cube.Pos): number;
        skyLight(pos: cube.Pos): number;
        biome(pos: cube.Pos): biome.Biome;
        setBiome(pos: cube.Pos, biome: biome.Biome): void;
        temperature(pos: cube.Pos): number;
        rainingAt(pos: cube.Pos): boolean;
        snowingAt(pos: cube.Pos): boolean;
        thunderingAt(pos: cube.Pos): boolean;
        addParticle(pos: mgl64.Vec3, particle: particle.Particle): void;
        playEntityAnimation(entity: entity.Entity, animation: EntityAnimation): void;
        playSound(pos: mgl64.Vec3, sound: sound.Sound): void;
        players(): (player: player.Player) => boolean;
    }

    interface Viewer {}

    class Difficulty {
        static readonly Peaceful: Difficulty;
        static readonly Easy: Difficulty;
        static readonly Normal: Difficulty;
        static readonly Hard: Difficulty;

        foodRegenerates(): boolean;
        starvationHealthLimit(): number;
        fireSpreadIncrease(): number;
    }
    function difficultyId(difficulty: Difficulty): number | null;
    function difficultyById(id: number): Difficulty | null;

    class GameMode {
        static readonly Survival: GameMode;
        static readonly Creative: GameMode;
        static readonly Adventure: GameMode;
        static readonly Spectator: GameMode;

        allowsEditing(): boolean;
        allowsTakingDamage(): boolean;
        creativeInventory(): boolean;
        hasCollision(): boolean;
        allowsFlying(): boolean;
        allowsInteraction(): boolean;
        visible(): boolean;
    }
    function gameModeId(gameMode: GameMode): number | null;
    function gameModeById(id: number): GameMode | null;

    class Dimension {
        static readonly Overworld: Dimension;
        static readonly Nether: Dimension;
        static readonly End: Dimension;

        range(): cube.Range;
        waterEvaporates(): boolean;
        lavaSpreadDuration(): Date;
        weatherCycle(): number;
        timeCycle(): number;
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
