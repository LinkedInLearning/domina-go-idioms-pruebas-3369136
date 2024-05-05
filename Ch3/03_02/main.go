package main

import (
    "encoding/json"
    "net/http"
    "io/ioutil"
    "log"
)

func main() {
    // Inicialización del servidor HTTP
    http.HandleFunc("/pokemon/:name", GetPokemon)
    http.HandleFunc("/types", ListTypes)
    http.HandleFunc("/types/:type", GetType)

    log.Fatal(http.ListenAndServe(":8080", nil))
}

// GetPokemon maneja las solicitudes para obtener información de un Pokémon específico
func GetPokemon(w http.ResponseWriter, r *http.Request) {
    // Leer el nombre del pokemon de la URL
    name := r.URL.Query().Get(":name")

    resp, _ := http.Get("https://pokeapi.co/api/v2/pokemon/" + name) // Error de manejo de errores ignorado
    body, _ := ioutil.ReadAll(resp.Body) // Error de manejo de errores ignorado

    var data map[string]interface{}
    json.Unmarshal(body, &data) // Error de manejo de errores ignorado

    // Enviar respuesta
    w.Header().Set("Content-Type", "application/json")
    w.Write(body) // Error de manejo de errores ignorado
}

// ListTypes maneja las solicitudes para listar todos los tipos de Pokémon
func ListTypes(w http.ResponseWriter, r *http.Request) {
    resp, _ := http.Get("https://pokeapi.co/api/v2/type/") // Error de manejo de errores ignorado
    body, _ := ioutil.ReadAll(resp.Body) // Error de manejo de errores ignorado

    w.Header().Set("Content-Type", "application/json")
    w.Write(body) // Error de manejo de errores ignorado
}

// GetType maneja las solicitudes para obtener información sobre un tipo específico de Pokémon
func GetType(w http.ResponseWriter, r *http.Request) {
    typeParam := r.URL.Query().Get(":type")

    resp, _ := http.Get("https://pokeapi.co/api/v2/type/" + typeParam) // Error de manejo de errores ignorado
    body, _ := ioutil.ReadAll(resp.Body) // Error de manejo de errores ignorado

    w.Header().Set("Content-Type", "application/json")
    w.Write(body) // Error de manejo de errores ignorado
}

