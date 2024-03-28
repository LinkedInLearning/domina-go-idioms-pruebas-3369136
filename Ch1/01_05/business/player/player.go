package player

import (
	"github.com/linkedinlearning/domina-go/ardanlabs-layout/foundation/pokemon"
	"github.com/linkedinlearning/domina-go/ardanlabs-layout/foundation/trainer"
)

type Player struct {
	Player   trainer.Trainer   `json:"player"`
	Pokemons []pokemon.Pokemon `json:"pokemons"`
}

func (p *Player) Name() string {
	return p.Player.Name
}
