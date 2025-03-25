package dfscript

import "github.com/dop251/goja"

type cancelable interface {
	Cancel()
	Canceled() bool
}

type cancelableEvent struct {
	cancel bool
}

func (c *cancelableEvent) Cancel() {
	c.cancel = true
}

func (c *cancelableEvent) Canceled() bool {
	return c.cancel
}

func (r *Runtime) setupEvent() error {
	obj := newObject().Method("clearHandle", func(c goja.FunctionCall) goja.Value {
		id, ok := c.Argument(0).Export().(int64)
		if !ok {
			panic(r.vm.NewTypeError("argument 0 is not an int64, got %T", c.Argument(0).Export()))
		}
		for name, handles := range r.eventHandles {
			for i, handle := range handles {
				if handle.id == int(id) {
					handles = append(handles[:i], handles[i+1:]...)
					r.eventHandles[name] = handles
					return nil
				}
			}
		}
		return nil
	})
	events := []string{
		// Player Events
		"onPlayerItemDrop", "onPlayerHeldSlotChange", "onPlayerMove", "onPlayerJump", "onTeleport",
		"onPlayerChangeWorld", "onPlayerToggleSprint", "onPlayerToggleSneak", "onPlayerCommandExecution",
		"onPlayerTransfer", "onPlayerChat", "onPlayerSkinChange", "onPlayerFireExtinguish", "onPlayerStartBreak",
		"onPlayerBlockBreak", "onPlayerBlockPlace", "onPlayerBlockPick", "onPlayerSignEdit", "onPlayerLecternPageTurn",
		"onPlayerItemPickup", "onPlayerItemUse", "onPlayerItemUseOnBlock", "onPlayerItemUseOnEntity", "onPlayerItemRelease",
		"onPlayerItemConsume", "onPlayerItemDamage", "onPlayerAttackEntity", "onPlayerExperienceGain", "onPlayerPunchAir",
		"onPlayerHurt", "onPlayerHeal", "onPlayerFoodLoss", "onPlayerDeath", "onPlayerRespawn", "onPlayerQuit",
		"onPlayerDiagnostics", "onPlayerJoin",

		// World Events
		"onWorldLiquidFlow", "onWorldLiquidDecay", "onWorldLiquidHarden", "onWorldSound", "onWorldFireSpread",
		"onWorldBlockBurn", "onWorldCropTrample", "onWorldLeavesDecay", "onWorldEntitySpawn", "onWorldEntityDespawn",
		"onWorldExplosion", "onWorldClose",
	}
	for _, name := range events {
		obj.Method(name, func(c goja.FunctionCall) goja.Value {
			f, ok := goja.AssertFunction(c.Argument(0))
			if !ok {
				panic(r.vm.NewTypeError("argument 0 is not a function, got %T", c.Argument(0).Export()))
			}
			id := r.nextEventID
			r.nextEventID++
			r.eventHandles[name] = append(r.eventHandles[name], callable{id, f, c.This})
			return r.vm.ToValue(id)
		})
	}
	return obj.Apply(r.vm, "event")
}
