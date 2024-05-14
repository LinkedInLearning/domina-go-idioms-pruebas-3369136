package main

import (
	"fmt"
	"runtime"
	"time"
)

// pokemonIncursion: simula una batalla de pokemons, en la que unos cuantos entrenadores intentan capturar pokemons.
// Cada entrenador es una goroutine que intenta capturar un pokemon y señala en un canal que ha capturado un pokemon.
// Una vez que todos los pokemons han sido capturados, la batalla termina.
func pokemonIncursion() {
	ch := make(chan Pokemon)

	// creamos un pool de entrenadores, una por cada CPU
	g := runtime.GOMAXPROCS(0)
	fmt.Printf("gimnasio: se han creado %d entrenadores\n", g)

	for t := 0; t < g; t++ {
		go func(trainer int) {
			// la goroutine espera a que un pokemon aparezca en el canal
			for pokemon := range ch {
				fmt.Printf("entrenador [%d] ha capturado un pokemon. Es un %s!\n", trainer, pokemon.Name)
			}
			fmt.Printf("entrenador [%d] se retira\n", trainer)
		}(t)
	}

	// van a aparecer 100 pokemons
	const work = 100
	for w := 1; w <= work; w++ {
		// señalamos en el canal que un pokemon ha aparecido
		ch <- PokemonAppears()
		fmt.Printf("gimnasio: un pokemon ha aparecido. Van %d.\n", w)
	}

	// cerramos el canal para que las goroutines puedan terminar
	close(ch)
	fmt.Println("gimnasio: termina la batalla de pokemons")

	// esperamos un segundo para que las goroutines puedan capturar los pokemons
	time.Sleep(time.Second)
}
