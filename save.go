package gorm_mptt

import (
	"fmt"
	"reflect"
)

func (db *Tree) SaveNode(o interface{}) (interface{}, error) {
	fmt.Printf("save inicial: %+v", o)

	rv := reflect.ValueOf(o)

	typeof := reflect.New(reflect.TypeOf(o))
	model := typeof.Interface()
	toSave := typeof.Elem()

	toSave.Set(rv)

	// rvp := reflect.ValueOf(&o).Elem()
	// rv := r.Elem()
	// original := reflect.New(reflect.TypeOf(o))
	// original.Elem().Set(rv)

	id := rv.FieldByName("ID")
	parent_id := rv.FieldByName("ParentId")

	if id.IsZero() && parent_id.IsZero() {
		edge := db.getMax(o)

		toSave.FieldByName("Lft").SetInt(int64(edge) + 1)
		toSave.FieldByName("Rght").SetInt(int64(edge) + 2)
	}
	if id.IsZero() && !parent_id.IsZero() {
		parent := db.getNodeByParentId(o)
		parent_rv := reflect.ValueOf(parent).Elem()

		edge := parent_rv.FieldByName("Rght").Int()

		toSave.FieldByName("Lft").SetInt(edge)
		toSave.FieldByName("Rght").SetInt(edge + 1)

		cond := fmt.Sprintf(">= %d", edge)

		db.sync(o, 2, "+", cond)
	}

	fmt.Printf("save antes de salvar: %+v", o)

	err := db.Statement.DB.Model(model).Create(toSave).Error
	return toSave, err
}
