package main

import "fmt"

func main() {
	go func() {
		fmt.Println("Hola, mundo!")
	}()
	fmt.Println("Adiós, mundo!")
}
