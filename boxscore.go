package gamedayapi

import "encoding/xml"

// Boxscore is the representation of the boxscore.xml file.
type Boxscore struct {
	XMLName       xml.Name `xml:"boxscore"`
	AwayFName     string   `xml:"away_fname,attr"`
	AwayID        string   `xml:"away_id,attr"`
	AwayLoss      string   `xml:"away_loss,attr"`
	AwaySName     string   `xml:"away_sname,attr"`
	AwayTeamCode  string   `xml:"away_team_code,attr"`
	AwayWins      string   `xml:"away_wins,attr"`
	Date          string   `xml:"date,attr"`
	GameID        string   `xml:"game_id,attr"`
	GamePk        string   `xml:"game_pk,attr"`
	HomeFName     string   `xml:"home_fname,attr"`
	HomeID        string   `xml:"home_id,attr"`
	HomeLoss      string   `xml:"home_loss,attr"`
	HomeSName     string   `xml:"home_sname,attr"`
	HomeSportCode string   `xml:"home_sport_code,attr"`
	HomeTeamCode  string   `xml:"home_team_code,attr"`
	HomeWins      string   `xml:"home_wins,attr"`
	StatusInd     string   `xml:"status_ind,attr"`
	VenueID       string   `xml:"venue_id,attr"`
	VenueName     string   `xml:"venue_name,attr"`

	Linescores []Linescore `xml:"linescore"`
}

// Linescore represents the linescore under the boxscore, not the individual linescore.xml file.
type Linescore struct {
	XMLName        xml.Name `xml:"linescore"`
	AwayInningRuns string   `xml:"away_inning_runs,attr"`
	HomeInningRuns string   `xml:"home_inning_runs,attr"`
}
