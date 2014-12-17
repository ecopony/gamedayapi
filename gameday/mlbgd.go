package main

import (
	"fmt"
	"github.com/ecopony/gamedayapi"
	"os"
	"strconv"
	"time"
)

var commands = map[string]func([]string) {}

func main() {
	initializeCommands()
	args := os.Args[1:]
	command := args[0]

	if function, ok := commands[command]; ok {
		function(args)
	} else {
		fmt.Println(fmt.Sprintf("%s is not a valid command. Valid commands:", command))
		printValidCommands()
		os.Exit(1)
	}
}

func initializeCommands() {
	commands["game"] = game
	commands["games-for-team-and-year"] = gameForTeamAndYears
	commands["games-for-year"] = gamesForYear
	commands["help"] = help
	commands["valid-teams-for-year"] = validTeamsForYear
}

func eagerLoadGame(game *gamedayapi.Game) {
	game.EagerLoad()
	fmt.Println("Game files saved to " + gamedayapi.BaseCachePath() + game.GameDataDirectory)
}

func printUsage() {
	fmt.Println("Usage: mlbgd <command> [<team code>] [<date|year(s)>]")
}

func printValidCommands() {
	printUsage()
	fmt.Println("Valid commands:")
	for validCommand, _ := range commands {
		fmt.Println(fmt.Sprintf("\t%s", validCommand))
	}
}

func validateArgLength(args []string, validLength int) {
	if len(args) <= validLength {
		printUsage()
		os.Exit(1)
	}
}

func game(args []string) {
	validateArgLength(args, 2)
	teamCode := args[1]
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
}

func gameForTeamAndYear(args []string) {
	validateArgLength(args, 1)
	yearArg := args[1]
	year, err := strconv.Atoi(yearArg)
	if err != nil {
		fmt.Println("Year is not valid")
	}
	teams := gamedayapi.TeamsForYear(year)
	for _, team := range teams {
		fmt.Println(team)
		gamedayapi.FetchByTeamAndYear(team, year, eagerLoadGame) // No goroutines here yet.
	}
}

func gameForTeamAndYears(args []string) {
	gamesFor(args)
}

func gamesForYear(args []string) {
	gamesFor(args)
}

func gamesFor(args []string) {
	validateArgLength(args, 2)
	teamCode := args[1]
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

func help(args []string) {
	printValidCommands()
	os.Exit(0)
}

func validTeamsForYear(args []string) {
	validateArgLength(args, 1)
	yearArg := args[1]
	year, err := strconv.Atoi(yearArg)
	if err != nil {
		fmt.Println("Year is not valid")
	}
	teams := gamedayapi.TeamsForYear(year)
	fmt.Println(teams)
}
