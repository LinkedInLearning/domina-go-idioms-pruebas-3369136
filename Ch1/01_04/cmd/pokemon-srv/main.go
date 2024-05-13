package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/linkedinlearning/domina-go/standard-layout/cmd/pokemon-srv/api"
)

var port int = 8080

func main() {
	err := run(port)
	if err != nil {
		log.Fatal(err)
	}
}

func run(port int) error {
	log.Printf("Starting Pokemon service at port %d...\n", port)

	http.HandleFunc("/pokemon", api.Welcome)

	return http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}
