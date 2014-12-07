package main

import (
	"github.com/ecopony/gamedayapi"
	"log"
	"os"
	"time"
	"fmt"
)

var validCommands = map[string]bool {
	"game": true,
//	"games-for-team-and-year": true,
//	"games-for-team-and-years": true,
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

		for k, _ := range validCommands {
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
		game, _ := gamedayapi.GameFor(teamCode, date)
		game.EagerLoad()
		fmt.Println("Game files saved to " + gamedayapi.BaseCachePath() + game.GameDataDirectory)
	}
}

func isCommandValid(command string) bool {
	return validCommands[command]
}
