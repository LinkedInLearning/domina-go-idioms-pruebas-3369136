package main

import (
	"fmt"
)

func main() {
	go func() {
		fmt.Println("Hola, mundo!")
	}()

	go hello()

	fmt.Println("Adiós, mundo!")
}

func hello() {
	fmt.Println("Hola, gophers!")
}
