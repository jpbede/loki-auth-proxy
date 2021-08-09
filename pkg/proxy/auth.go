package proxy

import (
	"bytes"
	"encoding/base64"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/valyala/fasthttp"
	"github.com/valyala/fasthttp/fasthttpadaptor"
)

// AuthHandler return handler func for fasthttp that performs authentication and then calls a sub handler passing the username
func (p *Proxy) AuthHandler(subHandler func(ctx *fasthttp.RequestCtx, username string)) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		if p.logger != nil {
			p.logger.Debug().
				Str("host", string(ctx.Request.Host())).
				Msg("Got request to proxy, checking auth first")
		}

		if auth := ctx.Request.Header.Peek("Authorization"); auth != nil {
			if bytes.HasPrefix(auth, basicAuthPrefix) {
				payload, err := base64.StdEncoding.DecodeString(string(auth[len(basicAuthPrefix):]))
				if err == nil {
					pair := bytes.SplitN(payload, []byte(":"), 2)
					if len(pair) == 2 && p.Authenticator.Authenticate(string(pair[0]), string(pair[1])) {
						subHandler(ctx, string(pair[0]))
						return
					}
					if p.logger != nil {
						p.logger.Debug().
							Str("host", string(ctx.Request.Host())).
							Str("username", string(pair[0])).
							Msg("Auth invalid, rejecting")
					}
				} else if p.logger != nil {
					p.logger.Error().
						Str("host", string(ctx.Request.Host())).
						Err(err).
						Msg("A error occurred while checking auth")
				}
			}
		}

		hdl := fasthttpadaptor.NewFastHTTPHandler(promhttp.Handler())
		hdl(ctx)

		// Request Basic Authentication otherwise
		ctx.Error(fasthttp.StatusMessage(fasthttp.StatusUnauthorized), fasthttp.StatusUnauthorized)
	}
}
