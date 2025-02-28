package main

import (
	"github.com/minekube/gate-plugin-template/plugins/bossbar"
	"github.com/minekube/gate-plugin-template/plugins/globalchat"
	"github.com/minekube/gate-plugin-template/plugins/ping"
	"github.com/minekube/gate-plugin-template/plugins/scalenet"
	"github.com/minekube/gate-plugin-template/plugins/tablist"
	"github.com/minekube/gate-plugin-template/plugins/titlecmd"
	"go.minekube.com/gate/cmd/gate"
	"go.minekube.com/gate/pkg/edition/java/proxy"
)

func main() {
	proxy.Plugins = append(proxy.Plugins,
		scalenet.Plugin,
		tablist.Plugin,
		globalchat.Plugin,
		bossbar.Plugin,
		ping.Plugin,
		titlecmd.Plugin,
	)
	
	gate.Execute()
}
