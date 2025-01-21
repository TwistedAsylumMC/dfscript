package dfscript

import (
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/recipe"
	"github.com/df-mc/dragonfly/server/world"
	"github.com/dop251/goja"
)

func (r *Runtime) setupRecipe() error {
	return newObject().
		Method("all", func(c goja.FunctionCall) goja.Value {
			return r.vm.ToValue(recipe.Recipes())
		}).
		Method("furnace", func(c goja.FunctionCall) goja.Value {
			input, ok := c.Argument(0).Export().(recipe.Item)
			if !ok {
				panic(r.vm.NewTypeError("argument 0 must be a recipe.Item"))
			}
			output, ok := c.Argument(1).Export().(item.Stack)
			if !ok {
				panic(r.vm.NewTypeError("argument 1 must be an item.Stack"))
			}
			block, ok := c.Argument(2).Export().(string)
			if !ok {
				panic(r.vm.NewTypeError("argument 2 must be a string"))
			}
			return r.vm.ToValue(recipe.NewFurnace(input, output, block))
		}).
		Method("itemTag", func(c goja.FunctionCall) goja.Value {
			tag, ok := c.Argument(0).Export().(string)
			if !ok {
				panic(r.vm.NewTypeError("argument 0 must be a string"))
			}
			count, ok := c.Argument(1).Export().(int64)
			if !ok {
				panic(r.vm.NewTypeError("argument 1 must be an int64"))
			}
			return r.vm.ToValue(recipe.NewItemTag(tag, int(count)))
		}).
		Method("perform", func(c goja.FunctionCall) goja.Value {
			block, ok := c.Argument(0).Export().(string)
			if !ok {
				panic(r.vm.NewTypeError("argument 0 must be a string"))
			}
			input, ok := c.Argument(1).Export().([]world.Item)
			if !ok {
				panic(r.vm.NewTypeError("argument 1 must be a []world.Item"))
			}
			output, ok := recipe.Perform(block, input...)
			if !ok {
				return nil
			}
			return r.vm.ToValue(output)
		}).
		Method("potion", func(c goja.FunctionCall) goja.Value {
			input, ok := c.Argument(0).Export().(recipe.Item)
			if !ok {
				panic(r.vm.NewTypeError("argument 0 must be a recipe.Item"))
			}
			reagent, ok := c.Argument(1).Export().(recipe.Item)
			if !ok {
				panic(r.vm.NewTypeError("argument 1 must be a recipe.Item"))
			}
			output, ok := c.Argument(2).Export().(item.Stack)
			if !ok {
				panic(r.vm.NewTypeError("argument 2 must be an item.Stack"))
			}
			return r.vm.ToValue(recipe.NewPotion(input, reagent, output))
		}).
		Method("potionContainerChange", func(c goja.FunctionCall) goja.Value {
			input, ok := c.Argument(0).Export().(world.Item)
			if !ok {
				panic(r.vm.NewTypeError("argument 0 must be a world.Item"))
			}
			reagent, ok := c.Argument(1).Export().(item.Stack)
			if !ok {
				panic(r.vm.NewTypeError("argument 1 must be an item.Stack"))
			}
			output, ok := c.Argument(2).Export().(world.Item)
			if !ok {
				panic(r.vm.NewTypeError("argument 2 must be a world.Item"))
			}
			return r.vm.ToValue(recipe.NewPotionContainerChange(input, output, reagent))
		}).
		Method("register", func(c goja.FunctionCall) goja.Value {
			rec, ok := c.Argument(0).Export().(recipe.Recipe)
			if !ok {
				panic(r.vm.NewTypeError("argument 0 must be a recipe.Recipe"))
			}
			recipe.Register(rec)
			return nil
		}).
		Method("shaped", func(c goja.FunctionCall) goja.Value {
			input, ok := c.Argument(0).Export().([]recipe.Item)
			if !ok {
				panic(r.vm.NewTypeError("argument 0 must be a []recipe.Item"))
			}
			output, ok := c.Argument(1).Export().(item.Stack)
			if !ok {
				panic(r.vm.NewTypeError("argument 1 must be an item.Stack"))
			}
			width, ok := c.Argument(2).Export().(int64)
			if !ok {
				panic(r.vm.NewTypeError("argument 2 must be an int64"))
			}
			height, ok := c.Argument(3).Export().(int64)
			if !ok {
				panic(r.vm.NewTypeError("argument 3 must be an int64"))
			}
			block, ok := c.Argument(4).Export().(string)
			if !ok {
				panic(r.vm.NewTypeError("argument 4 must be a string"))
			}
			return r.vm.ToValue(recipe.NewShaped(input, output, recipe.NewShape(int(width), int(height)), block))
		}).
		Method("shapeless", func(c goja.FunctionCall) goja.Value {
			input, ok := c.Argument(0).Export().([]recipe.Item)
			if !ok {
				panic(r.vm.NewTypeError("argument 0 must be a []recipe.Item"))
			}
			output, ok := c.Argument(1).Export().(item.Stack)
			if !ok {
				panic(r.vm.NewTypeError("argument 1 must be an item.Stack"))
			}
			block, ok := c.Argument(2).Export().(string)
			if !ok {
				panic(r.vm.NewTypeError("argument 2 must be a string"))
			}
			return r.vm.ToValue(recipe.NewShapeless(input, output, block))
		}).
		Method("smithingTransform", func(c goja.FunctionCall) goja.Value {
			base, ok := c.Argument(0).Export().(recipe.Item)
			if !ok {
				panic(r.vm.NewTypeError("argument 0 must be a recipe.Item"))
			}
			addition, ok := c.Argument(1).Export().(recipe.Item)
			if !ok {
				panic(r.vm.NewTypeError("argument 1 must be a recipe.Item"))
			}
			template, ok := c.Argument(2).Export().(recipe.Item)
			if !ok {
				panic(r.vm.NewTypeError("argument 2 must be a recipe.Item"))
			}
			output, ok := c.Argument(3).Export().(item.Stack)
			if !ok {
				panic(r.vm.NewTypeError("argument 3 must be an item.Stack"))
			}
			block, ok := c.Argument(4).Export().(string)
			if !ok {
				panic(r.vm.NewTypeError("argument 4 must be a string"))
			}
			return r.vm.ToValue(recipe.NewSmithingTransform(base, addition, template, output, block))
		}).
		Method("smithingTrim", func(c goja.FunctionCall) goja.Value {
			base, ok := c.Argument(0).Export().(recipe.Item)
			if !ok {
				panic(r.vm.NewTypeError("argument 0 must be a recipe.Item"))
			}
			addition, ok := c.Argument(1).Export().(recipe.Item)
			if !ok {
				panic(r.vm.NewTypeError("argument 1 must be a recipe.Item"))
			}
			template, ok := c.Argument(2).Export().(recipe.Item)
			if !ok {
				panic(r.vm.NewTypeError("argument 2 must be a recipe.Item"))
			}
			block, ok := c.Argument(3).Export().(string)
			if !ok {
				panic(r.vm.NewTypeError("argument 3 must be a string"))
			}
			return r.vm.ToValue(recipe.NewSmithingTrim(base, addition, template, block))
		}).
		Method("validBrewingReagent", func(c goja.FunctionCall) goja.Value {
			i, ok := c.Argument(0).Export().(world.Item)
			if !ok {
				panic(r.vm.NewTypeError("argument 0 must be a world.Item"))
			}
			return r.vm.ToValue(recipe.ValidBrewingReagent(i))
		}).
		Apply(r.vm, "recipe")
}
