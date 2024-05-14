package main

import (
	"fmt"
	"time"
)

// teamRocket: simula un ejercicio del team rocket, en el que Giovanni, líder del team rocket,
// prepara muchos pokemon para que el team rocket los capture, pero serán muchos más
// de los que son capaces de capturar. Por tanto se les escaparán bastantes.
func teamRocket() {
	// numero de bolas pokemon del team rocket, utilizadas para capturar pokemon.
	// Representa el tamaño del canal en el que se comunicarán la aparición
	// y captura de pokemons.
	const pokeballs = 10
	ch := make(chan Pokemon, pokeballs)

	go func() {
		count := 0
		// el team rocket se queda a la espera de que aparezca un pokemon
		for pokemon := range ch {
			count++
			fmt.Printf("team rocket: Hemos capturado al %s!\n", pokemon.Name)
		}

		fmt.Printf("team rocket: Hemos capturado un total de %d pokemons\n", count)
	}()

	// 200 es un número suficientemente grande en mi máquina para causar
	// que la goroutine que representa al team rocket no sea capaz de realizar el trabajo
	// con el tamaño definido para el canal (el número de pokeballs).
	const pokemons = 200
	escaped := 0
	for p := 1; p <= pokemons; p++ {
		select {
		case ch <- PokemonAppears():
			fmt.Printf("Giovanni: ha aparecido un pokemon, van %d\n", p)
		default:
			fmt.Printf("Giovanni: el pokemon %d ha escapado :(\n", p)
			escaped++
		}
	}

	// cerramos el canal para que las goroutines puedan terminar
	close(ch)
	fmt.Printf("Giovanni: entrenamiento terminado. Se han escapado %d pokemons\n", escaped)

	// esperamos un segundo para que el team rocket pueda capturar los pokemons
	time.Sleep(time.Second)
}
