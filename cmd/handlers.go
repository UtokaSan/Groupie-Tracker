package cmd

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
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

func getInput(r *http.Request, rangeData int) string {
	data, _ := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	input := string(data)[rangeData : len(data)-2]
	return input
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
