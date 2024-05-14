package main

import (
	"fmt"
	"math/rand"
	"time"
)

// pokemonGym: simula un gimnasio de pokemon donde varios entrenadores intentan capturar un pokemon.
// Cada entrenador es una goroutine que intenta capturar un pokemon y señala en un canal que ha capturado un pokemon.
// Una vez que todos los entrenadores han capturado un pokemon, el gimnasio imprime un mensaje.
func pokemonGym() {
	trainers := 100
	ch := make(chan Pokemon, trainers)

	for tr := 0; tr < trainers; tr++ {
		go func(trainer int) {
			time.Sleep(time.Duration(rand.Intn(200)) * time.Millisecond)
			ch <- PokemonAppears()
			fmt.Printf("entrenador [%d]: he capturado un pokemon!\n", trainer)
		}(tr)
	}

	// ahora esperamos a que todos los entrenadores hayan capturado un pokemon
	for trainers > 0 {
		// por cada entrenador, recibimos un pokemon del canal
		pokemon := <-ch
		// y decrementamos el número del entrenador en 1
		trainers--
		fmt.Printf("gimnasio: el entrenador [%d] ha capturado un %s!\n", trainers, pokemon.Name)
	}

	fmt.Println("gimnasio: todos los entrenadores han capturado un pokemon!")
}
