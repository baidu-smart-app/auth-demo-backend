package auth

type MemoryAuth struct {
	openID2sessionKey map[string]string
}

func NewMemoryAuth() Auth {
	return &MemoryAuth{
		openID2sessionKey: map[string]string{},
	}
}

func (ma *MemoryAuth) GetAllUser() (map[string]string, error) {
	return ma.openID2sessionKey, nil
}

func (ma *MemoryAuth) SetOpenID2SessionKey(openID, sessionKey string) error {
	ma.openID2sessionKey[openID] = sessionKey

	return nil
}
func (ma *MemoryAuth) GetSessionKeyByOpenID(openID string) (string, error) {
	return ma.openID2sessionKey[openID], nil
}
