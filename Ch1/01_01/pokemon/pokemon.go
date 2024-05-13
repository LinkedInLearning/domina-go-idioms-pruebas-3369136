package pokemon

// pokemonType es un tipo de dato que representa un tipo de pokemon.
// Lo estamos ocultando a propósito para que no sea accesible desde fuera del paquete,
// ya que es un detalle de implementación.
type pokemonType string

// Pokemon es una estructura de datos que representa un pokemon.
// Lo estamos exportando para que sea accesible desde fuera del paquete.
type Pokemon struct {
	Name    string
	Attack  int
	Defense int
	Life    int
	Types   []pokemonType
	Moves   []string
}

// Pikachu es un constructor de Pokemon que devuelve un Pikachu.
func Pikachu() Pokemon {
	return Pokemon{
		Name:    "Pikachu",
		Attack:  55,
		Defense: 40,
		Life:    35,
		Types:   []pokemonType{"Electric"},
		Moves:   []string{"Thunder Shock", "Quick Attack"},
	}
}

// Charmander es un constructor de Pokemon que devuelve un Charmander.
func Charmander() Pokemon {
	return Pokemon{
		Name:    "Charmander",
		Attack:  52,
		Defense: 43,
		Life:    39,
		Types:   []pokemonType{"Fire"},
		Moves:   []string{"Ember", "Scratch"},
	}
}

// Bulbasaur es un constructor de Pokemon que devuelve un Bulbasaur.
func Bulbasaur() Pokemon {
	return Pokemon{
		Name:    "Bulbasaur",
		Attack:  49,
		Defense: 49,
		Life:    45,
		Types:   []pokemonType{"Grass", "Poison"},
		Moves:   []string{"Vine Whip", "Tackle"},
	}
}

// Squirtle es un constructor de Pokemon que devuelve un Squirtle.
func Squirtle() Pokemon {
	return Pokemon{
		Name:    "Squirtle",
		Attack:  48,
		Defense: 65,
		Life:    44,
		Types:   []pokemonType{"Water"},
		Moves:   []string{"Water Gun", "Tackle"},
	}
}
