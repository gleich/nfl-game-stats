package data

import (
	"encoding/csv"
	"net/http"

	"github.com/gleich/lumber/v2"
)

func Fetch() ([][]string, error) {
	lumber.Info("Fetching raw data from Five Thirty-Eight")

	resp, err := http.Get("https://projects.fivethirtyeight.com/nfl-api/nfl_elo.csv")
	if err != nil {
		return [][]string{}, err
	}
	defer resp.Body.Close()

	records, err := csv.NewReader(resp.Body).ReadAll()
	if err != nil {
		return [][]string{}, err
	}
	records = records[1:]

	lumber.Success("Fetched raw data and did basic parse")
	return records, nil
}
