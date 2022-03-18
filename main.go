package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

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

	// spots, err := findSpots(db, 23.7249394, 37.9782996, 3000, "circle")
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(spots)

	handleRequests(db)
	
	res, err := http.Get("https://fakerapi.it/api/v1/custom?_quantity=1&latitude=latitude&longitude=longitude&radius=number")
	if err !=nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err !=nil {
		log.Fatal(err)
	}

	var response Response
	json.Unmarshal(body, &response)

	for p := range response.Coordinates {
		fmt.Println("Latitude:", p.Latitude)
		fmt.Println("Longitude:", p.Longitude)
		fmt.Println("Radius:", p.Radius)
	}
}

func homePage(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "Welcome to the HomePage!")
    fmt.Println("Endpoint Hit: homePage")
}

func returnSpotsWithDB(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request){
      fmt.Println("Endpoint Hit: returnSpots")
      spots, err := findSpots(db, 23.7249394, 37.9782996, 3000, "circle")
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
	Status string
	Code string
	Total int
	Coordinates []Coordinate `json:"data"`
}

type Coordinate struct {
	Latitude float64
	Longitude float64
	Radius int
}
