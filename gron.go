package gron

import (
	"fmt"
)

// c is a global internal gron instance
var c gron

type gron struct {
	tasks map[TaskName]*Task
}

// newGron is an internal call for testing purposes
func newGron() gron {
	return gron{
		tasks: make(map[TaskName]*Task),
	}
}

// Gron is an user interface for declaring periodic task
func Gron(ops ...Option) error {
	task := newTask()
	var err error
	for _, op := range ops {
		if task, err = op(task); err != nil {
			return err
		}
	}

	name := task.Name
	if _, ok := c.tasks[name]; ok {
		return fmt.Errorf("Duplicated task: %s", name)
	}
	c.tasks[name] = &task
	return nil
}

// Remove removes an inactive task from the scheduler
func Remove(name TaskName) {
	if t, ok := c.tasks[name]; ok {
		go disableTask(t)
		// Always keep tasks modification in main routine
		delete(c.tasks, name)
	}
}

// Disable puts the task inactive
func Disable(name TaskName) {
	if t, ok := c.tasks[name]; ok {
		go disableTask(t)
	}
}

// Enable puts the task active
func Enable(name TaskName) error {
	if t, ok := c.tasks[name]; ok {
		if t.c != nil {
			close(t.c)
			t.c = nil
		}
		t.c = make(chan struct{})
		go enableTask(t)
	}
	return fmt.Errorf("Task<%s> not exists", name)
}

// Notice: Block-call
func enableTask(task *Task) {
	publishEnabledEvent(*task)
	task.run()
}

// Notice: Block-call
func disableTask(task *Task) {
	publishDisabledEvent(*task)
	task.c <- struct{}{}
	<-task.c
}

func init() {
	c = newGron()
}
