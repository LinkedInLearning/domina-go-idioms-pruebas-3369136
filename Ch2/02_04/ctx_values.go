package main

import (
	"context"
	"fmt"
	"time"
)

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
		fmt.Println("Context with values: Error making request:", err)
		return
	}

	fmt.Println("Context with Values: Success!")
}
