package main

import (
	"database/sql"
)

type Spot struct {
	ID      string `json:"ID"`
	Name    string `json:"Name"`
	Website *string `json:"Website"`
	Coordinates string `json:"Coordinates"`
	Description *string `json:"Description"`
	Rating      float32 `json:"Rating"`
}

func findSpots(db *sql.DB, latitude float32, longitude float32, radius float32, radiusType string) ([]Spot, error) {
	query := `SELECT *
	FROM "MY_TABLE"
	WHERE ST_DWithin(coordinates, ST_SetSRID(ST_MakePoint($1, $2), 4326), $3)`
	rows, err := db.Query(query, latitude, longitude, radius)
	
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var spots []Spot
	for rows.Next() {
		var spot Spot
		if err := rows.Scan(&spot.ID, &spot.Name, &spot.Website, &spot.Coordinates,
			&spot.Description, &spot.Rating); err != nil {
			return spots, err
		}
		spots = append(spots, spot)
	}
	if err = rows.Err(); err != nil {
		return spots, err
	}
	return spots, nil
}

