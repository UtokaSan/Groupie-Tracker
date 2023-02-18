package cmd

import (
	"fmt"
	"net/http"
)

const port = ":3001"

func Runner() {
	server := http.NewServeMux()
	server.HandleFunc("/", Index)
	server.HandleFunc("/artistinfo", ArtistInfo)
	server.HandleFunc("/api", Api)
	fs := http.FileServer(http.Dir("templates/assets/"))
	server.Handle("/assets/", http.StripPrefix("/assets", fs))
	fmt.Println("(http://localhost:8080) on port ", port)
	err := http.ListenAndServe(port, server)
	if err != nil {
		fmt.Println("error :", err)
		return
	}
}
