package inheritence

import "fmt"

type TuringMachine struct {
	cpu  CPU
	tape Tape
}

func (turingMachine *TuringMachine) ConfigureSpecification(cpu CPU, tape Tape) {
	turingMachine.cpu.SetArchitecture(cpu.GetArchitecture())
	turingMachine.tape.SetType(tape.GetType())
}

func (turingMachine *TuringMachine) ShowSpecification() {
	fmt.Printf(
		"CPU: %s\nTape: %s\n",
		turingMachine.cpu.GetArchitecture(),
		turingMachine.tape.GetType(),
	)
}
