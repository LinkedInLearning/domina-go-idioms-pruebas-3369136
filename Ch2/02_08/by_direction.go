package main

func sendChannel() {
	// un channel únicamente de envío
	sendChannel := make(<-chan int)

	// el channel enviará un valor desde sendChannel a algún otro receptor
	<-sendChannel
	// sendChannel <- 1 // operación no permitida, por tanto no compilará
}

func receiveChannel() {
	// un channel únicamente de recepción
	receiveChannel := make(chan<- int)

	// el channel recibirá un valor desde algún otro emisor a receiveChannel
	receiveChannel <- 1
	// <-receiveChannel // operación no permitida, por tanto no compilará
}

func receiveSendChannel() {
	// un channel tanto de envío como de recepción
	receiveSendChannel := make(chan int)

	// el channel podrá tanto enviar como recibir un valor desde algún otro emisor o receptor

	receiveSendChannel <- 1
	<-receiveSendChannel
}
