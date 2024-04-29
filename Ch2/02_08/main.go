package main

// Este ejemplo provoca de manera intencionada un deadlock, ya que no hay un receptor para el valor enviado.
// Sirva como ejemplo de cómo se definen los channels y cómo se utilizan.
func main() {}

func unbufferedChannel() {
	// un channel sin buffer,
	// por lo que bloqueará si no hay un receptor
	channel := make(chan int)

	// vamos a enviar un valor al channel

	// esto bloqueará hasta que haya un receptor
	channel <- 1
}

func bufferedChannel() {
	// un channel con buffer de tamaño 10,
	// que únicamente bloqueará si se llena.
	bufferedChannel := make(chan int, 10)

	// vamos a enviar un valor al channel

	// esto no bloqueará, ya que hay espacio en el buffer
	bufferedChannel <- 1
}

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
