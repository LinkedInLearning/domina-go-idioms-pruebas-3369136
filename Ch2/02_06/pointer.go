package main

// PokemonPointerOption utiliza semántica de puntero para modificar la estructura Pokemon.
// Como veremos, ésto nos permite modificar la estructura original. Al finalizar
// la función, la estructura original se modifica.
type PokemonPointerOption func(*Pokemon)

func WithAttack(attack int) PokemonPointerOption {
	return func(p *Pokemon) {
		p.Attack = attack
	}
}

func WithDefense(defense int) PokemonPointerOption {
	return func(p *Pokemon) {
		p.Defense = defense
	}
}

func WithLife(life int) PokemonPointerOption {
	return func(p *Pokemon) {
		p.Life = life
	}
}

func WithMoves(moves []string) PokemonPointerOption {
	return func(p *Pokemon) {
		p.Moves = moves
	}
}

func WithTypes(types []string) PokemonPointerOption {
	return func(p *Pokemon) {
		p.Types = types
	}
}

func NewPokemonPointer(name string, opts ...PokemonPointerOption) *Pokemon {
	p := &Pokemon{
		Name: name,
	}

	for _, opt := range opts {
		opt(p)
	}

	return p
}
