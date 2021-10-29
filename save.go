package gorm_mptt

import "reflect"

func (db *Tree) SaveNode(o interface{}) (interface{}, error) {
	rv := reflect.ValueOf(o).Elem()

	id := rv.FieldByName("ID").String()
	parent_id := rv.FieldByName("ParentId").String()

	if id == "" && parent_id == "" {
		edge := db.getMax(o)

		rv.FieldByName("Lft").SetInt(int64(edge) + 1)
		rv.FieldByName("Rght").SetInt(int64(edge) + 2)
	}
	if id == "" && parent_id != "" {
		parent := db.getNodeByParentId(o)
		parent_rv := reflect.ValueOf(parent).Elem()

		rv.FieldByName("Lft").SetInt(parent_rv.FieldByName("Lft").Int())
		rv.FieldByName("Rght").SetInt(parent_rv.FieldByName("Lft").Int() + 1)
	}

	err := db.Statement.Create(o).Error
	return o, err
}
