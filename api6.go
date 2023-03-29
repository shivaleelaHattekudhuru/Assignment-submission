
/*To add a location filter for the existing statistics API, 
we can create a middleware function that checks whether the request is coming from an authorized location 
or not. If the location is not authorized, the middleware function will return an error response.*/
package main

func locationMiddleware(next http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        authorizedLocation := "Germany"
        userLocation := r.Header.Get("X-Location")

        if userLocation != authorizedLocation {
            http.Error(w, "Unauthorized location", http.StatusUnauthorized)
            return
        }

        next(w, r)
    }
}

/*n the above code, we're checking whether the user location in the request header (X-Location) matches the authorized location (germany). 
If it doesn't match, we're returning an unauthorized error response. If it matches, we're calling the next handler in the chain.
To use this middleware function with the /statistics endpoint, we can wrap the handler function with this middleware function:*/


func statisticsHandler(w http.ResponseWriter, r *http.Request) {
    // handle statistics request
}

http.Handle("/statistics", locationMiddleware(statisticsHandler))

type LocationRequest struct {
    Location string `json:"location"`
}

func setLocationHandler(w http.ResponseWriter, r *http.Request) {
    var locationReq LocationRequest
    err := json.NewDecoder(r.Body).Decode(&locationReq)
    if err != nil {
        http.Error(w, "Invalid request body", http.StatusBadRequest)
        return
    }

    // save the user's location in the session or database
    // ...

    fmt.Fprintf(w, "Location set to %s", locationReq.Location)
}

http.HandleFunc("/setlocation", setLocationHandler)
func resetLocationHandler(w http.ResponseWriter, r *http.Request) {
    // reset the user's location in the session or database
    // ...

    fmt.Fprint(w, "Location reset")
}

http.HandleFunc("/resetlocation", resetLocationHandler)


var statisticsData = map[string]int{
    "pageviews": 100,
    "clicks":    50,
}

func statisticsHandler(w http.ResponseWriter, r *http.Request) {
    json.NewEncoder(w).Encode(statisticsData)
}

http.Handle("/statistics", locationMiddleware(statisticsHandler))