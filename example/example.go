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
