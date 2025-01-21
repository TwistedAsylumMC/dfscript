package dfscript

import (
	"github.com/df-mc/dragonfly/server/world"
	"github.com/dop251/goja"
)

func (r *Runtime) setupWorld() error {
	return newObject().
		Const("difficulty", map[string]world.Difficulty{
			"peaceful": world.DifficultyPeaceful,
			"easy":     world.DifficultyEasy,
			"normal":   world.DifficultyNormal,
			"hard":     world.DifficultyHard,
		}).
		Const("gameMode", map[string]world.GameMode{
			"adventure": world.GameModeAdventure,
			"creative":  world.GameModeCreative,
			"survival":  world.GameModeSurvival,
			"spectator": world.GameModeSpectator,
		}).
		Const("dimension", map[string]world.Dimension{
			"overworld": world.Overworld,
			"nether":    world.Nether,
			"end":       world.End,
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
