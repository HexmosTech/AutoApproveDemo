package internal

import "sync"

// Store is a tiny in-memory key/value store used across the demo service.
// Methods now live in internal/store_*.go -- this file only holds the type.
type Store struct {
	mu   sync.RWMutex
	data map[string]string
}

// NewStore constructs an empty Store.
func NewStore() *Store {
	return &Store{data: make(map[string]string)}
}
