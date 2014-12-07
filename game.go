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
	"time"
)

// Game is the top-level abstraction, the starting point for clients.
// A game is obtained using the GameFor function.
// From a game, clients can navigate to all data: Innings, Boxscore, etc.
type Game struct {
	AwayAPMP        string `xml:"away_ampm,attr"`
	AwayCode        string `xml:"away_code,attr"`
	AwayLoss        string `xml:"away_loss,attr"`
	AwayTeamCity    string `xml:"away_team_city,attr"`
	AwayTeamID      string `xml:"away_team_id,attr"`
	AwayTeamName    string `xml:"away_team_name,attr"`
	AwayTime        string `xml:"away_time,attr"`
	AwayTimezone    string `xml:"away_time_zone,attr"`
	AwayWin         string `xml:"away_win,attr"`
	CalendarEventID string `xml:"calendar_event_id,attr"`
	HomeAMPM        string `xml:"home_ampm,attr"`
	HomeCode        string `xml:"home_code,attr"`
	HomeLoss        string `xml:"home_loss,attr"`
	HomeTeamCity    string `xml:"home_team_city,attr"`
	HomeTeamID      string `xml:"home_team_id,attr"`
	HomeTeamName    string `xml:"home_team_name,attr"`
	HomeTime        string `xml:"home_time,attr"`
	HomeTimezone    string `xml:"home_time_zone,attr"`
	HomeWin         string `xml:"home_win,attr"`
	ID              string `xml:"id,attr"`
	Gameday         string `xml:"gameday,attr"`
	GamePk          string `xml:"game_pk,attr"`
	GameType        string `xml:"game_type,attr"`
	TimeDate        string `xml:"time_date,attr"`
	Timezone        string `xml:"time_zone,attr"`
	Venue           string `xml:"venue,attr"`

	// GameDataDirectory does not always point to where the files are.
	// Use FetchableDataDirectory for building requests to gameday servers.
	GameDataDirectory string `xml:"game_data_directory,attr"`

	allInnings AllInnings
	boxscore   Boxscore
	hitChart   HitChart
}

// GameFor will return a pointer to a game instance for the team code and date provided.
// This is the place to start for interacting with a game.
// Does not account for doubleheaders. Use GamesFor if doubleheader support is needed.
func GameFor(teamCode string, date time.Time) (*Game, error) {
	epg := EpgFor(date)
	game, err := epg.GameForTeam(teamCode)
	if err != nil {
		return &Game{}, err
	}
	return game, nil
}

// GamesFor will return a collection of pointers to games for the team code and date provided.
// Accounts for doubleheaders. In most cases, the collection will only have one game in it.
func GamesFor(teamCode string, date time.Time) ([]*Game, error) {
	epg := EpgFor(date)
	games, err := epg.GamesForTeam(teamCode)
	if err != nil {
		return games, err
	}
	return games, nil
}

// AllInnings fetches the inning/innings_all.xml file from gameday servers and fills in all the
// structs beneath, all the way down to the pitches.
func (game *Game) AllInnings() *AllInnings {
	if len(game.allInnings.AtBat) == 0 {
		game.load("/inning/inning_all.xml", &game.allInnings)
	}
	return &game.allInnings
}

// Boxscore fetches the boxscore.xml file from the gameday servers and fills in all the structs beneath.
func (game *Game) Boxscore() *Boxscore {
	if len(game.boxscore.GameID) == 0 {
		game.load("/boxscore.xml", &game.boxscore)
	}
	return &game.boxscore
}

// HitChart fetches the inning/inning_hit.xml file from gameday servers
func (game *Game) HitChart() *HitChart {
	if len(game.hitChart.Hips) == 0 {
		game.load("/inning/inning_hit.xml", &game.hitChart)
	}
	return &game.hitChart
}

//func (game *Game) InningScores() *InningScores {}

// EagerLoad will eagerly load all of the files that the library pulls from the MLB gameday servers.
// Otherwise, files are lazily loaded as clients interact with the API.
func (game *Game) EagerLoad() {
	game.AllInnings()
	game.Boxscore()
	game.HitChart()
}

func (game Game) load(fileName string, val interface{}) {
	filePath := game.FetchableDataDirectory() + fileName
	localFilePath := BaseCachePath() + filePath
	if _, err := os.Stat(localFilePath); os.IsNotExist(err) {
		log.Println("Cache miss on " + localFilePath)
		fetchAndCache(filePath, val)
	} else {
		log.Println("Cache hit on " + localFilePath)
		body, _ := ioutil.ReadFile(localFilePath)
		xml.Unmarshal(body, val)
	}
}

// FetchableDataDirectory builds a data directory to the game files using the game's ID.
// GameDataDirectory is not always reliable (see epg on 2013-04-11 vs 2013-04-10 as an example).
func (game Game) FetchableDataDirectory() string {
	idPieces := s.Split(game.ID, "/")
	var buffer bytes.Buffer
	buffer.WriteString(GamedayBasePath)
	buffer.WriteString(fmt.Sprintf("/year_%s/month_%s/day_%s/gid_", idPieces[0], idPieces[1], idPieces[2]))
	buffer.WriteString(game.Gameday)
	return buffer.String()
}

func fetchAndCache(filePath string, val interface{}) {
	urlToFetch := GamedayHostname + filePath
	log.Println("Fetching " + urlToFetch)
	resp, err := http.Get(urlToFetch)
	check(err)
	if resp.StatusCode != 200 {
		log.Fatal(resp.Status + " fetching " + urlToFetch)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	check(err)
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
