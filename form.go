package dfscript

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/df-mc/dragonfly/server/player/form"
	"github.com/df-mc/dragonfly/server/world"
	"github.com/dop251/goja"
	"unicode/utf8"
)

type customForm struct {
	r *Runtime

	title    string
	elements []form.Element
	submit   func(p *goja.Object, values []any, closed bool)
}

func (f *customForm) Label(text string) *customForm {
	f.elements = append(f.elements, form.NewLabel(text))
	return f
}

func (f *customForm) Input(text, defaultValue, placeholder string) *customForm {
	f.elements = append(f.elements, form.NewInput(text, defaultValue, placeholder))
	return f
}

func (f *customForm) Toggle(text string, defaultValue bool) *customForm {
	f.elements = append(f.elements, form.NewToggle(text, defaultValue))
	return f
}

func (f *customForm) Slider(text string, min, max, step float64, defaultValue float64) *customForm {
	f.elements = append(f.elements, form.NewSlider(text, min, max, step, defaultValue))
	return f
}

func (f *customForm) Dropdown(text string, options []string, defaultValue int) *customForm {
	f.elements = append(f.elements, form.NewDropdown(text, options, defaultValue))
	return f
}

func (f *customForm) StepSlider(text string, steps []string, defaultValue int) *customForm {
	f.elements = append(f.elements, form.NewStepSlider(text, steps, defaultValue))
	return f
}

func (f *customForm) OnSubmit(submit func(p *goja.Object, values []any, closed bool)) *customForm {
	f.submit = submit
	return f
}

func (f *customForm) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]any{
		"type":    "custom_form",
		"title":   f.title,
		"content": f.elements,
	})
}

func (f *customForm) SubmitJSON(b []byte, submitter form.Submitter, tx *world.Tx) error {
	pl := submitter.(*player.Player)
	p := f.r.Player(pl.UUID())
	p.p = pl
	defer func() {
		p.p = nil
	}()
	if b == nil {
		if f.submit != nil {
			f.submit(p.obj, nil, true)
		}
		return nil
	}

	dec := json.NewDecoder(bytes.NewBuffer(b))
	dec.UseNumber()

	var data, values []any
	if err := dec.Decode(&data); err != nil {
		return fmt.Errorf("error decoding JSON data to slice: %w", err)
	}
	if len(data) != len(f.elements) {
		return fmt.Errorf("form JSON data array does not have the same length as the form elements: %v != %v", len(data), len(f.elements))
	}
	for i, element := range f.elements {
		v, err := f.parseValue(element, data[i])
		if err != nil {
			return fmt.Errorf("error parsing form response value: %w", err)
		}
		values = append(values, v)
	}
	if f.submit != nil {
		f.submit(p.obj, values, false)
	}
	return nil
}

func (f *customForm) parseValue(elem form.Element, s any) (any, error) {
	var ok bool
	var value any

	switch element := elem.(type) {
	case form.Label:
		value = element.Text
	case form.Input:
		value, ok = s.(string)
		if !ok {
			return value, fmt.Errorf("value %v is not allowed for input element", s)
		}
		if !utf8.ValidString(value.(string)) {
			return value, fmt.Errorf("value %v is not valid UTF8", s)
		}
	case form.Toggle:
		value, ok = s.(bool)
		if !ok {
			return value, fmt.Errorf("value %v is not allowed for toggle element", s)
		}
	case form.Slider:
		v, ok := s.(json.Number)
		f, err := v.Float64()
		if !ok || err != nil {
			return value, fmt.Errorf("value %v is not allowed for slider element", s)
		}
		if f > element.Max || f < element.Min {
			return value, fmt.Errorf("slider value %v is out of range %v-%v", f, element.Min, element.Max)
		}
		value = f
	case form.Dropdown:
		v, ok := s.(json.Number)
		f, err := v.Int64()
		if !ok || err != nil {
			return value, fmt.Errorf("value %v is not allowed for dropdown element", s)
		}
		if f < 0 || int(f) >= len(element.Options) {
			return value, fmt.Errorf("dropdown value %v is out of range %v-%v", f, 0, len(element.Options)-1)
		}
		value = int(f)
	case form.StepSlider:
		v, ok := s.(json.Number)
		f, err := v.Int64()
		if !ok || err != nil {
			return value, fmt.Errorf("value %v is not allowed for dropdown element", s)
		}
		if f < 0 || int(f) >= len(element.Options) {
			return value, fmt.Errorf("dropdown value %v is out of range %v-%v", f, 0, len(element.Options)-1)
		}
		value = int(f)
	}
	return value, nil
}

type menuForm struct {
	r *Runtime

	title       string
	body        string
	buttons     []form.Button
	buttonFuncs []func(p *goja.Object)
	submit      func(p *goja.Object, index int, closed bool)
}

func (f *menuForm) Button(text, image string, submit func(p *goja.Object)) *menuForm {
	f.buttons = append(f.buttons, form.NewButton(text, image))
	f.buttonFuncs = append(f.buttonFuncs, submit)
	return f
}

func (f *menuForm) OnSubmit(submit func(p *goja.Object, index int, closed bool)) *menuForm {
	f.submit = submit
	return f
}

func (f *menuForm) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]any{
		"type":    "form",
		"title":   f.title,
		"content": f.body,
		"buttons": f.buttons,
	})
}

func (f *menuForm) SubmitJSON(b []byte, submitter form.Submitter, tx *world.Tx) error {
	pl := submitter.(*player.Player)
	p := f.r.Player(pl.UUID())
	p.p = pl
	defer func() {
		p.p = nil
	}()
	if b == nil {
		if f.submit != nil {
			f.submit(p.obj, -1, true)
		}
		return nil
	}

	var index uint
	err := json.Unmarshal(b, &index)
	if err != nil {
		return fmt.Errorf("cannot parse button index as int: %w", err)
	}
	if index >= uint(len(f.buttons)) {
		return fmt.Errorf("button index points to inexistent button: %v (only %v buttons present)", index, len(f.buttons))
	}
	if f.buttonFuncs[index] != nil {
		f.buttonFuncs[index](p.obj)
	}
	if f.submit != nil {
		f.submit(p.obj, int(index), false)
	}
	return nil
}

type modalForm struct {
	r *Runtime

	title       string
	body        string
	buttons     [2]string
	buttonFuncs [2]func(p *goja.Object)
	submit      func(p *goja.Object, index int, closed bool)
}

func (f *modalForm) Yes(text string, submit func(p *goja.Object)) *modalForm {
	f.buttons[0] = text
	f.buttonFuncs[0] = submit
	return f
}

func (f *modalForm) No(text string, submit func(p *goja.Object)) *modalForm {
	f.buttons[1] = text
	f.buttonFuncs[1] = submit
	return f
}

func (f *modalForm) OnSubmit(submit func(p *goja.Object, index int, closed bool)) *modalForm {
	f.submit = submit
	return f
}

func (f *modalForm) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]any{
		"type":    "modal",
		"title":   f.title,
		"content": f.body,
		"button1": f.buttons[0],
		"button2": f.buttons[1],
	})
}

func (f *modalForm) SubmitJSON(b []byte, submitter form.Submitter, tx *world.Tx) error {
	pl := submitter.(*player.Player)
	p := f.r.Player(pl.UUID())
	p.p = pl
	defer func() {
		p.p = nil
	}()
	if b == nil {
		if f.submit != nil {
			f.submit(p.obj, -1, true)
		}
		return nil
	}

	var value bool
	if err := json.Unmarshal(b, &value); err != nil {
		return fmt.Errorf("error parsing JSON as bool: %w", err)
	}
	var index int
	if value {
		if f.buttonFuncs[0] != nil {
			f.buttonFuncs[0](p.obj)
		}
	} else {
		if f.buttonFuncs[1] != nil {
			f.buttonFuncs[1](p.obj)
		}
		index = 1
	}
	if f.submit != nil {
		f.submit(p.obj, index, false)
	}
	return nil
}

func (r *Runtime) setupForm() error {
	return newObject().
		Method("custom", func(c goja.FunctionCall) goja.Value {
			title, ok := c.Argument(0).Export().(string)
			if !ok {
				panic(r.vm.NewTypeError("argument 0 is not a string, got %T", c.Argument(0).Export()))
			}
			return r.vm.ToValue(&customForm{r: r, title: title})
		}).
		Method("menu", func(c goja.FunctionCall) goja.Value {
			title, ok := c.Argument(0).Export().(string)
			if !ok {
				panic(r.vm.NewTypeError("argument 0 is not a string, got %T", c.Argument(0).Export()))
			}
			body, _ := c.Argument(1).Export().(string)
			return r.vm.ToValue(&menuForm{r: r, title: title, body: body})
		}).
		Method("modal", func(c goja.FunctionCall) goja.Value {
			title, ok := c.Argument(0).Export().(string)
			if !ok {
				panic(r.vm.NewTypeError("argument 0 is not a string, got %T", c.Argument(0).Export()))
			}
			body, _ := c.Argument(1).Export().(string)
			return r.vm.ToValue(&modalForm{r: r, title: title, body: body, buttons: [2]string{"gui.yes", "gui.no"}})
		}).
		Apply(r.vm, "form")
}
