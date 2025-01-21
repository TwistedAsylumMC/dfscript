package dfscript

import (
	"github.com/dop251/goja"
	"github.com/sandertv/gophertunnel/minecraft/text"
)

func (r *Runtime) setupText() error {
	return newObject().
		Const("black", text.Black).
		Const("darkBlue", text.DarkBlue).
		Const("darkGreen", text.DarkGreen).
		Const("darkAqua", text.DarkAqua).
		Const("darkRed", text.DarkRed).
		Const("darkPurple", text.DarkPurple).
		Const("orange", text.Orange).
		Const("grey", text.Grey).
		Const("darkGrey", text.DarkGrey).
		Const("blue", text.Blue).
		Const("green", text.Green).
		Const("aqua", text.Aqua).
		Const("red", text.Red).
		Const("purple", text.Purple).
		Const("yellow", text.Yellow).
		Const("white", text.White).
		Const("darkYellow", text.DarkYellow).
		Const("quartz", text.Quartz).
		Const("iron", text.Iron).
		Const("netherite", text.Netherite).
		Const("obfuscated", text.Obfuscated).
		Const("bold", text.Bold).
		Const("redstone", text.Redstone).
		Const("copper", text.Copper).
		Const("italic", text.Italic).
		Const("gold", text.Gold).
		Const("emerald", text.Emerald).
		Const("reset", text.Reset).
		Const("diamond", text.Diamond).
		Const("lapis", text.Lapis).
		Const("amethyst", text.Amethyst).
		Const("resin", text.Resin).
		Method("ansi", func(c goja.FunctionCall) goja.Value {
			args := make([]any, len(c.Arguments))
			for i, arg := range c.Arguments {
				args[i] = arg.Export()
			}
			return r.vm.ToValue(text.ANSI(args...))
		}).
		Method("clean", func(c goja.FunctionCall) goja.Value {
			s := c.Argument(0).String()
			return r.vm.ToValue(text.Clean(s))
		}).
		Method("colourf", func(c goja.FunctionCall) goja.Value {
			format := c.Argument(0).String()
			var args []any
			if len(c.Arguments) > 1 {
				for _, argument := range c.Arguments[1:] {
					args = append(args, argument.Export())
				}
			}
			return r.vm.ToValue(text.Colourf(format, args...))
		}).
		Apply(r.vm, "text")
}
