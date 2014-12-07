package main

import (
	"fmt"
	"github.com/ecopony/gamedayapi"
	"os"
	"strconv"
	"time"
)

var validCommands = map[string]bool{
	"game": true,
	"games-for-team-and-year":  true,
	"games-for-team-and-years": true,
}

func main() {
	args := os.Args[1:]

	if len(args) <= 2 {
		fmt.Println("Usage: mlbgd <command> <team code> <date|year(s)>")
		os.Exit(1)
	}

	command := args[0]
	if !isCommandValid(command) {
		fmt.Println(fmt.Sprintf("%s is not a valid command. Valid commands:", command))

		for k := range validCommands {
			fmt.Println(fmt.Sprintf("\t%s", k))
		}

		os.Exit(1)
	}

	teamCode := args[1]

	if command == "game" {
		date, err := time.Parse("2006-01-02", args[2])
		if err != nil {
			fmt.Println("Date must be in the format 2006-01-02")
			os.Exit(1)
		}
		game, err := gamedayapi.GameFor(teamCode, date)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		game.EagerLoad()
		fmt.Println("Game files saved to " + gamedayapi.BaseCachePath() + game.GameDataDirectory)
	} else {
		yearArgs := args[2:]
		var years []int
		for i := 0; i < len(yearArgs); i++ {
			year, err := strconv.Atoi(yearArgs[i])
			if err != nil {
				fmt.Println("Year is not valid")
			}
			years = append(years, year)
		}
		gamedayapi.FetchByTeamAndYears(teamCode, years, eagerLoadGame)
	}
}

func isCommandValid(command string) bool {
	return validCommands[command]
}

func eagerLoadGame(game *gamedayapi.Game) {
	game.EagerLoad()
	fmt.Println("Game files saved to " + gamedayapi.BaseCachePath() + game.GameDataDirectory)
}
