package main

import (
	"fmt"
	"math/rand"
	"time"
)

// ambushPokemon: un entrenador es simulado por una goroutine que espera a que un pokemon aparezca en un canal.
// Una vez que el pokemon aparece, el entrenador imprime un mensaje.
// La función principal espera un tiempo aleatorio y señala en el canal que un pokemon ha aparecido.
func ambushPokemon() {
	ch := make(chan Pokemon)

	go func() {
		// la goroutine espera a que un pokemon aparezca en el canal
		pokemon := <-ch
		fmt.Printf("¡Pokemon capturado en goroutine: es un %s!\n", pokemon.Name)
	}()

	// esperamos un tiempo aleatorio para simular la aparición de un pokemon
	time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)

	// señalamos en el canal que un pokemon ha aparecido
	ch <- PokemonAppears()
	fmt.Println("hilo principal : señalamos en el canal que un pokemon ha aparecido")

	// esperamos un segundo para que la goroutine pueda capturar el pokemon
	time.Sleep(time.Second)
}
