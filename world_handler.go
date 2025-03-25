package dfscript

import (
	"github.com/df-mc/dragonfly/server/block/cube"
	"github.com/df-mc/dragonfly/server/event"
	"github.com/df-mc/dragonfly/server/world"
	"github.com/go-gl/mathgl/mgl64"
)

type WorldHandler struct {
	r *Runtime
}

func (w *WorldHandler) callEvent(name string, e any) {
	calls, ok := w.r.eventHandles[name]
	if ok {
		for _, c := range calls {
			_, err := c.f(c.this, w.r.vm.ToValue(e))
			if err != nil {
				panic(err)
			}
		}
	}
}

func (w *WorldHandler) callEventCtx(name string, ctx *event.Context[*world.Tx], e any) {
	calls, ok := w.r.eventHandles[name]
	if ok {
		for _, c := range calls {
			_, err := c.f(c.this, w.r.vm.ToValue(e))
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

type worldEvent struct {
	Tx *world.Tx
}

func weCtx(ctx *world.Context) worldEvent {
	return worldEvent{Tx: ctx.Val()}
}

func weTx(tx *world.Tx) worldEvent {
	return worldEvent{Tx: tx}
}

type liquidFlowEvent struct {
	worldEvent
	cancelableEvent
	From     cube.Pos
	Into     cube.Pos
	Liquid   world.Liquid
	Replaced world.Block
}

func (w *WorldHandler) HandleLiquidFlow(ctx *world.Context, from, into cube.Pos, liquid world.Liquid, replaced world.Block) {
	e := &liquidFlowEvent{worldEvent: weCtx(ctx), From: from, Into: into, Liquid: liquid, Replaced: replaced}
	w.callEventCtx("onWorldLiquidFlow", ctx, e)
}

type liquidDecayEvent struct {
	worldEvent
	cancelableEvent
	Pos    cube.Pos
	Before world.Liquid
	After  world.Liquid
}

func (w *WorldHandler) HandleLiquidDecay(ctx *world.Context, pos cube.Pos, before, after world.Liquid) {
	e := &liquidDecayEvent{worldEvent: weCtx(ctx), Pos: pos, Before: before, After: after}
	w.callEventCtx("onWorldLiquidDecay", ctx, e)
}

type liquidHardenEvent struct {
	worldEvent
	cancelableEvent
	Pos         cube.Pos
	Liquid      world.Block
	OtherLiquid world.Block
	NewBlock    world.Block
}

func (w *WorldHandler) HandleLiquidHarden(ctx *world.Context, pos cube.Pos, liquid, otherLiquid, newBlock world.Block) {
	e := &liquidHardenEvent{worldEvent: weCtx(ctx), Pos: pos, Liquid: liquid, OtherLiquid: otherLiquid, NewBlock: newBlock}
	w.callEventCtx("onWorldLiquidHarden", ctx, e)
}

type soundEvent struct {
	worldEvent
	cancelableEvent
	Pos   mgl64.Vec3
	Sound world.Sound
}

func (w *WorldHandler) HandleSound(ctx *world.Context, s world.Sound, pos mgl64.Vec3) {
	e := &soundEvent{worldEvent: weCtx(ctx), Pos: pos, Sound: s}
	w.callEventCtx("onWorldSound", ctx, e)
}

type fireSpreadEvent struct {
	worldEvent
	cancelableEvent
	From cube.Pos
	To   cube.Pos
}

func (w *WorldHandler) HandleFireSpread(ctx *world.Context, from, to cube.Pos) {
	e := &fireSpreadEvent{worldEvent: weCtx(ctx), From: from, To: to}
	w.callEventCtx("onWorldFireSpread", ctx, e)
}

type blockBurnEvent struct {
	worldEvent
	cancelableEvent
	Pos cube.Pos
}

func (w *WorldHandler) HandleBlockBurn(ctx *world.Context, pos cube.Pos) {
	w.callEventCtx("onWorldBlockBurn", ctx, &blockBurnEvent{worldEvent: weCtx(ctx), Pos: pos})
}

type cropTrampleEvent struct {
	worldEvent
	cancelableEvent
	Pos cube.Pos
}

func (w *WorldHandler) HandleCropTrample(ctx *world.Context, pos cube.Pos) {
	w.callEventCtx("onWorldCropTrample", ctx, &cropTrampleEvent{worldEvent: weCtx(ctx), Pos: pos})
}

type leavesDecayEvent struct {
	worldEvent
	cancelableEvent
	Pos cube.Pos
}

func (w *WorldHandler) HandleLeavesDecay(ctx *world.Context, pos cube.Pos) {
	w.callEventCtx("onWorldLeavesDecay", ctx, &leavesDecayEvent{worldEvent: weCtx(ctx), Pos: pos})
}

type entitySpawnEvent struct {
	worldEvent
	Entity world.Entity
}

func (w *WorldHandler) HandleEntitySpawn(tx *world.Tx, e world.Entity) {
	w.callEvent("onWorldEntitySpawn", &entitySpawnEvent{worldEvent: weTx(tx), Entity: e})
}

type entityDespawnEvent struct {
	worldEvent
	Entity world.Entity
}

func (w *WorldHandler) HandleEntityDespawn(tx *world.Tx, e world.Entity) {
	w.callEvent("onWorldEntityDespawn", &entityDespawnEvent{worldEvent: weTx(tx), Entity: e})
}

type explosionEvent struct {
	worldEvent
	cancelableEvent
	Pos            mgl64.Vec3
	Entities       []world.Entity
	Blocks         []cube.Pos
	ItemDropChance float64
	SpawnFire      bool
}

func (w *WorldHandler) HandleExplosion(ctx *world.Context, position mgl64.Vec3, entities *[]world.Entity, blocks *[]cube.Pos, itemDropChance *float64, spawnFire *bool) {
	e := &explosionEvent{worldEvent: weCtx(ctx), Pos: position, Entities: *entities, Blocks: *blocks,
		ItemDropChance: *itemDropChance, SpawnFire: *spawnFire}
	w.callEventCtx("onWorldExplosion", ctx, e)
	*entities = e.Entities
	*blocks = e.Blocks
	*itemDropChance = e.ItemDropChance
	*spawnFire = e.SpawnFire
}

type worldCloseEvent struct{ worldEvent }

func (w *WorldHandler) HandleClose(tx *world.Tx) {
	w.callEvent("onWorldClose", &worldCloseEvent{worldEvent: weTx(tx)})
}
