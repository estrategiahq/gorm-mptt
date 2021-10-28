package gorm_mptt

import (
	"fmt"
	"reflect"
)

func (db *Tree) GetNode(o interface{}) interface{} {
	rv := reflect.ValueOf(o)
	rv_id := rv.FieldByName("ID").String()

	var id string
	id = fmt.Sprintf("%s", rv_id)
	fmt.Println(id)
	db.Statement.First(&o, "id = ?", &id)
	return o

}

func (db *Tree) getMax(o interface{}) (interface{}, error) {
	return o, nil
}
