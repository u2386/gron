package gron

// bus is a global internal event stream
var bus chan Event

// EventType indicates the type of event
type eventType int

// Event presents the task event
type Event struct {
	TaskName
	E   eventType
	Msg string
}

const (
	// Enabled emitted when a task is enabled
	Enabled eventType = iota

	// Disabled emitted when a task has been disabled
	Disabled

	// Running emitted when a task is starting
	Running

	// Finished emitted when a task has finished successful
	Finished

	// Failed emitted when a task has failed
	Failed

	// Empty emitted when no task in crond
	Empty
)

// Subscribe returns the gron event channel
func Subscribe() <-chan Event {
	for name := range c.tasks {
		Enable(name)
	}
	return bus
}

func (et eventType) String() string {
	switch et {
	case Enabled:
		return "Enabled"
	case Disabled:
		return "Disabled"
	case Running:
		return "Running"
	case Finished:
		return "Finished"
	case Failed:
		return "Failed"
	case Empty:
		return "Empty"
	default:
		return "Undefined"
	}
}

func publishEnabledEvent(task Task) {
	bus <- Event{task.Name, Enabled, ""}
}

func publishDisabledEvent(task Task) {
	bus <- Event{task.Name, Disabled, ""}
}

func publishRunningEvent(task Task) {
	bus <- Event{task.Name, Running, ""}
}

func publishFinishedEvent(task Task) {
	bus <- Event{task.Name, Finished, ""}
}

func publishFailedEvent(task Task) {
	bus <- Event{task.Name, Failed, task.errMsg}
}

func publishEmptyEvent() {
	bus <- Event{"", Empty, ""}
}

func init() {
	bus = make(chan Event, 1)
}
