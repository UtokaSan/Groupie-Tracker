package cmd

import (
	"fmt"
	"net/http"
	"path/filepath"
)

const port = ":3001"

func Runner() {
	server := http.NewServeMux()
	server.HandleFunc("/", IndexHandlers)
	server.HandleFunc("/genre", IndexHandlers)
	server.HandleFunc("/artistinfo", ArtistInfo)
	server.HandleFunc("/api", Api)
	fs := http.FileServer(http.Dir("templates/assets"))
	server.Handle("/assets/", http.StripPrefix("/assets", fs))
	absPath, err1 := filepath.Abs("templates/assets")
	if err1 != nil {
		panic(err1)
	}
	fmt.Printf("Le chemin absolu du dossier est : %s\n", absPath)
	fmt.Println("(http://localhost:8080) on port ", port)
	err := http.ListenAndServe(port, server)
	if err != nil {
		fmt.Println("error :", err)
		return
	}
}
