package gamedayapi

import (
	"os/user"
	"log"
	"bytes"
	"os"
	"io/ioutil"
	"net/http"
	"encoding/xml"
	s "strings"
)

type Game struct {
	XMLName xml.Name `xml:"game"`
	GameType string `xml:"type,attr"`
	LocalGameTime string `xml:"local_game_time,attr"`
	Teams []Team `xml:"team"`
	Stadium Stadium `xml:"stadium"`
}

func Init(teamCode string, date string) {
	log.Println("Fetching game for " + teamCode + " on " + date)

	epgResp, err := http.Get(epgUrl(date))
	if err != nil {
		log.Fatal(err)
	}
	defer epgResp.Body.Close()
	epgBody, err := ioutil.ReadAll(epgResp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var epg Epg
	xml.Unmarshal(epgBody, &epg)
	gid := epg.GidForTeam(teamCode)

	game := fetchGame(&gid)
	log.Println(game)
}

type Team struct {
	XMLName xml.Name `xml:"team"`
	TeamType string `xml:"type,attr"`
	Code string `xml:"code,attr"`
	FileCode string `xml:"file_code,attr"`
}

type Stadium struct {
	XMLName xml.Name `xml:"stadium"`
	Id string `xml:"id,attr"`
	Name string `xml:"name,attr"`
}

func fetchGame(gid *Gid) Game {
	var game Game
	gameFileName := "game.xml"
	cachedFileName := cachePath(gid) + cacheFileName(gid, gameFileName)

	if _, err := os.Stat(cachedFileName); os.IsNotExist(err) {
		log.Println("No cache hit - go get it")

		resp, err := http.Get(gameUrl(gid))
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

	return game
}

func cacheResponse(gid *Gid, filename string, body []byte) {
	cachePath := cachePath(gid)
	os.MkdirAll(cachePath, (os.FileMode)(0775))
	f, err := os.Create(cachePath + cacheFileName(gid, filename))
	f.Write(body)
	check(err)
	defer f.Close()
}

func baseUrl() string {
	return "http://gd2.mlb.com/components/game/mlb/"
}

func dateUrl(date string) string {
	var buffer bytes.Buffer
	buffer.WriteString(baseUrl())
	buffer.WriteString(datePath(date))
	return buffer.String()
}

func epgUrl(date string) string {
	var buffer bytes.Buffer
	buffer.WriteString(dateUrl(date))
	buffer.WriteString("/epg.xml")
	return buffer.String()
}

func gameDirectoryUrl(gid *Gid) string {
	var buffer bytes.Buffer
	buffer.WriteString(baseUrl())
	buffer.WriteString(gid.DatePath())
	buffer.WriteString("/")
	buffer.WriteString(gid.String())
	buffer.WriteString("/")
	return buffer.String()
}

func gameUrl(gid *Gid) string {
	return gameDirectoryUrl(gid) + "game.xml"
}

func datePath(date string) string {
	// firx this to be date parsing, validating
	datePieces := s.Split(date, "-")
	var buffer bytes.Buffer
	buffer.WriteString("year_")
	buffer.WriteString(datePieces[0])
	buffer.WriteString("/month_")
	buffer.WriteString(datePieces[1])
	buffer.WriteString("/day_")
	buffer.WriteString(datePieces[2])
	return buffer.String()
}

func homeDir() string {
	usr, err := user.Current()
	if err != nil {
		log.Fatal( err )
	}
	return usr.HomeDir
}

func cachePath(gid *Gid) string {
	return homeDir() + "/go-gameday-cache/" + gid.Year + "/"
}

func cacheFileName(gid *Gid, filename string) string {
	return gid.String() + "-" + filename
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
