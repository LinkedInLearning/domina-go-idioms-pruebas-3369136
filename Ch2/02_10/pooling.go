package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

// pokemonIncursion: simula una batalla de pokemons, en la que unos cuantos entrenadores intentan capturar pokemons.
// Cada entrenador es una goroutine que intenta capturar un pokemon y señala en un canal que ha capturado un pokemon.
// Una vez que todos los pokemons han sido capturados, la battalla termina.
func pokemonIncursion() {
	ch := make(chan Pokemon)

	// creamos un pool de entrenadores, una por cada CPU
	g := runtime.GOMAXPROCS(0)
	fmt.Printf("gimnasio: se han creado %d entrenadores\n", g)
	for c := 0; c < g; c++ {
		go func(trainer int) {
			// la goroutine espera a que un pokemon aparezca en el canal
			for pokemon := range ch {
				fmt.Printf("entrenador [%d] ha capturado un pokemon. Es un %s!\n", trainer, pokemon.Name)
			}
			fmt.Printf("entrenador [%d] se retira\n", trainer)
		}(c)
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

// pokemonIncursionLimit: simulamos una batalla de pokemons, en la que unos cuantos entrenadores intentan capturar pokemons.
// La diferencia con pokemonIncursion es que en este caso, cada entrenador se retira una vez que ha capturado un pokemon.
// Si no hay pokemons que capturar, el entrenador se retira.
func pokemonIncursionLimit() {
	// todos los pokemons a capturar
	pokemons := []Pokemon{PokemonAppears(), PokemonAppears(), PokemonAppears(), PokemonAppears(), PokemonAppears()}

	// creamos un pool de entrenadores, una por cada CPU
	g := runtime.GOMAXPROCS(0)

	// el wait group nos va a permitir limitar el número de goroutines que pueden capturar pokemons
	var wg sync.WaitGroup

	// añadimos el número de goroutines al contador del WaitGroup
	wg.Add(g)

	ch := make(chan Pokemon, g)

	// para cada entrenador, creamos una goroutine que espera a que un pokemon aparezca en el canal
	for c := 0; c < g; c++ {
		go func(trainer int) {
			// cuando la goroutine termine, al retirarse, decrementamos el contador de WaitGroup
			defer wg.Done()

			// la goroutine espera a que un pokemon aparezca en el canal
			for pokemon := range ch {
				fmt.Printf("entrenador [%d] ha capturado un pokemon. Es un %s!\n", trainer, pokemon.Name)
			}
			fmt.Printf("entrenador [%d] se retira\n", trainer)
		}(c)
	}

	// señalamos en el canal que un pokemon ha aparecido
	for _, p := range pokemons {
		ch <- p
	}

	// cerramos el canal para que las goroutines puedan terminar
	close(ch)

	// esperamos a que todas las goroutines terminen
	wg.Wait()

	fmt.Println("gimnasio: termina la batalla de pokemons con límite de capturas")

	// esperamos un segundo para que las goroutines puedan capturar los pokemons
	time.Sleep(time.Second)
}
