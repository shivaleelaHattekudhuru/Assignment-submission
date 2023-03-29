package main

import (
	"fmt"
	"net/http"
)

func deleteTransactionsHandler(w http.ResponseWriter, r *http.Request) {
	// Delete all transactions here

	w.WriteHeader(http.StatusNoContent)
}

func main() {
	http.HandleFunc("/transactions", deleteTransactionsHandler)
	fmt.Println("Listening on port 3000...")
	http.ListenAndServe(":3000", nil)
}
