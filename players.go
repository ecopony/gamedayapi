package gamedayapi

import "encoding/xml"

// Players is the representation of the players.xml file.
type Players struct {
	XMLName xml.Name `xml:"game"`
	Date    string   `xml:"date,attr"`
	Teams   []Team   `xml:"team"`
}

// Team is the team fragment in the players.xml file.
type Team struct {
	XMLName xml.Name `xml:"team"`
	Type    string   `xml:"type,attr"`
	ID      string   `xml:"id,attr"`
	Name    string   `xml:"name,attr"`
	Players []Player `xml:"player"`
}

// Player is the player fragment in the players.xml file.
type Player struct {
	XMLName          xml.Name `xml:"player"`
	ID               string   `xml:"id,attr"`
	First            string   `xml:"first,attr"`
	Last             string   `xml:"last,attr"`
	Num              string   `xml:"num,attr"`
	Boxname          string   `xml:"boxname,attr"`
	Rl               string   `xml:"rl,attr"`
	Bats             string   `xml:"bats,attr"`
	Position         string   `xml:"position,attr"`
	CurrentPosition  string   `xml:"current_position,attr"`
	Status           string   `xml:"status,attr"`
	TeamAbbrev       string   `xml:"team_abbrev,attr"`
	TeamID           string   `xml:"team_id,attr"`
	ParentTeamAbbrev string   `xml:"parent_team_abbrev,attr"`
	ParentTeamID     string   `xml:"parent_team_id,attr"`
	BatOrder         string   `xml:"bat_order,attr"`
	GamePosition     string   `xml:"game_position,attr"`
	Avg              string   `xml:"avg,attr"`
	HR               string   `xml:"hr,attr"`
	RBI              string   `xml:"rbi,attr"`
	Wins             string   `xml:"wins,attr"`
	Losses           string   `xml:"losses,attr"`
	ERA              string   `xml:"era,attr"`
}
