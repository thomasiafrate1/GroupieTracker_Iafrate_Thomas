package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

type Character struct {
	Name            string   `json:"name"`
	Films           []string `json:"films"`
	ShortFilms      []string `json:"shortFilms"`
	TvShows         []string `json:"tvShows"`
	VideoGames      []string `json:"videoGames"`
	ParkAttractions []string `json:"parkAttractions"`
	Allies          []string `json:"allies"`
	Enemies         []string `json:"enemies"`
	ID              int      `json:"_id"`
	ImageURL        string   `json:"imageUrl"`
	URL             string   `json:"url"`
}

func getCharacter(id int) (*Character, error) {
	url := fmt.Sprintf("https://api.disneyapi.dev/characters/%d", id)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var character Character
	err = json.NewDecoder(resp.Body).Decode(&character)
	if err != nil {
		return nil, err
	}

	return &character, nil
}

func main() {
	tmpl := template.Must(template.ParseFiles("pages/personnages.html"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			err := r.ParseForm()
			if err != nil {
				http.Error(w, "Bad Request", http.StatusBadRequest)
				return
			}

			idStr := r.FormValue("id")
			id, err := strconv.Atoi(idStr)
			if err != nil {
				http.Error(w, "Invalid ID", http.StatusBadRequest)
				return
			}

			character, err := getCharacter(id)
			if err != nil {
				http.Error(w, "Character not found", http.StatusNotFound)
				return
			}

			data := struct {
				Name       string
				ImageURL   string
				Films      []string
				VideoGames []string
				Allies     []string
				Enemies    []string
				TvShows    []string
			}{
				Name:       character.Name,
				ImageURL:   character.ImageURL,
				Films:      character.Films,
				VideoGames: character.VideoGames,
				Allies:     character.Allies,
				Enemies:    character.Enemies,
				TvShows:    character.TvShows,
			}

			err = tmpl.Execute(w, data)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		} else {
			tmpl.Execute(w, nil)
		}
	})

	http.ListenAndServe(":8080", nil)
}
