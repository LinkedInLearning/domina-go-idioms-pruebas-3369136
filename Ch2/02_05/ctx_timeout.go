package main

import (
	"context"
	"fmt"
	"time"
)

func contextTimeout() {
	// ahora vamos a hacer una petición con un contexto
	// que expira pasados 1 milisegundo
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Millisecond)
	// aseguramos que se llame a cancel al final de la ejecución
	// y lo hacemos en un defer para que se ejecute siempre al final
	// de la ejecución de la función main
	defer cancel()

	// el mensaje de Request completed nunca se imprimirá,
	// con un mensaje de error "context deadline exceeded".
	err := httpCall(ctx)
	if err != nil {
		fmt.Println("Error making request:", err)
	}
}
