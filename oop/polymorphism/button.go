package polymorphism

type Button interface {
	turn()
}

func Turn(button Button) {
	button.turn()
}
