package main

import (
	"fmt"
	groupie_tracker "groupie-tracker/handlers"
	"net/http"
)

func main() {
	Port := ":3000"
	// Serve static files from the "web/css/" directory for URLs starting with "/Detailes/"
	http.Handle("/style/", http.StripPrefix("/style/", http.FileServer(http.Dir("./web/css/"))))

	// Handle requests for the home page "/"
	http.HandleFunc("/", groupie_tracker.HandleArtistsPage)

	// Handle requests for "/Detailes" (assuming it should be "/Details")
	http.HandleFunc("/Detailes/", groupie_tracker.HandleDetailesPage)

	fmt.Printf("Starting server on %s\n", Port)
	fmt.Println("http://localhost" + Port)
	if err := http.ListenAndServe(Port, nil); err != nil {
		fmt.Printf("Error starting server: %v\n", err)
	}
}
