package main

// PokemonValueOption utiliza semántica de valor para modificar la estructura Pokemon.
// Esta opción es exactamente igual a la opción de puntero, pero en lugar de recibir
// un puntero a la estructura, recibe la estructura directamente.
// Como veremos, ésto nos permite modificar la estructura original, pero
// al finalizar la función, la estructura original no se modifica, puesto
// que estamos pasando una copia de la estructura.
type PokemonValueOption func(Pokemon)

func WithAttackValue(attack int) PokemonValueOption {
	return func(p Pokemon) {
		p.Attack = attack
	}
}

func WithDefenseValue(defense int) PokemonValueOption {
	return func(p Pokemon) {
		p.Defense = defense
	}
}

func WithLifeValue(life int) PokemonValueOption {
	return func(p Pokemon) {
		p.Life = life
	}
}

func WithMovesValue(moves []string) PokemonValueOption {
	return func(p Pokemon) {
		p.Moves = moves
	}
}

func WithTypesValue(types []string) PokemonValueOption {
	return func(p Pokemon) {
		p.Types = types
	}
}

func NewPokemonValue(name string, opts ...PokemonValueOption) Pokemon {
	p := Pokemon{
		Name: name,
	}

	for _, opt := range opts {
		opt(p)
	}

	return p
}
