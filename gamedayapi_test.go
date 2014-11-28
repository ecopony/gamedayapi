package gamedayapi

import (
	"testing"
	"reflect"
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

func TestAtBatsOnInning(t *testing.T) {
	game := GameForTest()
	inning := game.AllInnings().Innings[0]
	atBatTotal := len(inning.Top.AtBats) + len(inning.Bottom.AtBats)
	assertEquals(t, len(inning.AtBats()), atBatTotal)
}
