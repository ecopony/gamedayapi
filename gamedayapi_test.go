package gamedayapi

import (
	"testing"
	"log"
	"reflect"
)

func assertEquals(t *testing.T, a interface{}, b interface{}) {
	if !reflect.DeepEqual(a, b) {
		t.Errorf("'%+v' != '%+v'", a, b)
	}
}

func GidForTest() *Gid {
	return &Gid{"2014", "06", "22", "chn", "sea", "1"}
}

func TestGidString(t *testing.T) {
	log.Print(GidForTest().String())
	assertEquals(t, GidForTest().String(), "gid_2014_06_22_chn_sea_1")
}

func TestGidDatePath(t *testing.T) {
	log.Print(GidForTest().String())
	assertEquals(t, GidForTest().DatePath(), "year_2014/month_06/day_22")
}
