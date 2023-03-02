package cmd

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func IndexHandlers(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		errorHandler(w, r, http.StatusNotFound)
	} else {
		t, err := template.ParseFiles("templates/index.html")
		if err != nil {
			fmt.Println(err)
		}
		t.Execute(w, r)
	}
}

func CategorieArtist(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/categorie" {
		errorHandler(w, r, http.StatusNotFound)
	} else {
		t, err := template.ParseFiles("templates/categorie.html")

		if err != nil {
			fmt.Println(err)
		}
		t.Execute(w, r)
	}
}

/*
func InformationArtist(w http.ResponseWriter, id string) {
	get, err := http.Get("https://groupietrackers.herokuapp.com/api/artists/" + id)
	if err != nil {
		fmt.Println(err)
	}
	resp, err := ioutil.ReadAll(get.Body)
	var artistInformation ArtistInformation
	err = json.Unmarshal(resp, &artistInformation)
	if err != nil {
		fmt.Println(err)
		return
	}
	json.NewEncoder(w).Encode(artistInformation)
	fmt.Println(artistInformation)
}
*/
func InformationArtistTag(name string) string {
	apikey := "471f5119aa12d32718ae05f982f745dc"
	escapeName := url.QueryEscape(name)
	get, err := http.Get("http://ws.audioscrobbler.com/2.0/?method=artist.getTopTags&artist=" + escapeName + "&api_key=" + apikey + "&format=json")
	if err != nil {
		fmt.Println(err)
	}
	resp, err := ioutil.ReadAll(get.Body)
	var artistTopTags AllTags
	err = json.Unmarshal(resp, &artistTopTags)
	if err != nil {
		fmt.Println(err)
	}
	if len(artistTopTags.Toptags.Tag) > 0 {
		artistGenre := strings.ToLower(artistTopTags.Toptags.Tag[0].Name)
		return artistGenre
	}
	return "bad"
}

/*
func InformationLocation(w http.ResponseWriter, id string) {
	get, err := http.Get("https://groupietrackers.herokuapp.com/api/locations/" + id)
	if err != nil {
		fmt.Println(err)
		return
	}
	resp, err := ioutil.ReadAll(get.Body)
	var artistLocations Location
	err = json.Unmarshal(resp, &artistLocations)
	if err != nil {
		fmt.Println(err)
		return
	}
	json.NewEncoder(w).Encode(artistLocations)
	fmt.Println(artistLocations)
}

func InformationDate(w http.ResponseWriter, id string) {
	get, err := http.Get("https://groupietrackers.herokuapp.com/api/dates/" + id)
	if err != nil {
		fmt.Println(err)
		return
	}
	resp, err := ioutil.ReadAll(get.Body)
	var artistDates Dates
	err = json.Unmarshal(resp, &artistDates)
	if err != nil {
		fmt.Println(err)
		return
	}
	json.NewEncoder(w).Encode(artistDates)
	fmt.Println(artistDates)
}
*/
func ApiGenre(w http.ResponseWriter, r *http.Request) {
	get, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		fmt.Println(err)
	}
	resp, err := ioutil.ReadAll(get.Body)
	data, _ := ioutil.ReadAll(r.Body)
	input := string(data)[7 : len(data)-2]
	fmt.Println(input)
	var artist []ImageID
	err = json.Unmarshal(resp, &artist)
	if err != nil {
		fmt.Println(err)
		return
	}
	var genreArtist []ImageID
	for _, values := range artist {
		genre := InformationArtistTag(values.Name)
		if strings.Contains(genre, input) {
			values.Genre = genre
			genreArtist = append(genreArtist, values)
		}
	}
	fmt.Println(genreArtist)
	json.NewEncoder(w).Encode(genreArtist)
}

func errorHandler(w http.ResponseWriter, r *http.Request, status int) {
	w.WriteHeader(status)
	if status == http.StatusNotFound {
		t, err := template.ParseFiles("templates/404" + ".html")
		if err != nil {
			fmt.Println(err)
		} else {
			t.Execute(w, r)
		}
	}
}
