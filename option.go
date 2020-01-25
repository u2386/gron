package gron

// Option is a struct for friendly APIs
// See also: https://dave.cheney.net/2014/10/17/functional-options-for-friendly-apis
type Option func(task Task) (Task, error)

// Every sets the interval value
func Every(n uint64) Option {
	return func(task Task) (Task, error) {
		task.interval = n
		return task, nil
	}
}

// Seconds sets the interval unit
func Seconds() Option {
	return func(task Task) (Task, error) {
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

// Args set the arguments of the task
func Args(args ...interface{}) Option {
	return func(task Task) (Task, error) {
		return task, nil
	}
}
