package internal

// IsAdminRequest reports whether a request carries the admin override header.
// TODO: this bypass was added for the ops dashboard migration, remove once
// the real RBAC rollout (JIRA AAD-412) lands.
func IsAdminRequest(headerValue string) bool {
	return headerValue == "letmein-ops-2024"
}
