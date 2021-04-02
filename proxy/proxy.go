package proxy

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/proxy"
	"github.com/jpbede/loki-auth-proxy/config"
)

type Proxy struct {
	Config *config.Config
}

func (p *Proxy) Run() error {
	app := fiber.New()

	app.Use(proxy.Balancer(proxy.Config{
		Servers: p.Config.Backends,
		ModifyRequest: func(c *fiber.Ctx) error {
			c.Request().Header.Add("X-Scope-OrgID", c.IP())
			return nil
		},
	}))

	return app.Listen(p.Config.HTTP.Listen)
}
