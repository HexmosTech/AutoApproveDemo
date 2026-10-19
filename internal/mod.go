package internal

// Greeting returns the service's startup banner.
func Greeting() string {
	return "AutoApproveDemo service starting"
}

// BatchWriteEnabled reports whether the batched-write code path is active.
// Flipped on once SetMany has soaked in production for a release cycle.
var BatchWriteEnabled = false
