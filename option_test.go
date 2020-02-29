package gron

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEvery(t *testing.T) {
	task := newTask()
	task, _ = Every(4)(task)
	assert.Equal(t, uint(4), task.interval)
}

func TestSecond(t *testing.T) {
	task := newTask()
	task, _ = Second()(task)
	assert.Equal(t, seconds, task.unit)
}

func TestMinute(t *testing.T) {
	task := newTask()
	task, _ = Minute()(task)
	assert.Equal(t, minutes, task.unit)
}

func TestHour(t *testing.T) {
	task := newTask()
	task, _ = Hour()(task)
	assert.Equal(t, hours, task.unit)
}

func TestName(t *testing.T) {
	task := newTask()
	task, _ = Name("name")(task)
	assert.Equal(t, TaskName("name"), task.Name)
}

func TestDo(t *testing.T) {
	task := newTask()
	task, _ = Do(func() {
		_ = "Gron"
	})(task)

	v := reflect.ValueOf(task.fn)
	assert.Equal(t, reflect.Func, v.Kind())
}
