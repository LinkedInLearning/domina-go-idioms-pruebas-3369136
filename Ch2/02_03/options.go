package main

import "fmt"

// tipo para opciones funcionales
type PokemonOption func(*Pokemon) error

func WithAttack(attack int) PokemonOption {
	return func(p *Pokemon) error {
		if attack < 0 {
			return fmt.Errorf("attack must be greater than 0")
		}

		p.Attack = attack

		return nil
	}
}

func WithDefense(defense int) PokemonOption {
	return func(p *Pokemon) error {
		if defense < 0 {
			return fmt.Errorf("defense must be greater than 0")
		}

		p.Defense = defense

		return nil
	}
}

func WithLife(life int) PokemonOption {
	return func(p *Pokemon) error {
		if life < 0 {
			return fmt.Errorf("life must be greater than 0")
		}

		p.Life = life

		return nil
	}
}

func WithMoves(moves []string) PokemonOption {
	return func(p *Pokemon) error {
		p.Moves = moves

		return nil
	}
}

func WithTypes(types []string) PokemonOption {
	return func(p *Pokemon) error {
		p.Types = types

		return nil
	}
}

func NewPokemonWithOptions(name string, opts ...PokemonOption) *Pokemon {
	p := &Pokemon{
		Name: name,
	}

	for _, opt := range opts {
		if err := opt(p); err != nil {
			return nil
		}
	}

	return p
}
