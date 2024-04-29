package main

// PokemonValueOption utiliza semántica de valor para modificar la estructura Pokemon.
// Esta opción es exactamente igual a la opción de puntero, pero en lugar de recibir
// un puntero a la estructura, recibe la estructura directamente.
// Como veremos, ésto nos permite modificar la estructura original, pero
// al finalizar la función, la estructura original no se modifica, puesto
// que estamos pasando una copia de la estructura.
type PokemonValueOption func(Pokemon) error

func WithAttackValue(attack int) PokemonValueOption {
	return func(p Pokemon) error {
		p.Attack = attack
		return nil
	}
}

func WithDefenseValue(defense int) PokemonValueOption {
	return func(p Pokemon) error {
		p.Defense = defense
		return nil
	}
}

func WithLifeValue(life int) PokemonValueOption {
	return func(p Pokemon) error {
		p.Life = life
		return nil
	}
}

func WithMovesValue(moves []string) PokemonValueOption {
	return func(p Pokemon) error {
		p.Moves = moves
		return nil
	}
}

func WithTypesValue(types []string) PokemonValueOption {
	return func(p Pokemon) error {
		p.Types = types
		return nil
	}
}

func NewPokemonValue(name string, opts ...PokemonValueOption) Pokemon {
	p := Pokemon{
		Name: name,
	}

	for _, opt := range opts {
		if err := opt(p); err != nil {
			return Pokemon{}
		}
	}

	return p
}
