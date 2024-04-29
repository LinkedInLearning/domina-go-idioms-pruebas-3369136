package main

import (
	"fmt"
)

func mutexCounter() {
	counter := Counter{}

	for i := 0; i < 10; i++ {
		go func() {
			counter.Increment()
		}()
	}

	fmt.Printf("Hemos capturado %d pokemons\n", counter.GetValue())
}
