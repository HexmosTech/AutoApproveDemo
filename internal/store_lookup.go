package internal

// Exists reports whether key is present in the store, without allocating a
// value copy -- useful for existence checks ahead of a Set/Delete.
func (s *Store) Exists(key string) bool {
	s.mu.RLock()
	defer s.mu.RUnlock()
	_, ok := s.data[key]
	return ok
}
