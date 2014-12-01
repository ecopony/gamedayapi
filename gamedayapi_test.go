package gamedayapi

import (
	"reflect"
	"testing"
	//	"log"
)

func GameForTest() *Game {
	game, _ := GameFor("sea", "2014-06-22")
	return game
}

func assertEquals(t *testing.T, a interface{}, b interface{}) {
	if !reflect.DeepEqual(a, b) {
		t.Errorf("'%+v' != '%+v'", a, b)
	}
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
	games, _ := GamesFor("sea", "2014-05-07")
	assertEquals(t, len(games), 2)
	assertEquals(t, games[0].Gameday, "2014_05_07_seamlb_oakmlb_1")
	assertEquals(t, games[1].Gameday, "2014_05_07_seamlb_oakmlb_2")
}
