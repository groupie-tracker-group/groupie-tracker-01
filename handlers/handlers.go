package groupie_tracker

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"sync"
)

// THE ARTIST STRUCT WILL HOLD THE DATA OF THE ARTIST
var (
	ArtistApi    = "https://groupietrackers.herokuapp.com/api/artists"
	RelationsApi = "https://groupietrackers.herokuapp.com/api/relation"
	DatesApi     = "https://groupietrackers.herokuapp.com/api/dates"
	LocationsApi = "https://groupietrackers.herokuapp.com/api/locations"
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

type datesLocations struct {
	ID             int                 `json:"id"`
	DatesLocations map[string][]string `json:"datesLocations"`
}

type locations struct {
	ID        int      `json:"id"`
	LocationS []string `json:"locations"`
	Dates     string   `json:"dates"`
}

type concertDates struct {
	ID    int      `json:"id"`
	Dates []string `json:"dates"`
}

// THE NEWDETAILS STRUCT WILL HOLD THE DATA OF THE ARTIST
type NewDetails struct {
	ArtistData artistData
	Dates      concertDates
	Locations  locations
	Relations  datesLocations
}

// THIS FUNCTION WILL HANDLE THE REQUEST TO THE HOME PAGE
func HandleArtistsPage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	// get the data from the api
	jsonData, err := http.Get(ArtistApi)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// decode the data into the artist struct
	var artistsData []artistData
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

// THIS GetDATA FUNCTION WILL MAKE A GET REQUEST TO THE API AND  DECODE THE DATA INTO THE DATAFORM STRUCT AND RETURN THE ERROR
func fetshData(apiEndpoint string, Id string, DataForm interface{}, wg *sync.WaitGroup) {
	// defer the done function to the end of the function
	defer wg.Done()
	// THE GET FUNCTION WILL MAKE A GET REQUEST TO THE API AND RETURN THE RESPONSE
	Response, err := http.Get(apiEndpoint + "/" + Id)
	if err != nil {
		log.Printf("\033[31m fetshing error \033[0m %s: \033[33m %v \033[0m", apiEndpoint, err)
	}
	// the response body will be closed after the function is done
	defer Response.Body.Close()

	if err := json.NewDecoder(Response.Body).Decode(DataForm); err != nil {
		log.Printf("\033[31m decoding error \033[0m %s: \033[33m %v \033[0m", apiEndpoint, err)
		return
	}
}

// THIS FUNCTION WILL HANDLE THE REQUEST TO THE DETAILS PAGE
func HandleDetailesPage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/Detailes/" {
		http.NotFound(w, r)
		return
	}
	// get the id from the url
	id := getIdFromURL(r)

	// check if the id is empty
	Id, _ := strconv.Atoi(id)
	if Id < 1 || Id > 52 {
		http.Error(w, "No id provided", http.StatusBadRequest)
		return

	}
	var NewDetails NewDetails

	wg := sync.WaitGroup{}
	// wait for the data to be fetched
	wg.Add(4)
	// go routines to fetch the data from the api
	go fetshData(ArtistApi, id, &NewDetails.ArtistData, &wg)
	go fetshData(DatesApi, id, &NewDetails.Dates, &wg)
	go fetshData(LocationsApi, id, &NewDetails.Locations, &wg)
	go fetshData(RelationsApi, id, &NewDetails.Relations, &wg)
	// wait still done for the 4 go routines to finish
	wg.Wait()

	// parse the template and pass the data to the front end
	tmpl, err := template.ParseFiles("./web/templates/Details.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// execute the template and pass the data to the front end
	if err := tmpl.Execute(w, NewDetails); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
