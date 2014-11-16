package gamedayapi

import (
	"log"
	"os"
	"io/ioutil"
	"net/http"
	"encoding/xml"
)

type Game struct {
	XMLName xml.Name `xml:"game"`
	GameType string `xml:"type,attr"`
	LocalGameTime string `xml:"local_game_time,attr"`
	Teams []Team `xml:"team"`
	Stadium Stadium `xml:"stadium"`
}

func GameFor(teamCode string, date string) *Game {
	epg := EpgFor(date)
	gid := epg.GidForTeam(teamCode)
	game := fetchGame(gid)
	return game
}

func fetchGame(gid *Gid) *Game {
	log.Println("Fetching game " + gid.String())
	var game Game
	gameFileName := "game.xml"
	cachedFileName := gid.CachePath() + cacheFileName(gid, gameFileName)

	if _, err := os.Stat(cachedFileName); os.IsNotExist(err) {
		log.Println("No cache hit - go get it")

		resp, err := http.Get(gameDirectoryUrl(gid) + gameFileName)
		if err != nil {
			log.Fatal(err)
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}

		xml.Unmarshal(body, &game)
		log.Println(resp.Status)
		log.Println(string(body))
		cacheResponse(gid, gameFileName, body)
	} else {
		log.Println("Cache hit - load it up")
		body, _ := ioutil.ReadFile(cachedFileName)
		log.Println(string(body))
		xml.Unmarshal(body, &game)
	}

	return &game
}
