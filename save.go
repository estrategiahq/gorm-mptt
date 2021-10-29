package gorm_mptt

import "reflect"

func (db *Tree) SaveNode(o interface{}) (interface{}, error) {
	rv := reflect.ValueOf(o).Elem()

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

		rv.FieldByName("Lft").SetInt(parent_rv.FieldByName("Lft").Int())
		rv.FieldByName("Rght").SetInt(parent_rv.FieldByName("Lft").Int() + 1)
	}

	err := db.Statement.Create(o).Error
	return o, err
}
