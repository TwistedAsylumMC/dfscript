/// <reference path="effect.d.ts" />

declare namespace potion {
    interface Potion {
        uint8(): number;
        effects(): effect.Effect[];
    }
    const all: Potion[];
    const awkward: Potion;
    const fireResistance: Potion;
    const harming: Potion;
    const healing: Potion;
    const invisibility: Potion;
    const leaping: Potion;
    const longFireResistance: Potion;
    const longInvisibility: Potion;
    const longLeaping: Potion;
    const longMundane: Potion;
    const longNightVision: Potion;
    const longPoison: Potion;
    const longRegeneration: Potion;
    const longSlowFalling: Potion;
    const longSlowness: Potion;
    const longStrength: Potion;
    const longSwiftness: Potion;
    const longTurtleMaster: Potion;
    const longWaterBreathing: Potion;
    const longWeakness: Potion;
    const mundane: Potion;
    const nightVision: Potion;
    const poison: Potion;
    const regeneration: Potion;
    const slowFalling: Potion;
    const slowness: Potion;
    const strength: Potion;
    const strongHarming: Potion;
    const strongHealing: Potion;
    const strongLeaping: Potion;
    const strongPoison: Potion;
    const strongRegeneration: Potion;
    const strongSlowness: Potion;
    const strongStrength: Potion;
    const strongSwiftness: Potion;
    const strongTurtleMaster: Potion;
    const swiftness: Potion;
    const thick: Potion;
    const turtleMaster: Potion;
    const water: Potion;
    const waterBreathing: Potion;
    const weakness: Potion;
    const wither: Potion;

    function from(id: number): Potion;
}