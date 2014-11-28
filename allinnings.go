package gamedayapi

import "encoding/xml"

// AllInnings is the representation of the innings/innings_all.xml file.
// From here, clients can iterate through all innings, at bats, and pitches in a game.
type AllInnings struct {
	XMLName xml.Name `xml:"game"`
	AtBat   string   `xml:"atBat,attr"`
	Deck    string   `xml:"deck,attr"`
	Hole    string   `xml:"hole,attr"`
	Ind     string   `xml:"ind,attr"`
	Innings []Inning `xml:"inning"`
}
