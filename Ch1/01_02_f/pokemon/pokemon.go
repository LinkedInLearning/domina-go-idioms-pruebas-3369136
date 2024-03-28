package pokemon

type Pokemon struct {
	Name string
}

func NewPikachu() Pokemon {
	return Pokemon{
		Name: "Pikachu",
	}
}
