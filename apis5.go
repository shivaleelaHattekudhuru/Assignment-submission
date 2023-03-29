package main

import (
	"fmt"
	"log"
	"net/http"
)
type Location struct {
	City string `json:"city"`
}
var userLocation *Location

func resetLocationHandler(w http.ResponseWriter, r *http.Request) {
	if userLocation == nil {
		fmt.Fprint(w, "Location not set")
		return
	}

	userLocation = nil
	fmt.Fprint(w, "Location reset")
}

func main() {
	http.HandleFunc("/location/reset", resetLocationHandler)
	log.Fatal(http.ListenAndServe(":5000", nil))
}
