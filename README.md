# GO-endpoint

This endpoint returns cool spots in the given radius of your choice (in meters) in London.  
It receives the coordinates (latitude and longitude), and a radius in meters and returns an array of objects containing all fields in the data set.

## How to Run

### Prerequisites
In order for this to run, you need to have a postgres with postgis extention running on localhost on port 5432,
and that you have imported the data from `spots.sql` file.  

### Clone the repo
```
git clone https://github.com/EviePalaiochorinou/GO-endpoint
```

### To start the program run

```go
go mod download
go run ./...
```

### To make a request against it run:
`curl http://localhost:10000/spots?latitude=51.5072&longitute=0.1276&radius=200&type=circle`

This will find all spots in a 200 meter CIRCLE radius from London.
