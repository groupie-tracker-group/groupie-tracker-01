package data

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"sync"
)

// first get the api keys from the main api giving
func init() {
	Unmarshal()
}

type Url struct {
	ArtistsApi   string `json:"artists"`
	LocationsApi string `json:"locations"`
	DatesApi     string `json:"dates"`
	RelationsApi string `json:"relation"`
}

var Urls Url

func Unmarshal() {
	API := "https://groupietrackers.herokuapp.com/api"
	resp, err := http.Get(API)
	if err != nil {
		log.Println("Error making request:", err)
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		log.Printf("Error: status code %d\n", resp.StatusCode)
		return
	}
	urlData, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error reading response body:", err)
		return
	}
	// get the data json in the form of []byte urlData
	json.Unmarshal(urlData, &Urls)
}

// this section is for getting the json data from the apis
type ArtistData struct {
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
	LocationS []string `json:"locations"`
	Dates     string   `json:"dates"`
}

type ConcertDates struct {
	ID    int      `json:"id"`
	Dates []string `json:"dates"`
}

// THE NewDetails STRUCT WILL HOLD THE DATA OF THE ARTIST
type NewDetails struct {
	Artist    ArtistData
	Dates     ConcertDates
	Location  Locations
	Relations DatesLocations
}

// THIS GetDATA FUNCTION WILL MAKE A GET REQUEST TO THE API AND  DECODE THE DATA INTO THE DATA FORM STRUCT AND RETURN THE ERROR
func FetchData(apiEndpoint string, Id string, DataForm interface{}, wg *sync.WaitGroup) {
	// defer the done function to the end of the function
	defer wg.Done()
	// THE GET FUNCTION WILL MAKE A GET REQUEST TO THE API AND RETURN THE RESPONSE
	Response, err := http.Get(apiEndpoint + "/" + Id)
	if err != nil {
		log.Printf("\033[31m fetching error \033[0m %s: \033[33m %v \033[0m", apiEndpoint, err)
	}
	// the response body will be closed after the function is done
	defer Response.Body.Close()

	if err := json.NewDecoder(Response.Body).Decode(DataForm); err != nil {
		log.Printf("\033[31m decoding error \033[0m %s: \033[33m %v \033[0m", apiEndpoint, err)
		return
	}
}
