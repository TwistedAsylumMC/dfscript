/// <reference path="cube.d.ts" />
/// <reference path="block.d.ts" />
/// <reference path="entity.d.ts" />
/// <reference path="item.d.ts" />
/// <reference path="mgl64.d.ts" />
/// <reference path="player.d.ts" />
/// <reference path="skin.d.ts" />
/// <reference path="sound.d.ts" />
/// <reference path="world.d.ts" />

declare namespace event {
    interface Event {}
    function clearHandle(id: number): void;

    interface Cancelable {
        cancel(): void;
        cancelled(): boolean;
    }

    interface PlayerEvent extends Event {
        tx: world.Tx;
        player: player.Player;
    }

    interface PlayerItemDropEvent extends PlayerEvent, Cancelable {
        readonly item: item.Stack;
    }
    interface PlayerHeldSlotChangeEvent extends PlayerEvent, Cancelable {
        readonly from: number;
        readonly to: number;
    }
    interface PlayerMoveEvent extends PlayerEvent, Cancelable {
        readonly newPos: mgl64.Vec3;
        readonly newRot: cube.Rotation;
    }
    interface PlayerJumpEvent extends PlayerEvent {}
    interface PlayerTeleportEvent extends PlayerEvent, Cancelable {
        readonly pos: mgl64.Vec3;
    }
    interface PlayerChangeWorldEvent extends PlayerEvent {
        readonly before: world.World;
        readonly after: world.World;
    }
    interface PlayerToggleSprintEvent extends PlayerEvent, Cancelable {
        readonly after: boolean;
    }
    interface PlayerToggleSneakEvent extends PlayerEvent, Cancelable {
        readonly after: boolean;
    }
    interface PlayerCommandExecutionEvent extends PlayerEvent, Cancelable {
        // TODO: readonly command: cmd.Command;
        readonly args: string[];
    }
    interface PlayerTransferEvent extends PlayerEvent, Cancelable {
        address: string;
    }
    interface PlayerChatEvent extends PlayerEvent, Cancelable {
        message: string;
    }
    interface PlayerSkinChangeEvent extends PlayerEvent, Cancelable {
        skin: skin.Skin;
    }
    interface PlayerFireExtinguishEvent extends PlayerEvent, Cancelable {
        readonly pos: cube.Pos;
    }
    interface PlayerStartBreakEvent extends PlayerEvent, Cancelable {
        readonly pos: cube.Pos;
    }
    interface PlayerBlockBreakEvent extends PlayerEvent, Cancelable {
        readonly pos: cube.Pos;
        drops: item.Stack[];
        xp: number;
    }
    interface PlayerBlockPlaceEvent extends PlayerEvent, Cancelable {
        readonly pos: cube.Pos;
        readonly block: block.Block;
    }
    interface PlayerBlockPickEvent extends PlayerEvent, Cancelable {
        readonly pos: cube.Pos;
        readonly block: block.Block;
    }
    interface PlayerSignEditEvent extends PlayerEvent, Cancelable {
        readonly pos: cube.Pos;
        readonly frontSide: boolean;
        readonly oldText: string;
        readonly newText: string;
    }
    interface PlayerLecternPageTurnEvent extends PlayerEvent, Cancelable {
        readonly pos: cube.Pos;
        readonly oldPage: number;
        newPage: number;
    }
    interface PlayerItemPickupEvent extends PlayerEvent, Cancelable {
        item: item.Stack;
    }
    interface PlayerItemUseEvent extends PlayerEvent, Cancelable {
        cancel(): void;
        canceled(): boolean;
    }
    interface PlayerItemUseOnBlockEvent extends PlayerEvent, Cancelable {
        readonly pos: cube.Pos;
        readonly face: cube.Face;
        readonly clickFace: mgl64.Vec3;
    }
    interface PlayerItemUseOnEntityEvent extends PlayerEvent, Cancelable {
        readonly entity: entity.Entity;
    }
    interface PlayerItemReleaseEvent extends PlayerEvent, Cancelable {
        readonly item: item.Stack;
        readonly duration: Date;
    }
    interface PlayerItemConsumeEvent extends PlayerEvent, Cancelable {
        readonly item: item.Stack;
    }
    interface PlayerItemDamageEvent extends PlayerEvent, Cancelable {
        readonly item: item.Stack;
        readonly damage: number;
    }
    interface PlayerAttackEntityEvent extends PlayerEvent, Cancelable {
        readonly entity: entity.Entity;
        force: number;
        height: number;
        critical: boolean;
    }
    interface PlayerExperienceGainEvent extends PlayerEvent, Cancelable {
        amount: number;
    }
    interface PlayerPunchAirEvent extends PlayerEvent {}
    interface PlayerHurtEvent extends PlayerEvent, Cancelable {
        damage: number;
        readonly immune: boolean;
        attackImmunity: Date;
        readonly source: entity.DamageSource;
    }
    interface PlayerHealEvent extends PlayerEvent, Cancelable {
        health: number;
        readonly source: entity.HealingSource;
    }
    interface PlayerFoodLossEvent extends PlayerEvent, Cancelable {
        readonly from: number;
        to: number;
    }
    interface PlayerDeathEvent extends PlayerEvent {
        readonly source: entity.DamageSource;
        keepInv: boolean;
    }
    interface PlayerRespawnEvent extends PlayerEvent {
        pos: mgl64.Vec3;
        world: world.World;
    }
    interface PlayerQuitEvent extends PlayerEvent {}
    interface PlayerDiagnosticsEvent extends PlayerEvent {
        readonly diagnostics: player.Diagnostics;
    }

    function onPlayerJoin(callback: (player: player.Player) => void): number;
    function onPlayerItemDrop(callback: (event: PlayerItemDropEvent) => void): number;
    function onPlayerHeldSlotChange(callback: (event: PlayerHeldSlotChangeEvent) => void): number;
    function onPlayerMove(callback: (event: PlayerMoveEvent) => void): number;
    function onPlayerJump(callback: (event: PlayerJumpEvent) => void): number;
    function onPlayerTeleport(callback: (event: PlayerTeleportEvent) => void): number;
    function onPlayerChangeWorld(callback: (event: PlayerChangeWorldEvent) => void): number;
    function onPlayerToggleSprint(callback: (event: PlayerToggleSprintEvent) => void): number;
    function onPlayerToggleSneak(callback: (event: PlayerToggleSneakEvent) => void): number;
    function onPlayerCommandExecution(callback: (event: PlayerCommandExecutionEvent) => void): number;
    function onPlayerTransfer(callback: (event: PlayerTransferEvent) => void): number;
    function onPlayerChat(callback: (event: PlayerChatEvent) => void): number;
    function onPlayerSkinChange(callback: (event: PlayerSkinChangeEvent) => void): number;
    function onPlayerFireExtinguish(callback: (event: PlayerFireExtinguishEvent) => void): number;
    function onPlayerStartBreak(callback: (event: PlayerStartBreakEvent) => void): number;
    function onPlayerBlockBreak(callback: (event: PlayerBlockBreakEvent) => void): number;
    function onPlayerBlockPlace(callback: (event: PlayerBlockPlaceEvent) => void): number;
    function onPlayerBlockPick(callback: (event: PlayerBlockPickEvent) => void): number;
    function onPlayerSignEdit(callback: (event: PlayerSignEditEvent) => void): number;
    function onPlayerLecternPageTurn(callback: (event: PlayerLecternPageTurnEvent) => void): number;
    function onPlayerItemPickup(callback: (event: PlayerItemPickupEvent) => void): number;
    function onPlayerItemUse(callback: (event: PlayerItemUseEvent) => void): number;
    function onPlayerItemUseOnBlock(callback: (event: PlayerItemUseOnBlockEvent) => void): number;
    function onPlayerItemUseOnEntity(callback: (event: PlayerItemUseOnEntityEvent) => void): number;
    function onPlayerItemRelease(callback: (event: PlayerItemReleaseEvent) => void): number;
    function onPlayerItemConsume(callback: (event: PlayerItemConsumeEvent) => void): number;
    function onPlayerItemDamage(callback: (event: PlayerItemDamageEvent) => void): number;
    function onPlayerAttackEntity(callback: (event: PlayerAttackEntityEvent) => void): number;
    function onPlayerExperienceGain(callback: (event: PlayerExperienceGainEvent) => void): number;
    function onPlayerPunchAir(callback: (event: PlayerPunchAirEvent) => void): number;
    function onPlayerHurt(callback: (event: PlayerHurtEvent) => void): number;
    function onPlayerHeal(callback: (event: PlayerHealEvent) => void): number;
    function onPlayerFoodLoss(callback: (event: PlayerFoodLossEvent) => void): number;
    function onPlayerDeath(callback: (event: PlayerDeathEvent) => void): number;
    function onPlayerRespawn(callback: (event: PlayerRespawnEvent) => void): number;
    function onPlayerQuit(callback: (event: PlayerQuitEvent) => void): number;
    function onPlayerDiagnostics(callback: (event: PlayerDiagnosticsEvent) => void): number;

    interface WorldEvent extends Event {
        tx: world.Tx;
    }

    interface WorldLiquidFlowEvent extends WorldEvent, Cancelable {
        from: cube.Pos;
        into: cube.Pos;
        liquid: block.Liquid;
        replaced: block.Block;
    }
    interface WorldLiquidDecayEvent extends WorldEvent, Cancelable {
        pos: cube.Pos;
        before: block.Liquid;
        after: block.Liquid;
    }
    interface WorldLiquidHardenEvent extends WorldEvent, Cancelable {
        pos: cube.Pos;
        liquid: block.Block;
        otherLiquid: block.Block;
        newBlock: block.Block;
    }
    interface WorldSoundEvent extends WorldEvent, Cancelable {
        pos: mgl64.Vec3;
        sound: sound.Sound;
    }
    interface WorldFireSpreadEvent extends WorldEvent, Cancelable {
        from: cube.Pos;
        to: cube.Pos;
    }
    interface WorldBlockBurnEvent extends WorldEvent, Cancelable {
        pos: cube.Pos;
    }
    interface WorldCropTrampleEvent extends WorldEvent, Cancelable {
        pos: cube.Pos;
    }
    interface WorldLeavesDecayEvent extends WorldEvent, Cancelable {
        pos: cube.Pos;
    }
    interface WorldEntitySpawnEvent extends WorldEvent {
        entity: entity.Entity;
    }
    interface WorldEntityDespawnEvent extends WorldEvent {
        entity: entity.Entity;
    }
    interface WorldCloseEvent extends WorldEvent {}

    function onWorldLiquidFlowEvent(callback: (event: WorldLiquidFlowEvent) => void): number;
    function onWorldLiquidDecayEvent(callback: (event: WorldLiquidDecayEvent) => void): number;
    function onWorldLiquidHardenEvent(callback: (event: WorldLiquidHardenEvent) => void): number;
    function onWorldSoundEvent(callback: (event: WorldSoundEvent) => void): number;
    function onWorldFireSpreadEvent(callback: (event: WorldFireSpreadEvent) => void): number;
    function onWorldBlockBurnEvent(callback: (event: WorldBlockBurnEvent) => void): number;
    function onWorldCropTrampleEvent(callback: (event: WorldCropTrampleEvent) => void): number;
    function onWorldLeavesDecayEvent(callback: (event: WorldLeavesDecayEvent) => void): number;
    function onWorldEntitySpawnEvent(callback: (event: WorldEntitySpawnEvent) => void): number;
    function onWorldEntityDespawnEvent(callback: (event: WorldEntityDespawnEvent) => void): number;
    function onWorldCloseEvent(callback: (event: WorldCloseEvent) => void): number;
}