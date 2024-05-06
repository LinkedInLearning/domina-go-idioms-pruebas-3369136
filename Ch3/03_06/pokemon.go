package main

import "math/rand"

type Pokemon struct {
	Name     string
	Friendly bool
	Attack   int
	Defense  int
	Life     int
	Types    []string
	Moves    []string
}

func Pikachu() Pokemon {
	return Pokemon{
		Name:     "Pikachu",
		Friendly: rand.Intn(80) > 20,
		Attack:   55,
		Defense:  40,
		Life:     35,
		Types:    []string{"Electric"},
		Moves:    []string{"Thunder Shock", "Quick Attack"},
	}
}

func Charmander() Pokemon {
	return Pokemon{
		Name:     "Charmander",
		Friendly: rand.Intn(80) > 30,
		Attack:   52,
		Defense:  43,
		Life:     39,
		Types:    []string{"Fire"},
		Moves:    []string{"Ember", "Scratch"},
	}
}

func Bulbasaur() Pokemon {
	return Pokemon{
		Name:     "Bulbasaur",
		Friendly: rand.Intn(80) > 40,
		Attack:   49,
		Defense:  49,
		Life:     45,
		Types:    []string{"Grass", "Poison"},
		Moves:    []string{"Vine Whip", "Tackle"},
	}
}

func Squirtle() Pokemon {
	return Pokemon{
		Name:     "Squirtle",
		Friendly: rand.Intn(80) > 50,
		Attack:   48,
		Defense:  65,
		Life:     44,
		Types:    []string{"Water"},
		Moves:    []string{"Water Gun", "Tackle"},
	}
}
