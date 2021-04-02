package authenticators

type Authenticator interface {
	Name() string
	Authenticate(username, password string) bool
	GetTenantID() string
}
