package groupie_tracker

import (
	"encoding/json"
	"fmt"
	groupie_tracker "groupie-tracker/utils"
	"html/template"
	"net/http"
)

// THE ARTIST STRUCT WILL HOLD THE DATA OF THE ARTIST
var (
	API = "https://groupietrackers.herokuapp.com/api"
	// RelationsApi = "https://groupietrackers.herokuapp.com/api/relation"
	// DatesApi = "https://groupietrackers.herokuapp.com/api/dates"
	// LocationsApi = "https://groupietrackers.herokuapp.com/api/locations"

)

type artistData struct {
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

type DatesLocations struct {
	ID             int                 `json:"id"`
	DatesLocations map[string][]string `json:"datesLocations"`
}

type Locations struct {
	ID        int      `json:"id"`
	LocationS []string `json:"Locations"`
	Dates     string   `json:"dates"`
}

type ConcertDates struct {
	ID    int      `json:"id"`
	Dates []string `json:"dates"`
}

// THE NEWDETAILS STRUCT WILL HOLD THE DATA OF THE ARTIST
type NewDetails struct {
	ArtistData artistData
	Dates      ConcertDates 
	Locations  Locations
	Relations  DatesLocations
}

// THIS FUNCTION WILL HANDLE THE REQUEST TO THE HOME PAGE
func HandleArtistsPage(w http.ResponseWriter, r *http.Request) {
	// get the data from the api
	jsonData, err := http.Get(groupie_tracker.API + "/artists")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// decode the data into the artist struct
	var artistsData []*artistData
	if err := json.NewDecoder(jsonData.Body).Decode(&artistsData); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// parse the template
	tmpl, err := template.ParseFiles("./web/templates/Home.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// execute the template and pass the data to the front end
	err = tmpl.Execute(w, artistsData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// THIS FUNCTION WILL GET THE ID FROM THE URL
func getIdFromURL(r *http.Request) string {
	id := r.URL.Query().Get("id")
	return id
}

func GetData(category string, Id string) (*http.Response, error) {
	// THE GET FUNCTION WILL MAKE A GET REQUEST TO THE API AND RETURN THE RESPONSE
	data, err := http.Get(groupie_tracker.API + "/" + category + "/" + Id)
	if err != nil {
		return nil, fmt.Errorf("error fetching %s: %v", API+"/"+category, err)
	}

	return data, nil
}

// THIS FUNCTION WILL HANDLE THE REQUEST TO THE DETAILS PAGE
func HandleDetailesPage(w http.ResponseWriter, r *http.Request) {
	// get the id from the url
	id := getIdFromURL(r)
	var NewDetails NewDetails
	// // get artist relations dates data from the api
	relationsData, err := GetData("relation", id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := json.NewDecoder(relationsData.Body).Decode(&NewDetails.Relations); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// fmt.Println(relations)

	// get artist concert dates data from the api
	concertDatesData, err := GetData("dates", id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := json.NewDecoder(concertDatesData.Body).Decode(&NewDetails.Dates); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// fmt.Println("DATES", concertDates)

	// get artist locations data from the api
	locationsData, err := GetData("locations", id)
	if err != nil {
		http.Error(w, "hna4", http.StatusInternalServerError)
		return
	}
	if err := json.NewDecoder(locationsData.Body).Decode(&NewDetails.Locations); err != nil {
		http.Error(w, "hna3", http.StatusInternalServerError)
		return
	}
	// fmt.Println("locations: ", locations)

	jsonData, err := GetData("artists", id)
	if err != nil {
		http.Error(w, "hna2", http.StatusInternalServerError)
		return
	}
	if err := json.NewDecoder(jsonData.Body).Decode(&NewDetails.ArtistData); err != nil {
		http.Error(w, "hna1", http.StatusInternalServerError)
		return
	}

	fmt.Println("Detailes data: ", NewDetails)

	// parse the template and pass the data to the front end
	tmpl, err := template.ParseFiles("./web/templates/Details.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// execute the template and pass the data to the front end
	err = tmpl.Execute(w, NewDetails)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
