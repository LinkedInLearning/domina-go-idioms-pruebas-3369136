package main

import (
	"context"
	"fmt"
	"testing"

	"github.com/cucumber/godog"
)

func TestFeatures(t *testing.T) {
	suite := godog.TestSuite{
		ScenarioInitializer: InitializeScenario,
		Options: &godog.Options{
			Format:   "pretty",
			Paths:    []string{"features"}, // ubicación de los feature files
			TestingT: t,                    // Testing instance that will run subtests.
		},
	}

	if suite.Run() != 0 {
		t.Fatal("non-zero status returned, failed to run feature tests")
	}
}

// PokemonScenario es una estructura que contiene el estado del escenario,
// en este caso un pokemon y un trainer para verificar la captura de pokemons.
type PokemonScenario struct {
	pokemon Pokemon
	trainer Trainer
}

// función que inicializa el escenario con los pasos que se van a ejecutar
// En ella se inicializa el trainer con 5 pokeballs y se resetean los pokemons capturados,
// y además definimos tanto los pasos del escenario como lo que va a ocurrir antes y después de cada escenario.
func InitializeScenario(ctx *godog.ScenarioContext) {
	// Inicializamos el escenario con un trainer y 5 pokeballs
	scenario := &PokemonScenario{
		trainer: Trainer{
			Name:             "Ash Ketchum",
			Pokeballs:        5,
			CapturedPokemons: []Pokemon{},
		},
	}

	// volvemos al estado inicial antes de cada escenario
	ctx.Before(func(ctx context.Context, sc *godog.Scenario) (context.Context, error) {
		scenario.trainer.Pokeballs = 5                  // volvemos a poner las 5 pokeballs al trainer
		scenario.trainer.CapturedPokemons = []Pokemon{} // reseteamos los pokemons capturados
		scenario.pokemon = Pokemon{}                    // reseteamos el pokemon

		return ctx, nil
	})

	// definimos los pasos del escenario, utilizando expresiones regulares,
	// que se corresponden con los pasos de los feature files, conectando
	// con el código Go que implementa cada paso.
	ctx.Given(`^a Pokemon appears$`, scenario.pokemonAppears)
	ctx.Given(`^the Pokemon is friendly$`, scenario.pokemonIsFriendly)
	ctx.Given(`^the Pokemon is not friendly$`, scenario.pokemonIsNotFriendly)
	ctx.Given(`^the trainer has no pokeballs$`, scenario.trainerHasNoPokeballs)
	ctx.When(`^the trainer tries to capture the Pokemon$`, scenario.trainerTriesToCapture)
	ctx.Then(`^the Pokemon is captured$`, scenario.pokemonIsCaptured)
	ctx.Then(`^the Pokemon escapes$`, scenario.pokemonEscapes)
}

// vamos a forzar que el pokemon sea amigable
func (s *PokemonScenario) pokemonIsFriendly() error {
	s.pokemon.Friendly = true
	return nil
}

// vamos a forzar que el pokemon NO sea amigable
func (s *PokemonScenario) pokemonIsNotFriendly() error {
	s.pokemon.Friendly = false
	return nil
}

// función que simula la captura de un pokemon por parte de un trainer,
// llamando al método Capture del trainer y comprobando si la captura es exitosa.
func (s *PokemonScenario) trainerTriesToCapture() error {
	friendly := s.pokemon.Friendly
	hasPokeballs := s.trainer.Pokeballs > 0

	err := s.trainer.Capture(s.pokemon)
	if err != nil {
		// la captura falla porque el trainer no tiene pokeballs
		if !hasPokeballs && err.Error() == "trainer has no pokeballs" {
			return nil
		}

		// la captura falla porque el pokemon no es amigable
		if !friendly && err.Error() == "pokemon is not friendly" {
			return nil
		}

		// la captura falla por otro motivo, por lo que devolvemos el error
		return err
	}

	// si no hubo error, pero el trainer no tenía pokeballs, la captura debería haber fallado
	if !hasPokeballs {
		return fmt.Errorf("capture should fail because trainer has no pokeballs")
	}

	// si no hubo error, pero el pokemon no era amigable, la captura debería haber fallado
	if !friendly {
		return fmt.Errorf("capture should fail because pokemon is not friendly")
	}

	return nil
}

// forzamos que el trainer no tenga pokeballs
func (s *PokemonScenario) trainerHasNoPokeballs() error {
	s.trainer.Pokeballs = 0
	return nil
}

// comprobamos que el pokemon capturado sea el mismo que el que apareció
func (s *PokemonScenario) pokemonIsCaptured() error {
	if s.trainer.CapturedPokemons[0].Name != s.pokemon.Name {
		return fmt.Errorf("captured pokemon is not the same as the one that appeared")
	}

	return nil
}

// aparece un pokemon nuevo
func (s *PokemonScenario) pokemonAppears() error {
	s.pokemon = PokemonAppears()
	return nil
}

// recorremos la lista de pokemons capturados para comprobar que no está el pokemon que apareció
func (s *PokemonScenario) pokemonEscapes() error {
	for _, p := range s.trainer.CapturedPokemons {
		if p.Name == s.pokemon.Name {
			return fmt.Errorf("captured pokemon should not be the same as the one that appeared")
		}
	}

	return nil
}
