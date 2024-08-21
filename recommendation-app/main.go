package main

import (
    "net/http"
    "recommendation-app/routes"
)

func main() {
 
    http.HandleFunc("/recommendations", routes.RecommendationsHandler)

    http.ListenAndServe(":8080", nil)
}
