package main

// Channels según el tipo de almacenamiento

func unbufferedChannel() {
	// un channel sin buffer que envia/recibe enteros,
	// por lo que bloqueará si no hay un receptor
	channel := make(chan int)

	// vamos a enviar un valor al channel

	// esto bloqueará hasta que haya un receptor
	channel <- 1
}

func bufferedChannel() {
	// un channel con buffer de tamaño 10, que envia/recibe enteros,
	// que únicamente bloqueará si se llena.
	bufferedChannel := make(chan int, 10)

	// vamos a enviar un valor al channel

	// esto no bloqueará, ya que hay espacio en el buffer
	bufferedChannel <- 1
}
