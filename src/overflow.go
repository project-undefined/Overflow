package overflow

import (
	"sync"

	"github.com/iotaledger/hive.go/core/node"
	"github.com/iotaledger/hive.go/node"
)

var (
	plugin     *node.Plugin
	pluginOnce sync.Once
)

func Plugin() *node.Plugin {
	pluginOnce.Do(func() {
		plugin = node.NewPlugin("projectundefined.org/overflow", node.Enabled, configure, run)
	})
	return plugin
}
func configure(plugin *node.Plugin) {

}

func run(plugin *node.Plugin) {

}
