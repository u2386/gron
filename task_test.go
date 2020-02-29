package gron

import (
	"testing"
	"time"

	"bou.ke/monkey"
	"github.com/stretchr/testify/assert"
)

func TestNavieInterval(t *testing.T) {
	interval := uint(5)
	tests := []struct {
		now      time.Time
		unit     timeUnit
		expected time.Time
	}{
		{
			now:      time.Date(2020, time.April, 1, 15, 24, 48, 0, time.Local),
			unit:     seconds,
			expected: time.Date(2020, time.April, 1, 15, 24, 53, 0, time.Local),
		},
		{
			now:      time.Date(2020, time.April, 1, 15, 24, 48, 0, time.Local),
			unit:     minutes,
			expected: time.Date(2020, time.April, 1, 15, 29, 48, 0, time.Local),
		},
		{
			now:      time.Date(2020, time.April, 1, 20, 24, 48, 0, time.Local),
			unit:     hours,
			expected: time.Date(2020, time.April, 2, 1, 24, 48, 0, time.Local),
		},
		{
			now:      time.Date(2020, time.February, 28, 15, 24, 48, 0, time.Local),
			unit:     days,
			expected: time.Date(2020, time.March, 4, 15, 24, 48, 0, time.Local),
		},
	}

	for i, tt := range tests {
		func() {
			patch := monkey.Patch(time.Now, func() time.Time { return tt.now })
			defer patch.Unpatch()

			task := newTask()
			task.interval = interval
			task.unit = tt.unit

			assert.Equal(t, time.Duration(0), willRunAfter(task))
			task.latestRun = time.Now()
			assert.Equal(t, tt.expected.Sub(time.Now()), willRunAfter(task), "case(%d)", i)
		}()
	}
}

func TestWillRunSpecifiedTime(t *testing.T) {
	tests := []struct {
		now      time.Time
		interval int
		at       string
		note     string
		expected time.Time
		weekday  time.Weekday
		unit     timeUnit
	}{
		{
			now:      time.Date(2020, time.April, 1, 15, 24, 48, 0, time.Local),
			interval: 1,
			at:       "16:30:50",
			note:     "Within this unit time",
			unit:     days,
			expected: time.Date(2020, time.April, 1, 16, 30, 50, 0, time.Local),
		},
		{
			now:      time.Date(2020, time.April, 1, 15, 24, 48, 0, time.Local),
			interval: 1,
			at:       ":30:50",
			note:     "Within this unit time",
			unit:     hours,
			expected: time.Date(2020, time.April, 1, 15, 30, 50, 0, time.Local),
		},
		{
			now:      time.Date(2020, time.April, 1, 15, 24, 48, 0, time.Local),
			interval: 1,
			at:       "::50",
			note:     "Within this unit time",
			unit:     minutes,
			expected: time.Date(2020, time.April, 1, 15, 24, 50, 0, time.Local),
		},
		{
			now:      time.Date(2020, time.April, 1, 16, 30, 48, 0, time.Local),
			interval: 1,
			at:       "16:30:04",
			note:     "Shoud wait for another unit",
			unit:     days,
			expected: time.Date(2020, time.April, 2, 16, 30, 04, 0, time.Local),
		},
		{
			now:      time.Date(2020, time.April, 1, 16, 30, 48, 0, time.Local),
			interval: 1,
			at:       ":30:04",
			note:     "Shoud wait for another unit",
			unit:     hours,
			expected: time.Date(2020, time.April, 1, 17, 30, 4, 0, time.Local),
		},
		{
			now:      time.Date(2020, time.April, 1, 16, 30, 48, 0, time.Local),
			interval: 1,
			at:       "::04",
			note:     "Shoud wait for another unit",
			unit:     minutes,
			expected: time.Date(2020, time.April, 1, 16, 31, 4, 0, time.Local),
		},

		// weekday
		{
			now:      time.Date(2020, time.February, 28, 16, 30, 48, 0, time.Local), // Friday
			interval: 1,
			note:     "Within this week and run right now",
			unit:     weekday,
			weekday:  time.Friday,
			expected: time.Date(2020, time.February, 28, 16, 30, 48, 0, time.Local),
		},
		{
			now:      time.Date(2020, time.February, 28, 16, 30, 48, 0, time.Local), // Friday
			interval: 1,
			note:     "Within this week",
			unit:     weekday,
			weekday:  time.Saturday,
			expected: time.Date(2020, time.February, 29, 16, 30, 48, 0, time.Local),
		},
		{
			now:      time.Date(2020, time.February, 28, 16, 30, 48, 0, time.Local), // Friday
			interval: 1,
			note:     "Within this week if specified `at`",
			at:       "00:00:00",
			unit:     weekday,
			weekday:  time.Saturday,
			expected: time.Date(2020, time.February, 29, 00, 00, 00, 0, time.Local),
		},
		{
			now:      time.Date(2020, time.February, 28, 16, 30, 48, 0, time.Local), // Friday
			interval: 1,
			note:     "Next week",
			unit:     weekday,
			weekday:  time.Sunday,
			expected: time.Date(2020, time.March, 1, 16, 30, 48, 0, time.Local),
		},
		{
			now:      time.Date(2020, time.February, 28, 16, 30, 48, 0, time.Local), // Friday
			interval: 1,
			note:     "Next week because specified `at`",
			at:       "16:29:00",
			unit:     weekday,
			weekday:  time.Friday,
			expected: time.Date(2020, time.March, 6, 16, 29, 00, 0, time.Local),
		},
	}

	var task Task
	for i, tt := range tests {
		func() {
			patch := monkey.Patch(time.Now, func() time.Time { return tt.now })
			defer patch.Unpatch()

			task = newTask()
			task.unit = tt.unit
			task.interval = uint(tt.interval)
			task.at = tt.at
			task.weekday = tt.weekday

			assert.Equal(t, tt.expected.Sub(time.Now()), willRunAfter(task), "case(%d) note(%s)", i, tt.note)
		}()
	}
}
