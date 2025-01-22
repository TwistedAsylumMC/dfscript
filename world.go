package dfscript

import (
	"github.com/df-mc/dragonfly/server/block/cube"
	"github.com/df-mc/dragonfly/server/world"
	"github.com/dop251/goja"
	"github.com/go-gl/mathgl/mgl64"
	"time"
)

type World struct {
	r   *Runtime
	obj *goja.Object

	w  *world.World
	tx *world.Tx

	events map[string][]callable
}

func NewWorld(r *Runtime, wo *world.World) *World {
	w := &World{
		r: r,
		w: wo,

		events: make(map[string][]callable),
	}
	w.setup()
	return w
}

func (w *World) setup() {
	vm := w.r.vm

	obj := newObject().
		Const("name", w.w.Name()).
		Const("dimension", w.w.Dimension()).
		Const("range", w.w.Range()).
		Method("time", func(c goja.FunctionCall) goja.Value {
			return vm.ToValue(w.w.Time())
		}).
		Method("setTime", func(c goja.FunctionCall) goja.Value {
			t, ok := c.Argument(0).Export().(int64)
			if !ok {
				panic(vm.NewTypeError("argument 0 must be an int64, got %T", c.Argument(0).Export()))
			}
			w.w.SetTime(int(t))
			return nil
		}).
		Method("stopTime", func(c goja.FunctionCall) goja.Value {
			w.w.StopTime()
			return nil
		}).
		Method("startTime", func(c goja.FunctionCall) goja.Value {
			w.w.StartTime()
			return nil
		}).
		Method("spawn", func(c goja.FunctionCall) goja.Value {
			return vm.ToValue(w.w.Spawn())
		}).
		Method("setSpawn", func(c goja.FunctionCall) goja.Value {
			pos, ok := c.Argument(0).Export().(cube.Pos)
			if !ok {
				panic(vm.NewTypeError("argument 0 must be a mgl64.Vec3, got %T", c.Argument(0).Export()))
			}
			w.w.SetSpawn(pos)
			return nil
		}).
		Method("playerSpawn", func(c goja.FunctionCall) goja.Value {
			id := w.r.exportUUID(c, 0)
			return vm.ToValue(w.w.PlayerSpawn(id))
		}).
		Method("setPlayerSpawn", func(c goja.FunctionCall) goja.Value {
			id := w.r.exportUUID(c, 0)
			pos, ok := c.Argument(1).Export().(cube.Pos)
			if !ok {
				panic(vm.NewTypeError("argument 1 must be a cube.Pos, got %T", c.Argument(1).Export()))
			}
			w.w.SetPlayerSpawn(id, pos)
			return nil
		}).
		Method("defaultGameMode", func(c goja.FunctionCall) goja.Value {
			return vm.ToValue(w.w.DefaultGameMode())
		}).
		Method("setDefaultGameMode", func(c goja.FunctionCall) goja.Value {
			gm, ok := c.Argument(0).Export().(world.GameMode)
			if !ok {
				panic(vm.NewTypeError("argument 0 must be a world.GameMode, got %T", c.Argument(0).Export()))
			}
			w.w.SetDefaultGameMode(gm)
			return nil
		}).
		Method("setTickRange", func(c goja.FunctionCall) goja.Value {
			r, ok := c.Argument(0).Export().(int64)
			if !ok {
				panic(vm.NewTypeError("argument 0 must be an int64, got %T", c.Argument(0).Export()))
			}
			w.w.SetTickRange(int(r))
			return nil
		}).
		Method("difficulty", func(c goja.FunctionCall) goja.Value {
			return vm.ToValue(w.w.Difficulty())
		}).
		Method("setDifficulty", func(c goja.FunctionCall) goja.Value {
			d, ok := c.Argument(0).Export().(world.Difficulty)
			if !ok {
				panic(vm.NewTypeError("argument 0 must be a world.Difficulty, got %T", c.Argument(0).Export()))
			}
			w.w.SetDifficulty(d)
			return nil
		}).
		Method("save", func(c goja.FunctionCall) goja.Value {
			w.w.Save()
			return nil
		}).
		Method("close", func(c goja.FunctionCall) goja.Value {
			delete(w.r.worlds, w.w.Name())
			return vm.ToValue(w.w.Close())
		})
	// TODO: portalDestination

	setExec := func(name string, f func(c goja.FunctionCall, tx *world.Tx) goja.Value) {
		obj.Method(name, func(c goja.FunctionCall) goja.Value {
			var ret goja.Value
			w.w.Exec(func(tx *world.Tx) {
				f(c, tx)
			})
			return ret
		})
	}
	setExec("block", func(c goja.FunctionCall, tx *world.Tx) goja.Value {
		pos, ok := c.Argument(0).Export().(cube.Pos)
		if !ok {
			panic(vm.NewTypeError("argument 0 must be a cube.Pos, got %T", c.Argument(0).Export()))
		}
		return vm.ToValue(tx.Block(pos))
	})
	setExec("setBlock", func(c goja.FunctionCall, tx *world.Tx) goja.Value {
		pos, ok := c.Argument(0).Export().(cube.Pos)
		if !ok {
			panic(vm.NewTypeError("argument 0 must be a cube.Pos, got %T", c.Argument(0).Export()))
		}
		b, ok := c.Argument(1).Export().(world.Block)
		if !ok {
			panic(vm.NewTypeError("argument 1 must be a world.Block, got %T", c.Argument(1).Export()))
		}
		opts, _ := c.Argument(2).Export().(*world.SetOpts)
		tx.SetBlock(pos, b, opts)
		return nil
	})
	setExec("liquid", func(c goja.FunctionCall, tx *world.Tx) goja.Value {
		pos, ok := c.Argument(0).Export().(cube.Pos)
		if !ok {
			panic(vm.NewTypeError("argument 0 must be a cube.Pos, got %T", c.Argument(0).Export()))
		}
		l, ok := tx.Liquid(pos)
		if !ok {
			return nil
		}
		return vm.ToValue(l)
	})
	setExec("setLiquid", func(c goja.FunctionCall, tx *world.Tx) goja.Value {
		pos, ok := c.Argument(0).Export().(cube.Pos)
		if !ok {
			panic(vm.NewTypeError("argument 0 must be a cube.Pos, got %T", c.Argument(0).Export()))
		}
		l, ok := c.Argument(1).Export().(world.Liquid)
		if !ok {
			panic(vm.NewTypeError("argument 1 must be a world.Liquid, got %T", c.Argument(1).Export()))
		}
		tx.SetLiquid(pos, l)
		return nil
	})
	setExec("buildStructure", func(c goja.FunctionCall, tx *world.Tx) goja.Value {
		pos, ok := c.Argument(0).Export().(cube.Pos)
		if !ok {
			panic(vm.NewTypeError("argument 0 must be a cube.Pos, got %T", c.Argument(0).Export()))
		}
		s, ok := c.Argument(1).Export().(world.Structure)
		if !ok {
			panic(vm.NewTypeError("argument 1 must be a world.Structure, got %T", c.Argument(1).Export()))
		}
		tx.BuildStructure(pos, s)
		return nil
	})
	setExec("scheduleBlockUpdate", func(c goja.FunctionCall, tx *world.Tx) goja.Value {
		pos, ok := c.Argument(0).Export().(cube.Pos)
		if !ok {
			panic(vm.NewTypeError("argument 0 must be a cube.Pos, got %T", c.Argument(0).Export()))
		}
		b, ok := c.Argument(1).Export().(world.Block)
		if !ok {
			panic(vm.NewTypeError("argument 1 must be a world.Block, got %T", c.Argument(1).Export()))
		}
		delay, ok := c.Argument(2).Export().(int64)
		if !ok {
			panic(vm.NewTypeError("argument 2 must be an int64, got %T", c.Argument(2).Export()))
		}
		tx.ScheduleBlockUpdate(pos, b, time.Millisecond*time.Duration(delay))
		return nil
	})
	setExec("highestLightBlocker", func(c goja.FunctionCall, tx *world.Tx) goja.Value {
		x, ok := c.Argument(0).Export().(int64)
		if !ok {
			panic(vm.NewTypeError("argument 0 must be an int64, got %T", c.Argument(0).Export()))
		}
		z, ok := c.Argument(1).Export().(int64)
		if !ok {
			panic(vm.NewTypeError("argument 1 must be an int64, got %T", c.Argument(1).Export()))
		}
		return vm.ToValue(tx.HighestLightBlocker(int(x), int(z)))
	})
	setExec("highestBlock", func(c goja.FunctionCall, tx *world.Tx) goja.Value {
		x, ok := c.Argument(0).Export().(int64)
		if !ok {
			panic(vm.NewTypeError("argument 0 must be an int64, got %T", c.Argument(0).Export()))
		}
		z, ok := c.Argument(1).Export().(int64)
		if !ok {
			panic(vm.NewTypeError("argument 1 must be an int64, got %T", c.Argument(1).Export()))
		}
		return vm.ToValue(tx.HighestBlock(int(x), int(z)))
	})
	setExec("light", func(c goja.FunctionCall, tx *world.Tx) goja.Value {
		pos, ok := c.Argument(0).Export().(cube.Pos)
		if !ok {
			panic(vm.NewTypeError("argument 0 must be a cube.Pos, got %T", c.Argument(0).Export()))
		}
		return vm.ToValue(tx.Light(pos))
	})
	setExec("skyLight", func(c goja.FunctionCall, tx *world.Tx) goja.Value {
		pos, ok := c.Argument(0).Export().(cube.Pos)
		if !ok {
			panic(vm.NewTypeError("argument 0 must be a cube.Pos, got %T", c.Argument(0).Export()))
		}
		return vm.ToValue(tx.SkyLight(pos))
	})
	setExec("biome", func(c goja.FunctionCall, tx *world.Tx) goja.Value {
		pos, ok := c.Argument(0).Export().(cube.Pos)
		if !ok {
			panic(vm.NewTypeError("argument 0 must be a cube.Pos, got %T", c.Argument(0).Export()))
		}
		return vm.ToValue(tx.Biome(pos))
	})
	setExec("setBiome", func(c goja.FunctionCall, tx *world.Tx) goja.Value {
		pos, ok := c.Argument(0).Export().(cube.Pos)
		if !ok {
			panic(vm.NewTypeError("argument 0 must be a cube.Pos, got %T", c.Argument(0).Export()))
		}
		b, ok := c.Argument(1).Export().(world.Biome)
		if !ok {
			panic(vm.NewTypeError("argument 1 must be a world.Biome, got %T", c.Argument(1).Export()))
		}
		tx.SetBiome(pos, b)
		return nil
	})
	setExec("temperature", func(c goja.FunctionCall, tx *world.Tx) goja.Value {
		pos, ok := c.Argument(0).Export().(cube.Pos)
		if !ok {
			panic(vm.NewTypeError("argument 0 must be a cube.Pos, got %T", c.Argument(0).Export()))
		}
		return vm.ToValue(tx.Temperature(pos))
	})
	setExec("rainingAt", func(c goja.FunctionCall, tx *world.Tx) goja.Value {
		pos, ok := c.Argument(0).Export().(cube.Pos)
		if !ok {
			panic(vm.NewTypeError("argument 0 must be a cube.Pos, got %T", c.Argument(0).Export()))
		}
		return vm.ToValue(tx.RainingAt(pos))
	})
	setExec("snowingAt", func(c goja.FunctionCall, tx *world.Tx) goja.Value {
		pos, ok := c.Argument(0).Export().(cube.Pos)
		if !ok {
			panic(vm.NewTypeError("argument 0 must be a cube.Pos, got %T", c.Argument(0).Export()))
		}
		return vm.ToValue(tx.SnowingAt(pos))
	})
	setExec("thunderingAt", func(c goja.FunctionCall, tx *world.Tx) goja.Value {
		pos, ok := c.Argument(0).Export().(cube.Pos)
		if !ok {
			panic(vm.NewTypeError("argument 0 must be a cube.Pos, got %T", c.Argument(0).Export()))
		}
		return vm.ToValue(tx.ThunderingAt(pos))
	})
	setExec("addParticle", func(c goja.FunctionCall, tx *world.Tx) goja.Value {
		pos, ok := c.Argument(0).Export().(mgl64.Vec3)
		if !ok {
			panic(vm.NewTypeError("argument 0 must be a mgl64.Vec3, got %T", c.Argument(0).Export()))
		}
		p, ok := c.Argument(1).Export().(world.Particle)
		if !ok {
			panic(vm.NewTypeError("argument 1 must be a world.Particle, got %T", c.Argument(1).Export()))
		}
		tx.AddParticle(pos, p)
		return nil
	})
	setExec("playEntityAnimation", func(c goja.FunctionCall, tx *world.Tx) goja.Value {
		e, ok := c.Argument(0).Export().(world.Entity)
		if !ok {
			panic(vm.NewTypeError("argument 0 must be a world.Entity, got %T", c.Argument(0).Export()))
		}
		a, ok := c.Argument(1).Export().(world.EntityAnimation)
		if !ok {
			panic(vm.NewTypeError("argument 1 must be a *world.EntityAnimation, got %T", c.Argument(1).Export()))
		}
		tx.PlayEntityAnimation(e, a)
		return nil
	})
	setExec("playSound", func(c goja.FunctionCall, tx *world.Tx) goja.Value {
		pos, ok := c.Argument(0).Export().(mgl64.Vec3)
		if !ok {
			panic(vm.NewTypeError("argument 0 must be a mgl64.Vec3, got %T", c.Argument(0).Export()))
		}
		s, ok := c.Argument(1).Export().(world.Sound)
		if !ok {
			panic(vm.NewTypeError("argument 1 must be a world.Sound, got %T", c.Argument(1).Export()))
		}
		tx.PlaySound(pos, s)
		return nil
	})
	// TODO: addEntity
	// TODO: removeEntity
	// TODO: entitiesWithin
	setExec("players", func(c goja.FunctionCall, tx *world.Tx) goja.Value {
		var players []*goja.Object
		for p := range tx.Players() {
			players = append(players, w.r.Player(p.H().UUID()).obj)
		}
		return vm.ToValue(players)
	})
	// TODO: viewers

	var err error
	w.obj, err = obj.Build(vm)
	if err != nil {
		panic(err)
	}
}

func (r *Runtime) setupWorld() error {
	return newObject().
		Const("Difficulty", map[string]world.Difficulty{
			"Peaceful": world.DifficultyPeaceful,
			"Easy":     world.DifficultyEasy,
			"Normal":   world.DifficultyNormal,
			"Hard":     world.DifficultyHard,
		}).
		Const("GameMode", map[string]world.GameMode{
			"Adventure": world.GameModeAdventure,
			"Creative":  world.GameModeCreative,
			"Survival":  world.GameModeSurvival,
			"Spectator": world.GameModeSpectator,
		}).
		Const("Dimension", map[string]world.Dimension{
			"Overworld": world.Overworld,
			"Nether":    world.Nether,
			"End":       world.End,
		}).
		Const("nopGenerator", world.NopGenerator{}).
		Const("nopProvider", world.NopProvider{}).
		Const("nopViewer", world.NopViewer{}).
		Method("biomeById", func(c goja.FunctionCall) goja.Value {
			id, ok := c.Argument(0).Export().(int64)
			if !ok {
				panic(r.vm.NewTypeError("argument 0 must be an int64"))
			}
			b, ok := world.BiomeByID(int(id))
			if !ok {
				return nil
			}
			return r.vm.ToValue(b)
		}).
		Method("biomeByName", func(c goja.FunctionCall) goja.Value {
			name, ok := c.Argument(0).Export().(string)
			if !ok {
				panic(r.vm.NewTypeError("argument 0 must be a string"))
			}
			b, ok := world.BiomeByName(name)
			if !ok {
				return nil
			}
			return r.vm.ToValue(b)
		}).
		Method("biomes", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(world.Biomes())
		}).
		Method("blockByName", func(c goja.FunctionCall) goja.Value {
			name, ok := c.Argument(0).Export().(string)
			if !ok {
				panic(r.vm.NewTypeError("argument 0 must be a string"))
			}
			props, ok := c.Argument(1).Export().(map[string]any)
			if !ok {
				panic(r.vm.NewTypeError("argument 1 must be a map[string]any"))
			}
			b, ok := world.BlockByName(name, props)
			if !ok {
				return nil
			}
			return r.vm.ToValue(b)
		}).
		Method("blockByRuntimeId", func(c goja.FunctionCall) goja.Value {
			rid, ok := c.Argument(0).Export().(int64)
			if !ok {
				panic(r.vm.NewTypeError("argument 0 must be an int64"))
			}
			b, ok := world.BlockByRuntimeID(uint32(rid))
			if !ok {
				return nil
			}
			return r.vm.ToValue(b)
		}).
		Method("blockHash", func(c goja.FunctionCall) goja.Value {
			b, ok := c.Argument(0).Export().(world.Block)
			if !ok {
				panic(r.vm.NewTypeError("argument 0 must be a world.Block"))
			}
			return r.vm.ToValue(world.BlockHash(b))
		}).
		Method("blockRuntimeId", func(c goja.FunctionCall) goja.Value {
			b, ok := c.Argument(0).Export().(world.Block)
			if !ok {
				panic(r.vm.NewTypeError("argument 0 must be a world.Block"))
			}
			return r.vm.ToValue(world.BlockRuntimeID(b))
		}).
		Method("chunkPos", func(c goja.FunctionCall) goja.Value {
			x, ok := c.Argument(0).Export().(int64)
			if !ok {
				panic(r.vm.NewTypeError("argument 0 must be an int64"))
			}
			z, ok := c.Argument(1).Export().(int64)
			if !ok {
				panic(r.vm.NewTypeError("argument 1 must be an int64"))
			}
			return r.vm.ToValue(world.ChunkPos{int32(x), int32(z)})
		}).
		Method("customBlocks", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(world.CustomBlocks())
		}).
		Method("customItems", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(world.CustomItems())
		}).
		Method("difficultyId", func(c goja.FunctionCall) goja.Value {
			diff, ok := c.Argument(0).Export().(world.Difficulty)
			if !ok {
				panic(r.vm.NewTypeError("argument 0 must be a world.Difficulty"))
			}
			id, ok := world.DifficultyID(diff)
			if !ok {
				return nil
			}
			return r.vm.ToValue(id)
		}).
		Method("difficultyById", func(c goja.FunctionCall) goja.Value {
			id, ok := c.Argument(0).Export().(int64)
			if !ok {
				panic(r.vm.NewTypeError("argument 0 must be an int64"))
			}
			diff, ok := world.DifficultyByID(int(id))
			if !ok {
				return nil
			}
			return r.vm.ToValue(diff)
		}).
		Method("dimensionId", func(c goja.FunctionCall) goja.Value {
			dim, ok := c.Argument(0).Export().(world.Dimension)
			if !ok {
				panic(r.vm.NewTypeError("argument 0 must be a world.Dimension"))
			}
			id, ok := world.DimensionID(dim)
			if !ok {
				return nil
			}
			return r.vm.ToValue(id)
		}).
		Method("dimensionById", func(c goja.FunctionCall) goja.Value {
			id, ok := c.Argument(0).Export().(int64)
			if !ok {
				panic(r.vm.NewTypeError("argument 0 must be an int64"))
			}
			dim, ok := world.DimensionByID(int(id))
			if !ok {
				return nil
			}
			return r.vm.ToValue(dim)
		}).
		Method("entityAnimation", func(c goja.FunctionCall) goja.Value {
			name, ok := c.Argument(0).Export().(string)
			if !ok {
				panic(r.vm.NewTypeError("argument 0 must be a string"))
			}
			return r.vm.ToValue(world.NewEntityAnimation(name))
		}).
		Method("gameModeId", func(c goja.FunctionCall) goja.Value {
			gm, ok := c.Argument(0).Export().(world.GameMode)
			if !ok {
				panic(r.vm.NewTypeError("argument 0 must be a world.GameMode"))
			}
			id, ok := world.GameModeID(gm)
			if !ok {
				return nil
			}
			return r.vm.ToValue(id)
		}).
		Method("gameModeById", func(c goja.FunctionCall) goja.Value {
			id, ok := c.Argument(0).Export().(int64)
			if !ok {
				panic(r.vm.NewTypeError("argument 0 must be an int64"))
			}
			gm, ok := world.GameModeByID(int(id))
			if !ok {
				return nil
			}
			return r.vm.ToValue(gm)
		}).
		Method("itemRuntimeId", func(c goja.FunctionCall) goja.Value {
			i, ok := c.Argument(0).Export().(world.Item)
			if !ok {
				panic(r.vm.NewTypeError("argument 0 must be a world.Item"))
			}
			rid, meta, ok := world.ItemRuntimeID(i)
			if !ok {
				return nil
			}
			return r.vm.ToValue(map[string]any{
				"runtimeId": rid,
				"meta":      meta,
			})
		}).
		Method("itemByName", func(c goja.FunctionCall) goja.Value {
			name, ok := c.Argument(0).Export().(string)
			if !ok {
				panic(r.vm.NewTypeError("argument 0 must be a string"))
			}
			meta, ok := c.Argument(1).Export().(int64)
			if !ok {
				panic(r.vm.NewTypeError("argument 1 must be an int64"))
			}
			i, ok := world.ItemByName(name, int16(meta))
			if !ok {
				return nil
			}
			return r.vm.ToValue(i)
		}).
		Method("itemByRuntimeId", func(c goja.FunctionCall) goja.Value {
			rid, ok := c.Argument(0).Export().(int64)
			if !ok {
				panic(r.vm.NewTypeError("argument 0 must be an int64"))
			}
			meta, ok := c.Argument(1).Export().(int64)
			if !ok {
				panic(r.vm.NewTypeError("argument 1 must be an int64"))
			}
			i, ok := world.ItemByRuntimeID(int32(rid), int16(meta))
			if !ok {
				return nil
			}
			return r.vm.ToValue(i)
		}).
		Method("items", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(world.Items())
		}).
		Method("loader", func(c goja.FunctionCall) goja.Value {
			radius, ok := c.Argument(0).Export().(int64)
			if !ok {
				panic(r.vm.NewTypeError("argument 0 must be an int64"))
			}
			w, ok := c.Argument(1).Export().(*world.World)
			if !ok {
				panic(r.vm.NewTypeError("argument 1 must be a *world.World"))
			}
			v, ok := c.Argument(2).Export().(world.Viewer)
			if !ok {
				panic(r.vm.NewTypeError("argument 2 must be a world.Viewer"))
			}
			return r.vm.ToValue(world.NewLoader(int(radius), w, v))
		}).
		Method("registerBiome", func(c goja.FunctionCall) goja.Value {
			b, ok := c.Argument(0).Export().(world.Biome)
			if !ok {
				panic(r.vm.NewTypeError("argument 0 must be a world.Biome"))
			}
			world.RegisterBiome(b)
			return nil
		}).
		Method("registerBlock", func(c goja.FunctionCall) goja.Value {
			b, ok := c.Argument(0).Export().(world.Block)
			if !ok {
				panic(r.vm.NewTypeError("argument 0 must be a world.Block"))
			}
			world.RegisterBlock(b)
			return nil
		}).
		Method("registerItem", func(c goja.FunctionCall) goja.Value {
			i, ok := c.Argument(0).Export().(world.Item)
			if !ok {
				panic(r.vm.NewTypeError("argument 0 must be a world.Item"))
			}
			world.RegisterItem(i)
			return nil
		}).
		Method("setOpts", func(c goja.FunctionCall) goja.Value {
			disableUpdates, _ := c.Argument(0).Export().(bool)
			disableDisplacement, _ := c.Argument(1).Export().(bool)
			return r.vm.ToValue(world.SetOpts{DisableBlockUpdates: disableUpdates, DisableLiquidDisplacement: disableDisplacement})
		}).
		Method("settings", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(&world.Settings{})
		}).
		Method("subChunkPos", func(c goja.FunctionCall) goja.Value {
			x, ok := c.Argument(0).Export().(int64)
			if !ok {
				panic(r.vm.NewTypeError("argument 0 must be an int64"))
			}
			y, ok := c.Argument(1).Export().(int64)
			if !ok {
				panic(r.vm.NewTypeError("argument 1 must be an int64"))
			}
			z, ok := c.Argument(2).Export().(int64)
			if !ok {
				panic(r.vm.NewTypeError("argument 2 must be an int64"))
			}
			return r.vm.ToValue(world.SubChunkPos{int32(x), int32(y), int32(z)})
		}).
		Method("config", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(world.Config{})
		}).
		Method("create", func(c goja.FunctionCall) goja.Value {
			conf, _ := c.Argument(0).Export().(world.Config)
			return r.vm.ToValue(conf.New())
		}).
		Apply(r.vm, "world")
}
