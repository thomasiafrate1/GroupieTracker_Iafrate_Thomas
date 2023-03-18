package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"math/rand"
	"net/http"
	"sort"
	"strings"
	"time"
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

func trierNoms(noms []string) []string {
	sort.Strings(noms)
	return noms
}

func trierNomsDecroissant(noms []string) []string {
	sort.Slice(noms, func(i, j int) bool {
		return noms[i] > noms[j]
	})
	return noms
}

func melangerNoms(noms []string) []string {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(noms), func(i, j int) {
		noms[i], noms[j] = noms[j], noms[i]
	})
	return noms
}

func listeNomsHandler(w http.ResponseWriter, r *http.Request) {
	noms := []string{
		"Mickey Mouse",
		"Minnie Mouse",
		"Donald Duck",
		"Daisy Duck",
		"Goofy",
		"Pluto",
		"Chip and Dale",
		"Goofy",
		"Winnie the Pooh",
		"Tigger",
		"Piglet",
		"Eeyore",
		"Kanga and Roo",
		"Pumbaa",
		"Timon",
		"Simba",
		"Scar",
		"Rafiki",
		"Mufasa",
		"Nala",
		"Zazu",
		"Aladdin",
		"Jasmine",
		"Genie",
		"Abu",
		"Jafar",
		"Tiana",
		"Naveen",
		"Louis",
		"Merida",
		"Brave",
		"Fergus",
		"Elinor",
		"Anna",
		"Elsa",
		"Olaf",
		"Kristoff",
		"Sven",
		"Hans",
		"Ariel",
		"Sebastian",
		"Flounder",
		"Ursula",
		"Belle",
		"Beast",
		"Lumiere",
		"Big Ben",
		"Mrs. Potts",
		"Chip",
		"Hercules",
		"Megara",
		"Philoctetes",
		"Hades",
		"Mulan",
		"Mushu",
		"Shan Yu",
		"Remy",
		"Linguini",
		"Colette",
		"Lilo",
		"Stitch",
		"Jumba",
		"Pleakley",
		"Peter Pan",
		"Wendy",
		"Captain Hook",
		"Tinker Bell",
		"Robin Hood",
		"Little John",
		"Rapunzel",
		"Flynn Rider",
		"Maximus",
		"Pascal",
		"Woody",
		"Buzz Lightyear",
		"Jessie",
		"Mr. Potato Head",
		"Rex",
		"Hamm",
		"Sulley",
		"Mike Wazowski",
		"Boo",
		"Cinderella",
		"Prince Charming",
		"Anastasia and Drizella",
		"Fairy Godmother",
		"Lady",
		"Tinker Bell",
		"Robin Hood",
		"Bernard and Bianca",
		"Sleeping Beauty",
		"Maleficent",
		"The Three Little Pigs",
		"Snow White and the Seven Dwarfs",
		"Cruella De Vil",
		"The 101 Dalmatians",
		"Mary Poppins",
		"Baloo",
		"King Louie",
		"Bagheera",
		"Moana",
		"Maui",
		"Pocahontas",
		"John Smith",
		"Meeko",
		"Flit",
		"Percy",
		"Esmeralda",
		"Quasimodo",
		"Judge Claude Frollo",
		"Phoebus",
		"Victor, Hugo, and Laverne",
		"Baymax",
		"Hiro Hamada",
		"Tadashi Hamada",
		"Gogo Tomago",
		"Honey Lemon",
		"Wasabi",
		"Fred",
		"Elsa Van Helsing",
		"Hercules",
		"Zeus",
		"Hera",
		"Hades",
		"Megara",
		"Tarzan",
		"Jane Porter",
		"Clayton",
		"Kuzco",
		"Yzma",
		"Kronk",
		"Milo Thatch",
		"Kida Nedakh",
		"Captain Amelia",
		"Jim Hawkins",
		"Long John Silver",
		"Vanellope von Schweetz",
		"Fix-It Felix Jr.",
		"Sergeant Calhoun",
		"Shank",
		"Jack Skellington",
		"Sally",
		"Oogie Boogie",
		"Lock, Shock, and Barrel",
		"Jack-Jack Parr",
		"Frozone",
		"Syndrome",
		"Edna Mode",
		"Violet Parr",
		"Dash Parr",
		"Mrs. Incredible",
		"Mr. Incredible",
		"Duke Caboom",
		"Forky",
		"Ducky and Bunny",
		"Bo Peep",
		"Giselle",
		"Robert Philip",
		"Nancy Tremaine",
		"King Stefan",
		"Diablo",
		"Prince Phillip",
		"The Evil Queen",
		"Robin Hood",
		"Prince John",
		"Sheriff of Nottingham",
		"Taran",
		"Eilonwy",
		"Gurgi",
		"The Horned King",
		"Basil of Baker Street",
		"Professor Ratigan",
		"Olivia Flaversham",
		"Flik",
		"Princess Atta",
		"Hopper",
		"Dot",
		"Francis",
		"Heimlich",
		"Tuck and Roll",
		"Lightning McQueen",
		"Mater",
		"Sally Carrera",
		"Doc Hudson",
		"Luigi and Guido",
		"Sheriff",
		"Finn McMissile",
		"Holley Shiftwell",
		"Dory",
		"Marlin",
		"Crush",
		"Squirt",
		"Nemo",
		"Bruce",
		"Darla",
		"Gill",
		"Hank",
		"Joy",
		"Sadness",
		"Bing Bong",
	}
	if r.Method == "POST" {
		tri := r.FormValue("tri")
		switch tri {
		case "az":
			// Tri de la liste de noms en ordre décroissant
			sort.Strings(noms)
		case "za":
			// Tri de la liste de noms en ordre décroissant
			sort.Sort(sort.Reverse(sort.StringSlice(noms)))
		case "random":
			noms = melangerNoms(noms)
			// Ajouter d'autres cas pour les autres types de tri
		}
	}
	// Génération du code HTML
	tmpl := template.Must(template.ParseFiles("pages/doc.html"))
	data := struct{ Noms []string }{noms}
	tmpl.Execute(w, data)
}
func main() {
	static := http.FileServer(http.Dir("css"))
	http.Handle("/css/", http.StripPrefix("/css/", static))
	tmpl := template.Must(template.ParseFiles("index.html"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			err := r.ParseForm()
			if err != nil {
				http.Error(w, "Bad Request", http.StatusBadRequest)
				return
			}

			Name := r.FormValue("Name")
			fmt.Println(Name)

			var id int
			var verif_id bool
			page := 1
			for !verif_id {
				test := fullCharacter(page)
				for i := 0; i < len(test.Data); i++ {
					if strings.EqualFold(Name, test.Data[i].Name) {
						id = test.Data[i].ID
						verif_id = true
						break
					}
				}
				if verif_id {
					data := getCharacter(id)
					data.Choix = 2
					err = tmpl.Execute(w, data)
					if err != nil {
						fmt.Println(err)
					}
					break
				} else if test.NextPage == "" {
					http.Error(w, "Character not found", http.StatusNotFound)
					break
				} else {
					page++
				}
			}
		} else {
			test := fullCharacter(1)
			test.Choix = 1
			err := tmpl.Execute(w, test)
			if err != nil {
				fmt.Println(err)
			}
		}
	})
	tmpl2 := template.Must(template.ParseFiles("pages/personnages.html"))
	http.HandleFunc("/personnages", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			err := r.ParseForm()
			if err != nil {
				http.Error(w, "Bad Request", http.StatusBadRequest)
				return
			}

			Name := r.FormValue("Name")
			fmt.Println(Name)

			var id int
			var verif_id bool
			page := 1
			for !verif_id {
				test := fullCharacter(page)
				for i := 0; i < len(test.Data); i++ {
					if strings.EqualFold(Name, test.Data[i].Name) {
						id = test.Data[i].ID
						verif_id = true
						break
					}
				}
				if verif_id {
					data := getCharacter(id)
					data.Choix = 2
					err = tmpl2.Execute(w, data)
					if err != nil {
						fmt.Println(err)
					}
					break
				} else if test.NextPage == "" {
					http.Error(w, "Character not found", http.StatusNotFound)
					break
				} else {
					page++
				}
			}
		} else {
			test := fullCharacter(1)
			test.Choix = 1
			err := tmpl2.Execute(w, test)
			if err != nil {
				fmt.Println(err)
			}
		}
	})

	http.HandleFunc("/doc", func(w http.ResponseWriter, r *http.Request) {
		listeNomsHandler(w, r)
	})

	http.ListenAndServe(":8888", nil)
}
