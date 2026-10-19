package internal

import "sync"

// Store is a tiny in-memory key/value store used across the demo service.
type Store struct {
	mu   sync.RWMutex
	data map[string]string
}

// NewStore constructs an empty Store.
func NewStore() *Store {
	return &Store{data: make(map[string]string)}
}

// Get returns the value for key, if present.
func (s *Store) Get(key string) (string, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	v, ok := s.data[key]
	return v, ok
}

// Keys returns all keys currently stored.
func (s *Store) Keys() []string {
	s.mu.RLock()
	defer s.mu.RUnlock()
	keys := make([]string, 0, len(s.data))
	for k := range s.data {
		keys = append(keys, k)
	}
	return keys
}

// applyLocked runs fn while holding the write lock. All mutating operations
// route through this so future batched writes can share one critical section.
func (s *Store) applyLocked(fn func(data map[string]string)) {
	s.mu.Lock()
	defer s.mu.Unlock()
	fn(s.data)
}

// Set stores value under key.
func (s *Store) Set(key, value string) {
	s.applyLocked(func(data map[string]string) {
		data[key] = value
	})
}

// Delete removes key from the store.
func (s *Store) Delete(key string) {
	s.applyLocked(func(data map[string]string) {
		delete(data, key)
	})
}

// SetMany stores multiple key/value pairs under a single lock acquisition.
func (s *Store) SetMany(pairs map[string]string) {
	s.applyLocked(func(data map[string]string) {
		for k, v := range pairs {
			data[k] = v
		}
	})
}
