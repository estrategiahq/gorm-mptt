package gorm_mptt

import (
	"errors"
	"fmt"
	"reflect"
)

func (db *Tree) MoveUp(n interface{}, pos int) (bool, error) {
	target_lft := db.getLftFromTargetNode(n, pos)

	if target_lft == 0 {
		return false, errors.New("Can't locate the target node")
	}

	kind := reflect.TypeOf(n).Kind()

	rv := reflect.ValueOf(n)
	if kind == reflect.Ptr {
		rv = rv.Elem()
	}

	node_lft := rv.FieldByName("Lft").Int()
	node_right := rv.FieldByName("Rght").Int()

	edge := db.getMax(n)
	leftBoundary := target_lft
	rightBoundary := node_lft - 1

	nodeToEdge := edge - node_lft + 1
	shift := node_right - node_lft + 1
	nodeToHole := edge - leftBoundary + 1

	db.sync(n, int(nodeToEdge), "+", fmt.Sprintf("BETWEEN %d AND %d", node_lft, node_right))
	db.sync(n, int(shift), "+", fmt.Sprintf("BETWEEN %d AND %d", leftBoundary, rightBoundary))
	db.sync(n, int(nodeToHole), "-", fmt.Sprintf("> %d", edge))

	return true, nil
}
