package dfscript

import (
	"github.com/df-mc/dragonfly/server/block/cube"
	"github.com/df-mc/dragonfly/server/world/chunk"
	"github.com/dop251/goja"
)

func (r *Runtime) setupChunk() error {
	return newObject().
		Const("subChunkVersion", chunk.SubChunkVersion).
		Const("currentBlockVersion", chunk.CurrentBlockVersion).
		Const("encoding", map[string]chunk.Encoding{
			"disk":    chunk.DiskEncoding,
			"network": chunk.NetworkEncoding,
		}).
		Const("paletteEncoding", map[string]any{
			"biome": chunk.BiomePaletteEncoding,
			"block": chunk.BlockPaletteEncoding,
		}).
		Const("light", map[string]any{
			"sky":   chunk.SkyLight,
			"block": chunk.BlockLight,
		}).
		Method("encodeBiomes", func(c goja.FunctionCall) goja.Value {
			ch, ok := c.Argument(0).Export().(*chunk.Chunk)
			if !ok {
				panic(r.vm.NewTypeError("argument 0 must be a *chunk.Chunk"))
			}
			e, ok := c.Argument(1).Export().(chunk.Encoding)
			if !ok {
				panic(r.vm.NewTypeError("argument 1 must be a chunk.Encoding"))
			}
			return r.vm.ToValue(chunk.EncodeBiomes(ch, e))
		}).
		Method("encodeSubChunk", func(c goja.FunctionCall) goja.Value {

			ch, ok := c.Argument(0).Export().(*chunk.Chunk)
			if !ok {
				panic(r.vm.NewTypeError("argument 0 must be a *chunk.Chunk"))
			}
			e, ok := c.Argument(1).Export().(chunk.Encoding)
			if !ok {
				panic(r.vm.NewTypeError("argument 1 must be a chunk.Encoding"))
			}
			ind, ok := c.Argument(2).Export().(int64)
			if !ok {
				panic(r.vm.NewTypeError("argument 2 must be an int64"))
			}
			return r.vm.ToValue(chunk.EncodeSubChunk(ch, e, int(ind)))
		}).
		Method("lightArea", func(c goja.FunctionCall) goja.Value {
			chunks, ok := c.Argument(0).Export().([]*chunk.Chunk)
			if !ok {
				panic(r.vm.NewTypeError("argument 0 must be a []*chunk.Chunk"))
			}
			baseX, ok := c.Argument(1).Export().(int64)
			if !ok {
				panic(r.vm.NewTypeError("argument 1 must be an int64"))
			}
			baseY, ok := c.Argument(2).Export().(int64)
			if !ok {
				panic(r.vm.NewTypeError("argument 2 must be an int64"))
			}
			return r.vm.ToValue(chunk.LightArea(chunks, int(baseX), int(baseY)))
		}).
		Method("runtimeIdToState", func(c goja.FunctionCall) goja.Value {
			rid, ok := c.Argument(0).Export().(int64)
			if !ok {
				panic(r.vm.NewTypeError("argument 0 must be an int64"))
			}
			name, props, ok := chunk.RuntimeIDToState(uint32(rid))
			if !ok {
				return nil
			}
			return r.vm.ToValue(map[string]any{
				"name":       name,
				"properties": props,
			})
		}).
		Method("stateToRuntimeId", func(c goja.FunctionCall) goja.Value {
			name, ok := c.Argument(0).Export().(string)
			if !ok {
				panic(r.vm.NewTypeError("argument 0 must be a string"))
			}
			props, ok := c.Argument(1).Export().(map[string]any)
			if !ok {
				panic(r.vm.NewTypeError("argument 1 must be a map[string]interface{}"))
			}
			rid, ok := chunk.StateToRuntimeID(name, props)
			if !ok {
				return nil
			}
			return r.vm.ToValue(rid)
		}).
		Method("diskDecode", func(c goja.FunctionCall) goja.Value {
			data, ok := c.Argument(0).Export().(chunk.SerialisedData)
			if !ok {
				panic(r.vm.NewTypeError("argument 0 must be a chunk.SerialisedData"))
			}
			ra, ok := c.Argument(1).Export().(cube.Range)
			if !ok {
				panic(r.vm.NewTypeError("argument 1 must be a cube.Range"))
			}
			ch, err := chunk.DiskDecode(data, ra)
			if err != nil {
				panic(r.vm.ToValue(err))
			}
			return r.vm.ToValue(ch)
		}).
		Method("networkDecode", func(c goja.FunctionCall) goja.Value {
			air, ok := c.Argument(0).Export().(int64)
			if !ok {
				panic(r.vm.NewTypeError("argument 0 must be an int64"))
			}
			data, ok := c.Argument(1).Export().([]byte)
			if !ok {
				panic(r.vm.NewTypeError("argument 1 must be a []byte"))
			}
			count, ok := c.Argument(2).Export().(int64)
			if !ok {
				panic(r.vm.NewTypeError("argument 2 must be an int64"))
			}
			ra, ok := c.Argument(3).Export().(cube.Range)
			if !ok {
				panic(r.vm.NewTypeError("argument 3 must be a cube.Range"))
			}
			ch, err := chunk.NetworkDecode(uint32(air), data, int(count), ra)
			if err != nil {
				panic(r.vm.ToValue(err))
			}
			return r.vm.ToValue(ch)
		}).
		Method("encode", func(c goja.FunctionCall) goja.Value {
			ch, ok := c.Argument(0).Export().(*chunk.Chunk)
			if !ok {
				panic(r.vm.NewTypeError("argument 0 must be a *chunk.Chunk"))
			}
			e, ok := c.Argument(1).Export().(chunk.Encoding)
			if !ok {
				panic(r.vm.NewTypeError("argument 1 must be a chunk.Encoding"))
			}
			return r.vm.ToValue(chunk.Encode(ch, e))
		}).
		Method("createSubChunk", func(c goja.FunctionCall) goja.Value {
			air, ok := c.Argument(0).Export().(int64)
			if !ok {
				panic(r.vm.NewTypeError("argument 0 must be an int64"))
			}
			return r.vm.ToValue(chunk.NewSubChunk(uint32(air)))
		}).
		Apply(r.vm, "chunk")
}
