package matrix

import (
	"github.com/mlange-42/arche/ecs"
	"github.com/mlange-42/arche/ecs/event"
	"github.com/mlange-42/arche/generic"
)

type GridManager struct {
	subs event.Subscription
	mask ecs.Mask

	grid   generic.Resource[LetterGrid]
	posMap generic.Map1[Position]
}

func NewGridManager(world *ecs.World) GridManager {
	return GridManager{
		subs:   event.Entities,
		mask:   ecs.All(ecs.ComponentID[Fader](world)),
		grid:   generic.NewResource[LetterGrid](world),
		posMap: generic.NewMap1[Position](world),
	}
}

// Notify the listener about a subscribed event.
func (m *GridManager) Notify(world *ecs.World, evt ecs.EntityEvent) {
	grid := m.grid.Get()
	pos := m.posMap.Get(evt.Entity)
	if evt.EventTypes.Contains(event.EntityCreated) {
		grid.Faders.Set(pos.X, pos.Y, evt.Entity)
	} else {
		grid.Faders.Set(pos.X, pos.Y, ecs.Entity{})
	}
}

// Subscriptions to one or more event types.
func (m *GridManager) Subscriptions() event.Subscription {
	return m.subs
}

// Components the listener subscribes to. Listening to all components indicated by nil.
func (m *GridManager) Components() *ecs.Mask {
	return &m.mask
}
