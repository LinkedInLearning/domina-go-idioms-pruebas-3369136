package trainer

import "github.com/linkedinlearning/domina-go/package-cycles/pokemon"

type Trainer struct {
	Name     string
	Pokemons []pokemon.Pokemon
}

func Ash() Trainer {
	return Trainer{
		Name: "Ash Ketchum",
	}
}
