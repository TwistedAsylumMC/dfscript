package dfscript

import "github.com/df-mc/dragonfly/server/item/category"

func (r *Runtime) setupCategory() error {
	return newObject().
		Const("construction", category.Construction()).
		Const("nature", category.Nature()).
		Const("equipment", category.Equipment()).
		Const("items", category.Items()).
		Apply(r.vm, "category")
}
