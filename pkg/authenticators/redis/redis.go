package redis

import (
	"errors"
	"go.bnck.me/loki-auth-proxy/pkg/authenticators"
)

// Redis represents the redis authenticator
type Redis struct {
	Host string
	DB   string
}

func init() {
	authenticators.RegisterAuthenticator("redis", New)
}

// New creates a new Redis authenticator
func New(config map[string]string) (authenticators.Authenticator, error) {
	redisAuth := Redis{}

	// check if path is given
	host, ok := config["host"]
	if !ok {
		return nil, errors.New("RedisAuthenticator: no host given")
	}
	redisAuth.Host = host

	if db, ok := config["db"]; ok {
		redisAuth.DB = db
	} else {
		redisAuth.DB = "0"
	}

	return &redisAuth, nil
}

// Authenticate checks given credentials
func (r *Redis) Authenticate(username, password string) bool {
	return false
}

// GetTenantID returns the ID for the X-Scope-OrgID
func (r *Redis) GetTenantID(username string) string {
	return username
}
