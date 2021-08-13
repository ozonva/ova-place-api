package models

import "fmt"

type Place struct {
	UserID uint64
	Memo   string
	Seat   string
}

func (p Place) String() string {
	return fmt.Sprintf("%s (%s)", p.Seat, p.Memo)
}
