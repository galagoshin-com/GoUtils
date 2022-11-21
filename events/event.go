package events

type EventName string

type Event struct {
	Name    EventName
	Execute func(...any)
}
