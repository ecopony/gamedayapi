package gamedayapi

import (
	"log"
	"strconv"
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

type FetchFunc func(game *Game)
