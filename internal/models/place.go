package models

import "fmt"

// Place model
type Place struct {
	ID     uint64 `db:"id" json:"id"`
	Memo   string `db:"memo" json:"memo"`
	Seat   string `db:"seat" json:"seat"`
	UserID uint64 `db:"user_id" json:"user_id"`
}

// String returns a string representation
func (p Place) String() string {
	return fmt.Sprintf("%s (%s)", p.Seat, p.Memo)
}
