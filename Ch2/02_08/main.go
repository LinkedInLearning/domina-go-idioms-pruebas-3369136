package main

func main() {
	mutexCounter()

	waitGroupCounter()

	for i := 0; i < 10; i++ {
		createFirstPokemon()
	}
}
