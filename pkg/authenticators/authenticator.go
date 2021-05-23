package authenticators

import (
	"errors"
	"fmt"
	"log"
)

type Authenticator interface {
	Authenticate(username, password string) bool
	GetTenantID(username string) string
}

type InitializerFunc func(map[string]string) (Authenticator, error)

var authenticators = map[string]InitializerFunc{}

func RegisterAuthenticator(name string, initializerFunc InitializerFunc) {
	if _, ok := authenticators[name]; ok {
		log.Fatalf("Cannot register authenticator %s multiple times", name)
	}
	authenticators[name] = initializerFunc
}

func GetAuthenticator(name string, config map[string]string) (Authenticator, error) {
	if initializerFunc, ok := authenticators[name]; ok {
		return initializerFunc(config)
	}
	return nil, errors.New(fmt.Sprintf("Authenticator %s not found", name))
}
