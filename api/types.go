package api

// MLBRoot is the root of an MLB JSON API response
type Schedule struct {
	Dates []ScheduleDate `json:"dates"`
}

// Date represents a data root from the MLB JSON API
type ScheduleDate struct {
	Date  string         `json:"date"`
	Games []ScheduleGame `json:"games"`
}

// Game is an individual game
type ScheduleGame struct {
	ID     int                `json:"gamePk"`
	Time   string             `json:"time"`
	Status ScheduleGameStatus `json:"status"`
	Teams  ScheduleGameTeams  `json:"teams"`
}

type ScheduleGameStatus struct {
	AbstractGameState string `json:"abstractGameState"`
	CodedGameState    string `json:"codedGameState"`
	DetailedState     string `json:"detailedState"`
	StatusCode        string `json:"statusCode"`
	AbstractGameCode  string `json:"abstractGameCode"`
}

type ScheduleGameTeams struct {
	Away GameTeam `json:"away"`
	Home GameTeam `json:"home"`
}

type GameTeam struct {
	LeagueRecord GameTeamLeagueRecord
	Score        int          `json:"score"`
	Team         GameTeamTeam `json:"team"`
	IsWinner     bool         `json:"isWinner"`
	SplitSquad   bool         `json:"splitSquad"`
	SeriesNumber int          `json:"seriesNumber"`
}

type GameTeamTeam struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Link string `json:"link"`
}

type GameTeamLeagueRecord struct {
	Wins   int    `json:"wins"`
	Losses int    `json:"losses"`
	Pct    string `json:"pct"`
}

type LineScore struct {
	CurrentInning        int               `json:"currentInning"`
	CurrentInningOrdinal string            `json:"currentInningOrdinal"`
	InningHalf           string            `json:"inningHalf"`
	IsTopInning          bool              `json:"isTopInning"`
	ScheduledInnings     int               `json:"scheduledInnings"`
	Innings              []LineScoreInning `json:"innings"`
	Teams                LineScoreTeams    `json:"teams"`
}

type LineScoreInning struct {
	Num        int                 `json:"num"`
	OrdinalNum string              `json:"ordinalNum"`
	Away       LineScoreInningTeam `json:"away"`
	Home       LineScoreInningTeam `json:"home"`
}

type LineScoreInningTeam struct {
	Runs int `json:"runs"`
}

type LineScoreTeams struct {
	Away LineScoreTeam `json:"away"`
	Home LineScoreTeam `json:"home"`
}

type LineScoreTeam struct {
	Runs   int `json:"runs"`
	Hits   int `json:"hits"`
	Errors int `json:"errors"`
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

type StandingsRecords struct {
	TeamRecords []StandingsTeamRecords `json:"teamRecords"`
}

type StandingsTeamRecords []StandingsTeamRecord

type StandingsTeamRecord struct {
	Team      StandingsTeamRecordTeam   `json:"team"`
	Streak    StandingsTeamRecordStreak `json:"streak"`
	GamesBack string                    `json:"gamesBack"`
}

type StandingsTeamRecordTeam struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type StandingsTeamRecordStreak struct {
	StreakCode string `json:"streakCode"`
}

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
