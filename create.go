package gorm_mptt

import (
	"fmt"
	"reflect"
)

func (db *Tree) CreateNode(n interface{}) error {

	sql := db.Statement
	rv := reflect.ValueOf(n).Elem()

	id := rv.FieldByName("ID")
	parent_id := rv.FieldByName("ParentID")

	if id.IsZero() {
		sql.Omit("id")
	}

	if parent_id.IsZero() {
		edge := db.getMax(n)

		rv.FieldByName("Lft").SetInt(edge + 1)
		rv.FieldByName("Rght").SetInt(edge + 2)
	}
	if !parent_id.IsZero() {
		parent := db.getNodeByParentID(n)

		edge := int64(parent["rght"].(int))

		rv.FieldByName("Lft").SetInt(edge)
		rv.FieldByName("Rght").SetInt(edge + 1)

		cond := fmt.Sprintf(">= %d", edge)

		db.sync(n, 2, "+", cond)
	}

	err := sql.Create(n).Error
	return err
}
