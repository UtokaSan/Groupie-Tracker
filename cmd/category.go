package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"sync"
)

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

func takeDataGenre() ([]ImageID, AllLocation, error) {
	getArtist, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		return nil, AllLocation{}, fmt.Errorf("Failed get Artist: %v", err)
	}
	getLocation, err := http.Get("https://groupietrackers.herokuapp.com/api/locations")
	if err != nil {
		return nil, AllLocation{}, fmt.Errorf("Failed get Artist: %v", err)

	}
	var artist []ImageID
	var locations AllLocation
	respArtist, err := ioutil.ReadAll(getArtist.Body)
	if err != nil {
		return nil, AllLocation{}, fmt.Errorf("Failed get Artist: %v", err)

	}
	respLocation, err := ioutil.ReadAll(getLocation.Body)
	if err != nil {
		return nil, AllLocation{}, fmt.Errorf("Failed get Artist: %v", err)
	}
	err = json.Unmarshal(respArtist, &artist)
	err = json.Unmarshal(respLocation, &locations)
	defer getArtist.Body.Close()
	defer getLocation.Body.Close()
	return artist, locations, nil
}

func ApiGenre(w http.ResponseWriter, r *http.Request) {
	input := getInput(r, 7)
	defer r.Body.Close()
	artist, locations, err := takeDataGenre()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	dataResult := AllArtistInformation{
		Artists:  artist,
		Location: locations,
	}
	var genreArtist []ImageID
	var wg sync.WaitGroup
	for _, values := range dataResult.Artists {
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
	dataResult.Artists = genreArtist
	json.NewEncoder(w).Encode(dataResult)
}
