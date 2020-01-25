package main

import (
	"fmt"

	. "github.com/u2386/gron"
)

func main() {
	Crontab(
		Every(2),
		Seconds(),
		Do(func() {
			fmt.Println("2 seconds elapsed...")
		}),
		Name("Fireball!"),
	)

	counter := 2
	// subscribe task events
	for ev := range Subscribe() {
		if ev.TaskName == "Fireball!" && ev.E == Running {
			counter--
		} else if ev.TaskName == "Fireball!" && ev.E == Finished && counter == 0 {
			Disable(ev.TaskName)
			Remove(ev.TaskName)
		} else if ev.TaskName == "Fireball!" && ev.E == Disabled {
			break
		}
	}
}
