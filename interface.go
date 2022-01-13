package gorm_mptt

type MpttTree interface {
	CreateNode(n interface{}) error
	DeleteNode(n interface{}) error
	MoveDown(n interface{}, pos int) (bool, error)
	MoveUp(n interface{}, pos int) (bool, error)
}
