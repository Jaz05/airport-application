package model

type Passenger struct {
	ID      int `gorm:"primaryKey"`
	Name    string
	SurName string
	Dni     int64
}
