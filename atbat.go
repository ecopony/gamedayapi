package gamedayapi

// AtBat represents each at bat in an inning, and will have all the pitches thrown.
type AtBat struct {
	Num string `xml:"num,attr"`

	Pitches []Pitch `xml:"pitch"`
}
