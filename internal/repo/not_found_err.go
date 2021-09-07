package repo

// NotFound is a structure for Repo methods.
type NotFound struct{}

// Error returns error message.
func (m *NotFound) Error() string {
	return "not found"
}
