package main

import "fmt"

type Trainer struct {
	Name             string
	Pokeballs        int
	CapturedPokemons []Pokemon
}

// Capture intenta capturar un pokemon, lanzando una pokeball y añadiendo el pokemon a la lista de pokemons capturados.
// Si el pokemon no es amigable, la captura falla, y si el pokemon es amigable, se añade a la lista de pokemons capturados.
// Si el entrenador no tiene pokeballs, la captura falla.
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

// ThrowBall lanza una pokeball, si el entrenador tiene pokeballs disponibles.
// Devuelve un error si el entrenador no tiene pokeballs.
func (t *Trainer) ThrowBall() error {
	if t.Pokeballs == 0 {
		return fmt.Errorf("trainer has no pokeballs")
	}

	t.Pokeballs--
	return nil
}
