package gamedayapi

import (
	"encoding/xml"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	s "strings"
)

// Game is the top-level abstraction, the starting point for clients.
// A game is obtained using the GameFor function.
// From a game, clients can navigate to all data: Innings, Boxscore, etc.
type Game struct {
	AwayAPMP          string `xml:"away_ampm,attr"`
	AwayCode          string `xml:"away_code,attr"`
	AwayLoss          string `xml:"away_loss,attr"`
	AwayTeamCity      string `xml:"away_team_city,attr"`
	AwayTeamID        string `xml:"away_team_id,attr"`
	AwayTeamName      string `xml:"away_team_name,attr"`
	AwayTime          string `xml:"away_time,attr"`
	AwayTimezone      string `xml:"away_time_zone,attr"`
	AwayWin           string `xml:"away_win,attr"`
	HomeAMPM          string `xml:"home_ampm,attr"`
	HomeCode          string `xml:"home_code,attr"`
	HomeLoss          string `xml:"home_loss,attr"`
	HomeTeamCity      string `xml:"home_team_city,attr"`
	HomeTeamID        string `xml:"home_team_id,attr"`
	HomeTeamName      string `xml:"home_team_name,attr"`
	HomeTime          string `xml:"home_time,attr"`
	HomeTimezone      string `xml:"home_time_zone,attr"`
	HomeWin           string `xml:"home_win,attr"`
	ID                string `xml:"id,attr"`
	GameDataDirectory string `xml:"game_data_directory,attr"`
	GamePk            string `xml:"game_pk,attr"`
	GameType          string `xml:"game_type,attr"`
	Timezone          string `xml:"time_zone,attr"`
	Venue             string `xml:"venue,attr"`

	boxScore   BoxScore
	allInnings AllInnings
}

func GameFor(teamCode string, date string) (*Game, error) {
	epg := EpgFor(date)
	game, err := epg.GameForTeam(teamCode)
	if err != nil {
		return &Game{}, err
	}
	return game, nil
}

func (game *Game) BoxScore() *BoxScore {
	if len(game.boxScore.GameID) == 0 {
		game.load("/boxscore.xml", &game.boxScore)
	}
	return &game.boxScore
}

// AllInnings fetches the inning/innings_all.xml file from gameday servers and hydrates all the structs beneath, all the
// way down to the pitches
func (game *Game) AllInnings() *AllInnings {
	if len(game.allInnings.AtBat) == 0 {
		game.load("/inning/inning_all.xml", &game.allInnings)
	}
	return &game.allInnings
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
