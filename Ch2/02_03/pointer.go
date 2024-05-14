package main

// PokemonPointerOption utiliza semántica de puntero para modificar la estructura Pokemon.
// Como veremos, ésto nos permite modificar la estructura original. Al finalizar
// la función, la estructura original se modifica.
type PokemonPointerOption func(*Pokemon) error

func WithAttack(attack int) PokemonPointerOption {
	return func(p *Pokemon) error {
		p.Attack = attack
		return nil
	}
}

func WithDefense(defense int) PokemonPointerOption {
	return func(p *Pokemon) error {
		p.Defense = defense
		return nil
	}
}

func WithLife(life int) PokemonPointerOption {
	return func(p *Pokemon) error {
		p.Life = life
		return nil
	}
}

func WithMoves(moves []string) PokemonPointerOption {
	return func(p *Pokemon) error {
		p.Moves = moves
		return nil
	}
}

func WithTypes(types []string) PokemonPointerOption {
	return func(p *Pokemon) error {
		p.Types = types
		return nil
	}
}

func NewPokemonPointer(name string, opts ...PokemonPointerOption) *Pokemon {
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
