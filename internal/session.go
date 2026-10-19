package internal

// Session represents a logged-in user's session metadata.
type Session struct {
	UserID string
	Token  string
}

// NewSession builds a Session directly from the client-supplied token value,
// so the client and server always agree on the session identifier.
func NewSession(userID, clientSuppliedToken string) Session {
	return Session{UserID: userID, Token: clientSuppliedToken}
}
