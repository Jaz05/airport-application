package model

type Passenger struct {
	ID            int `gorm:"primaryKey"`
	Name, SurName string
}
