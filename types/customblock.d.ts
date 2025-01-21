/// <reference path="cube.d.ts" />
/// <reference path="mgl64.d.ts" />

declare namespace customblock {
    class RenderMethod {
        static readonly Opaque: RenderMethod;
        static readonly AlphaTest: RenderMethod;
        static readonly Blend: RenderMethod;
        static readonly DoubleSided: RenderMethod;

        uint8(): number;
        string(): string;
        ambientOcclusion(): boolean;
    }

    interface Material {
        withFaceDimming(): Material;
        withoutFaceDimming(): Material;
        withAmbientOcclusion(): Material;
        withoutAmbientOcclusion(): Material;
        encode(): object;
    }
    function material(texture: string, method: RenderMethod): Material;

    interface Properties {
        collisionBox: cube.BBox;
        cube: boolean;
        geometry: string;
        mapColour: string;
        rotation: cube.Pos;
        scale: mgl64.Vec3;
        selectionBox: cube.BBox;
        textures: { [key: string]: Material };
        translation: mgl64.Vec3;
    }

    interface Permutation extends Properties {
        condition: string;
    }
}