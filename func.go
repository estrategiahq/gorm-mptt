package gorm_mptt

func (db *Tree) GetNode(o interface{}, id string) interface{} {
	db.Statement.First(o, "id = ?", id)
	return o
}

func (db *Tree) getMax(o interface{}) (interface{}, error) {
	return o, nil
}
