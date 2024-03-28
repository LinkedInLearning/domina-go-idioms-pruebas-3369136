package pokemon

import "github.com/linkedinlearning/domina-go/package-cycles/trainer"

type Pokemon struct {
	Name    string
	Trainer trainer.Trainer
}

func NewPikachu() Pokemon {
	return Pokemon{
		Name: "Pikachu",
	}
}
