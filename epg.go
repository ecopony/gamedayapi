package gamedayapi

import (
	s "strings"
)

type Epg struct {
	Date string `xml:"id,attr"`
	LastModified string `xml:"last_modified,attr"`
	DisplayTimeZone string `xml:"display_time_zone,attr"`
	EpgGames []EpgGame `xml:"game"`
}

/*
gids look like: gid_2014_07_22_nynmlb_seamlb_1

Doesn't yet handle doubleheader days. It'll just return the first match it finds for the team.
 */
func (e *Epg) GidForTeam(teamCode string) Gid {
	for _, game := range e.EpgGames {
		if s.Contains(game.Gameday, s.Join([]string{"_", teamCode, "mlb_"}, "")) {
			gamedayPieces := s.Split(game.Gameday, "_")
			return Gid{gamedayPieces[0], gamedayPieces[1], gamedayPieces[2], gamedayPieces[3], gamedayPieces[4], gamedayPieces[5]}
		}
	}
	return Gid{} // this should be an error
}

type EpgGame struct {
	CalendarEventId string `xml:"calendar_event_id,attr"`
	Start string `xml:"start,attr"`
	Id string `xml:"id,attr"`
	Gameday string `xml:"gameday,attr"`
}
