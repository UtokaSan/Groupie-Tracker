package cmd

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"sync"
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
func ArtistInfo(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/artistinfo" {
		errorHandler(w, r, http.StatusNotFound)
	} else {
		t, err := template.ParseFiles("templates/artistinfo.html")
		if err != nil {
			fmt.Println(err)
		}
		t.Execute(w, r)
	}
}

func InformationArtistTag(name string) string {
	apikey := "471f5119aa12d32718ae05f982f745dc"
	get, err := http.Get("http://ws.audioscrobbler.com/2.0/?method=artist.getTopTags&artist=" + url.QueryEscape(name) + "&api_key=" + apikey + "&format=json")
	if err != nil {
		fmt.Println(err)
	}
	defer get.Body.Close()
	resp, err := ioutil.ReadAll(get.Body)
	var artistTopTags AllTags
	err = json.Unmarshal(resp, &artistTopTags)
	if err != nil {
		fmt.Println(err)
	}
	if len(artistTopTags.Toptags.Tag) > 0 {
		return strings.ToLower(artistTopTags.Toptags.Tag[0].Name)
	}
	return "bad"
}

func ArtistInfoGet(w http.ResponseWriter, r *http.Request) {
	data, _ := ioutil.ReadAll(r.Body)
	input := string(data)[7 : len(data)-2]
	fmt.Println("test :", input)
	getArtists, err := http.Get("https://groupietrackers.herokuapp.com/api/artists/" + input)
	getLocation, err := http.Get("https://groupietrackers.herokuapp.com/api/locations/" + input)
	getDates, err := http.Get("https://groupietrackers.herokuapp.com/api/dates/" + input)
	var artist ImageID
	var location Location
	var dates Date
	respArtists, err := ioutil.ReadAll(getArtists.Body)
	respLocations, err := ioutil.ReadAll(getLocation.Body)
	respDates, err := ioutil.ReadAll(getDates.Body)
	err = json.Unmarshal(respArtists, &artist)
	err = json.Unmarshal(respLocations, &location)
	err = json.Unmarshal(respDates, &dates)
	fmt.Println("Structure requete :", artist)
	fmt.Println("Corps requete : ", string(respLocations))
	if err != nil {
		fmt.Println(err)
	}
	dataResult := ArtistInformation{
		Artist:   artist,
		Location: location,
		Dates:    dates,
	}
	fmt.Println(dates)
	json.NewEncoder(w).Encode(&dataResult)
}

func ApiGenre(w http.ResponseWriter, r *http.Request) {
	get, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		fmt.Println(err)
	}
	defer get.Body.Close()
	resp, err := ioutil.ReadAll(get.Body)
	if err != nil {
		fmt.Println(err)
	}
	data, _ := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
	}
	input := string(data)[7 : len(data)-2]
	fmt.Println(input)
	var artist []ImageID
	err = json.Unmarshal(resp, &artist)

	var genreArtist []ImageID
	var wg sync.WaitGroup
	for _, values := range artist {
		wg.Add(1)
		go func(values ImageID) {
			defer wg.Done()
			genre := InformationArtistTag(values.Name)
			if strings.Contains(genre, input) && input != "alternative" {
				values.Genre = genre
				genreArtist = append(genreArtist, values)
			}
			if input == "alternative" && genre == "alternative" {
				values.Genre = genre
				genreArtist = append(genreArtist, values)
			}
		}(values)
	}
	wg.Wait()
	fmt.Println(genreArtist)
	json.NewEncoder(w).Encode(genreArtist)
}

func SearchBar(w http.ResponseWriter, r *http.Request) {
	getArtists, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")
	getLocation, err := http.Get("https://groupietrackers.herokuapp.com/api/locations")
	getDates, err := http.Get("https://groupietrackers.herokuapp.com/api/dates")
	var artist []ImageID
	var location AllLocation
	var dates AllDates
	respArtists, err := ioutil.ReadAll(getArtists.Body)
	respLocations, err := ioutil.ReadAll(getLocation.Body)
	respDates, err := ioutil.ReadAll(getDates.Body)
	err = json.Unmarshal(respArtists, &artist)
	err = json.Unmarshal(respLocations, &location)
	err = json.Unmarshal(respDates, &dates)
	if err != nil {
		fmt.Println("erreur")
	}
	data := Test{
		Artists:  artist,
		Location: location,
		Dates:    dates,
	}
	json.NewEncoder(w).Encode(&data)
	if err != nil {
		fmt.Println(err)
	}
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
