package api

import (
	"fmt"
	"net/http"

	"github.com/linkedinlearning/domina-go/ardanlabs-layout/business/player"
	"github.com/linkedinlearning/domina-go/ardanlabs-layout/foundation/pokemon"
	"github.com/linkedinlearning/domina-go/ardanlabs-layout/foundation/trainer"
)

var mainPlayer = player.Player{
	Player:   trainer.Ash(),
	Pokemons: []pokemon.Pokemon{pokemon.Pikachu()},
}

func Welcome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Bienvenido %s a la captura de pokemons! Empezarás con %d pokemons: %v", mainPlayer.Name(), len(mainPlayer.Pokemons), mainPlayer.Pokemons)
}
