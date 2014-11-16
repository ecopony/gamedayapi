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

	var game gamedayapi.Game
	game.For(teamCode, date)
}
