package main

import (
	"../../../github.com/gorilla/mux"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

func getHour(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	t := time.Now()

	err := json.NewEncoder(w).Encode(t.Format(time.RFC3339))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("send okey")
}

// our main function
func main() {
	// define new router
	router := mux.NewRouter()

	// handler
	router.HandleFunc("/api/hour", getHour).Methods("GET")

	// take care of fatal error
	log.Fatal(http.ListenAndServe(":8001", router))
}
