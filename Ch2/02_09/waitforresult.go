package main

import (
	"fmt"
	"math/rand"
	"time"
)

// waitForPokemon: la ejecución de la función principal se bloquea hasta que se recibe un valor en el canal.
// En este caso, el canal es un canal sin buffer, por lo que la ejecución se bloqueará hasta que se envíe un valor al canal.
// Una vez que se envía un valor al canal, la ejecución se reanuda y se imprime el valor recibido.
// En este caso, el valor recibido es "¡pokemon capturado!".
func waitForPokemon() {
	ch := make(chan Pokemon)

	go func() {
		time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
		ch <- PokemonAppears()
		fmt.Println("goroutine : señalamos en el canal que hemos capturado un pokemon")
	}()

	pokemon := <-ch
	fmt.Printf("hilo principal : señal de pokemon capturado recibida, es un %s!\n", pokemon.Name)

	// esperamos un segundo para que la goroutine pueda producir el pokemon
	time.Sleep(time.Second)
}
