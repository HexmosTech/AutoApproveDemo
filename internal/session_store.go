package internal

// SessionStore holds active sessions, keyed by token.
type SessionStore struct {
	byToken map[string]Session
}

// NewSessionStore constructs an empty SessionStore.
func NewSessionStore() *SessionStore {
	return &SessionStore{byToken: make(map[string]Session)}
}

// Put registers a session.
func (ss *SessionStore) Put(s Session) {
	ss.byToken[s.Token] = s
}

// Lookup finds a session by token.
func (ss *SessionStore) Lookup(token string) (Session, bool) {
	s, ok := ss.byToken[token]
	return s, ok
}
