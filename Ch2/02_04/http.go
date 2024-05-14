package main

import (
	"context"
	"fmt"
	"net/http"
)

// es un idioma común en Go que las funciones que hacen llamadas a la red
// reciban un contexto siempre como primer argumento
func httpCall(ctx context.Context) error {
	httpCli := &http.Client{}
	req, err := http.NewRequest("GET", "https://www.google.com", nil)
	if err != nil {
		return fmt.Errorf("error creating request: %v", err)
	}

	res, err := httpCli.Do(req.WithContext(ctx))
	if err != nil {
		return fmt.Errorf("error making request: %v", err)
	}
	defer res.Body.Close()

	fmt.Println("Request completed:", res.Status)
	return nil
}

// definimos un tipo para el key del contexto
type key string

// definimos una key para el contexto
const currentPokemonKey key = key("current-pokemon")

// es un idioma común en Go que las funciones que hacen llamadas a la red
// reciban un contexto siempre como primer argumento
func httpCallWithContextValue(ctx context.Context) error {
	httpCli := &http.Client{}
	req, err := http.NewRequest("GET", "https://www.google.com", nil)
	if err != nil {
		return fmt.Errorf("error creating request: %v", err)
	}

	// recuperamos del contexto el valor que añadimos
	currentPokemon := ctx.Value(currentPokemonKey).(string)
	fmt.Println("Current pokemon is", currentPokemon)

	res, err := httpCli.Do(req.WithContext(ctx))
	if err != nil {
		return fmt.Errorf("error making request: %v", err)
	}
	defer res.Body.Close()

	fmt.Println("Request completed:", res.Status)
	return nil
}
