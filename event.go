package gron

import (
	"time"
)

const (
	// Enabled emitted when a task is enabled
	Enabled etype = iota

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

// bus is a global internal event stream
var bus chan Event

type etype int

// Event presents the task event
type Event struct {
	TaskName
	E   etype
	Msg string
	At  time.Time
}

type builder struct {
	ev Event
}

func newBuilder() *builder {
	return &builder{
		ev: Event{},
	}
}

func (b *builder) event(e etype) *builder {
	b.ev.E = e
	return b
}

func (b *builder) taskName(name TaskName) *builder {
	b.ev.TaskName = name
	return b
}

func (b *builder) message(m string) *builder {
	b.ev.Msg = m
	return b
}

func (b *builder) at(t time.Time) *builder {
	b.ev.At = t
	return b
}

func (b *builder) build() Event {
	return b.ev
}

// Subscribe returns the gron event channel
func Subscribe() <-chan Event {
	for name := range c.tasks {
		Enable(name)
	}
	return bus
}

func (et etype) String() string {
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
	publish(newBuilder().taskName(task.Name).event(Enabled).at(time.Now()).build())
}

func publishDisabledEvent(task Task) {
	publish(newBuilder().taskName(task.Name).event(Disabled).at(time.Now()).build())
}

func publishRunningEvent(task Task) {
	publish(newBuilder().taskName(task.Name).event(Running).at(time.Now()).build())
}

func publishFinishedEvent(task Task) {
	publish(newBuilder().taskName(task.Name).event(Finished).at(time.Now()).build())
}

func publishFailedEvent(task Task) {
	publish(newBuilder().taskName(task.Name).event(Failed).message(task.errMsg).at(time.Now()).build())
}

func publishEmptyEvent() {
	publish(newBuilder().event(Empty).at(time.Now()).build())
}

func publish(ev Event) {
	bus <- ev
}

func init() {
	bus = make(chan Event, 1)
}
