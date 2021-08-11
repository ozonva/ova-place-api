package models

import "testing"

func TestString_Place(t *testing.T) {
	place := Place{
		UserId: 1,
		Seat:   "34G",
		Memo:   "aeroflot 12.04.2022 09:00",
	}

	if "34G (aeroflot 12.04.2022 09:00)" != place.String() {
		t.Fatal("The strings not equals")
	}
}
