package main

import "fmt"
import "log"
import "io/ioutil"
import "net/http"
import "encoding/xml"

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
   fmt.Println(resp.Status)
   fmt.Println(string(body))
   fmt.Println(game)
}
