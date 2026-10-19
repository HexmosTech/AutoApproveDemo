package auth

import "testing"

func TestGenerateTokenLength(t *testing.T) {
	tok, err := GenerateToken()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(tok) != TokenLength {
		t.Fatalf("expected token of length %d, got %d", TokenLength, len(tok))
	}
}

func TestGenerateTokenUnique(t *testing.T) {
	a, _ := GenerateToken()
	b, _ := GenerateToken()
	equal := true
	for i := range a {
		if a[i] != b[i] {
			equal = false
			break
		}
	}
	if equal {
		t.Fatal("expected two generated tokens to differ")
	}
}
