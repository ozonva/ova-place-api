package metrics

// Counter is an interface for count different events.
type Counter interface {
	Inc()
	Add(value float64)
}
