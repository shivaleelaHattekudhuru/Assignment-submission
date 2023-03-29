package main

import (
	"encoding/json"
	"net/http"
	"strconv"
	"sync"
	"time"
)

type transaction struct {
	amount    float64
	timestamp time.Time
}

type stats struct {
	sum   float64
	avg   float64
	max   float64
	min   float64
	count uint64
}

type transactionManager struct {
	mtx          sync.RWMutex
	transactions []transaction
}

func (tm *transactionManager) addTransaction(amount float64, timestamp time.Time) {
	tm.mtx.Lock()
	defer tm.mtx.Unlock()

	tm.transactions = append(tm.transactions, transaction{amount: amount, timestamp: timestamp})
}

func (tm *transactionManager) getStats() stats {
	tm.mtx.RLock()
	defer tm.mtx.RUnlock()

	now := time.Now()
	var sum, avg, max, min float64
	var count uint64
	for i := len(tm.transactions) - 1; i >= 0; i-- {
		t := tm.transactions[i]
		if now.Sub(t.timestamp) > time.Second*60 {
			break
		}
		count++
		sum += t.amount
		if count == 1 || t.amount > max {
			max = t.amount
		}
		if count == 1 || t.amount < min {
			min = t.amount
		}
	}
	if count > 0 {
		avg = sum / float64(count)
	}
	return stats{sum: sum, avg: avg, max: max, min: min, count: count}
}

var tm transactionManager

func main() {
	http.HandleFunc("/transactions", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			amount, err := strconv.ParseFloat(r.FormValue("amount"), 64)
			if err != nil {
				http.Error(w, "Invalid amount", http.StatusBadRequest)
				return
			}
			timestamp, err := time.Parse(time.RFC3339Nano, r.FormValue("timestamp"))
			if err != nil {
				http.Error(w, "Invalid timestamp", http.StatusBadRequest)
				return
			}
			tm.addTransaction(amount, timestamp)
			w.WriteHeader(http.StatusCreated)
		} else if r.Method == http.MethodGet {
			s := tm.getStats()
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(s)
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})
	http.ListenAndServe(":8080", nil)
}
