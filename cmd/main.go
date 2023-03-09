package cmd

import (
	"fmt"
	"net/http"
)

const port = ":3001"

func Runner() {
	server := http.NewServeMux()
	server.HandleFunc("/", IndexHandlers)
	server.HandleFunc("/categorie", CategorieArtist)
	server.HandleFunc("/artistinfo", ArtistInfo)
	server.HandleFunc("/api/genre", ApiGenre)
	server.HandleFunc("/post/searchbar", SearchBar)
	server.HandleFunc("/get/artistinfo", ArtistInfoGet)
	fs := http.FileServer(http.Dir("templates/assets"))
	server.Handle("/assets/", http.StripPrefix("/assets", fs))
	fmt.Println("(http://localhost:8080", port)
	err := http.ListenAndServe(port, server)
	if err != nil {
		fmt.Println("error :", err)
		return
	}
}
