package internal

import (
	"sort"
	"testing"
)

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

func TestStoreGetMissing(t *testing.T) {
	s := NewStore()
	if _, ok := s.Get("missing"); ok {
		t.Fatal("expected ok=false for missing key")
	}
}

func TestStoreDelete(t *testing.T) {
	s := NewStore()
	s.Set("a", "1")
	s.Delete("a")
	if _, ok := s.Get("a"); ok {
		t.Fatal("expected key to be gone after delete")
	}
}

func TestStoreDeleteMissingIsNoop(t *testing.T) {
	s := NewStore()
	s.Delete("does-not-exist")
}

func TestStoreKeys(t *testing.T) {
	s := NewStore()
	s.Set("a", "1")
	s.Set("b", "2")
	s.Set("c", "3")

	keys := s.Keys()
	sort.Strings(keys)

	want := []string{"a", "b", "c"}
	if len(keys) != len(want) {
		t.Fatalf("expected %d keys, got %d (%v)", len(want), len(keys), keys)
	}
	for i := range want {
		if keys[i] != want[i] {
			t.Fatalf("expected keys %v, got %v", want, keys)
		}
	}
}

func TestStoreKeysEmpty(t *testing.T) {
	s := NewStore()
	if keys := s.Keys(); len(keys) != 0 {
		t.Fatalf("expected no keys, got %v", keys)
	}
}

func TestStoreOverwrite(t *testing.T) {
	s := NewStore()
	s.Set("a", "1")
	s.Set("a", "2")
	v, _ := s.Get("a")
	if v != "2" {
		t.Fatalf("expected overwritten value 2, got %q", v)
	}
}
