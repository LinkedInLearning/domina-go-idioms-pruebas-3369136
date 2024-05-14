package main

import (
	"fmt"
	"sync"
)

func waitGroupCounter() {
	counter := Counter{}

	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		// añadimos 1 goroutine al WaitGroup, antes de lanzarla
		wg.Add(1)
		go func() {
			// al finalizar la goroutine, quitamos 1 del WaitGroup
			defer wg.Done()
			counter.Increment()
		}()
	}

	// esperamos a que todas las goroutines terminen
	wg.Wait()

	fmt.Printf("Hemos capturado %d pokemons!\n", counter.GetValue())
}
