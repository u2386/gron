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

	start := time.Now()
	// subscribe task events
	for ev := range Subscribe() {
		if ev.E == Empty {
			break

		} else if ev.E == Disabled {
			fmt.Println("Disabled:", ev.TaskName)
			Remove(ev.TaskName)

		} else if time.Since(start) > 10*time.Second {
			Disable(ev.TaskName)

		} else {
			fmt.Println(ev.TaskName, ev.E)
		}
	}
}
