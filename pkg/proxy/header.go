package proxy

import (
	"github.com/valyala/fasthttp"
	"net"
)

// ref to: https://golang.org/src/net/http/httputil/reverseproxy.go#L169
var hopByHopHeaders = []string{
	"Connection",
	"Proxy-Connection",
	"Keep-Alive",
	"Proxy-Authenticate",
	"Proxy-Authorization",
	"Te",
	"Trailer",
	"Transfer-Encoding",
	"Upgrade",
}

func (p *Proxy) removeHopByHopHeaders(req *fasthttp.Request, resp *fasthttp.Response) {
	for _, h := range hopByHopHeaders {
		if req != nil {
			req.Header.Del(h)
		}
		if resp != nil {
			resp.Header.Del(h)
		}
	}
}

func (p *Proxy) addXForwardedHeader(ctx *fasthttp.RequestCtx) {
	if ip, _, err := net.SplitHostPort(ctx.RemoteAddr().String()); err == nil {
		ctx.Request.Header.Add("X-Forwarded-For", ip)
	}

	ctx.Request.Header.Add("X-Forwarded-Host", string(ctx.Request.Host()))
}
