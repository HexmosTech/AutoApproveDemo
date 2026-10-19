package internal

// Healthy always reports true for this demo service; a real service would
// check its dependencies here.
func Healthy() bool {
	return true
}
