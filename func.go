package gorm_mptt

import "reflect"

func (db *Tree) GetNode(o interface{}) interface{} {
	rv := reflect.ValueOf(&o).Elem()
	id := rv.FieldByName("ID")

	// var id string
	// id = rv_id.(string)
	db.Statement.Where("id = ?", id).First(&o)
	return o

}

func (db *Tree) getMax(o interface{}) (interface{}, error) {
	return o, nil
}
