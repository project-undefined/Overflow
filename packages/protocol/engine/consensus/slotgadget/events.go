package slotgadget

import (
	"github.com/iotaledger/hive.go/core/slot"
	"github.com/iotaledger/hive.go/runtime/event"
)

type Events struct {
	SlotConfirmed *event.Event1[slot.Index]

	event.Group[Events, *Events]
}

// NewEvents contains the constructor of the Events object (it is generated by a generic factory).
var NewEvents = event.CreateGroupConstructor(func() (newEvents *Events) {
	return &Events{
		SlotConfirmed: event.New1[slot.Index](),
	}
})
