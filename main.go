package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"strings"
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
	Choix           int
}

type AllCharacter struct {
	Data []struct {
		Films           []any  `json:"films"`
		ShortFilms      []any  `json:"shortFilms"`
		TvShows         []any  `json:"tvShows"`
		VideoGames      []any  `json:"videoGames"`
		ParkAttractions []any  `json:"parkAttractions"`
		Allies          []any  `json:"allies"`
		Enemies         []any  `json:"enemies"`
		ID              int    `json:"_id"`
		Name            string `json:"name"`
		ImageURL        string `json:"imageUrl"`
		URL             string `json:"url"`
	} `json:"data"`
	Count      int    `json:"count"`
	TotalPages int    `json:"totalPages"`
	NextPage   string `json:"nextPage"`
	Choix      int
}

func getCharacter(id int) Character {
	url := fmt.Sprintf("https://api.disneyapi.dev/characters/%d", id)

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	data, Rerr := ioutil.ReadAll(resp.Body)
	if Rerr != nil {
		fmt.Println(Rerr)
	}
	fmt.Println(data)
	character := Character{}
	err = json.Unmarshal(data, &character)
	if err != nil {
		fmt.Println(err)
	}

	if len(character.Films) > 2 {
		character.Films = character.Films[:2]
	}

	return character
}

func fullCharacter(page int) AllCharacter {
	url := fmt.Sprintf("https://api.disneyapi.dev/characters?page=%d", page)

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	data, Rerr := ioutil.ReadAll(resp.Body)
	if Rerr != nil {
		fmt.Println(Rerr)
	}

	character := AllCharacter{}
	err = json.Unmarshal(data, &character)
	if err != nil {
		fmt.Println(err)
	}

	return character
}

func main() {
	verif_id := false
	static := http.FileServer(http.Dir("css"))
	http.Handle("/css/", http.StripPrefix("/css/", static))
	tmpl := template.Must(template.ParseFiles("pages/personnages.html"))
	page := 1

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			err := r.ParseForm()
			if err != nil {
				http.Error(w, "Bad Request", http.StatusBadRequest)
				return
			}

			for page = 1; !verif_id; page++ {
				Name := r.FormValue("Name")
				fmt.Println(Name)
				var id int
				test := AllCharacter{}
				test = fullCharacter(page)
				for i := 0; i < len(test.Data); i++ {
					if strings.Compare(strings.ToLower(Name), strings.ToLower(test.Data[i].Name)) == 0 {
						id = test.Data[i].ID
						verif_id = true
					}
				}
				if verif_id {
					data := Character{}
					data = getCharacter(id)
					data.Choix = 2
					err = tmpl.Execute(w, data)
					if err != nil {
						fmt.Println(err)
						return
					}
					break
				}
			}
		} else {
			test := AllCharacter{}
			test = fullCharacter(page)
			test.Choix = 1
			tmpl.Execute(w, test)
		}
	})

	http.ListenAndServe(":8888", nil)
}
