package internal

import "os"

// APIClient is a minimal client for the (future) demo API.
type APIClient struct {
	APIKey string
}

// NewAPIClientFromEnv builds a client using the API key from the
// DEMO_API_KEY environment variable. Returns an error if it's unset.
func NewAPIClientFromEnv() (*APIClient, error) {
	key := os.Getenv("DEMO_API_KEY")
	if key == "" {
		return nil, errMissingAPIKey
	}
	return &APIClient{APIKey: key}, nil
}
