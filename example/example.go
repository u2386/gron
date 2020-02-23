package main

import (
	"fmt"
	"time"

	. "github.com/u2386/gron"
)

func main() {
	Gron(
		Every(2),
		Seconds(),
		Do(func() {
			fmt.Println("2 seconds elapsed...", time.Now().String())
		}),
		Name("Fireball!"),
	)

	Gron(
		Every(3),
		Seconds(),
		Do(func() {
			fmt.Println("3 seconds elapsed...", time.Now().String())
		}),
		Name("Blizzard!"),
	)

	Gron(
		Every(4),
		Seconds(),
		Do(func() {
			panic("Oops...")
		}),
		Name("Thunderclap!"),
	)

	start := time.Now()

LOOP:
	// subscribe task events
	for ev := range Subscribe() {
		if time.Since(start) > 10*time.Second {
			Disable(ev.TaskName)
		}

		switch ev.E {
		case Empty:
			fmt.Println(ev.E)
			break LOOP
		case Disabled:
			fmt.Println("Disabled:", ev.TaskName)
			Remove(ev.TaskName)
		case Failed:
			fmt.Println("Failed:", ev.TaskName, ev.Msg)
			Disable(ev.TaskName)
		default:
			fmt.Println(ev.TaskName, ev.E)
		}
	}
}
