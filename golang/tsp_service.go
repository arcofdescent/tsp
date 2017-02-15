// Travelling salesman problem
// ---------------------------
// Given a list of cities and the distances between each pair of cities, what is
// the shortest possible route that visits each city exactly once and returns to
// the origin city?
package main

import (
	"encoding/json"
	"fmt"
	"github.com/arcofdescent/golang/tsp"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"time"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/calcRoute", calcRoute)
	http.Handle("/", r)

	srv := &http.Server{
		Addr:    "127.0.0.1:8000",
		Handler: r,
	}

	srv.ListenAndServe()
}

func calcRoute(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Got request")
	startTime := time.Now()

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Printf("Error while reading body: %s\n", err)
	}
	//fmt.Printf("body: %s\n", body)

	points := make([]tsp.Point, 0)
	err = json.Unmarshal(body, &points)
	if err != nil {
		fmt.Printf("Error while unmarshalling: %s\n", err)
	}
	//fmt.Printf("struct: %#v", d)

	distances := tsp.CalcDistanceBetweenPoints(points)
	//fmt.Printf("distances: %#v\n", distances)

	res := tsp.CalcShortestRoute(len(points), distances)
	duration := time.Since(startTime)
	res.Duration = fmt.Sprintf("%v", duration)
	fmt.Printf("res: %#v\n", res)
	fmt.Printf("duration: %v\n", duration)
	/*
		res_json, err := json.Marshal(res)
		if err != nil {
			fmt.Printf("Error while marshaling: %s\n", err)
		}
		fmt.Printf("res_json: %s\n", res_json)
	*/

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
	//w.Write(res_json)
}
