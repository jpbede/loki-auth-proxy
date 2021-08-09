package proxy

import (
	"github.com/fasthttp/websocket"
	"github.com/valyala/fasthttp"
)

// lot of this code based on https://github.com/yeqown/fasthttp-reverse-proxy

// WebsocketProxy return a fasthttp handler for proxying websockets
func (p *Proxy) WebsocketProxy() func(ctx *fasthttp.RequestCtx, username string) {
	return func(ctx *fasthttp.RequestCtx, username string) {
		if success := websocket.FastHTTPIsWebSocketUpgrade(ctx); success {
			p.logger.Debug().Msg("Clients requests a upgrade to WebSocket")
		}

		dialer := websocket.DefaultDialer
		upgrader := &websocket.FastHTTPUpgrader{
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
		}

		connBackend, _, err := dialer.Dial(p.Querier, nil)
		if err != nil {
			ctx.Error(err.Error(), fasthttp.StatusServiceUnavailable)
		}

		err = upgrader.Upgrade(ctx, func(connPub *websocket.Conn) {
			defer connPub.Close()

			var (
				errClient  = make(chan error, 1)
				errBackend = make(chan error, 1)
				message    string
			)

			go replicateWebsocketMessage(connPub, connBackend, errClient)  // response
			go replicateWebsocketMessage(connBackend, connPub, errBackend) // request

			for {
				select {
				case err = <-errClient:
					message = "websocketproxy: Error when copying response: %v"
				case err = <-errBackend:
					message = "websocketproxy: Error when copying request: %v"
				}

				// log error except '*websocket.CloseError'
				if _, ok := err.(*websocket.CloseError); !ok {
					p.logger.Error().Msgf(message, err)
				}
			}
		})

	}
}

func replicateWebsocketMessage(src, dst *websocket.Conn, errChan chan error) {
	for {
		msgType, msg, err := src.ReadMessage()
		if err != nil {
			if ce, ok := err.(*websocket.CloseError); ok {
				msg = websocket.FormatCloseMessage(ce.Code, ce.Text)
			} else {
				msg = websocket.FormatCloseMessage(websocket.CloseAbnormalClosure, err.Error())
			}

			errChan <- err
			if err = dst.WriteMessage(websocket.CloseMessage, msg); err != nil {
			}
			break
		}

		err = dst.WriteMessage(msgType, msg)
		if err != nil {
			errChan <- err
			break
		}
	}
}
