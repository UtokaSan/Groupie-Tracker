package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"sync"
)

func InformationArtistAlbum(w http.ResponseWriter, r *http.Request) {
	input := getInput(r, 12)
	defer r.Body.Close()
	apikey := "471f5119aa12d32718ae05f982f745dc"
	getAlbum, err := http.Get("https://ws.audioscrobbler.com/2.0/?method=artist.gettopalbums&artist=" + url.QueryEscape(input) + "&api_key=" + "&api_key=" + apikey + "&format=json")
	if err != nil {
		fmt.Println(err)
	}
	defer getAlbum.Body.Close()
	getListeners, err := http.Get("https://ws.audioscrobbler.com/2.0/?method=artist.getinfo&artist=" + url.QueryEscape(input) + "&api_key=" + "&api_key=" + apikey + "&format=json")
	if err != nil {
		fmt.Println(err)
	}
	defer getListeners.Body.Close()
	defer getAlbum.Body.Close()
	var artistTopAlbum AllAlbum
	var artistListeners AllListeners
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		respAlbum, err := ioutil.ReadAll(getAlbum.Body)
		if err != nil {
			http.Error(w, "Failed request Artist", http.StatusInternalServerError)
			return
		}
		respListeners, err := ioutil.ReadAll(getListeners.Body)
		if err != nil {
			http.Error(w, "Failed request Artist", http.StatusInternalServerError)
			return
		}
		err = json.Unmarshal(respAlbum, &artistTopAlbum)
		err = json.Unmarshal(respListeners, &artistListeners)
		dataResult := AllInfoArtist{
			AllListeners: artistListeners,
			AllAlbum:     artistTopAlbum,
		}
		json.NewEncoder(w).Encode(&dataResult)
	}()
	wg.Wait()
}

func ArtistInfoGet(w http.ResponseWriter, r *http.Request) {
	input := getInput(r, 7)
	getArtist, err := http.Get("https://groupietrackers.herokuapp.com/api/artists/" + input)
	if err != nil {
		http.Error(w, "Failed get Artist", http.StatusInternalServerError)
		return
	}
	getLocation, err := http.Get("https://groupietrackers.herokuapp.com/api/locations/" + input)
	if err != nil {
		http.Error(w, "Failed get Location", http.StatusInternalServerError)
		return
	}
	getDate, err := http.Get("https://groupietrackers.herokuapp.com/api/dates/" + input)
	if err != nil {
		http.Error(w, "Failed get Dates", http.StatusInternalServerError)
		return
	}
	defer getArtist.Body.Close()
	defer getLocation.Body.Close()
	defer getDate.Body.Close()
	var artist ImageID
	var location Location
	var dates Date
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		respArtists, err := ioutil.ReadAll(getArtist.Body)
		if err != nil {
			http.Error(w, "Failed request Artist", http.StatusInternalServerError)
			return
		}
		respLocations, err := ioutil.ReadAll(getLocation.Body)
		if err != nil {
			http.Error(w, "Failed request Location", http.StatusInternalServerError)
			return
		}
		respDates, err := ioutil.ReadAll(getDate.Body)
		if err != nil {
			http.Error(w, "Failed request Date", http.StatusInternalServerError)
			return
		}
		err = json.Unmarshal(respArtists, &artist)
		err = json.Unmarshal(respLocations, &location)
		err = json.Unmarshal(respDates, &dates)
		dataResult := ArtistInformation{
			Artist:   artist,
			Location: location,
			Dates:    dates,
		}
		json.NewEncoder(w).Encode(&dataResult)
	}()
	wg.Wait()
}
