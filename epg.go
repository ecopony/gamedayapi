package gamedayapi

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

// Epg represents the epg.xml file that is in the root of each day's directory.
// It is essentially the schedule for the day.
type Epg struct {
	Date            string `xml:"date,attr"`
	LastModified    string `xml:"last_modified,attr"`
	DisplayTimeZone string `xml:"display_time_zone,attr"`
	Games           []Game `xml:"game"`
}

// EpgFor returns a pointer to the Epg for the given day.
// The Epg is how the API finds the game directory for a game on a given day.
func EpgFor(date time.Time) *Epg {
	var epg Epg
	cachedFilePath := BaseCachePath() + "/" + strconv.Itoa(date.Year()) + "/"
	cachedFileName := epgCacheFileName(date)

	if _, err := os.Stat(cachedFilePath + cachedFileName); os.IsNotExist(err) {
		log.Println("Fetching epg for " + date.Format("2006-01-02") + " from MLB")

		epgResp, err := http.Get(epgURL(date))
		if err != nil {
			log.Fatal(err)
		}
		defer epgResp.Body.Close()
		epgBody, err := ioutil.ReadAll(epgResp.Body)
		if err != nil {
			log.Fatal(err)
		}
		xml.Unmarshal(epgBody, &epg)
		cacheEpgResponse(cachedFilePath, cachedFileName, epgBody)
	} else {
		body, _ := ioutil.ReadFile(cachedFilePath + cachedFileName)
		xml.Unmarshal(body, &epg)
	}

	return &epg
}

// GameForTeam will return the first game for the given team based on the state of the Epg.
// Does not support doubleheaders. If support for doubleheaders is needed, use GamesForTeam.
func (epg *Epg) GameForTeam(teamCode string) (*Game, error) {
	for _, game := range epg.Games {
		if isGameForTeam(&game, teamCode) {
			return &game, nil
		}
	}
	return &Game{}, fmt.Errorf("[%s] doesn't have a game on [%s]", teamCode, epg.Date)
}

// GamesForTeam will return a collection of the team's games for the day. Supports days where
// the team played in a doubleheader.
func (epg *Epg) GamesForTeam(teamCode string) ([]*Game, error) {
	var games []*Game
	for i := 0; i < len(epg.Games); i++ {
		game := &epg.Games[i]
		if isGameForTeam(game, teamCode) {
			games = append(games, game)
		}
	}
	if len(games) == 0 {
		return games, fmt.Errorf("[%s] doesn't have a game on [%s]", teamCode, epg.Date)
	}
	return games, nil
}

func epgURL(date time.Time) string {
	return dateURL(date) + "/epg.xml"
}

func epgCacheFileName(date time.Time) string {
	return date.Format("2006-01-02") + "-epg.xml"
}

func cacheEpgResponse(path string, filename string, body []byte) {
	os.MkdirAll(path, (os.FileMode)(0775))
	f, err := os.Create(path + filename)
	f.Write(body)
	check(err)
	defer f.Close()
}

func isGameForTeam(game *Game, teamCode string) bool {
	return game.GameType == "R" && (game.HomeCode == teamCode || game.AwayCode == teamCode)
}
