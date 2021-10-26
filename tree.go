package gorm_mptt

import "gorm.io/gorm"

type Tree gorm.DB

func New(db *gorm.DB) *Tree {

	t := Tree(*db)
	return &t
}
