package dfscript

import (
	"fmt"
	"github.com/df-mc/dragonfly/server/block/cube"
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/df-mc/dragonfly/server/player/skin"
	"github.com/df-mc/dragonfly/server/session"
	"github.com/df-mc/dragonfly/server/world"
	"github.com/go-gl/mathgl/mgl64"
	"net"
	"time"
)

type itemDropEvent struct {
	cancelableEvent
	Item item.Stack
}

type heldSlotChangeEvent struct {
	cancelableEvent
	From int
	To   int
}

type moveEvent struct {
	cancelableEvent
	NewPos mgl64.Vec3
	NewRot cube.Rotation
}

type jumpEvent struct{}

type teleportEvent struct {
	cancelableEvent
	Pos mgl64.Vec3
}

type changeWorldEvent struct {
	Before *world.World
	After  *world.World
}

type toggleSprintEvent struct {
	cancelableEvent
	After bool
}

type toggleSneakEvent struct {
	cancelableEvent
	After bool
}

type commandExecutionEvent struct {
	cancelableEvent
	Command cmd.Command
	Args    []string
}

type transferEvent struct {
	cancelableEvent
	Address string
}

type chatEvent struct {
	cancelableEvent
	Message string
}

type skinChangeEvent struct {
	cancelableEvent
	Skin skin.Skin
}

type fireExtinguishEvent struct {
	cancelableEvent
	Pos cube.Pos
}

type startBreakEvent struct {
	cancelableEvent
	Pos cube.Pos
}

type blockBreakEvent struct {
	cancelableEvent
	Pos   cube.Pos
	Drops []item.Stack
	XP    int
}

type blockPlaceEvent struct {
	cancelableEvent
	Pos   cube.Pos
	Block world.Block
}

type blockPickEvent struct {
	cancelableEvent
	Pos   cube.Pos
	Block world.Block
}

type signEditEvent struct {
	cancelableEvent
	Pos       cube.Pos
	FrontSide bool
	OldText   string
	NewText   string
}

type lecternPageTurnEvent struct {
	cancelableEvent
	Pos     cube.Pos
	OldPage int
	NewPage int
}

type itemPickupEvent struct {
	cancelableEvent
	Item item.Stack
}

type itemUseEvent struct {
	cancelableEvent
}

type itemUseOnBlockEvent struct {
	cancelableEvent
	Pos       cube.Pos
	Face      cube.Face
	ClickFace mgl64.Vec3
}

type itemUseOnEntityEvent struct {
	cancelableEvent
	Entity world.Entity
}

type itemReleaseEvent struct {
	cancelableEvent
	Item     item.Stack
	Duration time.Duration
}

type itemConsumeEvent struct {
	cancelableEvent
	Item item.Stack
}

type itemDamageEvent struct {
	cancelableEvent
	Item   item.Stack
	Damage int
}

type attackEntityEvent struct {
	cancelableEvent
	Entity   world.Entity
	Force    float64
	Height   float64
	Critical bool
}

type experienceGainEvent struct {
	cancelableEvent
	Amount int
}

type punchAirEvent struct{}

type hurtEvent struct {
	cancelableEvent
	Damage         float64
	Immune         bool
	AttackImmunity time.Duration
	Source         world.DamageSource
}

type healEvent struct {
	cancelableEvent
	Health float64
	Source world.HealingSource
}

type foodLossEvent struct {
	cancelableEvent
	From int
	To   int
}

type deathEvent struct {
	Source  world.DamageSource
	KeepInv bool
}

type respawnEvent struct {
	Pos   mgl64.Vec3
	World *world.World
}

type quitEvent struct{}

type diagnosticsEvent struct {
	Diagnostics session.Diagnostics
}

func (p *Player) HandleItemDrop(ctx *player.Context, i item.Stack) {
	p.callEventCtx("onItemDrop", ctx, &itemDropEvent{Item: i})
}

func (p *Player) HandleHeldSlotChange(ctx *player.Context, from int, to int) {
	p.callEventCtx("onHeldSlotChange", ctx, &heldSlotChangeEvent{From: from, To: to})
}

func (p *Player) HandleMove(ctx *player.Context, newPos mgl64.Vec3, newRot cube.Rotation) {
	p.callEventCtx("onMove", ctx, &moveEvent{NewPos: newPos, NewRot: newRot})
}

func (p *Player) HandleJump(pl *player.Player) {
	p.callEvent("onJump", pl, &jumpEvent{})
}

func (p *Player) HandleTeleport(ctx *player.Context, pos mgl64.Vec3) {
	p.callEventCtx("onTeleport", ctx, &teleportEvent{Pos: pos})
}

func (p *Player) HandleChangeWorld(pl *player.Player, before, after *world.World) {
	p.callEvent("onChangeWorld", pl, &changeWorldEvent{Before: before, After: after})
}

func (p *Player) HandleToggleSprint(ctx *player.Context, after bool) {
	p.callEventCtx("onToggleSprint", ctx, &toggleSprintEvent{After: after})
}

func (p *Player) HandleToggleSneak(ctx *player.Context, after bool) {
	p.callEventCtx("onToggleSneak", ctx, &toggleSneakEvent{After: after})
}

func (p *Player) HandleCommandExecution(ctx *player.Context, command cmd.Command, args []string) {
	p.callEventCtx("onCommandExecution", ctx, &commandExecutionEvent{Command: command, Args: args})
}

func (p *Player) HandleTransfer(ctx *player.Context, addr *net.UDPAddr) {
	e := &transferEvent{Address: addr.String()}
	p.callEventCtx("onTransfer", ctx, e)
	add, err := net.ResolveUDPAddr("udp", e.Address)
	if err != nil {
		panic(fmt.Errorf("resolve %s: %w", e.Address, err))
	}
	*addr = *add
}

func (p *Player) HandleChat(ctx *player.Context, message *string) {
	e := &chatEvent{Message: *message}
	p.callEventCtx("onChat", ctx, e)
	*message = e.Message
}

func (p *Player) HandleSkinChange(ctx *player.Context, skin *skin.Skin) {
	e := &skinChangeEvent{Skin: *skin}
	p.callEventCtx("onSkinChange", ctx, e)
	*skin = e.Skin
}

func (p *Player) HandleFireExtinguish(ctx *player.Context, pos cube.Pos) {
	p.callEventCtx("onFireExtinguish", ctx, &fireExtinguishEvent{Pos: pos})
}

func (p *Player) HandleStartBreak(ctx *player.Context, pos cube.Pos) {
	p.callEventCtx("onStartBreak", ctx, &startBreakEvent{Pos: pos})
}

func (p *Player) HandleBlockBreak(ctx *player.Context, pos cube.Pos, drops *[]item.Stack, xp *int) {
	e := &blockBreakEvent{Pos: pos, Drops: *drops, XP: *xp}
	p.callEventCtx("onBlockBreak", ctx, e)
	*drops = e.Drops
	*xp = e.XP
}

func (p *Player) HandleBlockPlace(ctx *player.Context, pos cube.Pos, b world.Block) {
	p.callEventCtx("onBlockPlace", ctx, &blockPlaceEvent{Pos: pos, Block: b})
}

func (p *Player) HandleBlockPick(ctx *player.Context, pos cube.Pos, b world.Block) {
	p.callEventCtx("onBlockPick", ctx, &blockPickEvent{Pos: pos, Block: b})
}

func (p *Player) HandleSignEdit(ctx *player.Context, pos cube.Pos, frontSide bool, oldText string, newText string) {
	p.callEventCtx("onSignEdit", ctx, &signEditEvent{Pos: pos, FrontSide: frontSide, OldText: oldText, NewText: newText})
}

func (p *Player) HandleLecternPageTurn(ctx *player.Context, pos cube.Pos, oldPage int, newPage *int) {
	e := &lecternPageTurnEvent{Pos: pos, OldPage: oldPage, NewPage: *newPage}
	p.callEventCtx("onLecternPageTurn", ctx, e)
	*newPage = e.NewPage
}

func (p *Player) HandleItemPickup(ctx *player.Context, i *item.Stack) {
	e := itemPickupEvent{Item: *i}
	p.callEventCtx("onItemPickup", ctx, e)
	*i = e.Item
}

func (p *Player) HandleItemUse(ctx *player.Context) {
	p.callEventCtx("onItemUse", ctx, &itemUseEvent{})
}

func (p *Player) HandleItemUseOnBlock(ctx *player.Context, pos cube.Pos, face cube.Face, clickFace mgl64.Vec3) {
	p.callEventCtx("onItemUseOnBlock", ctx, &itemUseOnBlockEvent{Pos: pos, Face: face, ClickFace: clickFace})
}

func (p *Player) HandleItemUseOnEntity(ctx *player.Context, e world.Entity) {
	p.callEventCtx("onItemUseOnEntity", ctx, &itemUseOnEntityEvent{Entity: e})
}

func (p *Player) HandleItemRelease(ctx *player.Context, i item.Stack, dur time.Duration) {
	p.callEventCtx("onItemRelease", ctx, &itemReleaseEvent{Item: i, Duration: dur})
}

func (p *Player) HandleItemConsume(ctx *player.Context, i item.Stack) {
	p.callEventCtx("onItemConsume", ctx, &itemConsumeEvent{Item: i})
}

func (p *Player) HandleItemDamage(ctx *player.Context, i item.Stack, damage int) {
	p.callEventCtx("onItemDamage", ctx, &itemDamageEvent{Item: i, Damage: damage})
}

func (p *Player) HandleAttackEntity(ctx *player.Context, ent world.Entity, force *float64, height *float64, critical *bool) {
	e := &attackEntityEvent{Entity: ent, Force: *force, Height: *height, Critical: *critical}
	p.callEventCtx("onAttackEntity", ctx, e)
	*force = e.Force
	*height = e.Height
	*critical = e.Critical
}

func (p *Player) HandleExperienceGain(ctx *player.Context, amount *int) {
	e := &experienceGainEvent{Amount: *amount}
	p.callEventCtx("onExperienceGain", ctx, e)
	*amount = e.Amount
}

func (p *Player) HandlePunchAir(ctx *player.Context) {
	p.callEventCtx("onPunchAir", ctx, &punchAirEvent{})
}

func (p *Player) HandleHurt(ctx *player.Context, damage *float64, immune bool, attackImmunity *time.Duration, src world.DamageSource) {
	e := hurtEvent{Damage: *damage, Immune: immune, AttackImmunity: *attackImmunity, Source: src}
	p.callEventCtx("onHurt", ctx, e)
	*damage = e.Damage
	*attackImmunity = e.AttackImmunity
}

func (p *Player) HandleHeal(ctx *player.Context, health *float64, src world.HealingSource) {
	e := &healEvent{Health: *health, Source: src}
	p.callEventCtx("onHeal", ctx, e)
	*health = e.Health
}

func (p *Player) HandleFoodLoss(ctx *player.Context, from int, to *int) {
	e := &foodLossEvent{From: from, To: *to}
	p.callEventCtx("onFoodLoss", ctx, e)
	*to = e.To
}

func (p *Player) HandleDeath(pl *player.Player, src world.DamageSource, keepInv *bool) {
	e := &deathEvent{Source: src, KeepInv: *keepInv}
	p.callEvent("onDeath", pl, e)
	*keepInv = e.KeepInv
}

func (p *Player) HandleRespawn(pl *player.Player, pos *mgl64.Vec3, w **world.World) {
	e := &respawnEvent{Pos: *pos, World: *w}
	p.callEvent("onRespawn", pl, e)
	*pos = e.Pos
	*w = e.World
}

func (p *Player) HandleQuit(pl *player.Player) {
	p.callEvent("onQuit", pl, &quitEvent{})
	delete(p.r.players, pl.UUID())
}

func (p *Player) HandleDiagnostics(pl *player.Player, d session.Diagnostics) {
	p.callEvent("onDiagnostics", pl, &diagnosticsEvent{Diagnostics: d})
}
