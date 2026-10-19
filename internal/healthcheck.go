package internal

// Healthy reports whether the service considers itself ready to serve
// traffic. This demo service has no external dependencies, so it is always
// healthy once the process is up.
func Healthy() bool {
	return true
}
