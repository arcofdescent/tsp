// Travelling salesman problem
// ---------------------------
// Given a list of cities and the distances between each pair of cities, what is
// the shortest possible route that visits each city exactly once and returns to
// the origin city?
package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	//"io"
	"io/ioutil"
	"math"
	"net/http"
	"sort"
	"strconv"
	//"strings"
)

type Distance struct {
	Id1      string
	Id2      string
	Distance float64
}

type Point struct {
	Id string
	X  float64
	Y  float64
}

type Result struct {
	Route  []string
	Length string
}

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

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Printf("Error while reading body: %s\n", err)
	}
	//fmt.Printf("body: %s\n", body)

	points := make([]Point, 0)
	err = json.Unmarshal(body, &points)
	if err != nil {
		fmt.Printf("Error while unmarshalling: %s\n", err)
	}
	//fmt.Printf("struct: %#v", d)

	distances := calcDistanceBetweenPoints(points)
	//fmt.Printf("distances: %#v\n", distances)

	res := calcShortestRoute(len(points), distances)
	fmt.Printf("res: %#v\n", res)
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

func calcShortestRoute(numPoints int, distances []Distance) Result {

	idxs := make([]int, 0)
	for i := 0; i < numPoints; i++ {
		idxs = append(idxs, i)
	}
	//fmt.Printf("idxs: %v\n", idxs)

	perms := permutations(idxs)
	//fmt.Printf("perms: %v\n", perms)

	distance_ids := make([][]string, 0)
	for _, val := range perms {
		dst_ids := make([]string, 0)
		for _, v := range val {
			id := "P" + strconv.Itoa(v+1)
			dst_ids = append(dst_ids, id)
		}
		dst_ids = append(dst_ids, "P"+strconv.Itoa(val[0]+1))
		distance_ids = append(distance_ids, dst_ids)
	}
	//fmt.Printf("distance_ids: %v\n", distance_ids)

	var shortest_route_length float64
	var shortest_route []string

	for _, route := range distance_ids {
		route_pairs := make([][]string, 0)
		for idx := 0; idx < len(route)-1; idx++ {
			sorted_pair := []string{route[idx], route[idx+1]}
			sort.Strings(sorted_pair)
			//fmt.Printf("sorted_pair: %v\n", sorted_pair)
			route_pairs = append(route_pairs, sorted_pair)
		}

		//fmt.Printf("route: %v\n", route)
		//fmt.Printf("route_pairs: %v\n", route_pairs)

		dst := getRouteLength(route_pairs, distances)
		//fmt.Printf("dst: %.2f\n", dst)

		if shortest_route_length == 0 {
			shortest_route_length = dst
			shortest_route = route
		}

		if dst < shortest_route_length {
			shortest_route_length = dst
			shortest_route = route
		}
	}

	fmt.Printf("shortest_route: %v\n", shortest_route[:len(shortest_route)-1])
	fmt.Printf("shortest_route_length: %.2f\n", shortest_route_length)

	res := Result{Route: shortest_route[:len(shortest_route)-1], Length: fmt.Sprintf("%.2f", shortest_route_length)}
	return res
}

func getRouteLength(pairs [][]string, distances []Distance) float64 {

	var dst float64
	for _, pair := range pairs {
		dst += getDistance(pair[0], pair[1], distances)
	}

	return dst
}

func getDistance(p1 string, p2 string, distances []Distance) float64 {

	for _, d := range distances {
		if d.Id1 == p1 && d.Id2 == p2 {
			return d.Distance
		}
	}

	return 0
}

func calcDistanceBetweenPoints(points []Point) []Distance {

	distances := make([]Distance, 0)

	combinations(len(points), 2, func(c []int) {
		//fmt.Printf("c: %#v\n", c)
		id1 := points[c[0]].Id
		x1 := points[c[0]].X
		y1 := points[c[0]].Y
		id2 := points[c[1]].Id
		x2 := points[c[1]].X
		y2 := points[c[1]].Y

		dst := math.Sqrt(((y2 - y1) * (y2 - y1)) + ((x2 - x1) * (x2 - x1)))
		distance := Distance{Id1: id1, Id2: id2, Distance: dst}
		//fmt.Printf("dst: %#v\n", distance)
		distances = append(distances, distance)
	})

	return distances
}

func permutations(arr []int) [][]int {
	var helper func([]int, int)
	res := [][]int{}

	helper = func(arr []int, n int) {
		if n == 1 {
			tmp := make([]int, len(arr))
			copy(tmp, arr)
			res = append(res, tmp)
		} else {
			for i := 0; i < n; i++ {
				helper(arr, n-1)
				if n%2 == 1 {
					tmp := arr[i]
					arr[i] = arr[n-1]
					arr[n-1] = tmp
				} else {
					tmp := arr[0]
					arr[0] = arr[n-1]
					arr[n-1] = tmp
				}
			}
		}
	}

	helper(arr, len(arr))
	return res
}

func combinations(n, m int, emit func([]int)) {
	s := make([]int, m)
	last := m - 1
	var rc func(int, int)
	rc = func(i, next int) {
		for j := next; j < n; j++ {
			s[i] = j
			if i == last {
				emit(s)
			} else {
				rc(i+1, j+1)
			}
		}
		return
	}
	rc(0, 0)
}
