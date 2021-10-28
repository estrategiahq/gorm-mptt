package gorm_mptt

func (db *Tree) GetNode(o interface{}) interface{} {
	db.Statement.First(&o)
	return o
}

func (db *Tree) getMax(o interface{}) (interface{}, error) {
	return o, nil
}
