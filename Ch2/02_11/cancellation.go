package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

// quickCapture: simula una captura de un pokemon, en la que se define un tiempo
// máximo para la misma, o tiempo de escape. Si el entrenador tarda menos que este
// tiempo de escape, el pokemon será capturado. De lo contrario se escapará.
// El tiempo de escape controlará el contexto de cancelación, llamando a su función
// cancel si este tiempo es superado.
func quickCapture() {
	// tiempo en el que el pokemon escapará
	escapeTime := 150 * time.Millisecond
	ctx, cancel := context.WithTimeout(context.Background(), escapeTime)
	defer cancel()

	// el canal aceptará una única señal en forma de pokemon
	ch := make(chan Pokemon, 1)

	// intentaremos capturar al pokemon en esta goroutina
	go func() {
		// tiempo de batalla aleatorio, con un tiempo que podria ser superior
		// al de escape del pokemon
		time.Sleep(time.Duration(rand.Intn(200)) * time.Millisecond)
		// el pokemon aparece tras el tiempo de espera. Si es mayor que el tiempo
		// de escape del pokemon, en la cancelación del contexto, éste código no se
		// ejecutará
		ch <- PokemonAppears()
	}()

	// un bloque select bloquea hasta que alguno de sus casos pueden ser ejecutados,
	// en cuyo caso, los ejecuta.
	select {
	case p := <-ch:
		// el combate ha sido muy rápido y el pokemon ha sido capturado
		fmt.Printf("pokemon capturado a tiempo, es un %s!\n", p.Name)

	case <-ctx.Done():
		// el contexto ha sido cancelado porque ha ocurrido el timeout:
		// el entrenador ha tardado más que el pokemon en escapar
		fmt.Println("lo siento, pero el pokemon ha sido mucho más rápido que tú")
	}

	// esperamos un segundo para que el entrenador pueda capturar el pokemon
	time.Sleep(time.Second)
}
