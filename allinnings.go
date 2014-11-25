package gamedayapi

import "encoding/xml"

type AllInnings struct {
	XMLName xml.Name `xml:"game"`
	AtBat   string   `xml:"atBat,attr"`
	Innings []Inning `xml:"inning"`
}
