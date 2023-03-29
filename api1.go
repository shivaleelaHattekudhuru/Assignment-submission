package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
	"github.com/jinzhu/gorm"
	_ "github.com/go-sql-driver/mysql"
	
	_ "gorm.io/gorm"


)

var db *gorm.DB

type Transaction struct {
	Amount    string    `json:"amount"`
	Timestamp time.Time `json:"timestamp"`
}


func handleTransactions(w http.ResponseWriter, r *http.Request) {
	// Parse the request body into a Transaction struct
	var transaction Transaction
	err := json.NewDecoder(r.Body).Decode(&transaction)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Check if the transaction is older than 60 seconds
	diff := time.Now().Sub(transaction.Timestamp)
	if diff.Seconds() > 60 {
		http.Error(w, "transaction is older than 60 seconds", http.StatusNoContent)
		return
	}

	// Return an error if the transaction date is in the future
	if transaction.Timestamp.After(time.Now()) {
		http.Error(w, "transaction date is in the future", http.StatusUnprocessableEntity)
		return
	}

	// TODO: Do something with the transaction, like save it to a database

	// Return a success status code
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintln(w, "Transaction created successfully")
}

func init() {

	var err error

	db, err = gorm.Open("mysql", "root:Shiva@123@tcp(127.0.0.1:3306)/trainings?charset=utf8&parseTime=True")

	if err != nil {
		fmt.Print("jvt", err)
		panic("db not connected")

	}
	db.AutoMigrate(&Transaction{})

}

func main() {
	http.HandleFunc("/transactions", handleTransactions)
	http.ListenAndServe(":8080", nil)
}
