package cmd

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
)

type Artist struct {
	Id           int      `json:"id"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Locations    string   `json:"locations"`
	ConcertDates string   `json:"concertDates"`
	Relations    string   `json:"relations"`
}

type Location struct {
}

type Dates struct {
}

type Relation struct {
}

func Index(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/index.html")
	if err != nil {
		fmt.Println(err)
	}
	t.Execute(w, r)
}

func Api(w http.ResponseWriter, r *http.Request) {
	get, err := http.Get("https://groupietrackers.herokuapp.com/api/artists/1")
	if err != nil {
		fmt.Println(err)
	}
	data, err := ioutil.ReadAll(get.Body)
	var artist Artist
	err = json.Unmarshal(data, &artist)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("name : ", artist.Name)
	fmt.Println("Creation : ", artist.CreationDate)
	fmt.Println("members : ", artist.Members)
}
