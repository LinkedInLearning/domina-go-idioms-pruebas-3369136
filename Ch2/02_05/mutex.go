package main

import (
	"fmt"
)

func mutexCounter() {
	counter := Counter{}

	for i := 0; i < 10; i++ {
		go func() {
			// Incrementamos el contador dentro de una gorotuine
			counter.Increment()
		}()
	}

	// al terminar el bucle for, obtenemos el valor del contador.
	// Es muy posible que no todas las gorotuinas hayan terminado,
	// por lo que el valor será diferente cada vez.
	fmt.Printf("Hemos capturado %d pokemons\n", counter.GetValue())
}
