package gamedayapi

type Game struct {
	AwayAPMP			string		`xml:"away_ampm,attr"`
	AwayCode			string		`xml:"away_code,attr"`
	AwayLoss			string		`xml:"away_loss,attr"`
	AwayTeamCity		string		`xml:"away_team_city,attr"`
	AwayTeamId			string		`xml:"away_team_id,attr"`
	AwayTeamName		string		`xml:"away_team_name,attr"`
	AwayTime			string		`xml:"away_time,attr"`
	AwayTimezone		string		`xml:"away_time_zone,attr"`
	AwayWin				string		`xml:"away_win,attr"`
	HomeAMPM			string		`xml:"home_ampm,attr"`
	HomeCode			string		`xml:"home_code,attr"`
	HomeLoss			string		`xml:"home_loss,attr"`
	HomeTeamCity		string		`xml:"home_team_city,attr"`
	HomeTeamId			string		`xml:"home_team_id,attr"`
	HomeTeamName		string		`xml:"home_team_name,attr"`
	HomeTime			string		`xml:"home_time,attr"`
	HomeTimezone		string		`xml:"home_time_zone,attr"`
	HomeWin				string		`xml:"home_win,attr"`
	Id					string		`xml:"id,attr"`
	GamePk				string		`xml:"game_pk,attr"`
	Timezone			string		`xml:"time_zone,attr"`
	Venue				string		`xml:"venue,attr"`
	GameDataDirectory	string		`xml:"game_data_directory,attr"`
}

func GameFor(teamCode string, date string) *Game {
	epg := EpgFor(date)
	game := epg.GameForTeam(teamCode)
	return game
}
