package gamedayapi

type AtBat struct {
	Num				string		`xml:"num,attr"`

	Pitches []Pitch `xml:"pitch"`
}
