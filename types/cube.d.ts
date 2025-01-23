/// <reference path="mgl64.d.ts" />

declare namespace cube {
    class Axis {
        static readonly X: Axis;
        static readonly Y: Axis;
        static readonly Z: Axis;

        rotateLeft(): Axis;
        rotateRight(): Axis;
        vec3(): mgl64.Vec3;
        string(): string;
    }
    function axes(): Axis[];

    interface BBox {
        grow(x: number): BBox;
        growVec(vec: mgl64.Vec3): BBox;
        min(): mgl64.Vec3;
        max(): mgl64.Vec3;
        width(): number;
        length(): number;
        height(): number;
        extend(vec: mgl64.Vec3): BBox;
        extendTowards(face: Face, x: number): BBox;
        stretch(axis: Axis, x: number): BBox;
        translate(vec: mgl64.Vec3): BBox;
        translateTowards(face: Face, x: number): BBox;
        intersectsWith(other: BBox): boolean;
        vec3Within(vec: mgl64.Vec3): boolean;
        vec3WithinYZ(vec: mgl64.Vec3): boolean;
        vec3WithinXZ(vec: mgl64.Vec3): boolean;
        vec3WithinXY(vec: mgl64.Vec3): boolean;
        xOffset(nearby: BBox, deltaX: number): number;
        yOffset(nearby: BBox, deltaY: number): number;
        zOffset(nearby: BBox, deltaZ: number): number;
        corners(): mgl64.Vec3[];
        mul(val: number): BBox;
        volume(): number;
    }
    function box(x0: number, y0: number, z0: number, x1: number, y1: number, z1: number): BBox;
    function anyIntersections(boxes: BBox[], search: BBox): boolean;

    class Direction {
        static readonly North: Direction;
        static readonly East: Direction;
        static readonly South: Direction;
        static readonly West: Direction;

        face(): Face;
        opposite(): Direction;
        rotateLeft(): Direction;
        rotateRight(): Direction;
        string(): string;
    }
    function directions(): Direction[];

    class Face {
        static readonly Up: Face;
        static readonly Down: Face;
        static readonly North: Face;
        static readonly East: Face;
        static readonly South: Face;
        static readonly West: Face;

        axis(): Axis;
        direction(): Direction;
        opposite(): Face;
        rotateLeft(): Face;
        rotateRight(): Face;
        string(): string;
    }
    function faces(): Face[];
    function horizontalFaces(): Face[];

    interface Orientation {
        yaw(): number;
        opposite(): Orientation;
        rotateLeft(): Orientation;
        rotateRight(): Orientation;
    }
    function orientation(n: number): Orientation;
    function orientationFromYaw(yaw: number): Orientation;

    interface Pos {
        x(): number;
        y(): number;
        z(): number;
        add(pos: Pos): Pos;
        sub(pos: Pos): Pos;
        vec3(): mgl64.Vec3;
        vec3Middle(): mgl64.Vec3;
        vec3Centre(): mgl64.Vec3;
        side(face: Face): Pos;
        face(other: Pos): Face;
        neighbours(f: (neighbour: Pos) => void, range: Range): void;
        outOfBounds(range: Range): boolean;
        string(): string;
    }
    function pos(x: number, y: number, z: number): Pos;
    function posFromVec3(vec: mgl64.Vec3): Pos;

    interface Range {
        min(): number;
        max(): number;
        height(): number;
    }
    function range(min: number, max: number): Range;

    interface Rotation {
        yaw(): number;
        pitch(): number;
        elem(): [number, number];
        add(rot: Rotation): Rotation;
        opposite(): Rotation;
        neg(): Rotation;
        direction(): Direction;
        orientation(): Orientation;
        vec3(): mgl64.Vec3;
    }
    function rotation(yaw: number, pitch: number): Rotation;
}