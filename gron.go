package gron

import (
	"fmt"
)

// c is a global internal Gron instance
var c Gron

// Gron is an implementation of the task scheduler
type Gron struct {
	tasks map[TaskName]*Task
}

// newGron is an internal call for testing purposes
func newGron() Gron {
	return Gron{
		tasks: make(map[TaskName]*Task),
	}
}

// Crontab is an user interface for declaring periodic task
func Crontab(ops ...Option) error {
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

	enableTask(&task)
	return nil
}

// Remove removes an inactive task from the scheduler
func Remove(name TaskName) error {
	if t, ok := c.tasks[name]; ok {
		if t.status == active {
			return fmt.Errorf("Active task %s can not be removed", name)
		}
		delete(c.tasks, name)
	}
	return nil
}

// Disable puts the task inactive
func Disable(name TaskName) {
	if t, ok := c.tasks[name]; ok {
		go disableTask(t)
	}
}

// Enable puts the task active
func Enable(name TaskName) {
	if t, ok := c.tasks[name]; ok {
		go enableTask(t)
	}
}

func enableTask(task *Task) {
	bus <- Event{task.Name, Enabled, ""}
	task.run()
}

func disableTask(task *Task) {
	task.c <- struct{}{}
	<-task.c
	bus <- Event{task.Name, Disabled, ""}
}

func init() {
	c = newGron()
}
