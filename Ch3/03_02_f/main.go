package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"regexp"
)

var (
	PokemonRegexWithName = regexp.MustCompile(`^/pokemon/([a-z0-9]+(?:-[a-z0-9]+)?)$`)
	TypesRegex           = regexp.MustCompile(`^/types/*$`)
	TypesRegexWithName   = regexp.MustCompile(`^/types/([a-z0-9]+(?:-[a-z0-9]+)?)$`)
)

func main() {
	mux := http.NewServeMux()

	mux.Handle("/pokemon/", &PokemonHandler{})

	th := &PokemonTypesHandler{}
	mux.Handle("/types", th)
	mux.Handle("/types/", th)

	log.Fatal(http.ListenAndServe(":8080", mux))
}

type PokemonHandler struct{}

func (h *PokemonHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println(r.Method, r.URL.Path)

	if r.Method != http.MethodGet || !PokemonRegexWithName.MatchString(r.URL.Path) {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Extract the resource ID using a regex
	matches := PokemonRegexWithName.FindStringSubmatch(r.URL.Path)
	// Expect matches to be length >= 2 (full string + 1 matching group)
	if len(matches) < 2 {
		http.Error(w, "Invalid URL", http.StatusBadRequest)
		return
	}

	name := matches[1]

	resp, err := http.Get("https://pokeapi.co/api/v2/pokemon/" + name)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var data map[string]interface{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

type PokemonTypesHandler struct{}

func (h *PokemonTypesHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println(r.Method, r.URL.Path)

	switch {
	case r.Method == http.MethodGet && TypesRegex.MatchString(r.URL.Path):
		h.ListTypes(w, r)
		return
	case r.Method == http.MethodGet && TypesRegexWithName.MatchString(r.URL.Path):
		matches := TypesRegexWithName.FindStringSubmatch(r.URL.Path)
		if len(matches) < 2 {
			http.Error(w, "Invalid URL", http.StatusBadRequest)
			return
		}

		h.GetType(w, r, matches[1])
		return
	}

	http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
}

// ListTypes maneja las solicitudes para listar todos los tipos de Pokémon
func (h *PokemonTypesHandler) ListTypes(w http.ResponseWriter, _ *http.Request) {
	resp, err := http.Get("https://pokeapi.co/api/v2/type/")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// GetType maneja las solicitudes para obtener información sobre un tipo específico de Pokémon
func (h *PokemonTypesHandler) GetType(w http.ResponseWriter, _ *http.Request, pokemonType string) {
	resp, err := http.Get("https://pokeapi.co/api/v2/type/" + pokemonType)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
