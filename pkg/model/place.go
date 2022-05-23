package model

import "gorm.io/gorm"

type Place struct {
	gorm.Model
	Name string
}
