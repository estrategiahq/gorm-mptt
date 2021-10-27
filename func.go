package gorm_mptt

import (
	"fmt"
	"reflect"
)

func (db *Tree) GetNode(o interface{}) interface{} {
	rv := reflect.ValueOf(&o)
	id := rv.FieldByName("ID")

	// var id string
	// id = rv_id.(string)
	db.Statement.First(&o, "id = ?", fmt.Sprintf("%s", id))
	return o

}

func (db *Tree) getMax(o interface{}) (interface{}, error) {
	return o, nil
}
