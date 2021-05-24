package all

import (
	// import all existing authenticators
	_ "go.bnck.me/loki-auth-proxy/pkg/authenticators/file"
	_ "go.bnck.me/loki-auth-proxy/pkg/authenticators/redis"
)
