package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "psipsikos"
	dbname   = "postgres"
)

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")
	handleRequests(db)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func returnSpotsWithDB(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query()
		lat, _ := strconv.ParseFloat(query.Get("latitude"), 64)
		long, _ := strconv.ParseFloat(query.Get("logitude"), 64)
		radius, _ := strconv.ParseFloat(query.Get("radius"), 64)
		radType := query.Get("type")
		spots, err := findSpots(db, float32(lat), float32(long), float32(radius), string(radType))
		if err != nil {
			panic(err)
		}
		fmt.Println(spots)
		json.NewEncoder(w).Encode(spots)
	}
}

func handleRequests(db *sql.DB) {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/spots", returnSpotsWithDB(db))
	log.Fatal(http.ListenAndServe(":10000", nil))
}

type Response struct {
	Status      string
	Code        string
	Total       int
	Coordinates []Coordinate `json:"data"`
}

type Coordinate struct {
	Latitude  float64
	Longitude float64
	Radius    int
}
