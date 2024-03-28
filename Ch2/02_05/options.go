package main

// type for functional options
type PokemonOption func(*Pokemon)

func WithAttack(attack int) PokemonOption {
	return func(p *Pokemon) {
		p.Attack = attack
	}
}

func WithDefense(defense int) PokemonOption {
	return func(p *Pokemon) {
		p.Defense = defense
	}
}

func WithLife(life int) PokemonOption {
	return func(p *Pokemon) {
		p.Life = life
	}
}

func WithMoves(moves []string) PokemonOption {
	return func(p *Pokemon) {
		p.Moves = moves
	}
}

func WithTypes(types []string) PokemonOption {
	return func(p *Pokemon) {
		p.Types = types
	}
}

func NewPokemonWithOptions(name string, opts ...PokemonOption) *Pokemon {
	p := &Pokemon{
		Name: name,
	}

	for _, opt := range opts {
		opt(p)
	}

	return p
}
