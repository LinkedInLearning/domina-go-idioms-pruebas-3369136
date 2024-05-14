package main

import (
	"fmt"
	"sync"
)

var arceus Pokemon
var once sync.Once

func createFirstPokemon() {
	once.Do(func() {
		arceus = Pokemon{
			Name:    "Arceus",
			Attack:  120,
			Defense: 120,
			Life:    120,
			Types:   []string{"Normal"},
			Moves:   []string{"Hyper Beam"},
		}

		fmt.Println("Arceus ha sido creado!", arceus)
	})
}
