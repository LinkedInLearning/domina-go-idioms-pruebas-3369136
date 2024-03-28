package api

import (
	"fmt"
	"net/http"

	"github.com/linkedinlearning/domina-go/standard-layout/internal/player"
	"github.com/linkedinlearning/domina-go/standard-layout/pkg/pokemon"
	"github.com/linkedinlearning/domina-go/standard-layout/pkg/trainer"
)

var mainPlayer = player.Player{
	Player:   trainer.Ash(),
	Pokemons: []pokemon.Pokemon{pokemon.Pikachu()},
}

func Welcome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Bienvenido %s a la captura de pokemons! Empezarás con %d pokemons", mainPlayer.Name(), len(mainPlayer.Pokemons))
}
