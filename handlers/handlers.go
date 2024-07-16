package groupie_tracker

import (
	"encoding/json"
	"fmt"
	groupie_tracker "groupie_tracker/utils"
	"html/template"
	"net/http"
)

type Artist struct {
	ID           int      `json:"id"`
	Name         string   `json:"name"`
	Image        string   `json:"image"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Members      []string `json:"members"`
	Relations    string   `json:"relations"`
	ConcertDates string   `json:"concertDates"`
	Locations    string   `json:"locations"`
}

// that will huandle the /artists , will pass the data cards to the front end
func HandleArtistsPage(w http.ResponseWriter, r *http.Request) {
	jsonData, err := http.Get(groupie_tracker.API + "/artists")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	var artists []Artist
	if err := json.NewDecoder(jsonData.Body).Decode(&artists); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tmpl, err := template.ParseFiles("./web/templates/Home.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, artists)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}

func getIdFromURL(r *http.Request) string {
	id := r.URL.Query().Get("id")
	fmt.Println(id)
	return id
}

func HandleDetailesPage(w http.ResponseWriter, r *http.Request) {
	id := getIdFromURL(r)
	jsonData, err := http.Get(groupie_tracker.API + "/artists/" + id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	var artist Artist
	if err := json.NewDecoder(jsonData.Body).Decode(&artist); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tmpl, err := template.ParseFiles("./web/templates/Details.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = tmpl.Execute(w, artist)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// that func will handle the Artist details page
