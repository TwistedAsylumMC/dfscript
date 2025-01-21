/// <reference path="item.d.ts" />

declare namespace enchantment {
    interface Rarity {
        name(): string;
        cost(): number;
        weight(): number;
    }
    const rarity: {
        common: Rarity;
        uncommon: Rarity;
        rare: Rarity;
        veryRare: Rarity;
    }

    interface EnchantmentType {
        name(): string;
        maxLevel(): number;
        cost(level: number): [number, number];
        rarity(): Rarity;
        compatibleWithEnchantment(other: EnchantmentType): boolean;
        compatibleWithItem(item: item.Stack): boolean;
    }
    const aquaAffinity: EnchantmentType;
    const blastProtection: EnchantmentType;
    const curseOfVanishing: EnchantmentType;
    const depthStrider: EnchantmentType;
    const efficiency: EnchantmentType;
    const featherFalling: EnchantmentType;
    const fireAspect: EnchantmentType;
    const fireProtection: EnchantmentType;
    const flame: EnchantmentType;
    const infinity: EnchantmentType;
    const knockback: EnchantmentType;
    const mending: EnchantmentType;
    const power: EnchantmentType;
    const projectileProtection: EnchantmentType;
    const protection: EnchantmentType;
    const punch: EnchantmentType;
    const quickCharge: EnchantmentType;
    const respiration: EnchantmentType;
    const sharpness: EnchantmentType;
    const silkTouch: EnchantmentType;
    const soulSpeed: EnchantmentType;
    const swiftSneak: EnchantmentType;
    const thorns: EnchantmentType;
    const unbreaking: EnchantmentType;

    interface Enchantment {
        level(): number;
        type(): EnchantmentType;
    }
    function create(type: EnchantmentType, level: number): Enchantment;
    function register(id: number, type: EnchantmentType): void;
}