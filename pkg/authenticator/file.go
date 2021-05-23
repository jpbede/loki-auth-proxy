package authenticator

type File struct {
}

func (f *File) Name() string {
	return ""
}

func (f *File) Authenticate(username, password string) bool {
	return false
}

func (f *File) GetTenantID(username string) string {
	return username
}
