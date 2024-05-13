package main

import (
	"log"

	"github.com/linkedinlearning/domina-go/package-no-cycles/pokemon"
	"github.com/linkedinlearning/domina-go/package-no-cycles/trainer"
)

// La estructura Player representa un jugador de Pokemon, y en ella almacenamos
// un entrenador y una lista de pokemons, de manera desacoplada.
type Player struct {
	Player   trainer.Trainer
	Pokemons []pokemon.Pokemon
}

func main() {
	ash := Player{
		Player: trainer.Ash(),
		Pokemons: []pokemon.Pokemon{
			pokemon.NewPikachu(),
		},
	}

	log.Println(ash)
}
