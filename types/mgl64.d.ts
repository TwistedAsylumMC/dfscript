declare namespace mgl64 {
    interface Vec3 {
        x(): number;
        y(): number;
        z(): number;
        elem(): [number, number, number];
        cross(other: Vec3): Vec3;
        add(other: Vec3): Vec3;
        sub(other: Vec3): Vec3;
        mul(val: number): Vec3;
        dot(other: Vec3): number;
        len(): number;
        lenSqr(): number;
        normalize(): Vec3;
        approxEqual(other: Vec3): boolean;
        approxEqualThreshold(other: Vec3, threshold: number): boolean;
        approxFuncEqual(other: Vec3, func: (a: number, b: number) => boolean): boolean;
    }
    function vec3(x: number, y: number, z: number): Vec3;
}