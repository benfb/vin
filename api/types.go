package api

import "time"

// IDLink is another generic struct
type IDLink struct {
	ID   int    `json:"id"`
	Link string `json:"link"`
}

// IDNameLink is a generic struct that can represent several different types
type IDNameLink struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Link string `json:"link"`
}

// IDFullNameLink is another generic struct
type IDFullNameLink struct {
	ID       int    `json:"id"`
	FullName string `json:"fullName"`
	Link     string `json:"link"`
}

// LabelValue is another generic struct
type LabelValue struct {
	Label string `json:"label"`
	Value string `json:"value"`
}

// Standings is a wrapper for StandingsRecords
type Standings struct {
	Records []StandingsRecord `json:"records"`
}

// StandingsStreak is a winning or losing streak. You usually want StreakCode, which looks like `W7`.
type StandingsStreak struct {
	StreakType   string `json:"streakType"`
	StreakNumber int    `json:"streakNumber"`
	StreakCode   string `json:"streakCode"`
}

// StandingsDivisionRecord is a team's record against a division
type StandingsDivisionRecord struct {
	Wins     int        `json:"wins"`
	Losses   int        `json:"losses"`
	Pct      string     `json:"pct"`
	Division IDNameLink `json:"division"`
}

// StandingsGenericRecord is a generic representation of several more specialized types of records that come from the standings API
type StandingsGenericRecord struct {
	Wins   int    `json:"wins"`
	Losses int    `json:"losses"`
	Type   string `json:"type"`
	Pct    string `json:"pct"`
}

// StandingsLeagueRecord is a team's record against a league
type StandingsLeagueRecord struct {
	Wins   int        `json:"wins"`
	Losses int        `json:"losses"`
	Pct    string     `json:"pct"`
	League IDNameLink `json:"league"`
}

// StandingsDetailedRecords give more granular information about a team's record
type StandingsDetailedRecords struct {
	SplitRecords    []StandingsGenericRecord  `json:"splitRecords"`
	DivisionRecords []StandingsDivisionRecord `json:"divisionRecords"`
	OverallRecords  []StandingsGenericRecord  `json:"overallRecords"`
	LeagueRecords   []StandingsLeagueRecord   `json:"leagueRecords"`
	ExpectedRecords []StandingsGenericRecord  `json:"expectedRecords"`
}

// StandingsTeamRecord gives more detail about a team in the standings
type StandingsTeamRecord struct {
	Team                      IDNameLink               `json:"team"`
	Season                    string                   `json:"season"`
	Streak                    StandingsStreak          `json:"streak"`
	ClinchIndicator           string                   `json:"clinchIndicator,omitempty"`
	DivisionRank              string                   `json:"divisionRank"`
	LeagueRank                string                   `json:"leagueRank"`
	SportRank                 string                   `json:"sportRank"`
	GamesPlayed               int                      `json:"gamesPlayed"`
	GamesBack                 string                   `json:"gamesBack"`
	WildCardGamesBack         string                   `json:"wildCardGamesBack"`
	LeagueGamesBack           string                   `json:"leagueGamesBack"`
	SpringLeagueGamesBack     string                   `json:"springLeagueGamesBack"`
	SportGamesBack            string                   `json:"sportGamesBack"`
	DivisionGamesBack         string                   `json:"divisionGamesBack"`
	ConferenceGamesBack       string                   `json:"conferenceGamesBack"`
	LeagueRecord              StandingsLeagueRecord    `json:"leagueRecord"`
	LastUpdated               time.Time                `json:"lastUpdated"`
	Records                   StandingsDetailedRecords `json:"records"`
	RunsAllowed               int                      `json:"runsAllowed"`
	RunsScored                int                      `json:"runsScored"`
	DivisionChamp             bool                     `json:"divisionChamp"`
	DivisionLeader            bool                     `json:"divisionLeader"`
	HasWildcard               bool                     `json:"hasWildcard"`
	Clinched                  bool                     `json:"clinched"`
	EliminationNumber         string                   `json:"eliminationNumber"`
	WildCardEliminationNumber string                   `json:"wildCardEliminationNumber"`
	MagicNumber               string                   `json:"magicNumber,omitempty"`
	Wins                      int                      `json:"wins"`
	Losses                    int                      `json:"losses"`
	RunDifferential           int                      `json:"runDifferential"`
	WinningPercentage         string                   `json:"winningPercentage"`
	WildCardRank              string                   `json:"wildCardRank,omitempty"`
}

// StandingsRecord is an important standings API struct
type StandingsRecord struct {
	StandingsType string                `json:"standingsType"`
	League        IDLink                `json:"league"`
	Division      IDNameLink            `json:"division"`
	Sport         IDLink                `json:"sport"`
	LastUpdated   time.Time             `json:"lastUpdated"`
	TeamRecords   []StandingsTeamRecord `json:"teamRecords"`
}

// Schedule is a day's worth of games from the MLB API
type Schedule struct {
	TotalItems           int            `json:"totalItems"`
	TotalEvents          int            `json:"totalEvents"`
	TotalGames           int            `json:"totalGames"`
	TotalGamesInProgress int            `json:"totalGamesInProgress"`
	Dates                []ScheduleDate `json:"dates"`
}

// ScheduleGameStatus is a status of a game that is part of a Schedule
type ScheduleGameStatus struct {
	AbstractGameState string `json:"abstractGameState"`
	CodedGameState    string `json:"codedGameState"`
	DetailedState     string `json:"detailedState"`
	StatusCode        string `json:"statusCode"`
	Reason            string `json:"reason,omitempty"`
	AbstractGameCode  string `json:"abstractGameCode"`
}

// ScheduleTeam is a team that's part of a Schedule
type ScheduleTeam struct {
	LeagueRecord StandingsGenericRecord `json:"leagueRecord"`
	Score        int                    `json:"score"`
	Team         IDNameLink             `json:"team"`
	IsWinner     bool                   `json:"isWinner"`
	SplitSquad   bool                   `json:"splitSquad"`
	SeriesNumber int                    `json:"seriesNumber"`
}

// ScheduleTeams is a struct that maps a ScheduleTeam to both Away and Home
type ScheduleTeams struct {
	Away ScheduleTeam `json:"away"`
	Home ScheduleTeam `json:"home"`
}

// ScheduleGame is a game that is part of a schedule
type ScheduleGame struct {
	GamePk   int                `json:"gamePk"`
	Link     string             `json:"link"`
	GameType string             `json:"gameType"`
	Season   string             `json:"season"`
	GameDate time.Time          `json:"gameDate"`
	Status   ScheduleGameStatus `json:"status,omitempty"`
	Teams    ScheduleTeams      `json:"teams"`
	Venue    IDNameLink         `json:"venue"`
	Content  struct {
		Link string `json:"link"`
	} `json:"content"`
	IsTie                  bool      `json:"isTie,omitempty"`
	GameNumber             int       `json:"gameNumber"`
	PublicFacing           bool      `json:"publicFacing"`
	DoubleHeader           string    `json:"doubleHeader"`
	GamedayType            string    `json:"gamedayType"`
	Tiebreaker             string    `json:"tiebreaker"`
	CalendarEventID        string    `json:"calendarEventID"`
	SeasonDisplay          string    `json:"seasonDisplay"`
	DayNight               string    `json:"dayNight"`
	ScheduledInnings       int       `json:"scheduledInnings"`
	GamesInSeries          int       `json:"gamesInSeries"`
	SeriesGameNumber       int       `json:"seriesGameNumber"`
	SeriesDescription      string    `json:"seriesDescription"`
	RecordSource           string    `json:"recordSource"`
	IfNecessary            string    `json:"ifNecessary"`
	IfNecessaryDescription string    `json:"ifNecessaryDescription"`
	RescheduleDate         time.Time `json:"rescheduleDate,omitempty"`
}

// ScheduleDate is one day's worth of games
type ScheduleDate struct {
	Date                 string         `json:"date"`
	TotalItems           int            `json:"totalItems"`
	TotalEvents          int            `json:"totalEvents"`
	TotalGames           int            `json:"totalGames"`
	TotalGamesInProgress int            `json:"totalGamesInProgress"`
	Games                []ScheduleGame `json:"games"`
	Events               []interface{}  `json:"events"`
}

// LineScore represents an MLB API linescore
type LineScore struct {
	CurrentInning        int               `json:"currentInning"`
	CurrentInningOrdinal string            `json:"currentInningOrdinal"`
	InningState          string            `json:"inningState"`
	InningHalf           string            `json:"inningHalf"`
	IsTopInning          bool              `json:"isTopInning"`
	ScheduledInnings     int               `json:"scheduledInnings"`
	Innings              []LineScoreInning `json:"innings"`
	Teams                LineScoreTeams    `json:"teams"`
	Defense              LineScoreDefense  `json:"defense"`
	Offense              LineScoreOffense  `json:"offense"`
	Balls                int               `json:"balls"`
	Strikes              int               `json:"strikes"`
	Outs                 int               `json:"outs"`
}

// LineScoreInning represents an inning in a Linescore
type LineScoreInning struct {
	Num        int           `json:"num"`
	OrdinalNum string        `json:"ordinalNum"`
	Home       LineScoreTeam `json:"home,omitempty"`
	Away       LineScoreTeam `json:"away,omitempty"`
}

// LineScoreTeam is a team's representation in a LineScore
type LineScoreTeam struct {
	Runs       int `json:"runs"`
	Hits       int `json:"hits"`
	Errors     int `json:"errors"`
	LeftOnBase int `json:"leftOnBase"`
}

// LineScoreTeams maps a LineScoreTeam to Home and another to Away
type LineScoreTeams struct {
	Home LineScoreTeam `json:"home"`
	Away LineScoreTeam `json:"away"`
}

// LineScoreDefense gives information about players on defense
type LineScoreDefense struct {
	Pitcher   IDFullNameLink `json:"pitcher"`
	Catcher   IDFullNameLink `json:"catcher"`
	First     IDFullNameLink `json:"first"`
	Second    IDFullNameLink `json:"second"`
	Third     IDFullNameLink `json:"third"`
	Shortstop IDFullNameLink `json:"shortstop"`
	Left      IDFullNameLink `json:"left"`
	Center    IDFullNameLink `json:"center"`
	Right     IDFullNameLink `json:"right"`
	Team      IDNameLink     `json:"team"`
}

// LineScoreOffense gives information about players on offense
type LineScoreOffense struct {
	Batter  IDFullNameLink `json:"batter"`
	OnDeck  IDFullNameLink `json:"onDeck"`
	InHole  IDFullNameLink `json:"inHole"`
	Second  IDFullNameLink `json:"second"`
	Pitcher IDFullNameLink `json:"pitcher"`
	Team    IDNameLink     `json:"team"`
}

// RosterResponse is a response from /v1/teams/<teamid>/roster
type RosterResponse struct {
	Roster []struct {
		Person       IDFullNameLink `json:"person"`
		JerseyNumber string         `json:"jerseyNumber"`
		Position     struct {
			Code         string `json:"code"`
			Name         string `json:"name"`
			Type         string `json:"type"`
			Abbreviation string `json:"abbreviation"`
		} `json:"position"`
		Status struct {
			Code        string `json:"code"`
			Description string `json:"description"`
		} `json:"status"`
		ParentTeamID int `json:"parentTeamId"`
	} `json:"roster"`
	Link       string `json:"link"`
	TeamID     int    `json:"teamId"`
	RosterType string `json:"rosterType"`
}
