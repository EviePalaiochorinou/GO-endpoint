# GO-endpoint

## Question 1
1. Change the website field, so it only contains the domain
2. Count how many spots contain the same domain
3. Return spots which have a domain with a count greater than one
#### The answer to question one can be found in `query.sql`

## Question 2
### Prerequisites
In order for this to run, you need to have a postgres with postgis extention running on localhost on port 5432,
and that you have imported the data from `spots.sql` file.  

### To start the program run

```go
go mod download
go run ./...
```

### To make a request against it run:
`curl http://localhost:10000/spots?latitude=51.5072&longitute=0.1276&radius=200&type=circle`

This will find all spots in a 200 meter CIRCLE radius from London.
