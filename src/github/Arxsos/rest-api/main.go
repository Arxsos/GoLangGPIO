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

// Normal post
func tryPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")

	if err := r.ParseForm(); err != nil {
		panic(err)
	}
	fmt.Fprintf(w, "Post from website! r.PostFrom = %v\n", len(r.PostForm))
	test := r.FormValue("test")
	fmt.Fprintf(w, "Name = %s\n", test)
	fmt.Println(test)
}

// POST with JSON
func test(rw http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)
	var t test_struct
	err := decoder.Decode(&t)
	if err != nil {
		panic(err)
	}
	log.Println(t.test)
}

// our main function
func main() {
	// define new router
	router := mux.NewRouter()

	// handler
	router.HandleFunc("/api/hour", getHour).Methods("GET")

	//no JSON
	router.HandleFunc("/api/tryPost", tryPost).Methods("POST")

	//WITH JSON
	router.HandleFunc("/api/tryPostJSON", test).Methods("POST")

	// take care of fatal error
	log.Fatal(http.ListenAndServe(":8001", router))
}
