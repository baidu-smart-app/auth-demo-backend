package auth

type Auth interface {
	SetOpenID2SessionKey(openID, sessionKey string) error
	GetSessionKeyByOpenID(openID string) (string, error)
	GetAllUser() (map[string]string, error)
}

var DefaultAuth = NewMemoryAuth()
