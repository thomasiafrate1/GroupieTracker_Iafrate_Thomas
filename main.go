package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "test/index.html")
	})
	// Spécifiez le répertoire contenant les fichiers statiques (CSS, JS, images, etc.)
	staticDir := "test/css/"

	// Créez un routeur HTTP
	router := http.NewServeMux()

	// Servez les fichiers statiques en utilisant http.FileServer
	router.Handle("/test/", http.StripPrefix("/test/", http.FileServer(http.Dir(staticDir))))

	http.ListenAndServe(":8080", nil)
}
