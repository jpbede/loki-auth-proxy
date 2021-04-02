package authenticators

type FileAuthenticator struct {
}

func (fa *FileAuthenticator) Name() string {
	return ""
}

func (fa *FileAuthenticator) Authenticate(username, password string) bool {
	return false
}

func (fa *FileAuthenticator) GetTenantID() string {
	return ""
}
