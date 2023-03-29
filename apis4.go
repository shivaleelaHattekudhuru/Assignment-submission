package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	 "github.com/jinzhu/gorm"
	_ "github.com/go-sql-driver/mysql"
	
	_ "gorm.io/gorm"
)

var db *gorm.DB

type Location struct {
	City string `json:"city"`
}

var userLocation *Location

func setLocationHandler(w http.ResponseWriter, r *http.Request) {
	var location Location
	err := json.NewDecoder(r.Body).Decode(&location)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	userLocation = &location
	fmt.Fprintf(w, "Location set to %s", location.City)
}

func Search() {

	var err error

	db, err = gorm.Open("mysql", "root:Shiva@123@tcp(127.0.0.1:3306)/trainings?charset=utf8&parseTime=True")

	if err != nil {
		fmt.Print("jvt", err)
		panic("db not connected")

	}
	db.AutoMigrate(&Location{})

}

func main() {
	http.HandleFunc("/location/set", setLocationHandler)
	log.Fatal(http.ListenAndServe(":9090", nil))
}
