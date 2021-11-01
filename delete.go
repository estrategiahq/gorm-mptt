package gorm_mptt

import "fmt"

func (db *Tree) DeleteNode(n interface{}) error {
	var err error
	node := db.getNodeById(n)
	lft := node["lft"].(int)
	rght := node["rght"].(int)
	diff := rght - lft + 1

	if diff > 2 {
		err = db.Statement.DB.Where("lft BETWEEN ? AND ?", (lft + 1), (rght - 1)).Delete(n).Error
	}

	cond := fmt.Sprintf("> %d", rght)

	db.sync(n, diff, "-", cond)

	return err
}
