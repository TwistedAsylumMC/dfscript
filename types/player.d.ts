/// <reference path="block.d.ts" />
/// <reference path="bossbar.d.ts" />
/// <reference path="chat.d.ts" />
/// <reference path="cube.d.ts" />
/// <reference path="effect.d.ts" />
/// <reference path="entity.d.ts" />
/// <reference path="form.d.ts" />
/// <reference path="inventory.d.ts" />
/// <reference path="item.d.ts" />
/// <reference path="mgl64.d.ts" />
/// <reference path="particle.d.ts" />
/// <reference path="scoreboard.d.ts" />
/// <reference path="skin.d.ts" />
/// <reference path="sound.d.ts" />
/// <reference path="title.d.ts" />
/// <reference path="world.d.ts" />

declare namespace player {
    import Lectern = block.Lectern;

    interface ItemDropEvent {
        readonly item: item.Stack;

        cancel(): void;
        canceled(): boolean;
    }
    interface HeldSlotChangeEvent {
        readonly from: number;
        readonly to: number;

        cancel(): void;
        canceled(): boolean;
    }
    interface MoveEvent {
        readonly newPos: mgl64.Vec3;
        readonly newRot: cube.Rotation;

        cancel(): void;
        canceled(): boolean;
    }
    interface JumpEvent {}
    interface TeleportEvent {
        readonly pos: mgl64.Vec3;

        cancel(): void;
        canceled(): boolean;
    }
    interface ChangeWorldEvent {
        readonly before: world.World;
        readonly after: world.World;
    }
    interface ToggleSprintEvent {
        readonly after: boolean;

        cancel(): void;
        canceled(): boolean;
    }
    interface ToggleSneakEvent {
        readonly after: boolean;

        cancel(): void;
        canceled(): boolean;
    }
    interface CommandExecutionEvent {
        // TODO: readonly command: cmd.Command;
        readonly args: string[];

        cancel(): void;
        canceled(): boolean;
    }
    interface TransferEvent {
        address: string;

        cancel(): void;
        canceled(): boolean;
    }
    interface ChatEvent {
        message: string;

        cancel(): void;
        canceled(): boolean;
    }
    interface SkinChangeEvent {
        skin: skin.Skin;

        cancel(): void;
        canceled(): boolean;
    }
    interface FireExtinguishEvent {
        readonly pos: cube.Pos;

        cancel(): void;
        canceled(): boolean;
    }
    interface StartBreakEvent {
        readonly pos: cube.Pos;

        cancel(): void;
        canceled(): boolean;
    }
    interface BlockBreakEvent {
        readonly pos: cube.Pos;
        drops: item.Stack[];
        xp: number;

        cancel(): void;
        canceled(): boolean;
    }
    interface BlockPlaceEvent {
        readonly pos: cube.Pos;
        readonly block: block.Block;

        cancel(): void;
        canceled(): boolean;
    }
    interface BlockPickEvent {
        readonly pos: cube.Pos;
        readonly block: block.Block;

        cancel(): void;
        canceled(): boolean;
    }
    interface SignEditEvent {
        readonly pos: cube.Pos;
        readonly frontSide: boolean;
        readonly oldText: string;
        readonly newText: string;

        cancel(): void;
        canceled(): boolean;
    }
    interface LecternPageTurnEvent {
        readonly pos: cube.Pos;
        readonly oldPage: number;
        newPage: number;

        cancel(): void;
        canceled(): boolean;
    }
    interface ItemPickupEvent {
        item: item.Stack;

        cancel(): void;
        canceled(): boolean;
    }
    interface ItemUseEvent {
        cancel(): void;
        canceled(): boolean;
    }
    interface ItemUseOnBlockEvent {
        readonly pos: cube.Pos;
        readonly face: cube.Face;
        readonly clickFace: mgl64.Vec3;

        cancel(): void;
        canceled(): boolean;
    }
    interface ItemUseOnEntityEvent {
        readonly entity: entity.Entity;

        cancel(): void;
        canceled(): boolean;
    }
    interface ItemReleaseEvent {
        readonly item: item.Stack;
        readonly duration: Date;

        cancel(): void;
        canceled(): boolean;
    }
    interface ItemConsumeEvent {
        readonly item: item.Stack;

        cancel(): void;
        canceled(): boolean;
    }
    interface ItemDamageEvent {
        readonly item: item.Stack;
        readonly damage: number;

        cancel(): void;
        canceled(): boolean;
    }
    interface AttackEntityEvent {
        readonly entity: entity.Entity;
        force: number;
        height: number;
        critical: boolean;

        cancel(): void;
        canceled(): boolean;
    }
    interface ExperienceGainEvent {
        amount: number;

        cancel(): void;
        canceled(): boolean;
    }
    interface PunchAirEvent {}
    interface HurtEvent {
        damage: number;
        readonly immune: boolean;
        attackImmunity: Date;
        readonly source: entity.DamageSource;

        cancel(): void;
        canceled(): boolean;
    }
    interface HealEvent {
        health: number;
        readonly source: entity.HealingSource;

        cancel(): void;
        canceled(): boolean;
    }
    interface FoodLossEvent {
        readonly from: number;
        to: number;

        cancel(): void;
        canceled(): boolean;
    }
    interface DeathEvent {
        readonly source: entity.DamageSource;
        keepInv: boolean;
    }
    interface RespawnEvent {
        pos: mgl64.Vec3;
        world: world.World;
    }
    interface QuitEvent {}
    interface Diagnostics {
        readonly averageFramesPerSecond: number;
        readonly averageServerSimTickTime: number;
        readonly averageClientSimTickTime: number;
        readonly averageBeginFrameTime: number;
        readonly averageInputTime: number;
        readonly averageRenderTime: number;
        readonly averageEndFrameTime: number;
        readonly averageRemainderTimePercent: number;
        readonly averageUnaccountedTimePercent: number;
    }
    interface DiagnosticsEvent {
        readonly diagnostics: Diagnostics;
    }

    interface Player {
        readonly address: string;
        readonly deviceId: string;
        readonly deviceModel: string;
        readonly name: string;
        readonly selfSignedId: string;
        readonly uuid: string;
        readonly xuid: string;

        abortBreaking(): void;
        absorption(): number;
        addEffect(effect: effect.Effect): void;
        addExperience(amount: number): number;
        addFood(points: number): void;
        airSupply(): number;
        // TODO: armour(): void;
        attackEntity(entity: entity.Entity): boolean;
        beaconAffected(): boolean;
        breakBlock(pos: cube.Pos): void;
        breathing(): boolean;
        chat(...msg: any[]): void;
        close(): void;
        closeDialogue(): void;
        closeForm(): void;
        collect(item: item.Stack): number;
        collectExperience(value: number): boolean;
        continueBreaking(face: cube.Face): void;
        crawling(): boolean;
        // TODO: data(): void;
        dead(): boolean
        deathPosition(): void;
        disableInstantRespawn(): void;
        disconnect(...msg: any[]): void;
        drop(item: item.Stack): number;
        editSign(pos: cube.Pos, frontText: string, backText: string): void;
        effect(type: effect.Type): effect.Effect | null;
        effects(): effect.Effect[];
        enableInstantRespawn(): void;
        enchantmentSeed(): number;
        enderChestInventory(): inventory.Inventory;
        executeCommand(commandLine: string): void;
        exhaust(points: number): void;
        experience(): number;
        experienceLevel(): number;
        experienceProgress(): number;
        // TODO: explode(): void;
        extinguish(): void;
        eyeHeight(): number;
        fallDistance(): number;
        finalDamageFrom(damage: number, source: entity.DamageSource): number;
        finishBreaking(): void;
        fireProof(): boolean;
        flightSpeed(): number;
        flying(): boolean;
        food(): number;
        gameMode(): world.GameMode;
        gliding(): boolean;
        hasCooldown(item: item.Item): boolean;
        heal(health: number, source: entity.HealingSource): void;
        health(): number;
        heldItems(): [item.Stack, item.Stack];
        hideCoordinates(): void;
        hideEntity(entity: entity.Entity): void;
        hurt(damage: number, source: entity.DamageSource): void;
        immobile(): boolean;
        inventory(): inventory.Inventory;
        invisible(): boolean;
        jump(): void;
        knockBack(source: mgl64.Vec3, force: number, height: number): void;
        latency(): number;
        locale(): string;
        maxAirSupply(): number;
        maxHealth(): number;
        message(...msg: any[]): void;
        messagef(format: string, ...msg: any): void;
        messaget(translation: chat.Translation, ...msg: any): void;
        move(deltaPos: mgl64.Vec3, deltaYaw: number, deltaPitch: number): void;
        moveItemsToInventory(): void;
        nameTag(): string;
        onFireDuration(): Date;
        onGround(): boolean;
        openBlockContainer(pos: cube.Pos): void;
        openSign(pos: cube.Pos, frontSide: boolean): void;
        pickBlock(pos: cube.Pos): void;
        // TODO: placeBlock(): void;
        playSound(sound: sound.Sound): void;
        position(): mgl64.Vec3;
        punchAir(): void;
        releaseItem(): void;
        removeBossBar(): void;
        removeEffect(type: effect.Type): void;
        removeExperience(amount: number): void;
        removeScoreboard(): void;
        resetEnchantmentSeed(): void;
        resetFallDistance(): void;
        // TODO: respawn(): void;
        rotation(): cube.Rotation;
        saturate(food: number, saturation: number): void;
        scale(): number;
        scoreTag(): string;
        sendBossBar(bossBar: bossbar.BossBar): void;
        // TODO: sendCommandOutput(): void;
        // TODO: sendDialogue(): void;
        sendForm(form: form.Form): void;
        sendJukeboxPopup(...args: any[]): void;
        sendPopup(...args: any[]): void;
        sendScoreboard(scoreboard: scoreboard.Scoreboard): void;
        sendTip(...args: any[]): void;
        sendTitle(title: title.Title): void;
        sendToast(title: string, message: string): void;
        setAbsorption(health: number): void;
        setAirSupply(duration: Date): void;
        setCooldown(item: item.Item, cooldown: Date): void;
        setExperienceLevel(level: number): void;
        setExperienceProgress(progress: number): void;
        setFlightSpeed(speed: number): void;
        setFood(food: number): void;
        setGameMode(gameMode: world.GameMode): void;
        setHeldItems(mainHand: item.Stack, offHand: item.Stack): void;
        setHeldSlot(slot: number): void;
        setImmobile(): void;
        setInvisible(): void;
        setMaxAirSupply(duration: Date): void;
        setMaxHealth(health: number): void;
        setMobile(): void;
        setNameTag(name: string): void;
        setOnFire(duration: Date): void;
        setScale(scale: number): void;
        setScoreTag(...args: any[]): void;
        setSkin(skin: skin.Skin): void;
        setSpeed(speed: number): void;
        setVelocity(velocity: mgl64.Vec3): void;
        setVisible(): void;
        showCoordinates(): void;
        showEntity(entity: entity.Entity): void;
        showParticle(pos: mgl64.Vec3, particle: particle.Particle): void;
        skin(): skin.Skin;
        sneaking(): boolean;
        speed(): number;
        splash(pos: mgl64.Vec3): void;
        sprinting(): boolean;
        startBreaking(pos: cube.Pos, face: cube.Face): void;
        startCrawling(): void;
        startFlying(): void;
        startGliding(): void;
        startSneaking(): void;
        startSprinting(): void;
        startSwimming(): void;
        stopCrawling(): void;
        stopFlying(): void;
        stopGliding(): void;
        stopSneaking(): void;
        stopSprinting(): void;
        stopSwimming(): void;
        swimming(): void;
        swingArm(): void;
        teleport(pos: mgl64.Vec3): void;
        tick(current: number): void;
        transfer(address: string): void;
        turnLecternPage(pos: cube.Pos, page: number): void;
        updateDiagnostics(diagnostics: Diagnostics): void;
        useItem(): void;
        useItemOnBlock(pos: cube.Pos, face: cube.Face, clickPos: mgl64.Vec3): void;
        useItemOnEntity(entity: entity.Entity): void;
        usingItem(): boolean;
        velocity(): mgl64.Vec3;
        world(): world.World;

        onItemDrop(callback: (event: ItemDropEvent) => void): void;
        onHeldSlotChange(callback: (event: HeldSlotChangeEvent) => void): void;
        onMove(callback: (event: MoveEvent) => void): void;
        onJump(callback: (event: JumpEvent) => void): void;
        onTeleport(callback: (event: TeleportEvent) => void): void;
        onChangeWorld(callback: (event: ChangeWorldEvent) => void): void;
        onToggleSprint(callback: (event: ToggleSprintEvent) => void): void;
        onToggleSneak(callback: (event: ToggleSneakEvent) => void): void;
        onCommandExecution(callback: (event: CommandExecutionEvent) => void): void;
        onTransfer(callback: (event: TransferEvent) => void): void;
        onChat(callback: (event: ChatEvent) => void): void;
        onSkinChange(callback: (event: SkinChangeEvent) => void): void;
        onFireExtinguish(callback: (event: FireExtinguishEvent) => void): void;
        onStartBreak(callback: (event: StartBreakEvent) => void): void;
        onBlockBreak(callback: (event: BlockBreakEvent) => void): void;
        onBlockPlace(callback: (event: BlockPlaceEvent) => void): void;
        onBlockPick(callback: (event: BlockPickEvent) => void): void;
        onSignEdit(callback: (event: SignEditEvent) => void): void;
        onLecternPageTurn(callback: (event: LecternPageTurnEvent) => void): void;
        onItemPickup(callback: (event: ItemPickupEvent) => void): void;
        onItemUse(callback: (event: ItemUseEvent) => void): void;
        onItemUseOnBlock(callback: (event: ItemUseOnBlockEvent) => void): void;
        onItemUseOnEntity(callback: (event: ItemUseOnEntityEvent) => void): void;
        onItemRelease(callback: (event: ItemReleaseEvent) => void): void;
        onItemConsume(callback: (event: ItemConsumeEvent) => void): void;
        onItemDamage(callback: (event: ItemDamageEvent) => void): void;
        onAttackEntity(callback: (event: AttackEntityEvent) => void): void;
        onExperienceGain(callback: (event: ExperienceGainEvent) => void): void;
        onPunchAir(callback: (event: PunchAirEvent) => void): void;
        onHurt(callback: (event: HurtEvent) => void): void;
        onHeal(callback: (event: HealEvent) => void): void;
        onFoodLoss(callback: (event: FoodLossEvent) => void): void;
        onDeath(callback: (event: DeathEvent) => void): void;
        onRespawn(callback: (event: RespawnEvent) => void): void;
        onQuit(callback: (event: QuitEvent) => void): void;
        onDiagnostics(callback: (event: DiagnosticsEvent) => void): void;
    }
}