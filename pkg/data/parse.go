package data

import (
	"fmt"
	"math"
	"strconv"

	"github.com/gleich/lumber/v2"
)

type Game struct {
	Season     int
	Difference int
}

func Parse(data [][]string) ([]Game, error) {
	games := []Game{}
	for i, row := range data {
		if len(row) != 33 {
			return []Game{}, fmt.Errorf("Wrong number of columns in row #%v", i)
		}

		if row[28] == "" || row[29] == "" {
			continue
		}

		season, err := strconv.Atoi(row[1])
		if err != nil {
			return []Game{}, err
		}

		score1, err := strconv.Atoi(row[28])
		if err != nil {
			return []Game{}, err
		}

		score2, err := strconv.Atoi(row[29])
		if err != nil {
			return []Game{}, err
		}
		games = append(games, Game{
			Season:     season,
			Difference: int(math.Abs(float64(score1 - score2))),
		})
	}
	lumber.Success("Parsed data")
	return games, nil
}
