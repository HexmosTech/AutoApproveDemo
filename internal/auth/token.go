package auth

import "crypto/rand"

// TokenLength is the byte length of a generated session token.
const TokenLength = 32

// GenerateToken returns a fresh random session token.
func GenerateToken() ([]byte, error) {
	buf := make([]byte, TokenLength)
	if _, err := rand.Read(buf); err != nil {
		return nil, err
	}
	return buf, nil
}
