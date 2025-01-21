/// <reference path="item.d.ts" />

declare namespace inventory {
    interface Inventory {
        clone(func: (slot: number, before: item.Stack, after: item.Stack) => void): Inventory;
        slotFunc(func: (slot: number, before: item.Stack, after: item.Stack) => void): void;
        item(slot: number): [item.Stack];
        setItem(slot: number, item: item.Stack): void;
        slots(): item.Stack[];
        items(): item.Stack[];
        first(item: item.Stack): [number, boolean];
        firstFunc(func: (item: item.Stack) => boolean): [number, boolean];
        firstEmpty(): [number, boolean];
        swap(slot1: number, slot2: number): void;
        addItem(item: item.Stack): number;
        removeItem(item: item.Stack): void;
        removeItemFunc(n: number, func: (item: item.Stack) => boolean): void;
        containsItem(item: item.Stack): boolean;
        containsItemFunc(n: number, func: (item: item.Stack) => boolean): boolean;
        merge(other: Inventory, func: (slot: number, before: item.Stack, after: item.Stack) => void): Inventory;
        empty(): boolean;
        clear(): item.Stack[];
        size(): number;
        close(): void;
        string(): string;
    }
    // TODO: function create(size: number, func: (slot: number, before: item.Stack, after: item.Stack) => void): Inventory;
}