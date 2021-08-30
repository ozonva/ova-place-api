package repo

type NotFound struct{}

func (m *NotFound) Error() string {
	return "not found"
}
