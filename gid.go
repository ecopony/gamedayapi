package gamedayapi

import (
	"bytes"
)

type Gid struct {
	Year string
	Month string
	Day string
	Away string
	Home string
	GameNumber string
}

func (gid Gid) String() string {
	var buffer bytes.Buffer
	buffer.WriteString("gid_")
	buffer.WriteString(gid.Year)
	buffer.WriteString("_")
	buffer.WriteString(gid.Month)
	buffer.WriteString("_")
	buffer.WriteString(gid.Day)
	buffer.WriteString("_")
	buffer.WriteString(gid.Away)
	buffer.WriteString("_")
	buffer.WriteString(gid.Home)
	buffer.WriteString("_")
	buffer.WriteString(gid.GameNumber)
	return buffer.String()
}

func (gid Gid) DatePath() string {
	var buffer bytes.Buffer
	buffer.WriteString("year_")
	buffer.WriteString(gid.Year)
	buffer.WriteString("/month_")
	buffer.WriteString(gid.Month)
	buffer.WriteString("/day_")
	buffer.WriteString(gid.Day)
	return buffer.String()
}
