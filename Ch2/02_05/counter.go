package main

import "sync"

type Counter struct {
	value int
	mutex sync.Mutex
}

// Primero bloqueamos el mutex, incrementamos el valor
// y luego desbloqueamos el mutex en un defer para asegurarnos
// de que se desbloquee siempre.
func (c *Counter) Increment() {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.value++
}

// Al leer el valor, bloqueamos el mutex, leemos el valor
// y luego desbloqueamos el mutex en un defer para asegurarnos
// de que se desbloquee siempre.
func (c *Counter) GetValue() int {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	return c.value
}
