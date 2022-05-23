package model

import "gorm.io/gorm"

type Passenger struct {
	gorm.Model
	Name, SurName string
}
