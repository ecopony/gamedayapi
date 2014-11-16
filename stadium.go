package gamedayapi

import (
	"encoding/xml"
)

type Stadium struct {
	XMLName xml.Name `xml:"stadium"`
	Id string `xml:"id,attr"`
	Name string `xml:"name,attr"`
}
