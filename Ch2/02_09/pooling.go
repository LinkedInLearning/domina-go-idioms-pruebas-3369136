package main

import (
	"fmt"
	"runtime"
	"time"
)

// pokemonIncursion: simula una incursión, en la que unos cuantos entrenadores intentan capturar pokemons.
// Cada entrenador es una goroutine que intenta capturar un pokemon y señala en un canal que ha capturado un pokemon.
// Una vez que todos los pokemons han sido capturados, la incursión termina.
func pokemonIncursion() {
	ch := make(chan Pokemon)

	// creamos un pool de entrenadores, una por cada CPU
	g := runtime.GOMAXPROCS(0)
	fmt.Printf("gimnasio: se han creado %d entrenadores\n", g)

	for t := 0; t < g; t++ {
		go func(trainer int) {
			count := 0
			// la goroutine espera a que un pokemon aparezca en el canal
			for pokemon := range ch {
				fmt.Printf("entrenador [%d] ha capturado un pokemon. Es un %s!\n", trainer, pokemon.Name)
				count++
			}
			fmt.Printf("entrenador [%d] se retira con %d pokemons\n", trainer, count)
		}(t)
	}

	// En el hilo principal aparecerán 100 pokemons,
	// que serán enviados al canal de comunicación.
	const work = 100
	for w := 1; w <= work; w++ {
		// señalamos en el canal que un pokemon ha aparecido
		ch <- PokemonAppears()
		fmt.Printf("gimnasio: un pokemon ha aparecido. Van %d.\n", w)
	}

	// cerramos el canal para que las goroutines puedan terminar
	close(ch)
	fmt.Println("gimnasio: termina la incursión")

	// esperamos un segundo para que las goroutines puedan capturar los pokemons
	time.Sleep(time.Second)
}
