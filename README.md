# gron: Yet another task scheduling utility

![build](https://github.com/u2386/gron/workflows/Go/badge.svg?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/u2386/gron)](https://goreportcard.com/report/github.com/u2386/gron)
[![Coverage Status](https://coveralls.io/repos/github/u2386/gron/badge.svg?branch=master)](https://coveralls.io/github/u2386/gron?branch=master)

Gron is an elegant and simple periodic task scheduler, for human.

Inspired by Python pacakge [schedule](https://github.com/dbader/schedule) and Go package [goCron](https://github.com/jasonlvhit/gocron)

## Feature

* Parallel scheduling
* Isolated tasks
* Panic-proof
* EventBus

## Example

```go
package main

import (
	"fmt"
	"time"

	. "github.com/u2386/gron"
)

func main() {
	Gron(
		Every(10),
		Seconds(),
		Do(func() {
			fmt.Println("FireBall!", time.Now().Format("15:04:05"))
		}),
		Name("Fireball!"),
	)

	Gron(
		Every(1),
		Minutes(),
		At("::01"),
		Do(func() {
			fmt.Println("Cast Blizzard!", time.Now().Format("15:04:05"))
		}),
		Name("Blizzard!"),
	)

	Gron(
		Every(4),
		Seconds(),
		Do(func() {
			panic("Oops...") // panic-proof
		}),
		Name("Thunderclap!"),
	)

	Gron(
		Every(1),
		Saturday(),
		Do(func() {
			time.Sleep(4)
			fmt.Println("wubba lubba dub dub")
		}),
		Name("Rick"),
	)

	// subscribe task events
	for ev := range Subscribe() {
		switch ev.E {
		case Disabled:
			fmt.Println("Disabled:", ev.TaskName, ev.At.Format("15:04:05"))
		case Failed:
			fmt.Println("Failed:", ev.TaskName, ev.Msg, ev.At.Format("15:04:05"))
			Remove(ev.TaskName)
		default:
			fmt.Println(ev.TaskName, ev.E, ev.At.Format("15:04:05"))
		}
	}
}

```

## License

MIT
