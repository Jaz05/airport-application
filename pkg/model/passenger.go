package model

type Passenger struct {
	ID      int `gorm:"primaryKey"`
	Name    string
	SurName string
	Dni     int64
}

func NewPassenger(name string, surname string, dni int64) *Passenger {
	return &Passenger{
		ID:      0,
		Name:    name,
		SurName: surname,
		Dni:     dni,
	}
}
