package inheritence

type Tape struct {
	Type string
}

func (tape *Tape) SetType(t string) {
	tape.Type = t
}

func (tape *Tape) GetType() string {
	return tape.Type
}
