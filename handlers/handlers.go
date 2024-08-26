package groupie_tracker

import (
	"encoding/json"
	"fmt"
	i "groupie-tracker/internal"
	"html/template"
	"net/http"
	"strconv"
	"sync"
)

// THIS FUNCTION WILL HANDLE THE REQUEST TO THE HOME PAGE
func HandleArtistsPage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	// get the data from the api
	jsonData, err := http.Get(i.Urls.ArtistsApi)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// decode the data into the artist struct
	var artistsData []i.ArtistData
	if err := json.NewDecoder(jsonData.Body).Decode(&artistsData); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Search(w, r)

	// parse the template
	Template, err := template.ParseFiles("./web/templates/Home.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// execute the template and pass the data to the front end
	if err = Template.Execute(w, artistsData); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// THIS FUNCTION WILL HANDLE THE REQUEST TO THE DETAILS PAGE
func HandleDetailsPage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/Details/" {
		http.NotFound(w, r)
		return
	}
	// get the id from the url
	id := r.FormValue("id")
	fmt.Println(id)

	// check if the id is empty
	Id, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println(err)
		return
	}
	if Id < 1 || Id > 52 {
		http.Error(w, "No id provided", http.StatusBadRequest)
		return

	}

	var Details i.NewDetails
	wg := sync.WaitGroup{}
	// wait for the data to be fetched
	wg.Add(4)
	// go routines to fetch the data from the api
	go i.FetchData(i.Urls.ArtistsApi, id, &Details.Artist, &wg)
	go i.FetchData(i.Urls.DatesApi, id, &Details.Dates, &wg)
	go i.FetchData(i.Urls.LocationsApi, id, &Details.Location, &wg)
	go i.FetchData(i.Urls.RelationsApi, id, &Details.Relations, &wg)
	// wait still done for the 4 go routines to finish
	wg.Wait()

	// parse the template and pass the data to the front end
	Template, err := template.ParseFiles("./web/templates/Details.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// execute the template and pass the data to the front end
	if err := Template.Execute(w, Details); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func Search(w http.ResponseWriter, r *http.Request) {
	//parse form data
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	//parse form data
	fmt.Println("test")
	text := r.FormValue("text")
	category := r.FormValue("category")
	fmt.Println(text)
	fmt.Println(category)
}
