package gamedayapi

import "encoding/xml"

// HitChart is the representation of the innings/inning_hit.xml file.
type HitChart struct {
	XMLName xml.Name `xml:"hitchart"`
	Hips    []Hip    `xml:"hip"`
}
