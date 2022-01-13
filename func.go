package gorm_mptt

import (
	"fmt"
	"reflect"

	"gorm.io/gorm"
)

func (db *Tree) getNodeById(n interface{}) map[string]interface{} {
	kind := reflect.TypeOf(n).Kind()

	rv := reflect.ValueOf(n)

	if kind == reflect.Ptr {
		rv = rv.Elem()
	}

	id := rv.FieldByName("ID").String()

	result := map[string]interface{}{}

	db.Statement.DB.Model(n).First(&result, map[string]interface{}{"id": id})
	return result

}
func (db *Tree) getNodeByParentId(n interface{}) map[string]interface{} {
	kind := reflect.TypeOf(n).Kind()

	rv := reflect.ValueOf(n)

	if kind == reflect.Ptr {
		rv = rv.Elem()
	}

	parent_id := rv.FieldByName("ParentId").Addr().Interface()

	result := map[string]interface{}{}

	db.Statement.DB.Model(n).First(result, map[string]interface{}{"id": parent_id})

	return result

}

func (db *Tree) getMax(n interface{}) int64 {
	var rght int64
	db.Statement.Select("rght").Model(&n).Order("rght desc").Limit(1).Scan(&rght)
	return rght
}

func (db *Tree) getLftFromTargetNode(n interface{}, pos int) int64 {
	node := reflect.New(reflect.TypeOf(n)).Interface()

	rv := reflect.ValueOf(n)

	kind := reflect.TypeOf(n).Kind()

	if kind == reflect.Ptr {
		rv = rv.Elem()
	}

	node_lft := rv.FieldByName("Lft").Int()

	var lft int64

	query := db.Statement.DB.Model(node).Select("lft").
		// Where("parent_id IS NULL").
		Where("rght < ?", node_lft).
		Order("lft desc").
		Limit(1).
		Offset(pos - 1)

	if rv.FieldByName("ParentId").IsNil() {
		query = query.Where("parent_id IS NULL")
	} else {
		parent_id := rv.FieldByName("ParentId").Interface()
		query = query.Where("parent_id = ?", parent_id)
	}

	query.Scan(&lft)

	return lft
}

func (db *Tree) getRghtFromTargetNode(n interface{}, pos int) int64 {
	node := reflect.New(reflect.TypeOf(n)).Interface()

	rv := reflect.ValueOf(n)

	kind := reflect.TypeOf(n).Kind()

	if kind == reflect.Ptr {
		rv = rv.Elem()
	}

	node_rght := rv.FieldByName("Rght").Int()

	var rght int64

	query := db.Statement.DB.Model(node).Select("rght").
		// Where("parent_id IS NULL").
		Where("lft > ?", node_rght).
		Order("lft asc").
		Limit(1).
		Offset(pos - 1)

	if rv.FieldByName("ParentId").IsNil() {
		query = query.Where("parent_id IS NULL")
	} else {
		parent_id := rv.FieldByName("ParentId").Interface()
		query = query.Where("parent_id = ?", parent_id)
	}

	query.Scan(&rght)

	return rght
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
