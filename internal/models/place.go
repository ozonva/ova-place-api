package models

import "fmt"

type Place struct {
	ID     uint64 `db:"id"`
	Memo   string `db:"memo"`
	Seat   string `db:"seat"`
	UserID uint64 `db:"user_id"`
}

func (p Place) String() string {
	return fmt.Sprintf("%s (%s)", p.Seat, p.Memo)
}
