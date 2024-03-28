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

	// this pokemon has all the values
	pikachuWithOptions := NewPokemonWithOptions(
		"Pikachu",
		WithAttack(55), WithDefense(40), WithLife(35),
		WithTypes([]string{"Electric"}),
		WithMoves([]string{"Quick Attack", "Thunderbolt"}),
	)

	fmt.Println(pikachuWithOptions)

	// this pokemon has only the name
	charmander := NewPokemonWithOptions("Charmander")

	fmt.Println(charmander)

	// this pokemon has only the name and the points
	charizard := NewPokemonWithOptions(
		"Charizard",
		WithAttack(84), WithDefense(78), WithLife(78),
	)

	fmt.Println(charizard)
}

type Pokemon struct {
	Name    string
	Attack  int
	Defense int
	Life    int
	Types   []string
	Moves   []string
}

func NewPokemon(name string, attack int, defense int, life int,
	types []string, moves []string) *Pokemon {

	return &Pokemon{
		Name:    name,
		Attack:  attack,
		Defense: defense,
		Life:    life,
		Types:   types,
		Moves:   moves,
	}
}
