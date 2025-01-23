package dfscript

import (
	"fmt"
	"github.com/df-mc/dragonfly/server/block/cube"
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/event"
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/df-mc/dragonfly/server/player/skin"
	"github.com/df-mc/dragonfly/server/session"
	"github.com/df-mc/dragonfly/server/world"
	"github.com/go-gl/mathgl/mgl64"
	"net"
	"time"
)

type PlayerHandler struct {
	r *Runtime
}

func (p *PlayerHandler) callEvent(name string, e any) {
	calls, ok := p.r.eventHandles[name]
	if ok {
		for _, c := range calls {
			_, err := c.f(c.this, p.r.vm.ToValue(e))
			if err != nil {
				panic(err)
			}
		}
	}
}

func (p *PlayerHandler) callEventCtx(name string, ctx *event.Context[*player.Player], e any) {
	calls, ok := p.r.eventHandles[name]
	if ok {
		for _, c := range calls {
			_, err := c.f(c.this, p.r.vm.ToValue(e))
			if err != nil {
				panic(err)
			}
			if canc, ok := e.(cancelable); ok && canc.Canceled() {
				ctx.Cancel()
				return
			}
		}
	}
}

type playerEvent struct {
	Tx     *world.Tx
	Player *player.Player
}

func peCtx(ctx *player.Context) playerEvent {
	return playerEvent{Tx: ctx.Val().Tx(), Player: ctx.Val()}
}

func pePl(pl *player.Player) playerEvent {
	return playerEvent{Tx: pl.Tx(), Player: pl}
}

type itemDropEvent struct {
	playerEvent
	cancelableEvent
	Item item.Stack
}

func (p *PlayerHandler) HandleItemDrop(ctx *player.Context, i item.Stack) {
	p.callEventCtx("onPlayerItemDrop", ctx, &itemDropEvent{playerEvent: peCtx(ctx), Item: i})
}

type heldSlotChangeEvent struct {
	playerEvent
	cancelableEvent
	From int
	To   int
}

func (p *PlayerHandler) HandleHeldSlotChange(ctx *player.Context, from int, to int) {
	p.callEventCtx("onPlayerHeldSlotChange", ctx, &heldSlotChangeEvent{playerEvent: peCtx(ctx), From: from, To: to})
}

type moveEvent struct {
	playerEvent
	cancelableEvent
	NewPos mgl64.Vec3
	NewRot cube.Rotation
}

func (p *PlayerHandler) HandleMove(ctx *player.Context, newPos mgl64.Vec3, newRot cube.Rotation) {
	p.callEventCtx("onPlayerMove", ctx, &moveEvent{playerEvent: peCtx(ctx), NewPos: newPos, NewRot: newRot})
}

type jumpEvent struct{ playerEvent }

func (p *PlayerHandler) HandleJump(pl *player.Player) {
	p.callEvent("onPlayerJump", &jumpEvent{playerEvent: pePl(pl)})
}

type teleportEvent struct {
	playerEvent
	cancelableEvent
	Pos mgl64.Vec3
}

func (p *PlayerHandler) HandleTeleport(ctx *player.Context, pos mgl64.Vec3) {
	p.callEventCtx("onPlayerTeleport", ctx, &teleportEvent{playerEvent: peCtx(ctx), Pos: pos})
}

type changeWorldEvent struct {
	playerEvent
	Before *world.World
	After  *world.World
}

func (p *PlayerHandler) HandleChangeWorld(pl *player.Player, before, after *world.World) {
	p.callEvent("onPlayerChangeWorld", &changeWorldEvent{playerEvent: pePl(pl), Before: before, After: after})
}

type toggleSprintEvent struct {
	playerEvent
	cancelableEvent
	After bool
}

func (p *PlayerHandler) HandleToggleSprint(ctx *player.Context, after bool) {
	p.callEventCtx("onPlayerToggleSprint", ctx, &toggleSprintEvent{playerEvent: peCtx(ctx), After: after})
}

type toggleSneakEvent struct {
	playerEvent
	cancelableEvent
	After bool
}

func (p *PlayerHandler) HandleToggleSneak(ctx *player.Context, after bool) {
	p.callEventCtx("onPlayerToggleSneak", ctx, &toggleSneakEvent{playerEvent: peCtx(ctx), After: after})
}

type commandExecutionEvent struct {
	playerEvent
	cancelableEvent
	Command cmd.Command
	Args    []string
}

func (p *PlayerHandler) HandleCommandExecution(ctx *player.Context, command cmd.Command, args []string) {
	p.callEventCtx("onPlayerCommandExecution", ctx, &commandExecutionEvent{playerEvent: peCtx(ctx), Command: command, Args: args})
}

type transferEvent struct {
	playerEvent
	cancelableEvent
	Address string
}

func (p *PlayerHandler) HandleTransfer(ctx *player.Context, addr *net.UDPAddr) {
	e := &transferEvent{playerEvent: peCtx(ctx), Address: addr.String()}
	p.callEventCtx("onPlayerTransfer", ctx, e)
	add, err := net.ResolveUDPAddr("udp", e.Address)
	if err != nil {
		panic(fmt.Errorf("resolve %s: %w", e.Address, err))
	}
	*addr = *add
}

type chatEvent struct {
	playerEvent
	cancelableEvent
	Message string
}

func (p *PlayerHandler) HandleChat(ctx *player.Context, message *string) {
	e := &chatEvent{playerEvent: peCtx(ctx), Message: *message}
	p.callEventCtx("onPlayerChat", ctx, e)
	*message = e.Message
}

type skinChangeEvent struct {
	playerEvent
	cancelableEvent
	Skin skin.Skin
}

func (p *PlayerHandler) HandleSkinChange(ctx *player.Context, skin *skin.Skin) {
	e := &skinChangeEvent{playerEvent: peCtx(ctx), Skin: *skin}
	p.callEventCtx("onPlayerSkinChange", ctx, e)
	*skin = e.Skin
}

type fireExtinguishEvent struct {
	playerEvent
	cancelableEvent
	Pos cube.Pos
}

func (p *PlayerHandler) HandleFireExtinguish(ctx *player.Context, pos cube.Pos) {
	p.callEventCtx("onPlayerFireExtinguish", ctx, &fireExtinguishEvent{playerEvent: peCtx(ctx), Pos: pos})
}

type startBreakEvent struct {
	playerEvent
	cancelableEvent
	Pos cube.Pos
}

func (p *PlayerHandler) HandleStartBreak(ctx *player.Context, pos cube.Pos) {
	p.callEventCtx("onPlayerStartBreak", ctx, &startBreakEvent{playerEvent: peCtx(ctx), Pos: pos})
}

type blockBreakEvent struct {
	playerEvent
	cancelableEvent
	Pos   cube.Pos
	Drops []item.Stack
	XP    int
}

func (p *PlayerHandler) HandleBlockBreak(ctx *player.Context, pos cube.Pos, drops *[]item.Stack, xp *int) {
	e := &blockBreakEvent{playerEvent: peCtx(ctx), Pos: pos, Drops: *drops, XP: *xp}
	p.callEventCtx("onPlayerBlockBreak", ctx, e)
	*drops = e.Drops
	*xp = e.XP
}

type blockPlaceEvent struct {
	playerEvent
	cancelableEvent
	Pos   cube.Pos
	Block world.Block
}

func (p *PlayerHandler) HandleBlockPlace(ctx *player.Context, pos cube.Pos, b world.Block) {
	p.callEventCtx("onPlayerBlockPlace", ctx, &blockPlaceEvent{playerEvent: peCtx(ctx), Pos: pos, Block: b})
}

type blockPickEvent struct {
	playerEvent
	cancelableEvent
	Pos   cube.Pos
	Block world.Block
}

func (p *PlayerHandler) HandleBlockPick(ctx *player.Context, pos cube.Pos, b world.Block) {
	p.callEventCtx("onPlayerBlockPick", ctx, &blockPickEvent{playerEvent: peCtx(ctx), Pos: pos, Block: b})
}

type signEditEvent struct {
	playerEvent
	cancelableEvent
	Pos       cube.Pos
	FrontSide bool
	OldText   string
	NewText   string
}

func (p *PlayerHandler) HandleSignEdit(ctx *player.Context, pos cube.Pos, frontSide bool, oldText string, newText string) {
	p.callEventCtx("onPlayerSignEdit", ctx, &signEditEvent{playerEvent: peCtx(ctx), Pos: pos, FrontSide: frontSide, OldText: oldText, NewText: newText})
}

type lecternPageTurnEvent struct {
	playerEvent
	cancelableEvent
	Pos     cube.Pos
	OldPage int
	NewPage int
}

func (p *PlayerHandler) HandleLecternPageTurn(ctx *player.Context, pos cube.Pos, oldPage int, newPage *int) {
	e := &lecternPageTurnEvent{playerEvent: peCtx(ctx), Pos: pos, OldPage: oldPage, NewPage: *newPage}
	p.callEventCtx("onPlayerLecternPageTurn", ctx, e)
	*newPage = e.NewPage
}

type itemPickupEvent struct {
	playerEvent
	cancelableEvent
	Item item.Stack
}

func (p *PlayerHandler) HandleItemPickup(ctx *player.Context, i *item.Stack) {
	e := itemPickupEvent{playerEvent: peCtx(ctx), Item: *i}
	p.callEventCtx("onPlayerItemPickup", ctx, e)
	*i = e.Item
}

type itemUseEvent struct {
	playerEvent
	cancelableEvent
}

func (p *PlayerHandler) HandleItemUse(ctx *player.Context) {
	p.callEventCtx("onPlayerItemUse", ctx, &itemUseEvent{playerEvent: peCtx(ctx)})
}

type itemUseOnBlockEvent struct {
	playerEvent
	cancelableEvent
	Pos       cube.Pos
	Face      cube.Face
	ClickFace mgl64.Vec3
}

func (p *PlayerHandler) HandleItemUseOnBlock(ctx *player.Context, pos cube.Pos, face cube.Face, clickFace mgl64.Vec3) {
	p.callEventCtx("onPlayerItemUseOnBlock", ctx, &itemUseOnBlockEvent{playerEvent: peCtx(ctx), Pos: pos, Face: face, ClickFace: clickFace})
}

type itemUseOnEntityEvent struct {
	playerEvent
	cancelableEvent
	Entity world.Entity
}

func (p *PlayerHandler) HandleItemUseOnEntity(ctx *player.Context, e world.Entity) {
	p.callEventCtx("onPlayerItemUseOnEntity", ctx, &itemUseOnEntityEvent{playerEvent: peCtx(ctx), Entity: e})
}

type itemReleaseEvent struct {
	playerEvent
	cancelableEvent
	Item     item.Stack
	Duration time.Duration
}

func (p *PlayerHandler) HandleItemRelease(ctx *player.Context, i item.Stack, dur time.Duration) {
	p.callEventCtx("onPlayerItemRelease", ctx, &itemReleaseEvent{playerEvent: peCtx(ctx), Item: i, Duration: dur})
}

type itemConsumeEvent struct {
	playerEvent
	cancelableEvent
	Item item.Stack
}

func (p *PlayerHandler) HandleItemConsume(ctx *player.Context, i item.Stack) {
	p.callEventCtx("onPlayerItemConsume", ctx, &itemConsumeEvent{playerEvent: peCtx(ctx), Item: i})
}

type itemDamageEvent struct {
	playerEvent
	cancelableEvent
	Item   item.Stack
	Damage int
}

func (p *PlayerHandler) HandleItemDamage(ctx *player.Context, i item.Stack, damage int) {
	p.callEventCtx("onPlayerItemDamage", ctx, &itemDamageEvent{playerEvent: peCtx(ctx), Item: i, Damage: damage})
}

type attackEntityEvent struct {
	playerEvent
	cancelableEvent
	Entity   world.Entity
	Force    float64
	Height   float64
	Critical bool
}

func (p *PlayerHandler) HandleAttackEntity(ctx *player.Context, ent world.Entity, force *float64, height *float64, critical *bool) {
	e := &attackEntityEvent{playerEvent: peCtx(ctx), Entity: ent, Force: *force, Height: *height, Critical: *critical}
	p.callEventCtx("onPlayerAttackEntity", ctx, e)
	*force = e.Force
	*height = e.Height
	*critical = e.Critical
}

type experienceGainEvent struct {
	playerEvent
	cancelableEvent
	Amount int
}

func (p *PlayerHandler) HandleExperienceGain(ctx *player.Context, amount *int) {
	e := &experienceGainEvent{playerEvent: peCtx(ctx), Amount: *amount}
	p.callEventCtx("onPlayerExperienceGain", ctx, e)
	*amount = e.Amount
}

type punchAirEvent struct{ playerEvent }

func (p *PlayerHandler) HandlePunchAir(ctx *player.Context) {
	p.callEventCtx("onPlayerPunchAir", ctx, &punchAirEvent{})
}

type hurtEvent struct {
	playerEvent
	cancelableEvent
	Damage         float64
	Immune         bool
	AttackImmunity time.Duration
	Source         world.DamageSource
}

func (p *PlayerHandler) HandleHurt(ctx *player.Context, damage *float64, immune bool, attackImmunity *time.Duration, src world.DamageSource) {
	e := hurtEvent{playerEvent: peCtx(ctx), Damage: *damage, Immune: immune, AttackImmunity: *attackImmunity, Source: src}
	p.callEventCtx("onPlayerHurt", ctx, e)
	*damage = e.Damage
	*attackImmunity = e.AttackImmunity
}

type healEvent struct {
	playerEvent
	cancelableEvent
	Health float64
	Source world.HealingSource
}

func (p *PlayerHandler) HandleHeal(ctx *player.Context, health *float64, src world.HealingSource) {
	e := &healEvent{playerEvent: peCtx(ctx), Health: *health, Source: src}
	p.callEventCtx("onPlayerHeal", ctx, e)
	*health = e.Health
}

type foodLossEvent struct {
	playerEvent
	cancelableEvent
	From int
	To   int
}

func (p *PlayerHandler) HandleFoodLoss(ctx *player.Context, from int, to *int) {
	e := &foodLossEvent{playerEvent: peCtx(ctx), From: from, To: *to}
	p.callEventCtx("onPlayerFoodLoss", ctx, e)
	*to = e.To
}

type deathEvent struct {
	playerEvent
	Source  world.DamageSource
	KeepInv bool
}

func (p *PlayerHandler) HandleDeath(pl *player.Player, src world.DamageSource, keepInv *bool) {
	e := &deathEvent{playerEvent: pePl(pl), Source: src, KeepInv: *keepInv}
	p.callEvent("onPlayerDeath", e)
	*keepInv = e.KeepInv
}

type respawnEvent struct {
	playerEvent
	Pos   mgl64.Vec3
	World *world.World
}

func (p *PlayerHandler) HandleRespawn(pl *player.Player, pos *mgl64.Vec3, w **world.World) {
	e := &respawnEvent{playerEvent: pePl(pl), Pos: *pos, World: *w}
	p.callEvent("onPlayerRespawn", e)
	*pos = e.Pos
	*w = e.World
}

type quitEvent struct{ playerEvent }

func (p *PlayerHandler) HandleQuit(pl *player.Player) {
	p.callEvent("onPlayerQuit", &quitEvent{playerEvent: pePl(pl)})
}

type diagnosticsEvent struct {
	playerEvent
	Diagnostics session.Diagnostics
}

func (p *PlayerHandler) HandleDiagnostics(pl *player.Player, d session.Diagnostics) {
	p.callEvent("onPlayerDiagnostics", &diagnosticsEvent{playerEvent: pePl(pl), Diagnostics: d})
}
