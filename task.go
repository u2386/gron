package gron

import (
	"fmt"
	"reflect"
	"time"
)

// TaskName is the name of the task
type TaskName string

type timeUnit int

const (
	seconds timeUnit = iota
	minutes
	hours
	days
	weeks
)

type status int

const (
	active status = iota
	inactive
)

// Task is a periodic task
type Task struct {
	Name     TaskName
	interval uint64
	unit     timeUnit
	fn       interface{}
	c        chan struct{}
	status   status
	errMsg   string
}

func newTask() Task {
	return Task{
	}
}

func (t *Task) run() {
	t.status = active
	go func() {
		defer func() {
			t.status = inactive
		}()

		// TODO
		// Timer?
		ticker := time.NewTicker(time.Duration(t.interval) * time.Second)

		for {
			select {
			case <-t.c:
				t.c <- struct{}{}
				return
			case <-ticker.C:
				stepRun(t)
			default:
			}
		}
	}()
}

func stepRun(t *Task) {
	defer func() {
		if r := recover(); r != nil {
			t.errMsg = fmt.Sprintf("%v", r)
			publishFailedEvent(*t)
			return
		}
		publishFinishedEvent(*t)
	}()

	publishRunningEvent(*t)
	v := reflect.ValueOf(t.fn)
	v.Call([]reflect.Value{})
}
