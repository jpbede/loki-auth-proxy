package authenticators

import "log"

type Authenticator interface {
	Authenticate(username, password string) bool
	GetTenantID(username string) string
}

type InitializerFunc func(map[string]string) (Authenticator, error)

var authenticators map[string]InitializerFunc

func RegisterAuthenticator(name string, initializerFunc InitializerFunc) {
	if _, ok := authenticators[name]; ok {
		log.Fatalf("Cannot register authenticator %s multiple times", name)
	}
	authenticators[name] = initializerFunc
}

func GetAuthenticator(name string, config map[string]string) (Authenticator, error) {
	return authenticators[name](config)
}
