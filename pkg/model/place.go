package model

type Place struct {
	ID   int `gorm:"primaryKey"`
	Name string
}
