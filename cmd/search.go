package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func SearchBar(w http.ResponseWriter, r *http.Request) {
	getArtists, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		http.Error(w, "Failed get Artists", http.StatusInternalServerError)
		return
	}
	getLocations, err := http.Get("https://groupietrackers.herokuapp.com/api/locations")
	if err != nil {
		http.Error(w, "Failed get Locations", http.StatusInternalServerError)
		return
	}
	getDates, err := http.Get("https://groupietrackers.herokuapp.com/api/dates")
	if err != nil {
		http.Error(w, "Failed get Dates", http.StatusInternalServerError)
		return
	}
	var artists []ImageID
	var locations AllLocation
	var dates AllDates
	respArtists, err := ioutil.ReadAll(getArtists.Body)
	if err != nil {
		http.Error(w, "Failed request Artists", http.StatusInternalServerError)
		return
	}
	respLocations, err := ioutil.ReadAll(getLocations.Body)
	if err != nil {
		http.Error(w, "Failed request Locations", http.StatusInternalServerError)
		return
	}
	respDates, err := ioutil.ReadAll(getDates.Body)
	if err != nil {
		http.Error(w, "Failed request Dates", http.StatusInternalServerError)
		return
	}
	err = json.Unmarshal(respArtists, &artists)
	err = json.Unmarshal(respLocations, &locations)
	err = json.Unmarshal(respDates, &dates)

	data := AllArtistInformation{
		Artists:  artists,
		Location: locations,
		Dates:    dates,
	}
	json.NewEncoder(w).Encode(&data)
	if err != nil {
		fmt.Println(err)
	}
}
