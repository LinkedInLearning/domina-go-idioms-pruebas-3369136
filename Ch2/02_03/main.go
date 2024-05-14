package main

import "fmt"

func main() {
	pointerVariables()

	structVariables()
}

func structVariables() {
	// este pokemon se inicializa usando semántica de puntero
	pikachu := NewPokemonPointer(
		"Pikachu",
		WithAttack(55), WithDefense(40), WithLife(35),
		WithTypes([]string{"Electric"}),
		WithMoves([]string{"Quick Attack", "Thunderbolt"}),
	)

	fmt.Println(pikachu) // pikachu se modifica, puesto que se pasó un puntero

	// este pokemon se inicializa usando semántica de valor
	charmander := NewPokemonValue(
		"Charmander",
		WithAttackValue(52), WithDefenseValue(43), WithLifeValue(39),
		WithTypesValue([]string{"Fire"}),
		WithMovesValue([]string{"Ember", "Flamethrower"}),
	)

	fmt.Println(charmander) // charmander no se modifica, puesto que se pasó una copia
}

func pointerVariables() {
	// entero a almacenar
	var valor = 17

	// puntero apunta a la dirección de memoria de valor
	var puntero *int = &valor

	fmt.Println("valor almacenado:", valor)
	fmt.Println("dirección del valor entero:", &valor)
	fmt.Println("lo que el puntero almacena:", puntero)

	// actualiza el valor usando el operador de dereferenciación
	*puntero = 71

	fmt.Println("lo que el valorEntero almacena ahora:", valor) // cambia el valor
	fmt.Println("dirección del valor entero:", &valor)          // misma dirección de memoria
}
