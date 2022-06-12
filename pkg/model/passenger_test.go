package model

import (
	"testing"
)

func TestPassengerConstructorShouldReturnNewPassenger(t *testing.T) {
	expectedPassenger := Passenger{
		ID:      0,
		Name:    "agus",
		SurName: "legui",
		Dni:     38998262,
	}

	passenger := NewPassenger("agus", "legui", 38998262)

	if expectedPassenger != *passenger {
		t.Fatalf("Expected: %v, Got: %v", expectedPassenger, passenger)
	}

}
