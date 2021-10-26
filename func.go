package gorm_mptt

import "reflect"

func (db *Tree) GetNode(o interface{}) interface{} {
	rv := reflect.ValueOf(o)
	id := rv.FieldByName("ID").String()
	result := reflect.New(reflect.TypeOf(o))
	db.Statement.First(result, id)
	return result

}

func (db *Tree) getMax(o interface{}) (interface{}, error) {
	return o, nil
}
