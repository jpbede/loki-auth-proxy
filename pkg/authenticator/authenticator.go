package authenticator

type IAuthenticator interface {
	Name() string
	Authenticate(username, password string) bool
	GetTenantID(username string) string
}
