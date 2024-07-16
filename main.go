package main

import (
	"fmt"
	groupie_tracker "groupie-tracker/handlers"

	"net/http"
)

func main() {
	Port := ":3000"
	// THE HANDLEFUNC FUNCTION WILL HANDLE THE REQUESTS TO THE API
	// http.HandleFunc("/api/artists", groupie_tracker.HandleArtistsRequest)
	// http.HandleFunc("/api/dates", groupie_tracker.HandleDatesRequest)
	// http.HandleFunc("/api/locations", groupie_tracker.HandleLocationsRequest)
	// http.HandleFunc("/api/relation", groupie_tracker.HundelRelationRequest)

	http.HandleFunc("/", groupie_tracker.HandleArtistsPage)
	http.HandleFunc("/Detailes/", groupie_tracker.HandleDetailesPage)

	fmt.Printf("Starting server on %s\n", Port)
	fmt.Println("http://localhost" + Port)
	if err := http.ListenAndServe(Port, nil); err != nil {
		fmt.Printf("Error starting server: %v\n", err)
	}
}
