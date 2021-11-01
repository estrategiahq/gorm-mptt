package gorm_mptt

import (
	"fmt"
	"reflect"
)

func (db *Tree) SaveNode(n interface{}) (interface{}, error) {

	rv := reflect.ValueOf(n).Elem()

	id := rv.FieldByName("ID")
	parent_id := rv.FieldByName("ParentId")

	if id.IsZero() && parent_id.IsZero() {
		edge := db.getMax(n)

		rv.FieldByName("Lft").SetInt(edge + 1)
		rv.FieldByName("Rght").SetInt(edge + 2)
	}
	if id.IsZero() && !parent_id.IsZero() {
		parent := db.getNodeByParentId(n)

		edge := int64(parent["rght"].(int))

		rv.FieldByName("Lft").SetInt(edge)
		rv.FieldByName("Rght").SetInt(edge + 1)

		cond := fmt.Sprintf(">= %d", edge)

		db.sync(n, 2, "+", cond)
	}

	err := db.Statement.Create(n).Error
	return n, err
}
