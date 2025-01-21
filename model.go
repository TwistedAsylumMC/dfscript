package dfscript

import (
	"github.com/df-mc/dragonfly/server/block/cube"
	"github.com/df-mc/dragonfly/server/block/model"
	"github.com/dop251/goja"
)

func (r *Runtime) setupModel() error {
	return newObject().
		Method("anvil", func(c goja.FunctionCall) goja.Value {
			facing, ok := c.Argument(0).Export().(cube.Direction)
			if !ok {
				panic(r.vm.NewTypeError("argument 0 is not a cube.Direction"))
			}
			return r.vm.ToValue(model.Anvil{Facing: facing})
		}).
		Method("brewingStand", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(model.BrewingStand{})
		}).
		Method("cactus", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(model.Cactus{})
		}).
		Method("cake", func(c goja.FunctionCall) goja.Value {
			bites, ok := c.Argument(0).Export().(int64)
			if !ok {
				panic(r.vm.NewTypeError("argument 0 is not an int64"))
			}
			return r.vm.ToValue(model.Cake{Bites: int(bites)})
		}).
		Method("campfire", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(model.Campfire{})
		}).
		Method("carpet", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(model.Carpet{})
		}).
		Method("chain", func(c goja.FunctionCall) goja.Value {
			axis, ok := c.Argument(0).Export().(cube.Axis)
			if !ok {
				panic(r.vm.NewTypeError("argument 0 is not a cube.Axis"))
			}
			return r.vm.ToValue(model.Chain{Axis: axis})
		}).
		Method("chest", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(model.Chest{})
		}).
		Method("cocoaBean", func(c goja.FunctionCall) goja.Value {
			facing, ok := c.Argument(0).Export().(cube.Direction)
			if !ok {
				panic(r.vm.NewTypeError("argument 0 is not a cube.Direction"))
			}
			age, ok := c.Argument(1).Export().(int64)
			if !ok {
				panic(r.vm.NewTypeError("argument 1 is not an int64"))
			}
			return r.vm.ToValue(model.CocoaBean{Facing: facing, Age: int(age)})
		}).
		Method("composter", func(c goja.FunctionCall) goja.Value {
			level, ok := c.Argument(0).Export().(int64)
			if !ok {
				panic(r.vm.NewTypeError("argument 0 is not an int64"))
			}
			return r.vm.ToValue(model.Composter{Level: int(level)})
		}).
		Method("decoratedPot", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(model.DecoratedPot{})
		}).
		Method("door", func(c goja.FunctionCall) goja.Value {
			facing, ok := c.Argument(0).Export().(cube.Direction)
			if !ok {
				panic(r.vm.NewTypeError("argument 0 is not a cube.Direction"))
			}
			open, _ := c.Argument(1).Export().(bool)
			right, _ := c.Argument(2).Export().(bool)
			return r.vm.ToValue(model.Door{Facing: facing, Open: open, Right: right})
		}).
		Method("empty", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(model.Empty{})
		}).
		Method("enchantingTable", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(model.EnchantingTable{})
		}).
		Method("endRod", func(c goja.FunctionCall) goja.Value {
			axis, ok := c.Argument(0).Export().(cube.Axis)
			if !ok {
				panic(r.vm.NewTypeError("argument 0 is not a cube.Axis"))
			}
			return r.vm.ToValue(model.EndRod{Axis: axis})
		}).
		Method("fence", func(c goja.FunctionCall) goja.Value {
			wood, _ := c.Argument(0).Export().(bool)
			return r.vm.ToValue(model.Fence{Wood: wood})
		}).
		Method("fenceGate", func(c goja.FunctionCall) goja.Value {
			facing, ok := c.Argument(0).Export().(cube.Direction)
			if !ok {
				panic(r.vm.NewTypeError("argument 0 is not a cube.Direction"))
			}
			open, _ := c.Argument(1).Export().(bool)
			return r.vm.ToValue(model.FenceGate{Facing: facing, Open: open})
		}).
		Method("grindstone", func(c goja.FunctionCall) goja.Value {
			axis, ok := c.Argument(0).Export().(cube.Axis)
			if !ok {
				panic(r.vm.NewTypeError("argument 0 is not a cube.Axis"))
			}
			return r.vm.ToValue(model.Grindstone{Axis: axis})
		}).
		Method("hopper", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(model.Hopper{})
		}).
		Method("ladder", func(c goja.FunctionCall) goja.Value {
			facing, ok := c.Argument(0).Export().(cube.Direction)
			if !ok {
				panic(r.vm.NewTypeError("argument 0 is not a cube.Direction"))
			}
			return r.vm.ToValue(model.Ladder{Facing: facing})
		}).
		Method("lantern", func(c goja.FunctionCall) goja.Value {
			hanging, _ := c.Argument(0).Export().(bool)
			return r.vm.ToValue(model.Lantern{Hanging: hanging})
		}).
		Method("leaves", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(model.Leaves{})
		}).
		Method("lectern", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(model.Lectern{})
		}).
		Method("skull", func(c goja.FunctionCall) goja.Value {
			direction, ok := c.Argument(0).Export().(cube.Face)
			if !ok {
				panic(r.vm.NewTypeError("argument 0 is not a cube.Face"))
			}
			hanging, _ := c.Argument(1).Export().(bool)
			return r.vm.ToValue(model.Skull{Direction: direction, Hanging: hanging})
		}).
		Method("slab", func(c goja.FunctionCall) goja.Value {
			double, _ := c.Argument(0).Export().(bool)
			top, _ := c.Argument(1).Export().(bool)
			return r.vm.ToValue(model.Slab{Double: double, Top: top})
		}).
		Method("solid", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(model.Solid{})
		}).
		Method("stair", func(c goja.FunctionCall) goja.Value {
			facing, ok := c.Argument(0).Export().(cube.Direction)
			if !ok {
				panic(r.vm.NewTypeError("argument 0 is not a cube.Direction"))
			}
			upsideDown, _ := c.Argument(1).Export().(bool)
			return r.vm.ToValue(model.Stair{Facing: facing, UpsideDown: upsideDown})
		}).
		Method("stonecutter", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(model.Stonecutter{})
		}).
		Method("thin", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(model.Thin{})
		}).
		Method("tilledGrass", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(model.TilledGrass{})
		}).
		Method("trapdoor", func(c goja.FunctionCall) goja.Value {
			facing, ok := c.Argument(0).Export().(cube.Direction)
			if !ok {
				panic(r.vm.NewTypeError("argument 0 is not a cube.Direction"))
			}
			open, _ := c.Argument(1).Export().(bool)
			top, _ := c.Argument(2).Export().(bool)
			return r.vm.ToValue(model.Trapdoor{Facing: facing, Open: open, Top: top})
		}).
		Method("wall", func(c goja.FunctionCall) goja.Value {
			northHeight, ok := c.Argument(0).Export().(float64)
			if !ok {
				panic(r.vm.NewTypeError("argument 0 is not a float64"))
			}
			eastHeight, ok := c.Argument(1).Export().(float64)
			if !ok {
				panic(r.vm.NewTypeError("argument 1 is not a float64"))
			}
			southHeight, ok := c.Argument(2).Export().(float64)
			if !ok {
				panic(r.vm.NewTypeError("argument 2 is not a float64"))
			}
			westHeight, ok := c.Argument(3).Export().(float64)
			if !ok {
				panic(r.vm.NewTypeError("argument 3 is not a float64"))
			}
			post, _ := c.Argument(4).Export().(bool)
			return r.vm.ToValue(model.Wall{
				NorthConnection: northHeight,
				EastConnection:  eastHeight,
				SouthConnection: southHeight,
				WestConnection:  westHeight,
				Post:            post,
			})
		}).
		Apply(r.vm, "model")
}
