package tangle

import (
	"github.com/iotaledger/goshimmer/packages/protocol/engine/tangle/blockdag"
	"github.com/iotaledger/goshimmer/packages/protocol/engine/tangle/booker"
	"github.com/iotaledger/hive.go/runtime/event"
)

type Events struct {
	BlockDAG *blockdag.Events
	Booker   *booker.Events

	event.Group[Events, *Events]
}

// NewEvents contains the constructor of the Events object (it is generated by a generic factory).
var NewEvents = event.CreateGroupConstructor(func() (newEvents *Events) {
	return &Events{
		BlockDAG: blockdag.NewEvents(),
		Booker:   booker.NewEvents(),
	}
})