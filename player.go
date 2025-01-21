package dfscript

import (
	"github.com/df-mc/dragonfly/server/block"
	"github.com/df-mc/dragonfly/server/block/cube"
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/entity/effect"
	"github.com/df-mc/dragonfly/server/event"
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/df-mc/dragonfly/server/player/bossbar"
	"github.com/df-mc/dragonfly/server/player/chat"
	"github.com/df-mc/dragonfly/server/player/dialogue"
	"github.com/df-mc/dragonfly/server/player/form"
	"github.com/df-mc/dragonfly/server/player/scoreboard"
	"github.com/df-mc/dragonfly/server/player/skin"
	"github.com/df-mc/dragonfly/server/player/title"
	"github.com/df-mc/dragonfly/server/session"
	"github.com/df-mc/dragonfly/server/world"
	"github.com/dop251/goja"
	"github.com/go-gl/mathgl/mgl64"
	"time"
)

type Player struct {
	r   *Runtime
	obj *goja.Object

	h *world.EntityHandle
	p *player.Player

	events map[string][]callable
}

func NewPlayer(r *Runtime, pl *player.Player) *Player {
	p := &Player{
		r: r,

		h: pl.H(),
		p: pl,

		events: make(map[string][]callable),
	}
	p.setup()
	return p
}

func (p *Player) setup() {
	vm := p.r.vm

	obj := newObject().
		Const("address", p.p.Addr().String()).
		Const("deviceId", p.p.DeviceID()).
		Const("deviceModel", p.p.DeviceModel()).
		Const("name", p.p.Name()).
		Const("selfSignedId", p.p.SelfSignedID()).
		Const("uuid", p.p.UUID().String()).
		Const("xuid", p.p.XUID())
	setExec := func(name string, f func(c goja.FunctionCall, p *player.Player) goja.Value) {
		obj.Method(name, func(c goja.FunctionCall) goja.Value {
			var ret goja.Value
			p.exec(func(p *player.Player) {
				ret = f(c, p)
			})
			return ret
		})
	}

	setExec("abortBreaking", func(c goja.FunctionCall, pl *player.Player) goja.Value {
		pl.AbortBreaking()
		return nil
	})
	setExec("absorption", func(c goja.FunctionCall, pl *player.Player) goja.Value {
		return vm.ToValue(pl.Absorption())
	})
	setExec("addEffect", func(c goja.FunctionCall, pl *player.Player) goja.Value {
		e, ok := c.Argument(0).Export().(effect.Effect)
		if !ok {
			panic(vm.NewTypeError("argument 0 is not an effect.Effect"))
		}
		pl.AddEffect(e)
		return nil
	})
	setExec("addExperience", func(c goja.FunctionCall, pl *player.Player) goja.Value {
		amount, ok := c.Argument(0).Export().(int64)
		if !ok {
			panic(vm.NewTypeError("argument 0 is not an int64"))
		}
		return vm.ToValue(pl.AddExperience(int(amount)))
	})
	setExec("addFood", func(c goja.FunctionCall, pl *player.Player) goja.Value {
		points, ok := c.Argument(0).Export().(int64)
		if !ok {
			panic(vm.NewTypeError("argument 0 is not an int64"))
		}
		pl.AddFood(int(points))
		return nil
	})
	setExec("airSupply", func(c goja.FunctionCall, pl *player.Player) goja.Value {
		return vm.ToValue(pl.AirSupply())
	})
	setExec("armour", func(c goja.FunctionCall, pl *player.Player) goja.Value {
		return vm.ToValue(pl.Armour())
	})
	setExec("attackEntity", func(c goja.FunctionCall, pl *player.Player) goja.Value {
		e, ok := c.Argument(0).Export().(world.Entity)
		if !ok {
			panic(vm.NewTypeError("argument 0 is not a world.Entity, got %T", c.Argument(0).Export()))
		}
		return vm.ToValue(pl.AttackEntity(e))
	})
	setExec("beaconAffected", func(c goja.FunctionCall, pl *player.Player) goja.Value {
		return vm.ToValue(pl.BeaconAffected())
	})
	setExec("breakBlock", func(c goja.FunctionCall, pl *player.Player) goja.Value {
		pos, ok := c.Argument(0).Export().(cube.Pos)
		if !ok {
			panic(vm.NewTypeError("argument 0 is not a cube.Pos, got %T", c.Argument(0).Export()))
		}
		pl.BreakBlock(pos)
		return nil
	})
	setExec("breathing", func(c goja.FunctionCall, pl *player.Player) goja.Value {
		return vm.ToValue(pl.Breathing())
	})
	setExec("chat", func(c goja.FunctionCall, pl *player.Player) goja.Value {
		var args []any
		for _, argument := range c.Arguments {
			args = append(args, argument.Export())
		}
		pl.Chat(args...)
		return nil
	})
	setExec("close", func(c goja.FunctionCall, pl *player.Player) goja.Value {
		return vm.ToValue(pl.Close())
	})
	setExec("closeDialogue", func(c goja.FunctionCall, pl *player.Player) goja.Value {
		pl.CloseDialogue()
		return nil
	})
	setExec("closeForm", func(c goja.FunctionCall, pl *player.Player) goja.Value {
		pl.CloseForm()
		return nil
	})
	setExec("collect", func(c goja.FunctionCall, pl *player.Player) goja.Value {
		s, ok := c.Argument(0).Export().(item.Stack)
		if !ok {
			panic(vm.NewTypeError("argument 0 is not an item.Stack"))
		}
		n, _ := pl.Collect(s)
		return vm.ToValue(n)
	})
	setExec("collectExperience", func(c goja.FunctionCall, pl *player.Player) goja.Value {
		value, ok := c.Argument(0).Export().(int64)
		if !ok {
			panic(vm.NewTypeError("argument 0 is not an int64"))
		}
		return vm.ToValue(pl.CollectExperience(int(value)))
	})
	setExec("continueBreaking", func(c goja.FunctionCall, pl *player.Player) goja.Value {
		face, ok := c.Argument(0).Export().(cube.Face)
		if !ok {
			panic(vm.NewTypeError("argument 0 is not a cube.Face, got %T", c.Argument(0).Export()))
		}
		pl.ContinueBreaking(face)
		return nil
	})
	setExec("crawling", func(c goja.FunctionCall, pl *player.Player) goja.Value {
		return vm.ToValue(pl.Crawling())
	})
	setExec("data", func(c goja.FunctionCall, pl *player.Player) goja.Value {
		return vm.ToValue(pl.Data())
	})
	setExec("dead", func(c goja.FunctionCall, pl *player.Player) goja.Value {
		return vm.ToValue(pl.Dead())
	})
	setExec("deathPosition", func(c goja.FunctionCall, pl *player.Player) goja.Value {
		pl.DeathPosition()
		return nil
	})
	setExec("disableInstantRespawn", func(c goja.FunctionCall, pl *player.Player) goja.Value {
		pl.DisableInstantRespawn()
		return nil
	})
	setExec("disconnect", func(c goja.FunctionCall, pl *player.Player) goja.Value {
		var msg []any
		for _, argument := range c.Arguments {
			msg = append(msg, argument.Export())
		}
		pl.Disconnect(msg...)
		return nil
	})
	setExec("drop", func(c goja.FunctionCall, pl *player.Player) goja.Value {
		s, ok := c.Argument(0).Export().(item.Stack)
		if !ok {
			panic(vm.NewTypeError("argument 0 is not an item.Stack"))
		}
		return vm.ToValue(pl.Drop(s))
	})
	setExec("editSign", func(c goja.FunctionCall, pl *player.Player) goja.Value {
		pos, ok := c.Argument(0).Export().(cube.Pos)
		if !ok {
			panic(vm.NewTypeError("argument 0 is not a cube.Pos, got %T", c.Argument(0).Export()))
		}
		frontText, ok := c.Argument(1).Export().(string)
		if !ok {
			panic(vm.NewTypeError("argument 1 is not a string, got %T", c.Argument(1).Export()))
		}
		backText, ok := c.Argument(2).Export().(string)
		if !ok {
			panic(vm.NewTypeError("argument 2 is not a string, got %T", c.Argument(2).Export()))
		}
		return vm.ToValue(pl.EditSign(pos, frontText, backText))
	})
	setExec("effect", func(c goja.FunctionCall, pl *player.Player) goja.Value {
		t, ok := c.Argument(0).Export().(effect.Type)
		if !ok {
			panic(vm.NewTypeError("argument 0 is not an effect.Type"))
		}
		e, ok := pl.Effect(t)
		if !ok {
			return nil
		}
		return vm.ToValue(e)
	})
	setExec("effects", func(c goja.FunctionCall, pl *player.Player) goja.Value {
		return vm.ToValue(pl.Effects())
	})
	setExec("enableInstantRespawn", func(c goja.FunctionCall, pl *player.Player) goja.Value {
		pl.EnableInstantRespawn()
		return nil
	})
	setExec("enchantmentSeed", func(c goja.FunctionCall, pl *player.Player) goja.Value {
		return vm.ToValue(pl.EnchantmentSeed())
	})
	setExec("enderChestInventory", func(c goja.FunctionCall, pl *player.Player) goja.Value {
		return vm.ToValue(pl.EnderChestInventory())
	})
	setExec("executeCommand", func(c goja.FunctionCall, pl *player.Player) goja.Value {
		commandLine, ok := c.Argument(0).Export().(string)
		if !ok {
			panic(vm.NewTypeError("argument 0 is not a string, got %T", c.Argument(0).Export()))
		}
		pl.ExecuteCommand(commandLine)
		return nil
	})
	setExec("exhaust", func(c goja.FunctionCall, pl *player.Player) goja.Value {
		points, ok := c.Argument(0).Export().(float64)
		if !ok {
			panic(vm.NewTypeError("argument 0 is not a float64, got %T", c.Argument(0).Export()))
		}
		pl.Exhaust(points)
		return nil
	})
	setExec("experience", func(c goja.FunctionCall, pl *player.Player) goja.Value {
		return vm.ToValue(pl.Experience())
	})
	setExec("experienceLevel", func(c goja.FunctionCall, pl *player.Player) goja.Value {
		return vm.ToValue(pl.ExperienceLevel())
	})
	setExec("experienceProgress", func(c goja.FunctionCall, pl *player.Player) goja.Value {
		return vm.ToValue(pl.ExperienceProgress())
	})
	setExec("explode", func(c goja.FunctionCall, pl *player.Player) goja.Value {
		explosionPos, ok := c.Argument(0).Export().(mgl64.Vec3)
		if !ok {
			panic(vm.NewTypeError("argument 0 is not a mgl64.Vec3, got %T", c.Argument(0).Export()))
		}
		impact, ok := c.Argument(1).Export().(float64)
		if !ok {
			panic(vm.NewTypeError("argument 1 is not a float64, got %T", c.Argument(1).Export()))
		}
		conf, ok := c.Argument(2).Export().(block.ExplosionConfig)
		if !ok {
			panic(vm.NewTypeError("argument 2 is not a block.ExplosionConfig, got %T", c.Argument(2).Export()))
		}
		pl.Explode(explosionPos, impact, conf)
		return nil
	})
	setExec("extinguish", func(c goja.FunctionCall, pl *player.Player) goja.Value {
		pl.Extinguish()
		return nil
	})
	setExec("eyeHeight", func(c goja.FunctionCall, pl *player.Player) goja.Value {
		return vm.ToValue(pl.EyeHeight())
	})
	setExec("fallDistance", func(c goja.FunctionCall, pl *player.Player) goja.Value {
		return vm.ToValue(pl.FallDistance())
	})
	setExec("finalDamageFrom", func(c goja.FunctionCall, pl *player.Player) goja.Value {
		dmg, ok := c.Argument(0).Export().(float64)
		if !ok {
			panic(vm.NewTypeError("argument 0 is not a float64, got %T", c.Argument(0).Export()))
		}
		src, ok := c.Argument(1).Export().(world.DamageSource)
		if !ok {
			panic(vm.NewTypeError("argument 1 is not a world.DamageSource, got %T", c.Argument(1).Export()))
		}
		return vm.ToValue(pl.FinalDamageFrom(dmg, src))
	})
	setExec("finishBreaking", func(c goja.FunctionCall, pl *player.Player) goja.Value {
		pl.FinishBreaking()
		return nil
	})
	setExec("fireProof", func(c goja.FunctionCall, pl *player.Player) goja.Value {
		return vm.ToValue(pl.FireProof())
	})
	setExec("flightSpeed", func(c goja.FunctionCall, pl *player.Player) goja.Value {
		return vm.ToValue(pl.FlightSpeed())
	})
	setExec("flying", func(c goja.FunctionCall, pl *player.Player) goja.Value {
		return vm.ToValue(pl.Flying())
	})
	setExec("food", func(c goja.FunctionCall, pl *player.Player) goja.Value {
		return vm.ToValue(pl.Food())
	})
	setExec("gameMode", func(c goja.FunctionCall, pl *player.Player) goja.Value {
		return vm.ToValue(pl.GameMode())
	})
	setExec("gliding", func(c goja.FunctionCall, pl *player.Player) goja.Value {
		return vm.ToValue(pl.Gliding())
	})
	setExec("hasCooldown", func(c goja.FunctionCall, pl *player.Player) goja.Value {
		i, ok := c.Argument(0).Export().(world.Item)
		if !ok {
			panic(vm.NewTypeError("argument 0 is not a world.Item, got %T", c.Argument(0).Export()))
		}
		return vm.ToValue(pl.HasCooldown(i))
	})
	setExec("heal", func(c goja.FunctionCall, pl *player.Player) goja.Value {
		health, ok := c.Argument(0).Export().(float64)
		if !ok {
			panic(vm.NewTypeError("argument 0 is not a float64, got %T", c.Argument(0).Export()))
		}
		source, ok := c.Argument(1).Export().(world.HealingSource)
		if !ok {
			panic(vm.NewTypeError("argument 1 is not a world.HealingSource, got %T", c.Argument(1).Export()))
		}
		pl.Heal(health, source)
		return nil
	})
	setExec("health", func(c goja.FunctionCall, pl *player.Player) goja.Value {
		return vm.ToValue(pl.Health())
	})
	setExec("heldItems", func(c goja.FunctionCall, pl *player.Player) goja.Value {
		mainHand, offHand := pl.HeldItems()
		return vm.ToValue([]item.Stack{mainHand, offHand})
	})
	setExec("hideCoordinates", func(c goja.FunctionCall, pl *player.Player) goja.Value {
		pl.HideCoordinates()
		return nil
	})
	setExec("hideEntity", func(c goja.FunctionCall, pl *player.Player) goja.Value {
		e, ok := c.Argument(0).Export().(world.Entity)
		if !ok {
			panic(vm.NewTypeError("argument 0 is not a world.Entity, got %T", c.Argument(0).Export()))
		}
		pl.HideEntity(e)
		return nil
	})
	setExec("hurt", func(c goja.FunctionCall, pl *player.Player) goja.Value {
		dmg, ok := c.Argument(0).Export().(float64)
		if !ok {
			panic(vm.NewTypeError("argument 0 is not a float64, got %T", c.Argument(0).Export()))
		}
		src, ok := c.Argument(1).Export().(world.DamageSource)
		if !ok {
			panic(vm.NewTypeError("argument 1 is not a world.DamageSource, got %T", c.Argument(1).Export()))
		}
		pl.Hurt(dmg, src)
		return nil
	})
	setExec("immobile", func(c goja.FunctionCall, pl *player.Player) goja.Value {
		return vm.ToValue(pl.Immobile())
	})
	setExec("inventory", func(c goja.FunctionCall, pl *player.Player) goja.Value {
		return vm.ToValue(pl.Inventory())
	})
	setExec("invisible", func(c goja.FunctionCall, pl *player.Player) goja.Value {
		return vm.ToValue(pl.Invisible())
	})
	setExec("jump", func(c goja.FunctionCall, pl *player.Player) goja.Value {
		pl.Jump()
		return nil
	})
	setExec("knockBack", func(c goja.FunctionCall, pl *player.Player) goja.Value {
		src, ok := c.Argument(0).Export().(mgl64.Vec3)
		if !ok {
			panic(vm.NewTypeError("argument 0 is not a mgl64.Vec3, got %T", c.Argument(0).Export()))
		}
		force, ok := c.Argument(1).Export().(float64)
		if !ok {
			panic(vm.NewTypeError("argument 1 is not a float64, got %T", c.Argument(1).Export()))
		}
		height, ok := c.Argument(2).Export().(float64)
		if !ok {
			panic(vm.NewTypeError("argument 2 is not a float64, got %T", c.Argument(2).Export()))
		}
		pl.KnockBack(src, force, height)
		return nil
	})
	setExec("latency", func(c goja.FunctionCall, pl *player.Player) goja.Value {
		return vm.ToValue(pl.Latency())
	})
	setExec("locale", func(c goja.FunctionCall, p *player.Player) goja.Value {
		return vm.ToValue(p.Locale().String())
	})
	setExec("maxAirSupply", func(c goja.FunctionCall, pl *player.Player) goja.Value {
		return vm.ToValue(pl.MaxAirSupply())
	})
	setExec("maxHealth", func(c goja.FunctionCall, pl *player.Player) goja.Value {
		return vm.ToValue(pl.MaxHealth())
	})
	setExec("message", func(c goja.FunctionCall, p *player.Player) goja.Value {
		var args []any
		for _, argument := range c.Arguments {
			args = append(args, argument.Export())
		}
		p.Message(args...)
		return nil
	})
	setExec("messagef", func(c goja.FunctionCall, p *player.Player) goja.Value {
		format, ok := c.Argument(0).Export().(string)
		if !ok {
			panic(vm.NewTypeError("argument 0 is not a string, got %T", c.Argument(0).Export()))
		}
		var args []any
		if len(c.Arguments) > 1 {
			for _, argument := range c.Arguments[1:] {
				args = append(args, argument.Export())
			}
		}
		p.Messagef(format, args...)
		return nil
	})
	setExec("messaget", func(c goja.FunctionCall, p *player.Player) goja.Value {
		t, ok := c.Argument(0).Export().(chat.Translation)
		if !ok {
			panic(vm.NewTypeError("argument 0 is not a chat.Translation, got %T", c.Argument(0).Export()))
		}
		var args []any
		if len(c.Arguments) > 1 {
			for _, argument := range c.Arguments[1:] {
				args = append(args, argument.Export())
			}
		}
		p.Messaget(t, args...)
		return nil
	})
	setExec("move", func(c goja.FunctionCall, pl *player.Player) goja.Value {
		deltaPos, ok := c.Argument(0).Export().(mgl64.Vec3)
		if !ok {
			panic(vm.NewTypeError("argument 0 is not a mgl64.Vec3, got %T", c.Argument(0).Export()))
		}
		deltaYaw, ok := c.Argument(1).Export().(float64)
		if !ok {
			panic(vm.NewTypeError("argument 1 is not a float64, got %T", c.Argument(1).Export()))
		}
		deltaPitch, ok := c.Argument(2).Export().(float64)
		if !ok {
			panic(vm.NewTypeError("argument 2 is not a float64, got %T", c.Argument(2).Export()))
		}
		pl.Move(deltaPos, deltaYaw, deltaPitch)
		return nil
	})
	setExec("moveItemsToInventory", func(c goja.FunctionCall, pl *player.Player) goja.Value {
		pl.MoveItemsToInventory()
		return nil
	})
	setExec("nameTag", func(c goja.FunctionCall, pl *player.Player) goja.Value {
		return vm.ToValue(pl.NameTag())
	})
	setExec("onFireDuration", func(c goja.FunctionCall, pl *player.Player) goja.Value {
		return vm.ToValue(pl.OnFireDuration())
	})
	setExec("onGround", func(c goja.FunctionCall, pl *player.Player) goja.Value {
		return vm.ToValue(pl.OnGround())
	})
	setExec("openBlockContainer", func(c goja.FunctionCall, pl *player.Player) goja.Value {
		pos, ok := c.Argument(0).Export().(cube.Pos)
		if !ok {
			panic(vm.NewTypeError("argument 0 is not a cube.Pos, got %T", c.Argument(0).Export()))
		}
		tx, ok := c.Argument(1).Export().(*world.Tx)
		if !ok {
			panic(vm.NewTypeError("argument 1 is not a *world.Tx, got %T", c.Argument(1).Export()))
		}
		pl.OpenBlockContainer(pos, tx)
		return nil
	})
	setExec("openSign", func(c goja.FunctionCall, pl *player.Player) goja.Value {
		pos, ok := c.Argument(0).Export().(cube.Pos)
		if !ok {
			panic(vm.NewTypeError("argument 0 is not a cube.Pos, got %T", c.Argument(0).Export()))
		}
		frontSide, ok := c.Argument(1).Export().(bool)
		if !ok {
			panic(vm.NewTypeError("argument 1 is not a bool, got %T", c.Argument(1).Export()))
		}
		pl.OpenSign(pos, frontSide)
		return nil
	})
	setExec("pickBlock", func(c goja.FunctionCall, pl *player.Player) goja.Value {
		pos, ok := c.Argument(0).Export().(cube.Pos)
		if !ok {
			panic(vm.NewTypeError("argument 0 is not a cube.Pos, got %T", c.Argument(0).Export()))
		}
		pl.PickBlock(pos)
		return nil
	})
	setExec("placeBlock", func(c goja.FunctionCall, pl *player.Player) goja.Value {
		pos, ok := c.Argument(0).Export().(cube.Pos)
		if !ok {
			panic(vm.NewTypeError("argument 0 is not a cube.Pos, got %T", c.Argument(0).Export()))
		}
		b, ok := c.Argument(1).Export().(world.Block)
		if !ok {
			panic(vm.NewTypeError("argument 1 is not a block.Block, got %T", c.Argument(1).Export()))
		}
		ctx, ok := c.Argument(2).Export().(*item.UseContext)
		if !ok {
			panic(vm.NewTypeError("argument 2 is not a *item.UseContext, got %T", c.Argument(2).Export()))
		}
		pl.PlaceBlock(pos, b, ctx)
		return nil
	})
	setExec("playSound", func(c goja.FunctionCall, pl *player.Player) goja.Value {
		sound, ok := c.Argument(0).Export().(world.Sound)
		if !ok {
			panic(vm.NewTypeError("argument 0 is not a world.Sound, got %T", c.Argument(0).Export()))
		}
		pl.PlaySound(sound)
		return nil
	})
	setExec("position", func(c goja.FunctionCall, pl *player.Player) goja.Value {
		return vm.ToValue(pl.Position())
	})
	setExec("punchAir", func(c goja.FunctionCall, pl *player.Player) goja.Value {
		pl.PunchAir()
		return nil
	})
	setExec("releaseItem", func(c goja.FunctionCall, pl *player.Player) goja.Value {
		pl.ReleaseItem()
		return nil
	})
	setExec("removeBossBar", func(c goja.FunctionCall, pl *player.Player) goja.Value {
		pl.RemoveBossBar()
		return nil
	})
	setExec("removeEffect", func(c goja.FunctionCall, pl *player.Player) goja.Value {
		t, ok := c.Argument(0).Export().(effect.Type)
		if !ok {
			panic(vm.NewTypeError("argument 0 is not an effect.Type"))
		}
		pl.RemoveEffect(t)
		return nil
	})
	setExec("removeExperience", func(c goja.FunctionCall, pl *player.Player) goja.Value {
		amount, ok := c.Argument(0).Export().(int64)
		if !ok {
			panic(vm.NewTypeError("argument 0 is not an int64"))
		}
		pl.RemoveExperience(int(amount))
		return nil
	})
	setExec("removeScoreboard", func(c goja.FunctionCall, pl *player.Player) goja.Value {
		pl.RemoveScoreboard()
		return nil
	})
	setExec("resetEnchantmentSeed", func(c goja.FunctionCall, pl *player.Player) goja.Value {
		pl.ResetEnchantmentSeed()
		return nil
	})
	setExec("resetFallDistance", func(c goja.FunctionCall, pl *player.Player) goja.Value {
		pl.ResetFallDistance()
		return nil
	})
	setExec("respawn", func(c goja.FunctionCall, pl *player.Player) goja.Value {
		return vm.ToValue(pl.Respawn())
	})
	setExec("rotation", func(c goja.FunctionCall, pl *player.Player) goja.Value {
		return vm.ToValue(pl.Rotation())
	})
	setExec("saturate", func(c goja.FunctionCall, pl *player.Player) goja.Value {
		food, ok := c.Argument(0).Export().(int64)
		if !ok {
			panic(vm.NewTypeError("argument 0 is not an int64"))
		}
		saturation, ok := c.Argument(1).Export().(float64)
		if !ok {
			panic(vm.NewTypeError("argument 1 is not a float64, got %T", c.Argument(1).Export()))
		}
		pl.Saturate(int(food), saturation)
		return nil
	})
	setExec("scale", func(c goja.FunctionCall, pl *player.Player) goja.Value {
		return vm.ToValue(pl.Scale())
	})
	setExec("scoreTag", func(c goja.FunctionCall, pl *player.Player) goja.Value {
		return vm.ToValue(pl.ScoreTag())
	})
	setExec("sendBossBar", func(c goja.FunctionCall, pl *player.Player) goja.Value {
		bar, ok := c.Argument(0).Export().(bossbar.BossBar)
		if !ok {
			panic(vm.NewTypeError("argument 0 is not a bossbar.BossBar, got %T", c.Argument(0).Export()))
		}
		pl.SendBossBar(bar)
		return nil
	})
	setExec("sendCommandOutput", func(c goja.FunctionCall, pl *player.Player) goja.Value {
		output, ok := c.Argument(0).Export().(*cmd.Output)
		if !ok {
			panic(vm.NewTypeError("argument 0 is not a *cmd.Output, got %T", c.Argument(0).Export()))
		}
		pl.SendCommandOutput(output)
		return nil
	})
	setExec("sendDialogue", func(c goja.FunctionCall, pl *player.Player) goja.Value {
		d, ok := c.Argument(0).Export().(dialogue.Dialogue)
		if !ok {
			panic(vm.NewTypeError("argument 0 is not a dialogue.Dialogue, got %T", c.Argument(0).Export()))
		}
		e, ok := c.Argument(1).Export().(world.Entity)
		if !ok {
			panic(vm.NewTypeError("argument 1 is not a world.Entity, got %T", c.Argument(1).Export()))
		}
		pl.SendDialogue(d, e)
		return nil
	})
	setExec("sendForm", func(c goja.FunctionCall, pl *player.Player) goja.Value {
		f, ok := c.Argument(0).Export().(form.Form)
		if !ok {
			panic(vm.NewTypeError("argument 0 is not a form.Form, got %T", c.Argument(0).Export()))
		}
		pl.SendForm(f)
		return nil
	})
	setExec("sendJukeboxPopup", func(c goja.FunctionCall, pl *player.Player) goja.Value {
		var args []any
		for _, argument := range c.Arguments {
			args = append(args, argument.Export())
		}
		pl.SendJukeboxPopup(args)
		return nil
	})
	setExec("sendPopup", func(c goja.FunctionCall, pl *player.Player) goja.Value {
		var args []any
		for _, argument := range c.Arguments {
			args = append(args, argument.Export())
		}
		pl.SendPopup(args)
		return nil
	})
	setExec("sendScoreboard", func(c goja.FunctionCall, pl *player.Player) goja.Value {
		sb, ok := c.Argument(0).Export().(*scoreboard.Scoreboard)
		if !ok {
			panic(vm.NewTypeError("argument 0 is not a *scoreboard.Scoreboard, got %T", c.Argument(0).Export()))
		}
		pl.SendScoreboard(sb)
		return nil
	})
	setExec("sendTip", func(c goja.FunctionCall, pl *player.Player) goja.Value {
		var args []any
		for _, argument := range c.Arguments {
			args = append(args, argument.Export())
		}
		pl.SendTip(args)
		return nil
	})
	setExec("sendTitle", func(c goja.FunctionCall, pl *player.Player) goja.Value {
		t, ok := c.Argument(0).Export().(title.Title)
		if !ok {
			panic(vm.NewTypeError("argument 0 is not a title.Title, got %T", c.Argument(0).Export()))
		}
		pl.SendTitle(t)
		return nil
	})
	setExec("sendToast", func(c goja.FunctionCall, pl *player.Player) goja.Value {
		title, ok := c.Argument(0).Export().(string)
		if !ok {
			panic(vm.NewTypeError("argument 0 is not a string, got %T", c.Argument(0).Export()))
		}
		message, ok := c.Argument(1).Export().(string)
		if !ok {
			panic(vm.NewTypeError("argument 1 is not a string, got %T", c.Argument(1).Export()))
		}
		pl.SendToast(title, message)
		return nil
	})
	setExec("setAbsorption", func(c goja.FunctionCall, pl *player.Player) goja.Value {
		health, ok := c.Argument(0).Export().(float64)
		if !ok {
			panic(vm.NewTypeError("argument 0 is not a float64, got %T", c.Argument(0).Export()))
		}
		pl.SetAbsorption(health)
		return nil
	})
	setExec("setAirSupply", func(c goja.FunctionCall, pl *player.Player) goja.Value {
		duration, ok := c.Argument(0).Export().(time.Duration)
		if !ok {
			panic(vm.NewTypeError("argument 0 is not a time.Duration, got %T", c.Argument(0).Export()))
		}
		pl.SetAirSupply(duration)
		return nil
	})
	setExec("setCooldown", func(c goja.FunctionCall, pl *player.Player) goja.Value {
		item, ok := c.Argument(0).Export().(world.Item)
		if !ok {
			panic(vm.NewTypeError("argument 0 is not a world.Item, got %T", c.Argument(0).Export()))
		}
		cooldown, ok := c.Argument(1).Export().(time.Duration)
		if !ok {
			panic(vm.NewTypeError("argument 1 is not a time.Duration, got %T", c.Argument(1).Export()))
		}
		pl.SetCooldown(item, cooldown)
		return nil
	})
	setExec("setExperienceLevel", func(c goja.FunctionCall, pl *player.Player) goja.Value {
		level, ok := c.Argument(0).Export().(int64)
		if !ok {
			panic(vm.NewTypeError("argument 0 is not an int64"))
		}
		pl.SetExperienceLevel(int(level))
		return nil
	})
	setExec("setExperienceProgress", func(c goja.FunctionCall, pl *player.Player) goja.Value {
		progress, ok := c.Argument(0).Export().(float64)
		if !ok {
			panic(vm.NewTypeError("argument 0 is not a float64, got %T", c.Argument(0).Export()))
		}
		pl.SetExperienceProgress(progress)
		return nil
	})
	setExec("setFlightSpeed", func(c goja.FunctionCall, pl *player.Player) goja.Value {
		flightSpeed, ok := c.Argument(0).Export().(float64)
		if !ok {
			panic(vm.NewTypeError("argument 0 is not a float64, got %T", c.Argument(0).Export()))
		}
		pl.SetFlightSpeed(flightSpeed)
		return nil
	})
	setExec("setFood", func(c goja.FunctionCall, pl *player.Player) goja.Value {
		level, ok := c.Argument(0).Export().(int64)
		if !ok {
			panic(vm.NewTypeError("argument 0 is not an int64"))
		}
		pl.SetFood(int(level))
		return nil
	})
	setExec("setGameMode", func(c goja.FunctionCall, pl *player.Player) goja.Value {
		mode, ok := c.Argument(0).Export().(world.GameMode)
		if !ok {
			panic(vm.NewTypeError("argument 0 is not a world.GameMode, got %T", c.Argument(0).Export()))
		}
		pl.SetGameMode(mode)
		return nil
	})
	setExec("setHeldItems", func(c goja.FunctionCall, pl *player.Player) goja.Value {
		mainHand, ok := c.Argument(0).Export().(item.Stack)
		if !ok {
			panic(vm.NewTypeError("argument 0 is not an item.Stack"))
		}
		offHand, ok := c.Argument(1).Export().(item.Stack)
		if !ok {
			panic(vm.NewTypeError("argument 1 is not an item.Stack"))
		}
		pl.SetHeldItems(mainHand, offHand)
		return nil
	})
	setExec("setHeldSlot", func(c goja.FunctionCall, pl *player.Player) goja.Value {
		to, ok := c.Argument(0).Export().(int64)
		if !ok {
			panic(vm.NewTypeError("argument 0 is not an int64"))
		}
		return vm.ToValue(pl.SetHeldSlot(int(to)))
	})
	setExec("setImmobile", func(c goja.FunctionCall, pl *player.Player) goja.Value {
		pl.SetImmobile()
		return nil
	})
	setExec("setInvisible", func(c goja.FunctionCall, pl *player.Player) goja.Value {
		pl.SetInvisible()
		return nil
	})
	setExec("setMaxAirSupply", func(c goja.FunctionCall, pl *player.Player) goja.Value {
		duration, ok := c.Argument(0).Export().(time.Duration)
		if !ok {
			panic(vm.NewTypeError("argument 0 is not a time.Duration, got %T", c.Argument(0).Export()))
		}
		pl.SetMaxAirSupply(duration)
		return nil
	})
	setExec("setMaxHealth", func(c goja.FunctionCall, pl *player.Player) goja.Value {
		health, ok := c.Argument(0).Export().(float64)
		if !ok {
			panic(vm.NewTypeError("argument 0 is not a float64, got %T", c.Argument(0).Export()))
		}
		pl.SetMaxHealth(health)
		return nil
	})
	setExec("setMobile", func(c goja.FunctionCall, pl *player.Player) goja.Value {
		pl.SetMobile()
		return nil
	})
	setExec("setNameTag", func(c goja.FunctionCall, pl *player.Player) goja.Value {
		name, ok := c.Argument(0).Export().(string)
		if !ok {
			panic(vm.NewTypeError("argument 0 is not a string, got %T", c.Argument(0).Export()))
		}
		pl.SetNameTag(name)
		return nil
	})
	setExec("setOnFire", func(c goja.FunctionCall, pl *player.Player) goja.Value {
		duration, ok := c.Argument(0).Export().(time.Duration)
		if !ok {
			panic(vm.NewTypeError("argument 0 is not a time.Duration, got %T", c.Argument(0).Export()))
		}
		pl.SetOnFire(duration)
		return nil
	})
	setExec("setScale", func(c goja.FunctionCall, pl *player.Player) goja.Value {
		s, ok := c.Argument(0).Export().(float64)
		if !ok {
			panic(vm.NewTypeError("argument 0 is not a float64, got %T", c.Argument(0).Export()))
		}
		pl.SetScale(s)
		return nil
	})
	setExec("setScoreTag", func(c goja.FunctionCall, pl *player.Player) goja.Value {
		var args []any
		for _, argument := range c.Arguments {
			args = append(args, argument.Export())
		}
		pl.SetScoreTag(args...)
		return nil
	})
	setExec("setSkin", func(c goja.FunctionCall, pl *player.Player) goja.Value {
		s, ok := c.Argument(0).Export().(skin.Skin)
		if !ok {
			panic(vm.NewTypeError("argument 0 is not a skin.Skin, got %T", c.Argument(0).Export()))
		}
		pl.SetSkin(s)
		return nil
	})
	setExec("setSpeed", func(c goja.FunctionCall, pl *player.Player) goja.Value {
		speed, ok := c.Argument(0).Export().(float64)
		if !ok {
			panic(vm.NewTypeError("argument 0 is not a float64, got %T", c.Argument(0).Export()))
		}
		pl.SetSpeed(speed)
		return nil
	})
	setExec("setVelocity", func(c goja.FunctionCall, pl *player.Player) goja.Value {
		velocity, ok := c.Argument(0).Export().(mgl64.Vec3)
		if !ok {
			panic(vm.NewTypeError("argument 0 is not a mgl64.Vec3, got %T", c.Argument(0).Export()))
		}
		pl.SetVelocity(velocity)
		return nil
	})
	setExec("setVisible", func(c goja.FunctionCall, pl *player.Player) goja.Value {
		pl.SetVisible()
		return nil
	})
	setExec("showCoordinates", func(c goja.FunctionCall, pl *player.Player) goja.Value {
		pl.ShowCoordinates()
		return nil
	})
	setExec("showEntity", func(c goja.FunctionCall, pl *player.Player) goja.Value {
		e, ok := c.Argument(0).Export().(world.Entity)
		if !ok {
			panic(vm.NewTypeError("argument 0 is not a world.Entity, got %T", c.Argument(0).Export()))
		}
		pl.ShowEntity(e)
		return nil
	})
	setExec("showParticle", func(c goja.FunctionCall, pl *player.Player) goja.Value {
		pos, ok := c.Argument(0).Export().(mgl64.Vec3)
		if !ok {
			panic(vm.NewTypeError("argument 0 is not a mgl64.Vec3, got %T", c.Argument(0).Export()))
		}
		particle, ok := c.Argument(1).Export().(world.Particle)
		if !ok {
			panic(vm.NewTypeError("argument 1 is not a world.Particle, got %T", c.Argument(1).Export()))
		}
		pl.ShowParticle(pos, particle)
		return nil
	})
	setExec("skin", func(c goja.FunctionCall, pl *player.Player) goja.Value {
		return vm.ToValue(pl.Skin())
	})
	setExec("sneaking", func(c goja.FunctionCall, pl *player.Player) goja.Value {
		return vm.ToValue(pl.Sneaking())
	})
	setExec("speed", func(c goja.FunctionCall, pl *player.Player) goja.Value {
		return vm.ToValue(pl.Speed())
	})
	setExec("splash", func(c goja.FunctionCall, pl *player.Player) goja.Value {
		pos, ok := c.Argument(0).Export().(mgl64.Vec3)
		if !ok {
			panic(vm.NewTypeError("argument 0 is not a mgl64.Vec3, got %T", c.Argument(1).Export()))
		}
		pl.Splash(pl.Tx(), pos)
		return nil
	})
	setExec("sprinting", func(c goja.FunctionCall, pl *player.Player) goja.Value {
		return vm.ToValue(pl.Sprinting())
	})
	setExec("startBreaking", func(c goja.FunctionCall, pl *player.Player) goja.Value {
		pos, ok := c.Argument(0).Export().(cube.Pos)
		if !ok {
			panic(vm.NewTypeError("argument 0 is not a cube.Pos, got %T", c.Argument(0).Export()))
		}
		face, ok := c.Argument(1).Export().(cube.Face)
		if !ok {
			panic(vm.NewTypeError("argument 1 is not a cube.Face, got %T", c.Argument(1).Export()))
		}
		pl.StartBreaking(pos, face)
		return nil
	})
	setExec("startCrawling", func(c goja.FunctionCall, pl *player.Player) goja.Value {
		pl.StartCrawling()
		return nil
	})
	setExec("startFlying", func(c goja.FunctionCall, pl *player.Player) goja.Value {
		pl.StartFlying()
		return nil
	})
	setExec("startGliding", func(c goja.FunctionCall, pl *player.Player) goja.Value {
		pl.StartGliding()
		return nil
	})
	setExec("startSneaking", func(c goja.FunctionCall, pl *player.Player) goja.Value {
		pl.StartSneaking()
		return nil
	})
	setExec("startSprinting", func(c goja.FunctionCall, pl *player.Player) goja.Value {
		pl.StartSprinting()
		return nil
	})
	setExec("startSwimming", func(c goja.FunctionCall, pl *player.Player) goja.Value {
		pl.StartSwimming()
		return nil
	})
	setExec("stopCrawling", func(c goja.FunctionCall, pl *player.Player) goja.Value {
		pl.StopCrawling()
		return nil
	})
	setExec("stopFlying", func(c goja.FunctionCall, pl *player.Player) goja.Value {
		pl.StopFlying()
		return nil
	})
	setExec("stopGliding", func(c goja.FunctionCall, pl *player.Player) goja.Value {
		pl.StopGliding()
		return nil
	})
	setExec("stopSneaking", func(c goja.FunctionCall, pl *player.Player) goja.Value {
		pl.StopSneaking()
		return nil
	})
	setExec("stopSprinting", func(c goja.FunctionCall, pl *player.Player) goja.Value {
		pl.StopSprinting()
		return nil
	})
	setExec("stopSwimming", func(c goja.FunctionCall, pl *player.Player) goja.Value {
		pl.StopSwimming()
		return nil
	})
	setExec("swimming", func(c goja.FunctionCall, pl *player.Player) goja.Value {
		return vm.ToValue(pl.Swimming())
	})
	setExec("swingArm", func(c goja.FunctionCall, pl *player.Player) goja.Value {
		pl.SwingArm()
		return nil
	})
	setExec("teleport", func(c goja.FunctionCall, pl *player.Player) goja.Value {
		pos, ok := c.Argument(0).Export().(mgl64.Vec3)
		if !ok {
			panic(vm.NewTypeError("argument 0 is not a mgl64.Vec3, got %T", c.Argument(0).Export()))
		}
		pl.Teleport(pos)
		return nil
	})
	setExec("tick", func(c goja.FunctionCall, pl *player.Player) goja.Value {
		current, ok := c.Argument(0).Export().(int64)
		if !ok {
			panic(vm.NewTypeError("argument 0 is not an int64"))
		}
		pl.Tick(pl.Tx(), current)
		return nil
	})
	setExec("transfer", func(c goja.FunctionCall, pl *player.Player) goja.Value {
		address, ok := c.Argument(0).Export().(string)
		if !ok {
			panic(vm.NewTypeError("argument 0 is not a string, got %T", c.Argument(0).Export()))
		}
		return vm.ToValue(pl.Transfer(address))
	})
	setExec("turnLecternPage", func(c goja.FunctionCall, pl *player.Player) goja.Value {
		pos, ok := c.Argument(0).Export().(cube.Pos)
		if !ok {
			panic(vm.NewTypeError("argument 0 is not a cube.Pos, got %T", c.Argument(0).Export()))
		}
		page, ok := c.Argument(1).Export().(int64)
		if !ok {
			panic(vm.NewTypeError("argument 1 is not an int64"))
		}
		return vm.ToValue(pl.TurnLecternPage(pos, int(page)))
	})
	setExec("updateDiagnostics", func(c goja.FunctionCall, pl *player.Player) goja.Value {
		d, ok := c.Argument(0).Export().(session.Diagnostics)
		if !ok {
			panic(vm.NewTypeError("argument 0 is not a session.Diagnostics, got %T", c.Argument(0).Export()))
		}
		pl.UpdateDiagnostics(d)
		return nil
	})
	setExec("useItem", func(c goja.FunctionCall, pl *player.Player) goja.Value {
		pl.UseItem()
		return nil
	})
	setExec("useItemOnBlock", func(c goja.FunctionCall, pl *player.Player) goja.Value {
		pos, ok := c.Argument(0).Export().(cube.Pos)
		if !ok {
			panic(vm.NewTypeError("argument 0 is not a cube.Pos, got %T", c.Argument(0).Export()))
		}
		face, ok := c.Argument(1).Export().(cube.Face)
		if !ok {
			panic(vm.NewTypeError("argument 1 is not a cube.Face, got %T", c.Argument(1).Export()))
		}
		clickPos, ok := c.Argument(2).Export().(mgl64.Vec3)
		if !ok {
			panic(vm.NewTypeError("argument 2 is not a mgl64.Vec3, got %T", c.Argument(2).Export()))
		}
		pl.UseItemOnBlock(pos, face, clickPos)
		return nil
	})
	setExec("useItemOnEntity", func(c goja.FunctionCall, pl *player.Player) goja.Value {
		e, ok := c.Argument(0).Export().(world.Entity)
		if !ok {
			panic(vm.NewTypeError("argument 0 is not a world.Entity, got %T", c.Argument(0).Export()))
		}
		return vm.ToValue(pl.UseItemOnEntity(e))
	})
	setExec("usingItem", func(c goja.FunctionCall, pl *player.Player) goja.Value {
		return vm.ToValue(pl.UsingItem())
	})
	setExec("velocity", func(c goja.FunctionCall, pl *player.Player) goja.Value {
		return vm.ToValue(pl.Velocity())
	})

	events := []string{
		"onItemDrop", "onHeldSlotChange", "onMove", "onJump", "onTeleport", "onChangeWorld", "onToggleSprint",
		"onToggleSneak", "onCommandExecution", "onTransfer", "onChat", "onSkinChange", "onFireExtinguish",
		"onStartBreak", "onBlockBreak", "onBlockPlace", "onBlockPick", "onSignEdit", "onLecternPageTurn",
		"onItemPickup", "onItemUse", "onItemUseOnBlock", "onItemUseOnEntity", "onItemRelease", "onItemConsume",
		"onItemDamage", "onAttackEntity", "onExperienceGain", "onPunchAir", "onHurt", "onHeal", "onFoodLoss",
		"onDeath", "onRespawn", "onQuit", "onDiagnostics",
	}
	for _, name := range events {
		p.events[name] = make([]callable, 0)
		obj.Method(name, func(c goja.FunctionCall) goja.Value {
			f, ok := goja.AssertFunction(c.Argument(0))
			if !ok {
				panic(vm.NewTypeError("argument 0 is not a function, got %T", c.Argument(0).Export()))
			}
			p.events[name] = append(p.events[name], callable{f, c.This})
			return nil
		})
	}

	var err error
	p.obj, err = obj.Build(vm)
	if err != nil {
		panic(err)
	}
}

func (p *Player) exec(f func(p *player.Player)) {
	if p.p != nil {
		f(p.p)
		return
	}
	p.h.ExecWorld(func(tx *world.Tx, e world.Entity) {
		p.p = e.(*player.Player)
		f(p.p)
		p.p = nil
	})
}

func (p *Player) callEvent(name string, pl *player.Player, e any) {
	p.p = pl
	defer func() {
		p.p = nil
	}()

	for _, c := range p.events[name] {
		_, err := c.f(c.this, p.r.vm.ToValue(e))
		if err != nil {
			panic(err)
		}
	}
}

func (p *Player) callEventCtx(name string, ctx *event.Context[*player.Player], e any) {
	p.p = ctx.Val()
	defer func() {
		p.p = nil
	}()

	for _, c := range p.events[name] {
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
