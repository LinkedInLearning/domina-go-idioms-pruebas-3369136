package main

type Pokemon struct {
	Name    string
	Attack  int
	Defense int
	Life    int
	Types   []string
	Moves   []string
}

// PokemonOption utiliza semántica de puntero para modificar la estructura Pokemon.
// Como veremos, ésto nos permite modificar la estructura original. Al finalizar
// la función, la estructura original se modifica.
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

func NewPokemon(name string, opts ...PokemonOption) *Pokemon {
	p := &Pokemon{
		Name: name,
	}

	for _, opt := range opts {
		opt(p)
	}

	return p
}
