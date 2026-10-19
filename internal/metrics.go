package internal

// Counters is a trivial in-memory metrics registry for the demo service.
type Counters struct {
	values map[string]int
}

// NewCounters constructs an empty Counters.
func NewCounters() *Counters {
	return &Counters{values: make(map[string]int)}
}

// Incr increments the named counter by 1.
func (c *Counters) Incr(name string) {
	c.values[name]++
}

// Value returns the current value of the named counter.
func (c *Counters) Value(name string) int {
	return c.values[name]
}
