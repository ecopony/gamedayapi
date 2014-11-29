package gamedayapi

import "encoding/xml"

// Hip represents the fragments in the HitChart
type Hip struct {
	XMLName xml.Name `xml:"hip"`
	Des     string   `xml:"des,attr"`
	X       string   `xml:"x,attr"`
	Y       string   `xml:"y,attr"`
	Batter  string   `xml:"batter,attr"`
	Pitcher string   `xml:"pitcher,attr"`
	Type    string   `xml:"type,attr"`
	Team    string   `xml:"team,attr"`
	Inning  string   `xml:"inning,attr"`
}
