package main

import (
	"fmt"
	"log"
	"math/rand"
)

func main() {}

type Trainer struct {
	Name             string
	Pokeballs        int
	CapturedPokemons []Pokemon
}

func (t *Trainer) Capture(pokemon Pokemon) error {
	err := t.ThrowBall()
	if err != nil {
		return err
	}

	if pokemon.Friendly {
		t.CapturedPokemons = append(t.CapturedPokemons, pokemon)
		return nil
	}

	return fmt.Errorf("pokemon is not friendly")

}

func (t *Trainer) ThrowBall() error {
	if t.Pokeballs == 0 {
		return fmt.Errorf("trainer has no pokeballs")
	}

	t.Pokeballs--
	return nil
}

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
