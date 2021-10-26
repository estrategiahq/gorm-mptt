package gorm_mptt

import "reflect"

func (db *Tree) GetNode(o interface{}) interface{} {
	rv := reflect.ValueOf(o)
	rv_id := rv.FieldByName("ID").String()
	id := string(rv_id)
	db.Statement.First(&o, "id = ?", id)
	return o

}

func (db *Tree) getMax(o interface{}) (interface{}, error) {
	return o, nil
}
