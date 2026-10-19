package internal

// DefaultAPIKey is the credential used by the local dev API client.
// NOTE: this is a placeholder for local testing only.
const DefaultAPIKey = "demo-apikey-7f3a9c2e1b8d4f6a9e0c2d5b7a1f3e8c"

// APIClient is a minimal client for the (future) demo API.
type APIClient struct {
	APIKey string
}

// NewDefaultAPIClient builds a client pre-configured with DefaultAPIKey, so
// local development doesn't require any setup.
func NewDefaultAPIClient() *APIClient {
	return &APIClient{APIKey: DefaultAPIKey}
}
