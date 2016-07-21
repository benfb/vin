package api

// MLBRoot is the root of an MLB JSON API response
type MLBRoot struct {
	Data MLBData `json:"data"`
}

// MLBData represents a data root from the MLB JSON API
type MLBData struct {
	Games Games `json:"games"`
}

// Games represents a JSON root response from the API
type Games struct {
	Date     string `json:"date"`
	GameList []Game `json:"game"`
}

// Game is an individual game
type Game struct {
	ID                  string          `json:"id"`
	Time                string          `json:"time"`
	Status              string          `json:"status"`
	Inning              int             `json:"inning,string"`
	IsTop               bool            `json:"top_inning"`
	AwayTeam            string          `json:"away_team_name"`
	HomeTeam            string          `json:"home_team_name"`
	AwayTeamRuns        int             `json:"away_team_runs,string"`
	HomeTeamRuns        int             `json:"home_team_runs,string"`
	AwayTeamHits        int             `json:"away_team_hits,string"`
	HomeTeamHits        int             `json:"home_team_hits,string"`
	AwayTeamErrs        int             `json:"away_team_errors,string"`
	HomeTeamErrs        int             `json:"home_team_errors,string"`
	HomeTeamWins        int             `json:"home_win,string"`
	AwayTeamWins        int             `json:"away_win,string"`
	HomeTeamLosses      int             `json:"home_loss,string"`
	AwayTeamLosses      int             `json:"away_loss,string"`
	HomeProbablePitcher ProbablePitcher `json:"home_probable_pitcher"`
	AwayProbablePitcher ProbablePitcher `json:"away_probable_pitcher"`
}

// ProbablePitcher represents the likely pitcher for a game
type ProbablePitcher struct {
	ID        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Wins      int    `json:"wins,string"`
	Losses    int    `json:"losses,string"`
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
	Streak        string  `json:"streak"`
}
