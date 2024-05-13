package main

import (
	"fmt"

	"github.com/linkedinlearning/domina-go/package-layout/pokemon"
)

func main() {
	pikachu := pokemon.Pikachu()
	fmt.Println(pikachu)

	charmander := pokemon.Charmander()
	fmt.Println(charmander)

	bulbasaur := pokemon.Bulbasaur()
	fmt.Println(bulbasaur)

	squirtle := pokemon.Squirtle()
	fmt.Println(squirtle)

	pokeType := pokemon.pokemonType("Electric")
	fmt.Println(pokeType)
}
