package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

type Pokemon struct {
	ID              int      `json:"_id"`
	Films           []string `json:"films"`
	ShortFilms      []string `json:"shortFilms"`
	TvShows         []string `json:"tvShows"`
	VideoGames      []string `json:"videoGames"`
	ParkAttractions []string `json:"parkAttractions"`
	Allies          []any    `json:"allies"`
	Enemies         []any    `json:"enemies"`
	Name            string   `json:"name"`
	ImageURL        string   `json:"imageUrl"`
	URL             string   `json:"url"`
}

func idpokedex(pokemonid int) Pokemon {

	url := fmt.Sprintf("https://api.disneyapi.dev/characters/%d", pokemonid)

	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	var pokemon Pokemon
	err = json.NewDecoder(resp.Body).Decode(&pokemon)
	if err != nil {
		panic(err)
	}

	return pokemon

}

func main() {
	static := http.FileServer(http.Dir("assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", static))

	tmpl := template.Must(template.ParseFiles("index.html"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			err := r.ParseForm()
			if err != nil {
				http.Error(w, "Bad Request", http.StatusBadRequest)
				return
			}

			pokedexID := r.FormValue("pokedexID")
			id, err := strconv.Atoi(pokedexID)
			if err != nil {
				http.Error(w, "NUL", http.StatusBadRequest)
				return
			}

			pokemon := idpokedex(id)

			err = tmpl.Execute(w, pokemon)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		} else {
			tmpl.Execute(w, nil)
		}
	})

	http.ListenAndServe(":8888", nil)
}
