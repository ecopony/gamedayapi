package gamedayapi

import (
	"log"
	"bytes"
	"io/ioutil"
	"net/http"
	"encoding/xml"
	"os"
	s "strings"
)

type Epg struct {
	Date string `xml:"id,attr"`
	LastModified string `xml:"last_modified,attr"`
	DisplayTimeZone string `xml:"display_time_zone,attr"`
	EpgGames []EpgGame `xml:"game"`
}

func EpgFor(date string) *Epg {
	var epg Epg
	log.Println("Fetching epg for " + date)
	year := s.Split(date, "-")[0]
	cachedFilePath := BaseCachePath() + year + "/"
	cachedFileName := EpgCacheFileName(date)

	if _, err := os.Stat(cachedFilePath + cachedFileName); os.IsNotExist(err) {
		log.Println("No epg cache hit - go get it")
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
		log.Println("EPG cache hit - load it up")
		body, _ := ioutil.ReadFile(cachedFilePath + cachedFileName)
		xml.Unmarshal(body, &epg)
	}

	return &epg
}

/*
gids look like: gid_2014_07_22_nynmlb_seamlb_1

Doesn't yet handle doubleheader days. It'll just return the first match it finds for the team.
 */
func (e *Epg) GidForTeam(teamCode string) *Gid {
	for _, game := range e.EpgGames {
		if s.Contains(game.Gameday, s.Join([]string{"_", teamCode, "mlb_"}, "")) {
			gamedayPieces := s.Split(game.Gameday, "_")
			return &Gid{gamedayPieces[0], gamedayPieces[1], gamedayPieces[2], gamedayPieces[3], gamedayPieces[4], gamedayPieces[5]}
		}
	}
	return &Gid{} // this should be an error
}

type EpgGame struct {
	CalendarEventId string `xml:"calendar_event_id,attr"`
	Start string `xml:"start,attr"`
	Id string `xml:"id,attr"`
	Gameday string `xml:"gameday,attr"`
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

