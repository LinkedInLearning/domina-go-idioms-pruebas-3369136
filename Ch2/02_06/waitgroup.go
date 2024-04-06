package main

import (
	"fmt"
	"sync"
)

type WaitGroupCounter struct {
	value int
	mutex sync.Mutex
}

func (c *WaitGroupCounter) Increment() {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.value++
}

func (c *WaitGroupCounter) GetValue() int {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	return c.value
}

func waitGroupCounter() {
	counter := WaitGroupCounter{}

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
