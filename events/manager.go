package events

var events = make(map[EventName]map[int]Event)

func RegisterEvent(event Event) {
	if events[event.Name] == nil {
		events[event.Name] = make(map[int]Event)
	}
	events[event.Name][len(events[event.Name])-1] = event
}

func CallAllEvents(eventType EventName, args ...any) {
	mp := events[eventType]
	for _, event := range mp {
		event.Execute(args...)
	}
}
