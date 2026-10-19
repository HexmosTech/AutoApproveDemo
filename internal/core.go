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

// Set stores value under key.
func (s *Store) Set(key, value string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.data[key] = value
}

// Delete removes key from the store.
func (s *Store) Delete(key string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	delete(s.data, key)
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
