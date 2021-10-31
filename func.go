package gorm_mptt

import (
	"fmt"
	"reflect"

	"gorm.io/gorm"
)

func (db *Tree) getNodeById(o interface{}) interface{} {
	// rv := reflect.ValueOf(o).Elem()
	rv := reflect.ValueOf(o)
	id := rv.FieldByName("ID").String()

	db.Statement.First(&o, map[string]interface{}{"id": id})
	return o

}
func (db *Tree) getNodeByParentId(o interface{}) interface{} {
	// rv := reflect.ValueOf(o).Elem()
	rv := reflect.ValueOf(o).Elem()
	parent_id := rv.FieldByName("ParentId").String()

	db.Statement.First(&o, map[string]interface{}{"id": parent_id})
	return o

}

func (db *Tree) getMax(o interface{}) int {
	var rght int
	db.Statement.Select("rght").Model(&o).Order("rght desc").Scan(&rght)
	return rght
}

func (db *Tree) sync(o interface{}, shift int, dir, conditions string) {
	fmt.Printf("sync: %+v", o)
	fields := map[int]string{
		0: "lft",
		1: "rght",
	}

	newObj := reflect.New(reflect.TypeOf(o)).Interface()

	for _, v := range fields {
		exp := fmt.Sprintf("%s %s ?", v, dir)
		where := fmt.Sprintf("%s %s", v, conditions)

		// gorm.Expr("? ? ?", v, dir, shift)

		db.Statement.DB.Model(&newObj).Where(where).Update(v, gorm.Expr(exp, shift))
		// db.Statement.UpdateColumn(v, 2).Where(where).Model(&o)
	}
	fmt.Printf("sync update: %+v", o)
}
