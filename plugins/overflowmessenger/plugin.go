package overflow

import (
	"sync"

	"github.com/iotaledger/goshimmer/packages/app/blockissuer"
	"github.com/iotaledger/goshimmer/packages/node"
	"github.com/iotaledger/goshimmer/packages/protocol"
	"go.uber.org/dig"
)

var (
	plugin     *node.Plugin
	pluginOnce sync.Once
)

type dependencies struct {
	dig.In
	BlockIssuer *blockissuer.BlockIssuer
	Protocol    *protocol.Protocol
}

func Plugin() *node.Plugin {
	pluginOnce.Do(func() {
		plugin = node.NewPlugin("OverflowMessenger", new(dependencies), node.Enabled, configure, run)
	})
	return plugin
}

func configure(plugin *node.Plugin) {

}

func run(plugin *node.Plugin) {

}
