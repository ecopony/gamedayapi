package gamedayapi

import (
	"encoding/xml"
)

type Team struct {
	XMLName xml.Name `xml:"team"`
	TeamType string `xml:"type,attr"`
	Code string `xml:"code,attr"`
	FileCode string `xml:"file_code,attr"`
}
