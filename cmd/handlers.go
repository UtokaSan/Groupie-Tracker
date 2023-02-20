package cmd

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"net/url"
)

type ImageID struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Image string `json:"image"`
	Genre string
}

type ArtistInformation struct {
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
}

type Location struct {
	ID        int      `json:"id"`
	Locations []string `json:"locations"`
}
type Dates struct {
	ID    int      `json:"id"`
	Dates []string `json:"dates"`
}

type Relations struct {
	ID             int                 `json:"id"`
	DatesLocations map[string][]string `json:"datesLocations"`
}

type Tag struct {
	Count int    `json:"count"`
	Name  string `json:"name"`
	URL   string `json:"url"`
}

type Toptags struct {
	Tag []Tag `json:"tag"`
}

type AllTags struct {
	Toptags Toptags `json:"toptags"`
}

var artistGenre ImageID

func IndexHandlers(w http.ResponseWriter, r *http.Request) {
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
		t, err := template.ParseFiles("templates/artistinfo.html")
		data, _ := ioutil.ReadAll(r.Body)
		input := string(data)[6 : len(data)-1]
		InformationArtist(w, input)
		InformationLocation(w, input)
		InformationDate(w, input)
		if err != nil {
			fmt.Println(err)
		}
		t.Execute(w, r)
	}
}

func InformationArtist(w http.ResponseWriter, id string) {
	get, err := http.Get("https://groupietrackers.herokuapp.com/api/artists/" + id)
	if err != nil {
		fmt.Println(err)
	}
	data, err := ioutil.ReadAll(get.Body)
	var artistInformation ArtistInformation
	err = json.Unmarshal(data, &artistInformation)
	if err != nil {
		fmt.Println(err)
		return
	}
	json.NewEncoder(w).Encode(artistInformation)
	fmt.Println(artistInformation)
}

func InformationArtistTag(name string) string {
	apikey := "471f5119aa12d32718ae05f982f745dc"
	escapeName := url.QueryEscape(name)
	get, err := http.Get("http://ws.audioscrobbler.com/2.0/?method=artist.getTopTags&artist=" + escapeName + "&api_key=" + apikey + "&format=json")
	if err != nil {
		fmt.Println(err)
	}
	data, err := ioutil.ReadAll(get.Body)
	var artistTopTags AllTags
	err = json.Unmarshal(data, &artistTopTags)
	if err != nil {
		fmt.Println(err)
	}
	if len(artistTopTags.Toptags.Tag[0].Name) > 0 {
		artistGenre := artistTopTags.Toptags.Tag[0].Name
		return artistGenre
	}
	return "bad"
}

func InformationLocation(w http.ResponseWriter, id string) {
	get, err := http.Get("https://groupietrackers.herokuapp.com/api/locations/" + id)
	if err != nil {
		fmt.Println(err)
		return
	}
	data, err := ioutil.ReadAll(get.Body)
	var artistLocations Location
	err = json.Unmarshal(data, &artistLocations)
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
	data, err := ioutil.ReadAll(get.Body)
	var artistDates Dates
	err = json.Unmarshal(data, &artistDates)
	if err != nil {
		fmt.Println(err)
		return
	}
	json.NewEncoder(w).Encode(artistDates)
	fmt.Println(artistDates)
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
	for i, _ := range artist {
		artist[i].Genre = InformationArtistTag(artist[i].Name)
	}
	json.NewEncoder(w).Encode(artist)
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
