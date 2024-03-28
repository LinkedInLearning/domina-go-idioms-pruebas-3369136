package main

import (
	"log"

	"github.com/linkedinlearning/domina-go/package-no-cycles/pokemon"
	"github.com/linkedinlearning/domina-go/package-no-cycles/trainer"
)

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
