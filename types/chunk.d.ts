/// <reference path="cube.d.ts" />

declare namespace chunk {
    interface Chunk {}
    interface SubChunk {}
    function createSubChunk(air: number): SubChunk;

    const subChunkVersion: number;
    const currentBlockVersion: number;

    class Encoding {
        static readonly Disk: Encoding;
        static readonly Network: Encoding;
    }

    class PaletteEncoding {
        static readonly Biome: PaletteEncoding;
        static readonly Block: PaletteEncoding;
    }

    class Light {
        static readonly Block: Light;
        static readonly sky: Light;
    }

    interface LightArea {
        fill(): void;
        spread(): void;
    }
    function lightArea(chunks: Chunk[], baseX: number, baseY: number): LightArea;

    function runtimeIdToState(runtimeId: number): { name: string, properties: object } | null;
    function stateToRuntimeId(name: string, properties: object): number | null;

    interface SerialisedData {
        subChunks: Uint8Array[];
        biomes: Uint8Array;
    }
    function diskDecode(data: SerialisedData, range: cube.Range): Chunk;
    function networkDecode(air: number, data: Uint8Array, count: number, range: cube.Range): Chunk;

    function encode(chunk: Chunk, encoding: Encoding): SerialisedData;
    function encodeBiomes(chunk: Chunk, encoding: Encoding): Uint8Array;
    function encodeSubChunks(chunk: Chunk, encoding: Encoding, index: number): Uint8Array;
}