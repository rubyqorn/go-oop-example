package polymorphism

import "fmt"

type CalculationButton struct {
	Color string
}

func (button *CalculationButton) turn() {
	fmt.Printf(
		"Button is: %s, You are gonna run calculation...\n",
		button.Color,
	)
}
