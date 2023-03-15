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
	defer r.Body.Close()
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

func ApiGenre(w http.ResponseWriter, r *http.Request) {
	getArtist, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		http.Error(w, "Failed get Artist", http.StatusInternalServerError)
		return
	}
	getLocation, err := http.Get("https://groupietrackers.herokuapp.com/api/locations")
	if err != nil {
		http.Error(w, "Failed get Location", http.StatusInternalServerError)
		return
	}
	var artist []ImageID
	var locations AllLocation
	respArtist, err := ioutil.ReadAll(getArtist.Body)
	if err != nil {
		http.Error(w, "Failed request Artist", http.StatusInternalServerError)
		return
	}
	respLocation, err := ioutil.ReadAll(getLocation.Body)
	if err != nil {
		http.Error(w, "Failed request Location", http.StatusInternalServerError)
		return
	}
	defer getArtist.Body.Close()
	defer getLocation.Body.Close()
	data, _ := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	input := string(data)[7 : len(data)-2]
	fmt.Println(input)
	err = json.Unmarshal(respArtist, &artist)
	err = json.Unmarshal(respLocation, &locations)
	dataResult := Test{
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
	fmt.Println(genreArtist)
	dataResult.Artists = genreArtist
	json.NewEncoder(w).Encode(dataResult)
}

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

	data := Test{
		Artists:  artists,
		Location: locations,
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
