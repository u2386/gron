package gron

import (
	"fmt"
	"math/rand"
	"reflect"
	"strconv"
	"strings"
	"time"
)

// TaskName is the name of the task
type TaskName string

type timeUnit int

const shifting = 50

const (
	seconds timeUnit = iota
	minutes
	hours
	days
	weekday
)

type status int

const (
	active status = iota
	inactive
)

// Task is a periodic task
type Task struct {
	Name      TaskName
	interval  uint
	unit      timeUnit
	at        string
	weekday   time.Weekday
	fn        interface{} // TODO: support args channel
	c         chan struct{}
	status    status
	errMsg    string
	latestRun time.Time
}

func newTask() Task {
	return Task{}
}

func (t *Task) run() {
	defer func() {
		t.status = inactive
		publishDisabledEvent(*t)
	}()

	t.status = active
	publishEnabledEvent(*t)

LOOP:
	for {
		r := rand.Intn(shifting)
		timer := time.NewTimer(willRunAfter(*t) + +time.Duration(r)*time.Millisecond)

		select {
		case <-t.c:
			timer.Stop()
			break LOOP
		case <-timer.C:
			t.latestRun = time.Now()
			stepRun(t)
		}
	}
}

func willRunAfter(t Task) time.Duration {
	interval := time.Duration(t.interval)
	unit := t.unit

	now := time.Now()
	y, m, d := now.Date()
	rt := time.Date(y, m, d, 0, 0, 0, now.Nanosecond(), time.Local)

	// at specific time
	ts := make([]string, 3)
	if len(t.at) > 0 {
		ts = strings.Split(t.at, ":")
	}
	if v, err := strconv.Atoi(ts[0]); err == nil {
		rt = rt.Add(time.Duration(v) * time.Hour)
	} else {
		rt = rt.Add(time.Duration(now.Hour()) * time.Hour)
	}
	if v, err := strconv.Atoi(ts[1]); err == nil {
		rt = rt.Add(time.Duration(v) * time.Minute)
	} else {
		rt = rt.Add(time.Duration(now.Minute()) * time.Minute)
	}
	if v, err := strconv.Atoi(ts[2]); err == nil {
		rt = rt.Add(time.Duration(v) * time.Second)
	} else {
		rt = rt.Add(time.Duration(now.Second()) * time.Second)
	}

	// on specific day
	if unit == weekday {
		wd := rt.Weekday()
		rt = rt.AddDate(0, 0, int(t.weekday+7-wd)%7)
		// wait for a week if we have run it today yet
		// or we have missed that time we specified today
		if t.weekday == wd && !t.latestRun.IsZero() || rt.Before(now) {
			rt = rt.AddDate(0, 0, 7)
		}

	} else {
		if now.Before(rt) {
			return rt.Sub(now)
		}

		// purely periodic
		// run immediate if not specified
		if len(t.at) == 0 && t.latestRun.IsZero() {
			return time.Duration(0)
		}
		switch unit {
		case seconds:
			rt = rt.Add(time.Duration(interval * time.Second))
		case minutes:
			rt = rt.Add(time.Duration(interval * time.Minute))
		case hours:
			rt = rt.Add(time.Duration(interval * time.Hour))
		case days:
			rt = rt.Add(time.Duration(interval * time.Hour * 24))
		}
	}

	return rt.Sub(now)
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
