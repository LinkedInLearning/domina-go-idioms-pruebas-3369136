package main

import (
	"log"

	"github.com/linkedinlearning/domina-go/package-cycles/pokemon"
	"github.com/linkedinlearning/domina-go/package-cycles/trainer"
)

func main() {
	ash := trainer.Ash()
	pikachu := pokemon.NewPikachu()

	log.Println(ash)
	log.Println(pikachu)
}
