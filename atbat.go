package gamedayapi

// AtBat represents each at bat in an inning, and will have all the pitches thrown.
type AtBat struct {
	Num string `xml:"num,attr"`
	B string `xml:"b,attr"`
	S string `xml:"s,attr"`
	O string `xml:"o,attr"`
	StartTFS string `xml:"start_tfs,attr"`
	StartTFSZulu string `xml:"start_tfs_zulu,attr"`
	Batter string `xml:"batter,attr"`
	Stand string `xml:"stand,attr"`
	BHeight string `xml:"b_height,attr"`
	Pitcher string `xml:"pitcher,attr"`
	PThrows string `xml:"p_throws,attr"`
	Des string `xml:"des,attr"`
	DesEs string `xml:"des_es,attr"`
	Event string `xml:"event,attr"`

	Pitches []Pitch `xml:"pitch"`
}
