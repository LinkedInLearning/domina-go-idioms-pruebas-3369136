package trainer

type Trainer struct {
	Name string `json:"name"`
}

func Ash() Trainer {
	return Trainer{
		Name: "Ash Ketchum",
	}
}
