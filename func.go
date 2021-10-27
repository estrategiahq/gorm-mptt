package gorm_mptt

import "reflect"

func (db *Tree) GetNode(o interface{}) interface{} {
	rv := reflect.ValueOf(o)
	rv_id := rv.FieldByName("ID").String()
	var id string
	id = string(rv_id)
	db.Statement.Where("id = ?", id).First(&o)
	return o

}

func (db *Tree) getMax(o interface{}) (interface{}, error) {
	return o, nil
}
