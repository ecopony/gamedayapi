package gamedayapi

import (
	"fmt"
	"log"
	"os/user"
	"time"
)

const (
	// GamedayHostname is the hostname of the MLB gameday site
	GamedayHostname = "http://gd2.mlb.com"

	// GamedayBaseURL is the base URL of the MLB gameday files
	GamedayBaseURL = "http://gd2.mlb.com/components/game/mlb"

	// GamedayBasePath is the base path of the MLB gameday files
	GamedayBasePath = "/components/game/mlb"
)

// BaseCachePath returns the local directory where gameday files are cached.
func BaseCachePath() string {
	return homeDir() + "/go-gameday-cache"
}

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

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func appendIfMissing(slice []string, i string) []string {
	for _, ele := range slice {
		if ele == i {
			return slice
		}
	}
	return append(slice, i)
}
