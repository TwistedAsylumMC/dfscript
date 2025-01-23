/// <reference path="world.d.ts">

declare namespace entity {
    interface Entity {}

    interface Handle<T extends Entity> {
        // TODO: type
        // TODO: uuid
        entity(tx: world.Tx): [T, boolean];
        execWorld(tx: (tx: world.Tx, entity: T) => void): void;
        close(): void;
    }

    interface DamageSource {}
    interface HealingSource {}
}