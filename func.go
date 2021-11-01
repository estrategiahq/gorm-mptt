package gorm_mptt

import (
	"fmt"
	"reflect"

	"gorm.io/gorm"
)

func (db *Tree) getNodeById(o interface{}) map[string]interface{} {
	rv := reflect.ValueOf(o).Elem()
	id := rv.FieldByName("ID").String()

	result := map[string]interface{}{}

	db.Statement.DB.Model(o).First(&result, map[string]interface{}{"id": id})
	return result

}
func (db *Tree) getNodeByParentId(o interface{}) map[string]interface{} {
	rv := reflect.ValueOf(o).Elem()
	parent_id := rv.FieldByName("ParentId").String()

	result := map[string]interface{}{}

	db.Statement.DB.Model(o).First(result, map[string]interface{}{"id": parent_id})

	return result

}

func (db *Tree) getMax(o interface{}) int64 {
	var rght int64
	db.Statement.Select("rght").Model(&o).Order("rght desc").Limit(1).Scan(&rght)
	return rght
}

func (db *Tree) sync(o interface{}, shift int, dir, conditions string) {
	fmt.Printf("sync: %+v", o)
	fields := map[int]string{
		0: "lft",
		1: "rght",
	}

	// newObj := reflect.New(reflect.TypeOf(o)).Interface()

	for _, v := range fields {
		exp := fmt.Sprintf("%s %s ?", v, dir)
		where := fmt.Sprintf("%s %s", v, conditions)

		// gorm.Expr("? ? ?", v, dir, shift)

		db.Statement.DB.Model(o).Where(where).Update(v, gorm.Expr(exp, shift))
		// db.Statement.UpdateColumn(v, 2).Where(where).Model(&o)
	}
	fmt.Printf("sync update: %+v", o)
}
