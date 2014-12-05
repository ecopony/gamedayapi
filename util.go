package gamedayapi

import (
	"log"
	"os/user"
	"time"
	"fmt"
)

const (
	// GamedayHostname is the hostname of the MLB gameday site
	GamedayHostname = "http://gd2.mlb.com"

	// GamedayBaseURL is the base URL of the MLB gameday files
	GamedayBaseURL = "http://gd2.mlb.com/components/game/mlb"

	// GamedayBasePath is the base path of the MLB gameday files
	GamedayBasePath = "/components/game/mlb"
)

func datePath(date time.Time) string {
	return fmt.Sprintf("/year_%02d/month_%02d/day_%02d", date.Year(), date.Month(), date.Day())
}

func dateURL(date time.Time) string {
	return GamedayBaseURL + datePath(date)
}

func homeDir() string {
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	return usr.HomeDir
}

func baseCachePath() string {
	return homeDir() + "/go-gameday-cache"
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
