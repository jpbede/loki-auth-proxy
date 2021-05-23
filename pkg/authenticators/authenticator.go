package authenticators

import (
	"fmt"
	"log"
)

// Authenticator performs authentication
type Authenticator interface {
	// Authenticate checks given credentials
	Authenticate(username, password string) bool
	// GetTenantID returns the ID for the X-Scope-OrgID
	GetTenantID(username string) string
}

type InitializerFunc func(map[string]string) (Authenticator, error)

var authenticators = map[string]InitializerFunc{}

// RegisterAuthenticator add a authenticator to the registry
func RegisterAuthenticator(name string, initializerFunc InitializerFunc) {
	if _, ok := authenticators[name]; ok {
		log.Fatalf("Cannot register authenticator %s multiple times", name)
	}
	authenticators[name] = initializerFunc
}

// GetAuthenticator returns a authenticator by name
func GetAuthenticator(name string, config map[string]string) (Authenticator, error) {
	if initializerFunc, ok := authenticators[name]; ok {
		return initializerFunc(config)
	}
	return nil, fmt.Errorf("Authenticator %s not found", name)
}
