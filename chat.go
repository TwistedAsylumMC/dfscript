package dfscript

import (
	"github.com/df-mc/dragonfly/server/player/chat"
	"github.com/dop251/goja"
	"golang.org/x/text/language"
)

type translationStr string

// Resolve returns the translation identifier as a string.
func (s translationStr) Resolve(language.Tag) string { return string(s) }

func (r *Runtime) setupChat() error {
	return newObject().
		Const("global", chat.Global).
		Method("translate", func(c goja.FunctionCall) goja.Value {
			if len(c.Arguments) < 3 {
				panic("translate expects at least 3 arguments")
			}
			str := translationStr(c.Argument(0).String())
			params := int(c.Argument(1).ToInteger())
			fallback := c.Argument(2).String()
			return r.vm.ToValue(chat.Translate(str, params, fallback))
		}).
		Apply(r.vm, "chat")
}
