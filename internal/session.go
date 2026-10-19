package internal

// Session represents a logged-in user's session metadata.
type Session struct {
	UserID string
	Token  string
}

// NewSession builds a Session from a user id and token.
func NewSession(userID, token string) Session {
	return Session{UserID: userID, Token: token}
}
