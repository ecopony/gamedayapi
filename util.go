package gamedayapi

import (
	"os/user"
	"log"
	"bytes"
	s "strings"
)

const (
	GamedayHostname = "http://gd2.mlb.com"
	GamedayBaseUrl = "http://gd2.mlb.com/components/game/mlb"

)

func datePath(date string) string {
	// firx this to be date parsing, validating
	datePieces := s.Split(date, "-")
	var buffer bytes.Buffer
	buffer.WriteString("/year_")
	buffer.WriteString(datePieces[0])
	buffer.WriteString("/month_")
	buffer.WriteString(datePieces[1])
	buffer.WriteString("/day_")
	buffer.WriteString(datePieces[2])
	return buffer.String()
}

func dateUrl(date string) string {
	var buffer bytes.Buffer
	buffer.WriteString(GamedayBaseUrl)
	buffer.WriteString(datePath(date))
	return buffer.String()
}

func homeDir() string {
	usr, err := user.Current()
	if err != nil {
		log.Fatal( err )
	}
	return usr.HomeDir
}

func BaseCachePath() string {
	return homeDir() + "/go-gameday-cache"
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
