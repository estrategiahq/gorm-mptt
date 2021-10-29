package gorm_mptt

import (
	"fmt"
	"reflect"

	"gorm.io/gorm"
)

func (db *Tree) getNodeById(o interface{}) interface{} {
	rv := reflect.ValueOf(o).Elem()
	id := rv.FieldByName("ID").String()

	db.Statement.First(&o, map[string]interface{}{"id": id})
	return o

}
func (db *Tree) getNodeByParentId(o interface{}) interface{} {
	rv := reflect.ValueOf(o).Elem()
	parent_id := rv.FieldByName("ParentId").String()

	db.Statement.First(&o, map[string]interface{}{"id": parent_id})
	return o

}

func (db *Tree) getMax(o interface{}) int {
	var rght int
	db.Statement.Select("rght").Model(o).Order("rght desc").Scan(&rght)
	return rght
}

func (db *Tree) sync(o interface{}, shift int, dir, conditions string) {
	fields := map[int]string{
		0: "lft",
		1: "rght",
	}

	for _, v := range fields {

		fmt.Println(v, "<<<<<")

		where := fmt.Sprintf("%s %s", v, conditions)

		db.Statement.Update(v, gorm.Expr("? ? ?", v, dir, shift)).Model(o).Where(where)
	}

}
