package api

// Games represents an XML root response from the API
type Games struct {
	Date     string `xml:"date"`
	GameList []Game `xml:"game"`
}

// Game is an individual game
type Game struct {
	ID           string `xml:"id,attr"`
	Time         string `xml:"time,attr"`
	Status       string `xml:"status,attr"`
	Inning       int    `xml:"inning,attr"`
	IsTop        bool   `xml:"top_inning,attr"`
	AwayTeam     string `xml:"away_team_name,attr"`
	HomeTeam     string `xml:"home_team_name,attr"`
	AwayTeamRuns int    `xml:"away_team_runs,attr"`
	HomeTeamRuns int    `xml:"home_team_runs,attr"`
	AwayTeamHits int    `xml:"away_team_hits,attr"`
	HomeTeamHits int    `xml:"home_team_hits,attr"`
	AwayTeamErrs int    `xml:"away_team_errors,attr"`
	HomeTeamErrs int    `xml:"home_team_errors,attr"`
}
