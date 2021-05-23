package file

import (
	"go.bnck.me/loki-auth-proxy/pkg/authenticators"
)

type File struct {
}

func init() {
	authenticators.RegisterAuthenticator("file", New)
}

func New(config map[string]string) (authenticators.Authenticator, error) {
	return &File{}, nil
}

func (f *File) Authenticate(username, password string) bool {
	return false
}

func (f *File) GetTenantID(username string) string {
	return username
}
