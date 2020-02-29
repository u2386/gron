package gron

import (
	"fmt"
	"sync"
)

// c is a global internal Gron instance
var (
	c        gron
	gronOnce sync.Once
)

type gron struct {
	tasks map[TaskName]*Task
}

// newGron is an internal call for creating a singleton gron
func newGron() gron {
	gronOnce.Do(func() {
		c = gron{
			tasks: make(map[TaskName]*Task),
		}
	})
	return c
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

	if err := verifyTask(task); err != nil {
		return err
	}
	c.tasks[task.Name] = &task
	return nil
}

func verifyTask(t Task) error {
	name := t.Name
	if _, ok := c.tasks[name]; ok {
		return fmt.Errorf("Duplicated task: %s", name)
	}
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
	// TODO: sync.Once?
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

func enableTask(task *Task) {
	task.run()
}

func disableTask(task *Task) {
	task.c <- struct{}{}
}

func init() {
	c = newGron()
}
