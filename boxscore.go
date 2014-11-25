package gamedayapi

import "encoding/xml"

type BoxScore struct {
	XMLName xml.Name `xml:"boxscore"`
	GameID  string   `xml:"game_id,attr"`
}
