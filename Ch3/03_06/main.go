package main

import (
	"log"
	"math/rand"
)

func main() {}

// PokemonAppears simula la aparición de un pokemon de los cuatro posibles,
// devolviendo un pokemon aleatorio.
func PokemonAppears() Pokemon {
	pokemons := []Pokemon{Pikachu(), Charmander(), Bulbasaur(), Squirtle()}
	p := pokemons[rand.Intn(4)]

	if p.Friendly {
		log.Printf("A friendly %s appears!\n", p.Name)
		return p
	}

	log.Printf("A wild %s appears!\n", p.Name)
	return p
}
