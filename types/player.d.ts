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
    interface Player extends entity.Entity {
        h(): entity.Handle<Player>;
        tx(): world.Tx;

        abortBreaking(): void;
        absorption(): number;
        addEffect(effect: effect.Effect): void;
        addExperience(amount: number): number;
        addFood(points: number): void;
        // TODO: address
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
        deviceID(): string;
        deviceModel(): string;
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
        name(): string;
        nameTag(): string;
        onFireDuration(): Date;
        onGround(): boolean;
        openBlockContainer(pos: cube.Pos, tx: world.Tx): void;
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
        selfSignedID(): string;
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
        splash(tx: world.Tx, pos: mgl64.Vec3): void;
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
        tick(tx: world.Tx, current: number): void;
        transfer(address: string): void;
        turnLecternPage(pos: cube.Pos, page: number): void;
        updateDiagnostics(diagnostics: Diagnostics): void;
        useItem(): void;
        useItemOnBlock(pos: cube.Pos, face: cube.Face, clickPos: mgl64.Vec3): void;
        useItemOnEntity(entity: entity.Entity): void;
        usingItem(): boolean;
        velocity(): mgl64.Vec3;
        world(): world.World;
        xuid(): string;
    }

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
}