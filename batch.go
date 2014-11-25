package gamedayapi

import (
	"log"
	"time"
	"strconv"
)

func FetchByYearAndTeam(year int, teamCode string, fetchFunc FetchFunc) {
	log.Println("Batchin it in " + strconv.Itoa(year) + " for " + teamCode)
	openingDay, finalDay := OpeningAndFinalDatesForYear(year)
	currentDay := openingDay

	for {
		game := GameFor(teamCode, currentDay.Format("2006-01-02"))
		log.Println(game.GameDataDirectory)

		currentDay = currentDay.Add(time.Hour*24)
		if currentDay.After(finalDay) {
			break
		}
	}
}

type FetchFunc func(game *Game)
