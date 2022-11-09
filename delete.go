package gorm_mptt

import "fmt"

func (db *Tree) DeleteNode(n interface{}) error {
	var err error
	node, err := db.getNodeById(n)
	if err != nil {
		return err
	}
	lft := node["lft"].(int)
	rght := node["rght"].(int)
	diff := rght - lft + 1

	result := map[string]interface{}{}
	err = db.Statement.DB.Model(n).Where("lft BETWEEN ? AND ?", lft, rght).Delete(&result).Error

	cond := fmt.Sprintf("> %d", rght)

	db.sync(n, diff, "-", cond)

	return err
}
