package main

import (
	"log"

	"github.com/linkedinlearning/domina-go/package-access/internal"
)

func main() {
	log.Println("Bienvenido a la captura de pokemons!")
	internal.CapturePokemon()
	log.Println("Bye!")
}
