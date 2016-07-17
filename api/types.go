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

// StandingsResponse is a json root response from the API
type StandingsResponse struct {
	Date         string    `json:"standings_date"`
	StandingList Standings `json:"standing"`
}

// Standings is a slice of multiple Standings
type Standings []Standing

// Standing is an individual standing
type Standing struct {
	ID            string  `json:"team_id"`
	Rank          int     `json:"rank"`
	OrdinalRank   string  `json:"ordinal_rank"`
	Won           int     `json:"won"`
	Lost          int     `json:"lost"`
	FirstName     string  `json:"first_name"`
	LastName      string  `json:"last_name"`
	League        string  `json:"conference"`
	Division      string  `json:"division"`
	GamesBack     float64 `json:"games_back"`
	GamesPlayed   int     `json:"games_played"`
	WinPercentage string  `json:"win_percentage"`
}

func (slice Standings) Len() int {
	return len(slice)
}

func (slice Standings) Less(i, j int) bool {
	return slice[i].WinPercentage > slice[j].WinPercentage
}

func (slice Standings) Swap(i, j int) {
	slice[i], slice[j] = slice[j], slice[i]
}
