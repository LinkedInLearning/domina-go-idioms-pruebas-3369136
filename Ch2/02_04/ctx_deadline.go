package main

import (
	"context"
	"fmt"
	"time"
)

func contextDeadline() {
	// ahora vamos a hacer una petición con un contexto
	// que expira pasados 1 milisegundo
	deadline := time.Now().Add(1 * time.Millisecond)
	ctx, cancel := context.WithDeadline(context.Background(), deadline)
	// aseguramos que se llame a cancel al final de la ejecución
	// y lo hacemos en un defer para que se ejecute siempre al final
	// de la ejecución de la función main
	defer cancel()

	// el mensaje de Request completed nunca se imprimirá,
	// con un mensaje de error "context deadline exceeded".
	err := httpCall(ctx)
	if err != nil {
		fmt.Println("Context with Deadline: Error making request:", err)
		return
	}

	fmt.Println("Context with Deadline: Success!")
}
