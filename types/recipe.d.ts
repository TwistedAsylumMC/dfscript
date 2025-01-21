/// <reference path="item.d.ts" />

declare namespace recipe {
    interface Recipe {
        input(): item.Stack;
        output(): item.Stack[];
        block(): string;
        priority(): number;
    }
    function all(): Recipe[];
    function register(recipe: Recipe): void;
    function perform(block: string, input: item.Stack[]): item.Stack[] | null;

    interface Furnace extends Recipe {}
    function furnace(input: item.Stack, output: item.Stack, block: string): Furnace;

    interface ItemTag extends Recipe {}
    function itemTag(tag: string, count: number): ItemTag;

    interface Potion extends Recipe {}
    function potion(input: item.Stack, reagent: item.Stack, output: item.Stack): Potion;

    interface PotionContainerChange extends Recipe {}
    function potionContainerChange(input: item.Stack, reagent: item.Stack, output: item.Stack): PotionContainerChange;

    interface Shaped extends Recipe {}
    function shaped(input: item.Stack[], output: item.Stack, width: number, height: number, block: string): Shaped;

    interface Shapeless extends Recipe {}
    function shapeless(input: item.Stack[], output: item.Stack, block: string): Shapeless;

    interface SmithingTransform extends Recipe {}
    function smithingTransform(base: item.Stack, addition: item.Stack, template: item.Stack, output: item.Stack, block: string): SmithingTransform;

    interface SmithingTrim extends Recipe {}
    function smithingTrim(base: item.Stack, addition: item.Stack, template: item.Stack, block: string): SmithingTrim;
}