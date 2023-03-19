package main

import (
	"github.com/rubyqorn/go-oop-example/oop/encapsulation"
	"github.com/rubyqorn/go-oop-example/oop/inheritence"
	"github.com/rubyqorn/go-oop-example/oop/polymorphism"
)

func main() {
	people := encapsulation.People{}

	people.SetPeople(1, "Alan", "Turing", "Mathematician")
	people.Introduce()

	cpu := inheritence.CPU{}
	tape := inheritence.Tape{}
	turingMachine := inheritence.TuringMachine{}

	cpu.SetArchitecture("Own based architecture")
	tape.SetType("Magnetic")
	turingMachine.ConfigureSpecification(cpu, tape)
	turingMachine.ShowSpecification()

	calculationBtn := polymorphism.CalculationButton{Color: "green"}
	powerOffBtn := polymorphism.PowerOffButton{Color: "red"}

	polymorphism.Turn(&calculationBtn)
	polymorphism.Turn(&powerOffBtn)
}
