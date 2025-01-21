package dfscript

import (
	"fmt"
	"github.com/dop251/goja"
	"go.uber.org/multierr"
)

type objectBuilder struct {
	fields  []objectBuilderField
	methods []objectBuilderMethod
}

type objectBuilderField struct {
	name  string
	value any

	writable, configurable, enumerable bool
}

type objectBuilderMethod struct {
	name string
	fn   func(c goja.FunctionCall) goja.Value
}

func newObject() *objectBuilder {
	return &objectBuilder{}
}

func (b *objectBuilder) Const(name string, value any) *objectBuilder {
	return b.Field(name, value, false, false, true)
}

func (b *objectBuilder) Field(name string, value any, writable, configurable, enumerable bool) *objectBuilder {
	b.fields = append(b.fields, objectBuilderField{
		name:         name,
		value:        value,
		writable:     writable,
		configurable: configurable,
		enumerable:   enumerable,
	})
	return b
}

func (b *objectBuilder) Method(name string, fn func(c goja.FunctionCall) goja.Value) *objectBuilder {
	b.methods = append(b.methods, objectBuilderMethod{
		name: name,
		fn:   fn,
	})
	return b
}

func (b *objectBuilder) Apply(vm *goja.Runtime, name string) error {
	obj, err := b.Build(vm)
	if err != nil {
		return err
	} else if err = vm.Set(name, obj); err != nil {
		return err
	}
	return nil
}

func (b *objectBuilder) Build(vm *goja.Runtime) (*goja.Object, error) {
	var errs error
	obj := vm.NewObject()
	for _, field := range b.fields {
		w, c, e := goja.FLAG_FALSE, goja.FLAG_FALSE, goja.FLAG_FALSE
		if field.writable {
			w = goja.FLAG_TRUE
		}
		if field.configurable {
			c = goja.FLAG_TRUE
		}
		if field.enumerable {
			e = goja.FLAG_TRUE
		}
		if err := obj.DefineDataProperty(field.name, vm.ToValue(field.value), w, c, e); err != nil {
			multierr.AppendInto(&errs, fmt.Errorf("define %s: %w", field.name, err))
		}
	}
	for _, method := range b.methods {
		if err := obj.Set(method.name, method.fn); err != nil {
			multierr.AppendInto(&errs, fmt.Errorf("set %s: %w", method.name, err))
		}
	}
	return obj, errs
}
