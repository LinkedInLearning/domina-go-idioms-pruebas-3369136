package trainer

type Trainer struct {
	Name string
}

func Ash() Trainer {
	return Trainer{
		Name: "Ash Ketchum",
	}
}
