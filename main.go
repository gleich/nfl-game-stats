package main

import (
	"github.com/gleich/lumber/v2"
	"github.com/gleich/nfl-game-stats/pkg/data"
)

func main() {
	rawData, err := data.Fetch()
	if err != nil {
		lumber.Fatal(err, "Failed to fetch data")
	}

	games, err := data.Parse(rawData)
	if err != nil {
		lumber.Fatal(err, "Failed to parse data")
	}

	data.Categorize(games)
}
