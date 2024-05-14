package main

import (
	"context"
	"fmt"
)

func contextCancellable() {
	// ahora vamos a hacer una petición con un contexto que se cancela
	ctx, cancel := context.WithCancel(context.Background())
	cancel() // forzamos la cancelación del contexto para simular un error

	// el mensaje de Request completed nunca se imprimirá,
	// con un mensaje de error "context canceled".
	err := httpCall(ctx)
	if err != nil {
		fmt.Println("Context cancellable: Error making request:", err)
	}
}
