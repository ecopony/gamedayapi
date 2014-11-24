package gamedayapi

import (
	"log"
	"time"
	"strconv"
)

func FetchByYearAndTeam(year int, team string, fetchFunc FetchFunc) {
	log.Println("Batchin it in " + strconv.Itoa(year) + " for " + team)
	openingDay, finalDay := OpeningAndFinalDatesForYear(year)
	currentDay := openingDay

	log.Println(openingDay)
	log.Println(finalDay)
	currentDay = currentDay.Add(time.Hour*24)
	log.Println(currentDay)
}

type FetchFunc func(game *Game)
