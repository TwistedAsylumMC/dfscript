declare namespace effect {
    interface Type {} // TODO: Add type declarations
    interface Effect {
        withoutParticles(): Effect;
        particlesHidden(): boolean;
        level(): number;
        duration(): Date;
        ambient(): boolean;
        type(): Type;
        tickDuration(): Effect;
        tick(): number;
    }

    const absorption: Type;
    const blindness: Type;
    const conduitPower: Type;
    const darkness: Type;
    const fatalPoison: Type;
    const fireResistance: Type;
    const haste: Type;
    const healthBoost: Type;
    const hunger: Type;
    const instantDamage: Type;
    const instantHealth: Type;
    const invisibility: Type;
    const jumpBoost: Type;
    const levitation: Type;
    const miningFatigue: Type;
    const nausea: Type;
    const nightVision: Type;
    const poison: Type;
    const regeneration: Type;
    const resistance: Type;
    const saturation: Type;
    const slowFalling: Type;
    const slowness: Type;
    const speed: Type;
    const strength: Type;
    const waterBreathing: Type;
    const weakness: Type;
    const wither: Type;

    function byId(id: number): Type | null;
    function id(type: Type): number | null;
    function create(type: Type, level: number, duration: Date): Effect;
    function createAmbient(type: Type, level: number, duration: Date): Effect;
    function createInstant(type: Type, level: number, potency?: number): Effect;
}