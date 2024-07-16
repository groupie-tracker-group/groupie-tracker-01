package groupie_tracker

import (
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

// that will handle the /api/dates route

func HandleDatesRequest(w http.ResponseWriter, r *http.Request) {
	address := r.RemoteAddr
	fmt.Fprintf(w, "Hello %s\n", address)
	fmt.Println()
	jsonData, err := groupie_tracker.GetData("dates")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprint(w, template.HTML(jsonData))
}

// that will handle the /api/artists route

func HandleArtistsRequest(w http.ResponseWriter, r *http.Request) {
	jsonData, err := groupie_tracker.GetData("artists")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprint(w, template.HTML(jsonData))
}

// that will handle the /api/locations route

func HandleLocationsRequest(w http.ResponseWriter, r *http.Request) {
	jsonData, err := groupie_tracker.GetData("locations")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprint(w, template.HTML(jsonData))
}

// that will handle the /api/relation route

func HundelRelationRequest(w http.ResponseWriter, r *http.Request) {
	jsonData, err := groupie_tracker.GetData("relation")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprint(w, template.HTML(jsonData))
}

// that will huandle the /artists , will pass the data cards to the front end
// func HandleArtistsPage(w http.ResponseWriter, r *http.Request) {
// 	jsonData, err := http.Get(groupie_tracker.API + "/artists")
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}
// 	var artists []Artist
// 	if err := json.NewDecoder(jsonData.Body).Decode(&artists); err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}
// 	tmpl, err := template.ParseFiles("./web/templates/Home.html")
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}
// 	err = tmpl.Execute(w, artists)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}

// }

// func HandleDetailesPage(w http.ResponseWriter, r *http.Request) {
// 	id := r.URL.Path
// 	fmt.Print( id)

// }
