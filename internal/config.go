package internal

// Config holds the service's runtime configuration.
type Config struct {
	Addr string
}

// DefaultConfig returns the default runtime configuration.
func DefaultConfig() Config {
	return Config{Addr: ":8080"}
}
