package internal

import "testing"

func TestGreeting(t *testing.T) {
	if Greeting() == "" {
		t.Fatal("expected non-empty greeting")
	}
}

func TestStoreSetGet(t *testing.T) {
	s := NewStore()
	s.Set("a", "1")
	v, ok := s.Get("a")
	if !ok || v != "1" {
		t.Fatalf("expected a=1, got %q ok=%v", v, ok)
	}
}
