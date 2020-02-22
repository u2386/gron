# gron: Yet another crontab-alike utility

Gron is an elegant and simple periodic Golang task scheduler, built for human beings.
And, yes, it is `parallel`.

## Example

```go
package main

import (
	"fmt"

	. "github.com/u2386/gron"
)

func main() {
	Gron(
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

```

## License

MIT
