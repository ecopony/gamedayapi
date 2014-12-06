package gamedayapi

import (
	"log"
	"strconv"
	"sync"
	"time"
)

// FetchByTeamAndYear takes a year and a team code and will roll through all the games for that season.
// The fetchFunc will be passed each game for the year so clients can pull data, compute stats, etc.
func FetchByTeamAndYear(teamCode string, year int, fetchFunc FetchFunc) {
	log.Println("Batchin it in " + strconv.Itoa(year) + " for " + teamCode)
	openingDay, finalDay := OpeningAndFinalDatesForYear(year)
	currentDay := openingDay

	for {
		games, err := GamesFor(teamCode, currentDay)
		if err != nil {
			log.Println(err)
		} else {
			for i := 0; i < len(games); i++ {
				fetchFunc(games[i])
			}
		}

		currentDay = currentDay.Add(time.Hour * 24)

		if currentDay.After(finalDay) {
			break
		}
	}
}

// FetchByTeamAndYears takes a collection of years and a team code and will concurrently roll through all the games for
// the seasons.
// The fetchFunc will be passed each game for the year so clients can pull data, compute stats, etc.
func FetchByTeamAndYears(teamCode string, years []int, fetchFunc FetchFunc) {
	var wg sync.WaitGroup
	for _, year := range years {
		wg.Add(1)
		go func(year int) {
			defer wg.Done()
			FetchByTeamAndYear(teamCode, year, fetchFunc)
		}(year)
	}
	wg.Wait()
}

// FetchFunc is a function passed into fetchers in order to operate on the games fetched.
type FetchFunc func(game *Game)
