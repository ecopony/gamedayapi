gamedayapi - an mlb gameday api for go
======================================

gamedayapi is an API for interacting with MLB Gameday files using the go programming language.

Included is a simple command-line utility for fetching and saving Gameday files to your local filesystem.

Prerequisites: go

To build and run from the command line
---------

In your go workspace, under a directory github.com/ecopony/

    git clone git@github.com:ecopony/gamedayapi.git
    cd gamedayapi
    go build gameday/mlbgd.go

To fetch files for a single game:

    ./mlbgd game sea 2014-07-22

To fetch files for a full season for a team:

    ./mlbgd games-for-team-and-year sea 2014 

To fetch files for multiple seasons for a team:

    ./mlbgd games-for-team-and-years sea 2012 2013 2014


Using the API from your own go code
---------

Import the API.

    import "github.com/ecopony/gamedayapi"

To work with a game:

    date, _ := time.Parse("2006-01-02", "2014-06-02")
    game, _ := gamedayapi.GameFor("sea", date)
    fmt.Println(game.Venue)
    fmt.Println(game.HomeTeamName)
    ...

    // to come... a list of valid team codes

To operate on all games in a season:

    gamedayapi.FetchByTeamAndYears("sea", []int{2012, 2013, 2014}, exampleOfNavigatingAllPitches)
    
    func exampleOfNavigatingAllPitches(game *gamedayapi.Game) {
         log.Println(">>>> " + game.ID + " <<<<")
         for _, inning := range game.AllInnings().Innings {
              for _, atBat := range inning.AtBats() {
                   for _, pitch := range atBat.Pitches {
                        log.Println("> " + pitch.Des)
                   }
              }
         }
    }

The function passed into FetchByTeamAndYears will be passed each game. This could be used to do analysis to save data to
an alternate format or something else fun and entertaining and within the limitations of the usage restrictions.


Usage Restriction
=================

All documents fetched by this API clearly state: Copyright 2011 MLB Advanced Media, L.P. Use of any content on this page
acknowledges agreement to the terms posted here http://gdx.mlb.com/components/copyright.txt

Which furthermore states: "The accounts, descriptions, data and presentation in the referring page (the “Materials”) are
proprietary content of MLB Advanced Media, L.P (“MLBAM”). Only individual, non-commercial, non-bulk use of the Materials
is permitted and any other use of the Materials is prohibited without prior written authorization from MLBAM. Authorized
users of the Materials are prohibited from using the Materials in any commercial manner other than as
expressly authorized by MLBAM."

Naturally, these terms are passed on to any who use this API. It is your responsibility to abide by them.