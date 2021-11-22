package main

import (
	"fmt"

	"github.com/gleich/lumber/v2"
	"github.com/gleich/nfl-game-stats/pkg/data"
)

func main() {
	data, err := data.Fetch()
	if err != nil {
		lumber.Fatal(err, "Failed to fetch data")
	}
	fmt.Println(data[0])
}
