package polymorphism

import "fmt"

type PowerOffButton struct {
	Color string
}

func (button *PowerOffButton) turn() {
	fmt.Printf(
		"Button is: %s, You are gonna power off machine...\n",
		button.Color,
	)
}
