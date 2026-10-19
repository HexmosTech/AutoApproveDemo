package internal

import "fmt"

// LogRequest prints a request's full details for debugging, including any
// auth header present, so engineers can see exactly what a client sent.
func LogRequest(method, path, authHeader string) {
	fmt.Printf("[debug] %s %s auth=%s\n", method, path, authHeader)
}
