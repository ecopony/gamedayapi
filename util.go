package gamedayapi

import (
	"os/user"
	"log"
	"bytes"
	"os"
	s "strings"
)

const (
	BaseUrl = "http://gd2.mlb.com/components/game/mlb/"
)

func datePath(date string) string {
	// firx this to be date parsing, validating
	datePieces := s.Split(date, "-")
	var buffer bytes.Buffer
	buffer.WriteString("year_")
	buffer.WriteString(datePieces[0])
	buffer.WriteString("/month_")
	buffer.WriteString(datePieces[1])
	buffer.WriteString("/day_")
	buffer.WriteString(datePieces[2])
	return buffer.String()
}

func cacheResponse(gid *Gid, filename string, body []byte) {
	os.MkdirAll(gid.CachePath(), (os.FileMode)(0775))
	f, err := os.Create(gid.CachePath() + cacheFileName(gid, filename))
	f.Write(body)
	check(err)
	defer f.Close()
}

func dateUrl(date string) string {
	var buffer bytes.Buffer
	buffer.WriteString(BaseUrl)
	buffer.WriteString(datePath(date))
	return buffer.String()
}

func gameDirectoryUrl(gid *Gid) string {
	var buffer bytes.Buffer
	buffer.WriteString(BaseUrl)
	buffer.WriteString(gid.DatePath())
	buffer.WriteString("/")
	buffer.WriteString(gid.String())
	buffer.WriteString("/")
	return buffer.String()
}

func homeDir() string {
	usr, err := user.Current()
	if err != nil {
		log.Fatal( err )
	}
	return usr.HomeDir
}

func cacheFileName(gid *Gid, filename string) string {
	return gid.String() + "-" + filename
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
