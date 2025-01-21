/// <reference path="biome.d.ts" />
/// <reference path="block.d.ts" />
/// <reference path="chunk.d.ts" />
/// <reference path="world.d.ts" />

declare namespace generator {
    interface Generator {
        generateChunk(pos: world.ChunkPos, chunk: chunk.Chunk): void;
    }
    function flat(biome: biome.Biome, layers: block.Block[]): Generator;
}