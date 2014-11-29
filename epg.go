package gamedayapi

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	s "strings"
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
func EpgFor(date string) *Epg {
	var epg Epg
	year := s.Split(date, "-")[0]
	cachedFilePath := baseCachePath() + "/" + year + "/"
	cachedFileName := epgCacheFileName(date)

	if _, err := os.Stat(cachedFilePath + cachedFileName); os.IsNotExist(err) {
		log.Println("Fetching epg for " + date + " from MLB")

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

// GameForTeam will find the game for the given team based on the state of the Epg.
// Does not yet support doubleheaders, for which it might need to return a collection of games.
func (epg *Epg) GameForTeam(teamCode string) (*Game, error) {
	for _, game := range epg.Games {
		if game.GameType == "R" && (game.HomeCode == teamCode || game.AwayCode == teamCode) {
			return &game, nil
		}
	}
	return &Game{}, fmt.Errorf("[%s] doesn't have a game on [%s]", teamCode, epg.Date)
}

func epgURL(date string) string {
	var buffer bytes.Buffer
	buffer.WriteString(dateURL(date))
	buffer.WriteString("/epg.xml")
	return buffer.String()
}

func epgCacheFileName(date string) string {
	return date + "-" + "epg.xml"
}

func cacheEpgResponse(path string, filename string, body []byte) {
	os.MkdirAll(path, (os.FileMode)(0775))
	f, err := os.Create(path + filename)
	f.Write(body)
	check(err)
	defer f.Close()
}
