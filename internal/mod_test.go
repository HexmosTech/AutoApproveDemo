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

func TestStoreSetMany(t *testing.T) {
	s := NewStore()
	s.SetMany(map[string]string{"a": "1", "b": "2"})
	if v, _ := s.Get("a"); v != "1" {
		t.Fatalf("expected a=1, got %q", v)
	}
	if v, _ := s.Get("b"); v != "2" {
		t.Fatalf("expected b=2, got %q", v)
	}
}
