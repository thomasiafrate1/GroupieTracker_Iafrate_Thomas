package main

import (
	"log"
	"net/http"
)

func main() {
	// Créer un gestionnaire de fichiers statiques pour servir les fichiers du dossier "test"
	fs := http.FileServer(http.Dir("test"))

	// Définir les routes pour servir les fichiers statiques et les pages HTML
	http.Handle("/", fs)
	http.Handle("/pages/", fs)
	http.Handle("/css/", fs)

	// Démarrer le serveur web
	log.Println("Serveur démarré sur le port 8080")
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		log.Fatal("Erreur de démarrage du serveur: ", err)
	}
}
