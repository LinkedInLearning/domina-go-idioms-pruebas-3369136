package main

import (
	"fmt"
	"sync"
)

type MutexCounter struct {
	value int
	mutex sync.Mutex
}

func (c *MutexCounter) Increment() {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.value++
}

func (c *MutexCounter) GetValue() int {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	return c.value
}

func mutexCounter() {
	counter := MutexCounter{}

	for i := 0; i < 10; i++ {
		go func() {
			counter.Increment()
		}()
	}

	fmt.Printf("Hemos capturado %d pokemons\n", counter.GetValue())
}
