package gamedayapi

import "encoding/xml"

type Inning struct {
	XMLName			xml.Name	`xml:"inning"`
	AwayTeam		string		`xml:"away_team,attr"`
	HomeTeam		string		`xml:"home_team,attr"`
	Next			string		`xml:"next,attr"`
	Num				string		`xml:"num,attr"`

	Top Top `xml:"top"`
	Bottom Bottom `xml:"bottom"`
}

type Top struct {
	XMLName xml.Name `xml:"top"`
	AtBats []AtBat `xml:"atbat"`
}

type Bottom struct {
	XMLName xml.Name `xml:"bottom"`
	AtBats []AtBat `xml:"atbat"`
}
