package gorm_mptt

import (
	"reflect"
)

func (db *Tree) getNodeById(o interface{}) interface{} {
	rv := reflect.ValueOf(o).Elem()
	id := rv.FieldByName("ID").String()

	db.Statement.First(&o, map[string]interface{}{"id": id})
	return o

}
func (db *Tree) getNodeByParentId(o interface{}) interface{} {
	rv := reflect.ValueOf(o).Elem()
	parent_id := rv.FieldByName("Parent_id").String()

	db.Statement.First(&o, map[string]interface{}{"id": parent_id})
	return o

}

func (db *Tree) getMax(o interface{}) int {
	var rght int
	db.Statement.Select("rght").Model(o).Order("rght desc").Scan(&rght)
	return rght
}
