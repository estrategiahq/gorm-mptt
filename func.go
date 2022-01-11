package gorm_mptt

import (
	"fmt"
	"reflect"

	"gorm.io/gorm"
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

func (db *Tree) getLftFromParentNode(n interface{}, pos int) int64 {
	rv := reflect.ValueOf(n).Elem()
	parent_id := rv.FieldByName("ParentId").String()
	node_lft := rv.FieldByName("Lft").Int()

	var lft int64

	query := db.Statement.DB.Model(n).Select("lft").
		Where("parent_id IS ?", parent_id).
		Where("rght < ?", node_lft).
		Order("lft asc").
		Limit(1)

	if pos > 0 {
		query = query.Offset(pos - 1)
	}

	query.Scan(lft)

	return lft
}

func (db *Tree) sync(n interface{}, shift int, dir, conditions string) {

	node := reflect.New(reflect.TypeOf(n)).Interface()

	fields := map[int]string{
		0: "lft",
		1: "rght",
	}

	for _, v := range fields {
		exp := fmt.Sprintf("%s %s ?", v, dir)
		where := fmt.Sprintf("%s %s", v, conditions)

		db.Statement.DB.Model(node).Select(v).Where(where).Update(v, gorm.Expr(exp, shift))
	}
}
