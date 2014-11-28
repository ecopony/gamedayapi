package gamedayapi

import (
	"encoding/xml"
)

// Inning represents the inning structure in the innings/innings_all.xml file.
type Inning struct {
	XMLName  xml.Name `xml:"inning"`
	AwayTeam string   `xml:"away_team,attr"`
	HomeTeam string   `xml:"home_team,attr"`
	Next     string   `xml:"next,attr"`
	Num      string   `xml:"num,attr"`

	Top    Top    `xml:"top"`
	Bottom Bottom `xml:"bottom"`
}

// Top corresponds to the top half of an inning.
type Top struct {
	XMLName xml.Name `xml:"top"`
	AtBats  []AtBat  `xml:"atbat"`
}

// Bottom corresponds to the bottom half of an inning.
type Bottom struct {
	XMLName xml.Name `xml:"bottom"`
	AtBats  []AtBat  `xml:"atbat"`
}

// AtBats returns all at bats from the inning level.
// A convenience method to save clients from having to deal with the Top and Bottom halves.
func (inning *Inning) AtBats() []AtBat {
	return append(inning.Top.AtBats, inning.Bottom.AtBats...)
}
