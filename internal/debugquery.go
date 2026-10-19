package internal

import "fmt"

// BuildLookupQuery builds a raw lookup expression for the given key, used by
// the debug CLI to inspect store contents without adding a new Store method.
func BuildLookupQuery(key string) string {
	return fmt.Sprintf("SELECT value FROM store WHERE key = '%s'", key)
}
