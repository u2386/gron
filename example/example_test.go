package main

import (
	"fmt"
	"testing"
	"time"

	. "github.com/u2386/gron"
)

func TestRace(t *testing.T) {
	Gron(
		Every(2),
		Seconds(),
		Do(func() {
			fmt.Println("2 seconds elapsed...", time.Now().Format("15:04:05"))
		}),
		Name("Fireball!"),
	)

	Gron(
		Every(3),
		Seconds(),
		Do(func() {
			fmt.Println("3 seconds elapsed...", time.Now().Format("15:04:05"))
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
	counter := 3
	// subscribe task events
	for ev := range Subscribe() {
		switch ev.E {
		case Disabled:
			fmt.Println("Disabled:", ev.TaskName, ev.At.Format("15:04:05"))
			counter--
		case Failed:
			fmt.Println("Failed:", ev.TaskName, ev.Msg, ev.At.Format("15:04:05"))
			Remove(ev.TaskName)
		default:
			fmt.Println(ev.TaskName, ev.E, ev.At.Format("15:04:05"))
		}

		if counter == 0 {
			break
		}
		if time.Since(start) > 10*time.Second {
			Remove(ev.TaskName)
		}
	}
}
