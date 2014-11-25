package gamedayapi

import (
	"log"
	"bytes"
	"io/ioutil"
	"net/http"
	"encoding/xml"
	"os"
	s "strings"
	"fmt"
)

type Epg struct {
	Date string `xml:"date,attr"`
	LastModified string `xml:"last_modified,attr"`
	DisplayTimeZone string `xml:"display_time_zone,attr"`
	Games []Game `xml:"game"`
}

func EpgFor(date string) *Epg {
	var epg Epg
	year := s.Split(date, "-")[0]
	cachedFilePath := BaseCachePath() + "/" + year + "/"
	cachedFileName := EpgCacheFileName(date)

	if _, err := os.Stat(cachedFilePath + cachedFileName); os.IsNotExist(err) {
		log.Println("Fetching epg for " + date + " from MLB")

		epgResp, err := http.Get(EpgUrl(date))
		if err != nil {
			log.Fatal(err)
		}
		defer epgResp.Body.Close()
		epgBody, err := ioutil.ReadAll(epgResp.Body)
		if err != nil {
			log.Fatal(err)
		}
		xml.Unmarshal(epgBody, &epg)
		CacheEpgResponse(cachedFilePath, cachedFileName, epgBody)
	} else {
		body, _ := ioutil.ReadFile(cachedFilePath + cachedFileName)
		xml.Unmarshal(body, &epg)
	}

	return &epg
}

func (epg *Epg) GameForTeam(teamCode string) (*Game, error) {
	for _, game := range epg.Games {
		if (game.GameType == "R" && (game.HomeCode == teamCode || game.AwayCode == teamCode)) {
			return &game, nil
		}
	}
	return &Game{}, fmt.Errorf("[%s] doesn't have a game on [%s]", teamCode, epg.Date)
}

func EpgUrl(date string) string {
	var buffer bytes.Buffer
	buffer.WriteString(dateUrl(date))
	buffer.WriteString("/epg.xml")
	return buffer.String()
}

func EpgCacheFileName(date string) string {
	return date + "-" + "epg.xml"
}

func CacheEpgResponse(path string, filename string, body []byte) {
	os.MkdirAll(path, (os.FileMode)(0775))
	f, err := os.Create(path + filename)
	f.Write(body)
	check(err)
	defer f.Close()
}

