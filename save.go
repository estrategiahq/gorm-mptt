package gorm_mptt

import (
	"fmt"
	"reflect"
)

func (db *Tree) SaveNode(o interface{}) (interface{}, error) {
	fmt.Printf("save inicial: %+v", o)

	r := reflect.ValueOf(o)
	rv := r.Elem()
	original := reflect.New(reflect.TypeOf(o))

	reflect.Copy(original, r)

	id := rv.FieldByName("ID")
	parent_id := rv.FieldByName("ParentId")

	if id.IsZero() && parent_id.IsZero() {
		edge := db.getMax(o)

		rv.FieldByName("Lft").SetInt(int64(edge) + 1)
		rv.FieldByName("Rght").SetInt(int64(edge) + 2)
	}
	if id.IsZero() && !parent_id.IsZero() {
		parent := db.getNodeByParentId(o)
		parent_rv := reflect.ValueOf(parent).Elem()

		edge := parent_rv.FieldByName("Rght").Int()

		rv.FieldByName("Lft").SetInt(edge)
		rv.FieldByName("Rght").SetInt(edge + 1)

		cond := fmt.Sprintf(">= %d", edge)

		db.sync(o, 2, "+", cond)
	}

	fmt.Printf("save antes de salvar: %+v", original)

	err := db.Statement.Create(original.Interface()).Error
	return o, err
}
