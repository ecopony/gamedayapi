package main

import (
	_ "github.com/lib/pq"
	"log"
	"os"
	"io/ioutil"
	"net/http"
	"encoding/xml"
	"database/sql"
)

type Game struct{
	XMLName xml.Name `xml:"game"`
	GameType string `xml:"type,attr"`
	LocalGameTime string `xml:"local_game_time,attr"`
	Teams []Team `xml:"team"`
	Stadium Stadium `xml:"stadium"`
}

type Team struct{
	XMLName xml.Name `xml:"team"`
	TeamType string `xml:"type,attr"`
	Code string `xml:"code,attr"`
	FileCode string `xml:"file_code,attr"`
}

type Stadium struct{
	XMLName xml.Name `xml:"stadium"`
	Id string `xml:"id,attr"`
	Name string `xml:"name,attr"`
}

func main() {
	args := os.Args[1:]
	if (len(args) != 2) {
		log.Fatal("Usage: gameday teamCode date")
	}

	teamCode := args[0]
	date := args[1]

	log.Println(teamCode)
	log.Println(date)

	resp, err := http.Get("http://gd2.mlb.com/components/game/mlb/year_2014/month_07/day_06/gid_2014_07_06_seamlb_chamlb_1/game.xml")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var game Game
	xml.Unmarshal(body, &game)
	log.Println(resp.Status)
	log.Println(string(body))
	log.Println(game)

	/*
	Assumes a pg database exists named go-gameday, a role that can access it.
	Assumes a table called pitches with a character column called code.
	 */
	db, err := sql.Open("postgres", "user=go-gameday dbname=go-gameday sslmode=disable")
//	issues := db.Ping()
//	log.Println(issue)

	if err != nil {
		log.Fatal(err)
	}

	rows, err :=  db.Query("SELECT code FROM pitches")

	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	for rows.Next() {
		var code string
		err = rows.Scan(&code)
		log.Println(code)
	}
}
