package gamedayapi

import (
	"log"
	"strconv"
	"sync"
	"time"
)

// FetchByYearAndTeam takes a year and a team code and will roll through all the games for that season.
// The fetchFunc will be passed each game for the year so clients can pull data, compute stats, etc.
func FetchByYearAndTeam(year int, teamCode string, fetchFunc FetchFunc) {
	log.Println("Batchin it in " + strconv.Itoa(year) + " for " + teamCode)
	openingDay, finalDay := OpeningAndFinalDatesForYear(year)
	currentDay := openingDay

	for {
		game, err := GameFor(teamCode, currentDay.Format("2006-01-02"))
		if err != nil {
			log.Println(err)
		} else {
			fetchFunc(game)
		}

		currentDay = currentDay.Add(time.Hour * 24)
		if currentDay.After(finalDay) {
			break
		}
	}
}

// FetchByYearsAndTeam takes a collection of years and a team code and will concurrently roll through all the games for
// the seasons.
// The fetchFunc will be passed each game for the year so clients can pull data, compute stats, etc.
func FetchByYearsAndTeam(years []int, teamCode string, fetchFunc FetchFunc) {
	var wg sync.WaitGroup
	for _, year := range years {
		wg.Add(1)
		go func(year int) {
			defer wg.Done()
			FetchByYearAndTeam(year, teamCode, fetchFunc)
		}(year)
	}
	wg.Wait()
}

// FetchFunc is a function passed into fetchers in order to operate on the games fetched.
type FetchFunc func(game *Game)
