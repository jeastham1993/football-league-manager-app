package infrastructure

import "team-service/domain"

// MockEventBus is a mock implementation of an event bus.
type MockEventBus struct {
}

// Publish sends a new message to the event bus.
func (ev MockEventBus) Publish(evt domain.Event) error {
	println(string(evt.AsEvent()))

	return nil
}
