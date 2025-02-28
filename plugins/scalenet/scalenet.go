package scalenet

import (
	"context"
	"github.com/robinbraemer/event"
	"go.minekube.com/common/minecraft/component"
	"go.minekube.com/gate/pkg/edition/java/proxy"
)

var Plugin = proxy.Plugin{
	Name: "scale-net",
	Init: func(ctx context.Context, p *proxy.Proxy) error {
		var redis = ConnectRedis()
		event.Subscribe(p.Event(), 0, func(e *proxy.PlayerChooseInitialServerEvent) {
			server, err := redis.GetValues("scale-net-default-server")
			if err != nil {
				e.Player().Disconnect(&component.Text{Content: "scaleNet no server found!"})
				return
			}

			e.SetInitialServer(p.Server(server[0]))
		})

		return nil
	},
}
