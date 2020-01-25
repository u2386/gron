package gron

// bus is a global internal event stream
var bus chan Event

// Subscribe returns the crond event channel
func Subscribe() <-chan Event {
	return bus
}

// EventType indicates the type of event
type EventType int

const (
	// Enabled emitted when a task is enabled
	Enabled EventType = iota

	// Disabled emitted when a task has been disabled
	Disabled

	// Running emitted when a task is starting
	Running

	// Finished emitted when a task has finished successful
	Finished

	// Failed emitted when a task has failed
	Failed
)

// Event presents the task event
type Event struct {
	TaskName
	E   EventType
	Msg string
}

func init() {
	bus = make(chan Event, 1)
}
