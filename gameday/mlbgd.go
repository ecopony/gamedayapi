package main

import (
	"github.com/ecopony/gamedayapi"
	"log"
	"os"
)

func main() {
	args := os.Args[1:]
	if (len(args) != 2) {
		log.Fatal("Usage: gameday teamCode date")
	}

	teamCode := args[0]
	date := args[1]

	game, _ := gamedayapi.GameFor(teamCode, date)
	log.Println(game.GameDataDirectory)
	log.Println(game.BoxScore().GameId)
	log.Println(game.AllInnings().Innings[0].Top.AtBats[0].Pitches[0].Des)

//	gamedayapi.FetchByYearAndTeam(2014, "sea", func(game *gamedayapi.Game) { log.Println("Do something with game " + game.Id)})
}
