package main

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
