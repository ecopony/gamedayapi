package gamedayapi

import "encoding/xml"

// Boxscore is the representation of the boxscore.xml file.
type Boxscore struct {
	XMLName       xml.Name `xml:"boxscore"`
	AwayFName     string   `xml:"away_fname,attr"`
	AwayID        string   `xml:"away_id,attr"`
	AwayLoss      string   `xml:"away_loss,attr"`
	AwaySName     string   `xml:"away_sname,attr"`
	AwayTeamCode  string   `xml:"away_team_code,attr"`
	AwayWins      string   `xml:"away_wins,attr"`
	Date          string   `xml:"date,attr"`
	GameID        string   `xml:"game_id,attr"`
	GamePk        string   `xml:"game_pk,attr"`
	HomeFName     string   `xml:"home_fname,attr"`
	HomeID        string   `xml:"home_id,attr"`
	HomeLoss      string   `xml:"home_loss,attr"`
	HomeSName     string   `xml:"home_sname,attr"`
	HomeSportCode string   `xml:"home_sport_code,attr"`
	HomeTeamCode  string   `xml:"home_team_code,attr"`
	HomeWins      string   `xml:"home_wins,attr"`
	StatusInd     string   `xml:"status_ind,attr"`
	VenueID       string   `xml:"venue_id,attr"`
	VenueName     string   `xml:"venue_name,attr"`

	Linescore Linescore  `xml:"linescore"`
	Batting   []Batting  `xml:"batting"`
	Pitching  []Pitching `xml:"pitching"`
}

// Linescore represents the linescore under the boxscore, not the individual linescore.xml file.
type Linescore struct {
	XMLName      xml.Name `xml:"linescore"`
	AwayTeamRuns string   `xml:"away_team_runs,attr"`
	HomeTeamRuns string   `xml:"home_team_runs,attr"`

	InningLineScores []InningLineScore `xml:"inning_line_score"`
}

// InningLineScore represents the individual innings in the linescore.
type InningLineScore struct {
	XMLName xml.Name `xml:"inning_line_score"`
	Away    string   `xml:"away,attr"`
	Home    string   `xml:"home,attr"`
	Inning  string   `xml:"inning,attr"`
}

// Batting represents the batting elements in the boxscore.
type Batting struct {
	XMLName  xml.Name `xml:"batting"`
	TeamFlag string   `xml:"team_flag,attr"`
	AB       string   `xml:"ab,attr"`
	R        string   `xml:"r,attr"`
	H        string   `xml:"h,attr"`
	D        string   `xml:"d,attr"`
	T        string   `xml:"t,attr"`
	HR       string   `xml:"hr,attr"`
	RBI      string   `xml:"rbi,attr"`
	BB       string   `xml:"bb,attr"`
	PO       string   `xml:"po,attr"`
	DA       string   `xml:"da,attr"`
	SO       string   `xml:"so,attr"`
	LOB      string   `xml:"lob,attr"`
	AVG      string   `xml:"avg,attr"`

	Batters []Batter `xml:"batter"`
}

// Batter represents the batter elements in the boxscore.
type Batter struct {
	XMLName              xml.Name `xml:"batter"`
	ID                   string   `xml:"id,attr"`
	Name                 string   `xml:"name,attr"`
	NameDisplayFirstLast string   `xml:"name_display_first_last,attr"`
	Pos                  string   `xml:"pos,attr"`
	BO                   string   `xml:"bo,attr"`
	AB                   string   `xml:"ab,attr"`
	PO                   string   `xml:"po,attr"`
	R                    string   `xml:"r,attr"`
	A                    string   `xml:"a,attr"`
	BB                   string   `xml:"bb,attr"`
	SAC                  string   `xml:"sac,attr"`
	T                    string   `xml:"t,attr"`
	SF                   string   `xml:"sf,attr"`
	H                    string   `xml:"h,attr"`
	E                    string   `xml:"e,attr"`
	D                    string   `xml:"d,attr"`
	HBP                  string   `xml:"hbp,attr"`
	SO                   string   `xml:"so,attr"`
	HR                   string   `xml:"hr,attr"`
	RBI                  string   `xml:"rbi,attr"`
	LOB                  string   `xml:"lob,attr"`
	FLDG                 string   `xml:"fldg,attr"`
	SB                   string   `xml:"sb,attr"`
	CS                   string   `xml:"cs,attr"`
	SHR                  string   `xml:"s_hr,attr"`
	SRBI                 string   `xml:"s_rbi,attr"`
	SH                   string   `xml:"s_h,attr"`
	SBB                  string   `xml:"s_bb,attr"`
	SR                   string   `xml:"s_r,attr"`
	SSO                  string   `xml:"s_so,attr"`
	AVG                  string   `xml:"savg,attr"`
	GO                   string   `xml:"go,attr"`
	AO                   string   `xml:"ao,attr"`
}

// Pitching represents the pitching elements in the boxscore.
type Pitching struct {
	XMLName  xml.Name `xml:"pitching"`
	TeamFlag string   `xml:"team_flag,attr"`
	Out      string   `xml:"out,attr"`
	H        string   `xml:"h,attr"`
	R        string   `xml:"r,attr"`
	ER       string   `xml:"er,attr"`
	BB       string   `xml:"bb,attr"`
	SO       string   `xml:"so,attr"`
	HR       string   `xml:"hr,attr"`
	BF       string   `xml:"bf,attr"`
	ERA      string   `xml:"era,attr"`

	Pitchers []Pitcher `xml:"pitcher"`
}

// Pitcher represents the pitcher elements in the boxscore.
type Pitcher struct {
	XMLName              xml.Name `xml:"pitcher"`
	ID                   string   `xml:"id,attr"`
	Name                 string   `xml:"name,attr"`
	NameDisplayFirstLast string   `xml:"name_display_first_last,attr"`
	Pos                  string   `xml:"pos,attr"`
	Out                  string   `xml:"out,attr"`
	BF                   string   `xml:"bf,attr"`
	ER                   string   `xml:"er,attr"`
	R                    string   `xml:"r,attr"`
	H                    string   `xml:"h,attr"`
	SO                   string   `xml:"so,attr"`
	HR                   string   `xml:"hr,attr"`
	BB                   string   `xml:"bb,attr"`
	NP                   string   `xml:"np,attr"`
	S                    string   `xml:"s,attr"`
	W                    string   `xml:"w,attr"`
	L                    string   `xml:"l,attr"`
	SV                   string   `xml:"sv,attr"`
	BS                   string   `xml:"bs,attr"`
	HLD                  string   `xml:"hld,attr"`
	SIP                  string   `xml:"s_ip,attr"`
	SH                   string   `xml:"s_h,attr"`
	SR                   string   `xml:"s_r,attr"`
	SER                  string   `xml:"s_er,attr"`
	SBB                  string   `xml:"s_bb,attr"`
	SSO                  string   `xml:"s_so,attr"`
	ERA                  string   `xml:"era,attr"`
	Win                  string   `xml:"win,attr"`
	Note                 string   `xml:"note,attr"`
}
