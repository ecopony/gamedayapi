package gamedayapi

import (
	"reflect"
	"testing"
	"time"
	"fmt"
)

func GameForTest() *Game {
	date, _ := time.Parse("2006-01-02", "2014-06-22")
	game, _ := GameFor("sea", date)
	return game
}

func assertEquals(t *testing.T, a interface{}, b interface{}) {
	if !reflect.DeepEqual(a, b) {
		t.Errorf("'%+v' != '%+v'", a, b)
	}
}

func TestYear(t *testing.T) {
	game := GameForTest()
	assertEquals(t, game.Year(), 2014)
}

func TestFetchableDataDirectory(t *testing.T) {
	game := GameForTest()
	assertEquals(t, game.FetchableDataDirectory(), "/components/game/mlb/year_2014/month_06/day_22/gid_2014_06_22_seamlb_kcamlb_1")
}

func TestAtBatsOnInning(t *testing.T) {
	game := GameForTest()
	inning := game.AllInnings().Innings[0]
	atBatTotal := len(inning.Top.AtBats) + len(inning.Bottom.AtBats)
	assertEquals(t, len(inning.AtBats()), atBatTotal)
}

func TestDoubleheaders(t *testing.T) {
	date, _ := time.Parse("2006-01-02", "2014-05-07")
	games, _ := GamesFor("sea", date)
	assertEquals(t, len(games), 2)
	assertEquals(t, games[0].Gameday, "2014_05_07_seamlb_oakmlb_1")
	assertEquals(t, games[1].Gameday, "2014_05_07_seamlb_oakmlb_2")
}

func TestNoGamesThatDay(t *testing.T) {
	date, _ := time.Parse("2006-01-02", "2014-03-22")
	games, err := GamesFor("sea", date)
	assertEquals(t, fmt.Sprint(err), "[sea] doesn't have a game on [20140322]")
	assertEquals(t, len(games), 0)
}

func TestTeamsForScheduleYear(t *testing.T) {
	teams := TeamsForYear(2014)
	assertEquals(t, len(teams), 30)
	assertEquals(t, teams[0], "LAN")
	assertEquals(t, teams[29], "MIA")
}
