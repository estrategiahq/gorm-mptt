package gorm_mptt

import "gorm.io/gorm"

type Tree gorm.DB

func New(db interface{}) Tree {

	return db.(Tree)
}
