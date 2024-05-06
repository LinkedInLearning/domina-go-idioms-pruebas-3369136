package main

import "fmt"

func main() {
	pikachu := NewPokemon(
		"Pikachu",
		55, 40, 35,
		[]string{"Electric"},
		[]string{"Quick Attack", "Thunderbolt"},
	)

	fmt.Println(pikachu)

	// este pokemon tiene todos los valores
	pikachuWithOptions := NewPokemonWithOptions(
		"Pikachu",
		WithAttack(55), WithDefense(40), WithLife(35),
		WithTypes([]string{"Electric"}),
		WithMoves([]string{"Quick Attack", "Thunderbolt"}),
	)

	fmt.Println(pikachuWithOptions)

	// este pokemon tiene solo el nombre
	charmander := NewPokemonWithOptions("Charmander")

	fmt.Println(charmander)

	// este pokemon tiene solo el nombre y los puntos
	charizard := NewPokemonWithOptions(
		"Charizard",
		WithAttack(84), WithDefense(78), WithLife(78),
	)

	fmt.Println(charizard)
}
