module go.bnck.me/loki-auth-proxy

go 1.16

require (
	github.com/carousell/fasthttp-prometheus-middleware v1.0.3
	github.com/fasthttp/router v1.3.13
	github.com/jinzhu/configor v1.2.1
	github.com/prometheus/client_golang v1.10.0
	github.com/rs/zerolog v1.22.0
	github.com/stretchr/testify v1.7.0
	github.com/urfave/cli/v2 v2.3.0
	github.com/valyala/fasthttp v1.25.0
	gopkg.in/yaml.v2 v2.3.0
)

replace github.com/carousell/fasthttp-prometheus-middleware v1.0.3 => github.com/jpbede/fasthttp-prometheus-middleware v1.1.0
