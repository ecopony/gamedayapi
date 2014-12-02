package main

import (
	"github.com/ecopony/gamedayapi"
	"log"
	"os"
)

func main() {
	args := os.Args[1:]
	if (len(args) != 2) {
		log.Fatal("Usage: mlbgd teamCode date")
	}

	teamCode := args[0]
	date := args[1]

	game, _ := gamedayapi.GameFor(teamCode, date)
	log.Println(game.GameDataDirectory)
	log.Println(game.Boxscore().GameID)
	log.Println(game.AllInnings().Innings[0].Top.AtBats[0].Pitches[0].Des)
	log.Println(game.HitChart().Hips[0].X)

//  Uncommenting these will execute batch fetch operations. These will be moving to their own commands at some point.
//	gamedayapi.FetchByYearAndTeam(2014, "sea", ExampleOfPullingDownAllFilesForGame)
//	gamedayapi.FetchByYearsAndTeam([]int{2012, 2013, 2014}, "sea", ExampleOfNavigatingAllPitches)
}

func ExampleOfNavigatingAllPitches(game *gamedayapi.Game) {
	log.Println(">>>> " + game.ID + " <<<<")
	for _, inning := range game.AllInnings().Innings {
		for _, atBat := range inning.AtBats() {
			for _, pitch := range atBat.Pitches {
				log.Println("> " + pitch.Des)
			}
		}
	}
}

func ExampleOfPullingDownAllFilesForGame(game *gamedayapi.Game) {
	game.EagerLoad()
}
