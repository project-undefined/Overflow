package plugins

import (
	"github.com/iotaledger/goshimmer/packages/node"
	"github.com/iotaledger/goshimmer/plugins/dagsvisualizer"
	"github.com/iotaledger/goshimmer/plugins/dashboard"
)

// UI contains the user interface plugins of a GoShimmer node.
var UI = node.Plugins(
	dagsvisualizer.Plugin,
	dashboard.Plugin,
)
