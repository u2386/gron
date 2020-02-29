package gron

import "time"

// Option is a struct for friendly APIs
// See also: https://dave.cheney.net/2014/10/17/functional-options-for-friendly-apis
type Option func(task Task) (Task, error)

// Every sets the interval value
func Every(n uint) Option {
	return func(task Task) (Task, error) {
		task.interval = n
		return task, nil
	}
}

// Seconds sets the interval unit
func Seconds() Option {
	return func(task Task) (Task, error) {
		task.unit = seconds
		return task, nil
	}
}

// Second is a singular Seconds
func Second() Option {
	return Seconds()
}

// Minutes sets the interval unit
func Minutes() Option {
	return func(task Task) (Task, error) {
		task.unit = minutes
		return task, nil
	}
}

// Minute is a singular minutes
func Minute() Option {
	return Minutes()
}

// Hours sets the interval unit
func Hours() Option {
	return func(task Task) (Task, error) {
		task.unit = hours
		return task, nil
	}
}

// Hour is a singular Hours
func Hour() Option {
	return Hours()
}

// At sets the specific time of a day
// Support one of the following spec:
//   * HH:MM:SS
//   * :MM:SS
//   * ::SS
func At(ts string) Option {
	return func(task Task) (Task, error) {
		task.at = ts
		return task, nil
	}
}

// Sunday makes sure that this task will be run on Sunday
func Sunday() Option {
	return func(task Task) (Task, error) {
		task.weekday = time.Sunday
		task.unit = weekday
		return task, nil
	}
}

// Monday makes sure that this task will be run on Monday
func Monday() Option {
	return func(task Task) (Task, error) {
		task.weekday = time.Monday
		task.unit = weekday
		return task, nil
	}
}

// Tuesday makes sure that this task will be run on Tuesday
func Tuesday() Option {
	return func(task Task) (Task, error) {
		task.weekday = time.Tuesday
		task.unit = weekday
		return task, nil
	}
}

// Wednesday makes sure that this task will be run on Wednesday
func Wednesday() Option {
	return func(task Task) (Task, error) {
		task.weekday = time.Wednesday
		task.unit = weekday
		return task, nil
	}
}

// Thursday makes sure that this task will be run on Thursday
func Thursday() Option {
	return func(task Task) (Task, error) {
		task.weekday = time.Thursday
		task.unit = weekday
		return task, nil
	}
}

// Friday makes sure that this task will be run on Friday
func Friday() Option {
	return func(task Task) (Task, error) {
		task.weekday = time.Friday
		task.unit = weekday
		return task, nil
	}
}

// Saturday makes sure that this task will be run on Saturday
func Saturday() Option {
	return func(task Task) (Task, error) {
		task.weekday = time.Saturday
		task.unit = weekday
		return task, nil
	}
}

// Do assigns the periodic task
func Do(fn interface{}) Option {
	return func(task Task) (Task, error) {
		task.fn = fn
		return task, nil
	}
}

// Name sets the identifier of the task
func Name(name string) Option {
	return func(task Task) (Task, error) {
		task.Name = TaskName(name)
		return task, nil
	}
}
