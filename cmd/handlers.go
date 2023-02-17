package cmd

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
)

// Faire structure ici pour image et id de l'artist

type ImageID struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Image string `json:"image"`
}

type ArtistInformation struct {
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Locations    string   `json:"locations"`
	ConcertDates string   `json:"concertDates"`
	Relations    string   `json:"relations"`
}

type DatesLocation struct {
	LocationAndDate []string `json:"dateLocation"`
}

type IndexLocationItem struct {
	Id            int `json:"id"`
	DatesLocation []DatesLocation
}

type IndexLocationAllItem struct {
	IndexLocationItem []IndexLocationItem
}

type Dates struct {
}

type Relation struct {
}

func Index(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		errorHandler(w, r, http.StatusNotFound)
	} else {
		t, err := template.ParseFiles("templates/index.html")
		/*data, _ := ioutil.ReadAll(r.Body)
		request := string(data)
		fmt.Println(request)*/
		if err != nil {
			fmt.Println(err)
		}
		t.Execute(w, r)
	}
}

func ArtistInfo(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/artistinfo" {
		errorHandler(w, r, http.StatusNotFound)
	} else {
		t, err := template.ParseFiles("templates/ArtistInfo")

		if err != nil {
			fmt.Println(err)
		}
		t.Execute(w, r)
	}
}

func Api(w http.ResponseWriter, r *http.Request) {
	get, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		fmt.Println(err)
	}
	data, err := ioutil.ReadAll(get.Body)
	var artist []ImageID
	err = json.Unmarshal(data, &artist)
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, values := range artist {
		json.NewEncoder(w).Encode(values)
		fmt.Println(values)
	}
}

func errorHandler(w http.ResponseWriter, r *http.Request, status int) {
	w.WriteHeader(status)
	if status == http.StatusNotFound {
		t, err := template.ParseFiles("./hangman-web/templates/404" + ".html")
		if err != nil {
			fmt.Println(err)
		} else {
			t.Execute(w, r)
		}
	}
}
