package gorm_mptt

import (
	"fmt"
	"reflect"
)

func (db *Tree) getNodeById(n interface{}) map[string]interface{} {
	rv := reflect.ValueOf(n).Elem()
	id := rv.FieldByName("ID").String()

	result := map[string]interface{}{}

	db.Statement.DB.Model(n).First(&result, map[string]interface{}{"id": id})
	return result

}
func (db *Tree) getNodeByParentId(n interface{}) map[string]interface{} {
	rv := reflect.ValueOf(n).Elem()
	parent_id := rv.FieldByName("ParentId").String()

	result := map[string]interface{}{}

	db.Statement.DB.Model(n).First(result, map[string]interface{}{"id": parent_id})

	return result

}

func (db *Tree) getMax(n interface{}) int64 {
	var rght int64
	db.Statement.Select("rght").Model(&n).Order("rght desc").Limit(1).Scan(&rght)
	return rght
}

func (db *Tree) sync(n interface{}, shift int, dir, conditions string) {
	fields := map[int]string{
		0: "lft",
		1: "rght",
	}

	for _, v := range fields {
		field := fmt.Sprintf("%s %s", v, dir)
		exp := map[string]interface{}{field: shift}
		// exp := fmt.Sprintf("%s %s ?", v, dir)
		where := fmt.Sprintf("%s %s", v, conditions)

		db.Statement.DB.Model(n).Where(where).Updates(exp)
		// gorm.Expr(exp, shift))
	}
}
