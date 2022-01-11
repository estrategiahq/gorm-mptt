package gorm_mptt

import (
	"fmt"
	"reflect"
)

func (db *Tree) MoveUp(n interface{}, pos int) (bool, error) {

	target_lft := db.getLftFromParentNode(n, pos)

	rv := reflect.ValueOf(n).Elem()
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
