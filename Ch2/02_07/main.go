package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	contextBackground()

	contextCancellable()

	contextTimeout()

	contextDeadline()

	contextWithValues()
}

func contextBackground() {
	err := httpCall(context.Background())
	if err != nil {
		fmt.Println("Error making request:", err)
	}
}

func contextCancellable() {
	// ahora vamos a hacer una petición con un contexto que se cancela
	ctx, cancel := context.WithCancel(context.Background())
	cancel() // forzamos la cancelación del contexto para simular un error

	// el mensaje de Request completed nunca se imprimirá,
	// con un mensaje de error "context canceled".
	err := httpCall(ctx)
	if err != nil {
		fmt.Println("Error making request:", err)
	}
}

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
		fmt.Println("Error making request:", err)
	}
}

func contextWithValues() {
	// ahora vamos a hacer una petición con un contexto
	// que expira pasados 10 segundos
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	// aseguramos que se llame a cancel al final de la ejecución
	// y lo hacemos en un defer para que se ejecute siempre al final
	// de la ejecución de la función main
	defer cancel()

	// vamos a añadir un valor al contexto
	ctx = context.WithValue(ctx, currentPokemonKey, "pikachu")

	err := httpCallWithContextValue(ctx)
	if err != nil {
		fmt.Println("Error making request:", err)
	}
}
