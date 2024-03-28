package pokemon

type Pokemon struct {
	Name string `json:"name"`
}

func Pikachu() Pokemon {
	return Pokemon{
		Name: "Pikachu",
	}
}
