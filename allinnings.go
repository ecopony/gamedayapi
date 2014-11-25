package gamedayapi

import "encoding/xml"

// AllInnings is the representation of the innings/innings_all.xml file.
// From here, clients can iterate through all innings, at bats, and pitches in a game.
type AllInnings struct {
	XMLName xml.Name `xml:"game"`
	AtBat   string   `xml:"atBat,attr"`
	Innings []Inning `xml:"inning"`
}
