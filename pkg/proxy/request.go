package proxy

import "github.com/valyala/fasthttp"

func (p *Proxy) prepareContext(ctx *fasthttp.RequestCtx, proxyClient *fasthttp.HostClient) {
	// add X-Forwarded header to the request
	p.addXForwardedHeader(ctx)

	// Remove hop-by-hop headers from request
	p.removeHopByHopHeaders(&ctx.Request, nil)

	// set Host header to configured one so we can support virtual hosts
	// and not only IP addresses
	ctx.Request.SetHost(proxyClient.Addr)
}

func (p *Proxy) postProcessContext(ctx *fasthttp.RequestCtx) {
	// Remove hop-by-hop headers from response
	p.removeHopByHopHeaders(nil, &ctx.Response)
}
