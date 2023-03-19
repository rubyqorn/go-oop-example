package inheritence

type CPU struct {
	Architecture string
}

func (cpu *CPU) SetArchitecture(architecture string) {
	cpu.Architecture = architecture
}

func (cpu *CPU) GetArchitecture() string {
	return cpu.Architecture
}
