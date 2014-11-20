package gamedayapi

import (
	"log"
	"os"
	"net/http"
	"io/ioutil"
	"encoding/xml"
	s "strings"
)

type Game struct {
	AwayAPMP			string		`xml:"away_ampm,attr"`
	AwayCode			string		`xml:"away_code,attr"`
	AwayLoss			string		`xml:"away_loss,attr"`
	AwayTeamCity		string		`xml:"away_team_city,attr"`
	AwayTeamId			string		`xml:"away_team_id,attr"`
	AwayTeamName		string		`xml:"away_team_name,attr"`
	AwayTime			string		`xml:"away_time,attr"`
	AwayTimezone		string		`xml:"away_time_zone,attr"`
	AwayWin				string		`xml:"away_win,attr"`
	HomeAMPM			string		`xml:"home_ampm,attr"`
	HomeCode			string		`xml:"home_code,attr"`
	HomeLoss			string		`xml:"home_loss,attr"`
	HomeTeamCity		string		`xml:"home_team_city,attr"`
	HomeTeamId			string		`xml:"home_team_id,attr"`
	HomeTeamName		string		`xml:"home_team_name,attr"`
	HomeTime			string		`xml:"home_time,attr"`
	HomeTimezone		string		`xml:"home_time_zone,attr"`
	HomeWin				string		`xml:"home_win,attr"`
	Id					string		`xml:"id,attr"`
	GamePk				string		`xml:"game_pk,attr"`
	Timezone			string		`xml:"time_zone,attr"`
	Venue				string		`xml:"venue,attr"`
	GameDataDirectory	string		`xml:"game_data_directory,attr"`

	boxScore BoxScore
}

func GameFor(teamCode string, date string) *Game {
	epg := EpgFor(date)
	game := epg.GameForTeam(teamCode)
	return game
}

func (game *Game) BoxScore() *BoxScore {
	if len(game.boxScore.GameId) == 0 {
		game.load("/boxscore.xml", &game.boxScore)
	}
	return &game.boxScore
}

func (game Game) load(fileName string, val interface{}) {
	filePath := game.GameDataDirectory + fileName
	localFilePath := BaseCachePath() + filePath
	if _, err := os.Stat(localFilePath); os.IsNotExist(err) {
		log.Println("Cache miss on " + localFilePath)
		fetchAndCache(filePath, val)
	} else {
		body, _ := ioutil.ReadFile(localFilePath)
		xml.Unmarshal(body, val)
	}
}

func fetchAndCache(filePath string, val interface{}) {
	log.Println("Fetching " + filePath + " from MLB")
	resp, err := http.Get(GamedayHostname + filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	xml.Unmarshal(body, val)
	cacheFile(filePath, body)
}

func cacheFile(filePath string, body []byte) {
	localCachePath := BaseCachePath() + filePath[0:s.LastIndex(filePath, "/")]
	os.MkdirAll(localCachePath, (os.FileMode)(0775))
	f, err := os.Create(localCachePath + filePath[s.LastIndex(filePath, "/"):])
	f.Write(body)
	check(err)
	defer f.Close()
}
