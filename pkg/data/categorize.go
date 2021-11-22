package data

import "fmt"

func Categorize(games []Game) map[int]map[string]int {
	categorizedData := map[int]map[string]int{}

	for _, game := range games {
		if categorizedData[game.Season] == nil {
			categorizedData[game.Season] = map[string]int{}
		}
		if game.Difference == 0 {
			categorizedData[game.Season]["0"]++
		}

		for i := 5; i < 50; i += 5 {
			if i-5 < game.Difference && game.Difference <= i {
				categorizedData[game.Season][fmt.Sprintf("%v < %v \\leq %v", i-5, game.Difference, i)]++
			}
		}
	}

	fmt.Println(categorizedData)

	return categorizedData
}
