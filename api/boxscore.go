package api

import (
	"fmt"
	"os"
	"strconv"

	"github.com/olekukonko/tablewriter"
)

type BoxScore struct {
	Teams         BoxScoreTeams       `json:"teams"`
	Officials     []BoxScoreOfficials `json:"officials"`
	Info          []LabelValue        `json:"info"`
	PitchingNotes []interface{}       `json:"pitchingNotes"`
}
type BoxScoreTeams struct {
	Away BoxScoreTeam `json:"away"`
	Home BoxScoreTeam `json:"home"`
}
type BoxScoreOfficials struct {
	Official     IDFullNameLink `json:"official"`
	OfficialType string         `json:"officialType"`
}
type BoxScoreRecord struct {
	GamesPlayed           int                    `json:"gamesPlayed"`
	WildCardGamesBack     string                 `json:"wildCardGamesBack"`
	LeagueGamesBack       string                 `json:"leagueGamesBack"`
	SpringLeagueGamesBack string                 `json:"springLeagueGamesBack"`
	SportGamesBack        string                 `json:"sportGamesBack"`
	DivisionGamesBack     string                 `json:"divisionGamesBack"`
	ConferenceGamesBack   string                 `json:"conferenceGamesBack"`
	LeagueRecord          StandingsGenericRecord `json:"leagueRecord"`
	Records               interface{}            `json:"records"`
	DivisionLeader        bool                   `json:"divisionLeader"`
	Wins                  int                    `json:"wins"`
	Losses                int                    `json:"losses"`
	WinningPercentage     string                 `json:"winningPercentage"`
}
type SpringLeague struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Link         string `json:"link"`
	Abbreviation string `json:"abbreviation"`
}
type BoxScoreDetailedTeam struct {
	ID              int            `json:"id"`
	Name            string         `json:"name"`
	Link            string         `json:"link"`
	Season          int            `json:"season"`
	Venue           IDNameLink     `json:"venue"`
	TeamCode        string         `json:"teamCode"`
	FileCode        string         `json:"fileCode"`
	Abbreviation    string         `json:"abbreviation"`
	TeamName        string         `json:"teamName"`
	LocationName    string         `json:"locationName"`
	FirstYearOfPlay string         `json:"firstYearOfPlay"`
	League          IDNameLink     `json:"league"`
	Division        IDNameLink     `json:"division"`
	Sport           IDNameLink     `json:"sport"`
	ShortName       string         `json:"shortName"`
	Record          BoxScoreRecord `json:"record"`
	SpringLeague    SpringLeague   `json:"springLeague"`
	AllStarStatus   string         `json:"allStarStatus"`
	Active          bool           `json:"active"`
}
type BoxScoreBattingStats struct {
	FlyOuts              int    `json:"flyOuts"`
	GroundOuts           int    `json:"groundOuts"`
	Runs                 int    `json:"runs"`
	Doubles              int    `json:"doubles"`
	Triples              int    `json:"triples"`
	HomeRuns             int    `json:"homeRuns"`
	StrikeOuts           int    `json:"strikeOuts"`
	BaseOnBalls          int    `json:"baseOnBalls"`
	IntentionalWalks     int    `json:"intentionalWalks"`
	Hits                 int    `json:"hits"`
	HitByPitch           int    `json:"hitByPitch"`
	Avg                  string `json:"avg"`
	AtBats               int    `json:"atBats"`
	Obp                  string `json:"obp"`
	Slg                  string `json:"slg"`
	Ops                  string `json:"ops"`
	CaughtStealing       int    `json:"caughtStealing"`
	StolenBases          int    `json:"stolenBases"`
	GroundIntoDoublePlay int    `json:"groundIntoDoublePlay"`
	GroundIntoTriplePlay int    `json:"groundIntoTriplePlay"`
	TotalBases           int    `json:"totalBases"`
	Rbi                  int    `json:"rbi"`
	LeftOnBase           int    `json:"leftOnBase"`
	SacBunts             int    `json:"sacBunts"`
	SacFlies             int    `json:"sacFlies"`
	CatchersInterference int    `json:"catchersInterference"`
	Pickoffs             int    `json:"pickoffs"`
}
type BoxScorePitchingStats struct {
	GroundOuts             int    `json:"groundOuts"`
	Runs                   int    `json:"runs"`
	Doubles                int    `json:"doubles"`
	Triples                int    `json:"triples"`
	HomeRuns               int    `json:"homeRuns"`
	StrikeOuts             int    `json:"strikeOuts"`
	BaseOnBalls            int    `json:"baseOnBalls"`
	IntentionalWalks       int    `json:"intentionalWalks"`
	Hits                   int    `json:"hits"`
	AtBats                 int    `json:"atBats"`
	CaughtStealing         int    `json:"caughtStealing"`
	StolenBases            int    `json:"stolenBases"`
	Era                    string `json:"era"`
	InningsPitched         string `json:"inningsPitched"`
	SaveOpportunities      int    `json:"saveOpportunities"`
	EarnedRuns             int    `json:"earnedRuns"`
	Whip                   string `json:"whip"`
	BattersFaced           int    `json:"battersFaced"`
	Outs                   int    `json:"outs"`
	CompleteGames          int    `json:"completeGames"`
	Shutouts               int    `json:"shutouts"`
	HitBatsmen             int    `json:"hitBatsmen"`
	WildPitches            int    `json:"wildPitches"`
	Pickoffs               int    `json:"pickoffs"`
	AirOuts                int    `json:"airOuts"`
	Rbi                    int    `json:"rbi"`
	InheritedRunners       int    `json:"inheritedRunners"`
	InheritedRunnersScored int    `json:"inheritedRunnersScored"`
	CatchersInterference   int    `json:"catchersInterference"`
	SacBunts               int    `json:"sacBunts"`
	SacFlies               int    `json:"sacFlies"`
}
type TeamStats struct {
	Batting  BoxScoreBattingStats        `json:"batting"`
	Pitching BoxScorePitchingStats       `json:"pitching"`
	Fielding BoxScorePlayerFieldingStats `json:"fielding"`
}
type SeasonStats struct {
	Batting  BoxScoreBattingStats        `json:"batting"`
	Pitching BoxScorePitchingStats       `json:"pitching"`
	Fielding BoxScorePlayerFieldingStats `json:"fielding"`
}
type Position struct {
	Code         string `json:"code"`
	Name         string `json:"name"`
	Type         string `json:"type"`
	Abbreviation string `json:"abbreviation"`
}

type BoxScorePlayerStats struct {
	Batting  BoxScorePlayerBattingStats  `json:"batting"`
	Pitching BoxScorePlayerPitchingStats `json:"pitching"`
	Fielding BoxScorePlayerFieldingStats `json:"fielding"`
}
type BoxScorePlayerStatus struct {
	Code        string `json:"code"`
	Description string `json:"description"`
}
type BoxScorePlayerBattingStats struct {
	Note                 string `json:"note"`
	GamesPlayed          int    `json:"gamesPlayed"`
	FlyOuts              int    `json:"flyOuts"`
	GroundOuts           int    `json:"groundOuts"`
	Runs                 int    `json:"runs"`
	Doubles              int    `json:"doubles"`
	Triples              int    `json:"triples"`
	HomeRuns             int    `json:"homeRuns"`
	StrikeOuts           int    `json:"strikeOuts"`
	BaseOnBalls          int    `json:"baseOnBalls"`
	IntentionalWalks     int    `json:"intentionalWalks"`
	Hits                 int    `json:"hits"`
	HitByPitch           int    `json:"hitByPitch"`
	Avg                  string `json:"avg"`
	AtBats               int    `json:"atBats"`
	Obp                  string `json:"obp"`
	Slg                  string `json:"slg"`
	Ops                  string `json:"ops"`
	CaughtStealing       int    `json:"caughtStealing"`
	StolenBases          int    `json:"stolenBases"`
	StolenBasePercentage string `json:"stolenBasePercentage"`
	GroundIntoDoublePlay int    `json:"groundIntoDoublePlay"`
	GroundIntoTriplePlay int    `json:"groundIntoTriplePlay"`
	TotalBases           int    `json:"totalBases"`
	Rbi                  int    `json:"rbi"`
	LeftOnBase           int    `json:"leftOnBase"`
	SacBunts             int    `json:"sacBunts"`
	SacFlies             int    `json:"sacFlies"`
	CatchersInterference int    `json:"catchersInterference"`
	Pickoffs             int    `json:"pickoffs"`
}
type BoxScorePlayerPitchingStats struct {
	Note                   string `json:"note"`
	GamesPlayed            int    `json:"gamesPlayed"`
	GamesStarted           int    `json:"gamesStarted"`
	GroundOuts             int    `json:"groundOuts"`
	Runs                   int    `json:"runs"`
	Doubles                int    `json:"doubles"`
	Triples                int    `json:"triples"`
	HomeRuns               int    `json:"homeRuns"`
	StrikeOuts             int    `json:"strikeOuts"`
	BaseOnBalls            int    `json:"baseOnBalls"`
	IntentionalWalks       int    `json:"intentionalWalks"`
	Hits                   int    `json:"hits"`
	AtBats                 int    `json:"atBats"`
	CaughtStealing         int    `json:"caughtStealing"`
	StolenBases            int    `json:"stolenBases"`
	StolenBasePercentage   string `json:"stolenBasePercentage"`
	Era                    string `json:"era"`
	InningsPitched         string `json:"inningsPitched"`
	Wins                   int    `json:"wins"`
	Losses                 int    `json:"losses"`
	Saves                  int    `json:"saves"`
	SaveOpportunities      int    `json:"saveOpportunities"`
	Holds                  int    `json:"holds"`
	BlownSaves             int    `json:"blownSaves"`
	EarnedRuns             int    `json:"earnedRuns"`
	Whip                   string `json:"whip"`
	Outs                   int    `json:"outs"`
	GamesPitched           int    `json:"gamesPitched"`
	CompleteGames          int    `json:"completeGames"`
	Shutouts               int    `json:"shutouts"`
	HitBatsmen             int    `json:"hitBatsmen"`
	WildPitches            int    `json:"wildPitches"`
	Pickoffs               int    `json:"pickoffs"`
	AirOuts                int    `json:"airOuts"`
	Rbi                    int    `json:"rbi"`
	WinPercentage          string `json:"winPercentage"`
	GamesFinished          int    `json:"gamesFinished"`
	StrikeoutWalkRatio     string `json:"strikeoutWalkRatio"`
	StrikeoutsPer9Inn      string `json:"strikeoutsPer9Inn"`
	WalksPer9Inn           string `json:"walksPer9Inn"`
	HitsPer9Inn            string `json:"hitsPer9Inn"`
	InheritedRunners       int    `json:"inheritedRunners"`
	InheritedRunnersScored int    `json:"inheritedRunnersScored"`
	CatchersInterference   int    `json:"catchersInterference"`
	SacBunts               int    `json:"sacBunts"`
	SacFlies               int    `json:"sacFlies"`
}
type BoxScorePlayerFieldingStats struct {
	Assists              int    `json:"assists"`
	PutOuts              int    `json:"putOuts"`
	Errors               int    `json:"errors"`
	Chances              int    `json:"chances"`
	Fielding             string `json:"fielding,omitempty"`
	CaughtStealing       int    `json:"caughtStealing"`
	PassedBall           int    `json:"passedBall"`
	StolenBases          int    `json:"stolenBases"`
	StolenBasePercentage string `json:"stolenBasePercentage"`
	Pickoffs             int    `json:"pickoffs"`
}
type BoxScorePlayerGameStatus struct {
	IsCurrentBatter  bool `json:"isCurrentBatter"`
	IsCurrentPitcher bool `json:"isCurrentPitcher"`
	IsOnBench        bool `json:"isOnBench"`
	IsSubstitute     bool `json:"isSubstitute"`
}
type AllPositions struct {
	Code         string `json:"code"`
	Name         string `json:"name"`
	Type         string `json:"type"`
	Abbreviation string `json:"abbreviation"`
}
type BoxScoreTeamInfo struct {
	Title     string       `json:"title"`
	FieldList []LabelValue `json:"fieldList"`
}
type BoxScoreTeam struct {
	Team         BoxScoreDetailedTeam      `json:"team"`
	TeamStats    TeamStats                 `json:"teamStats"`
	Players      map[string]BoxScorePlayer `json:"players"`
	Batters      []int                     `json:"batters"`
	Pitchers     []int                     `json:"pitchers"`
	Bench        []int                     `json:"bench"`
	Bullpen      []int                     `json:"bullpen"`
	BattingOrder []int                     `json:"battingOrder"`
	Info         []BoxScoreTeamInfo        `json:"info"`
	Note         []LabelValue              `json:"note"`
}
type BoxScorePlayer struct {
	Person       IDFullNameLink           `json:"person"`
	JerseyNumber string                   `json:"jerseyNumber"`
	Position     Position                 `json:"position"`
	Stats        BoxScorePlayerStats      `json:"stats"`
	Status       BoxScorePlayerStatus     `json:"status"`
	ParentTeamID int                      `json:"parentTeamId"`
	BattingOrder string                   `json:"battingOrder"`
	SeasonStats  SeasonStats              `json:"seasonStats"`
	GameStatus   BoxScorePlayerGameStatus `json:"gameStatus"`
	AllPositions []AllPositions           `json:"allPositions,omitempty"`
}

// FetchBoxScore gets the latest game data from the MLB API
// and returns a list of games on the day specified by `t`
func FetchBoxScore(id string) *BoxScore {
	c := NewClient()
	req, _ := c.NewRequest("GET", "game/"+id+"/boxscore", map[string]string{})

	boxScoreStruct := &BoxScore{}
	c.Do(req, boxScoreStruct)

	return boxScoreStruct
}

func (p BoxScorePlayer) batterBoxScoreTableData() []string {
	var data []string
	var battingOrderString string
	if p.BattingOrder[len(p.BattingOrder)-1] == '0' {
		battingOrderString = string(p.BattingOrder[0]) + "."
	} else {
		battingOrderString = "   "
	}
	var posString string
	for i, pos := range p.AllPositions {
		sep := "-"
		if len(p.AllPositions) == i-1 || i == 0 {
			sep = ""
		}
		posString = posString + sep + pos.Abbreviation
	}
	var noteString string
	if p.Stats.Batting.Note == "" {
		noteString = " "
	} else {
		noteString = " " + p.Stats.Batting.Note
	}
	data = append(data, battingOrderString+noteString+p.Person.FullName+" "+posString, strconv.Itoa(p.Stats.Batting.AtBats), strconv.Itoa(p.Stats.Batting.Runs), strconv.Itoa(p.Stats.Batting.Hits), strconv.Itoa(p.Stats.Batting.Rbi), strconv.Itoa(p.Stats.Batting.BaseOnBalls), strconv.Itoa(p.Stats.Batting.StrikeOuts), strconv.Itoa(p.Stats.Batting.LeftOnBase), p.SeasonStats.Batting.Avg, p.SeasonStats.Batting.Ops)
	return data
}

func (p BoxScorePlayer) pitcherBoxScoreTableData() []string {
	var data []string
	data = append(
		data,
		p.Person.FullName+" "+p.Stats.Pitching.Note,
		p.Stats.Pitching.InningsPitched,
		strconv.Itoa(p.Stats.Pitching.Hits),
		strconv.Itoa(p.Stats.Pitching.Runs),
		strconv.Itoa(p.Stats.Pitching.EarnedRuns),
		strconv.Itoa(p.Stats.Pitching.BaseOnBalls),
		strconv.Itoa(p.Stats.Pitching.StrikeOuts),
		strconv.Itoa(p.Stats.Pitching.HomeRuns),
		p.SeasonStats.Pitching.Era,
	)
	return data
}

// BattersTableData formats a team's boxscore data for table printing
func (bst *BoxScoreTeam) BattersTableData() [][]string {
	var data [][]string
	for _, b := range bst.Batters {
		id := strconv.Itoa(b)
		player := bst.Players["ID"+id]
		if player.BattingOrder != "" {
			data = append(data, player.batterBoxScoreTableData())
		}
	}
	data = append(data)
	return data
}

// PitchersTableData formats a team's pitcher boxscore data for table printing
func (bst *BoxScoreTeam) PitchersTableData() [][]string {
	var data [][]string
	for _, b := range bst.Pitchers {
		id := strconv.Itoa(b)
		player := bst.Players["ID"+id]
		if player.BattingOrder != "" {
			data = append(data, player.pitcherBoxScoreTableData())
		}
	}
	data = append(data)
	return data
}

// PrintBattingTable formats a team's battinng boxscore data for table printing
func (bst *BoxScoreTeam) PrintBattingTable() {
	data := bst.BattersTableData()
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{bst.Team.LocationName + " Batters", "AB", "R", "H", "RBI", "BB", "SO", "LOB", "AVG", "OPS"})
	// table.SetBorder(false)
	table.AppendBulk(data)
	table.SetFooter([]string{"TOTALS", strconv.Itoa(bst.TeamStats.Batting.AtBats), strconv.Itoa(bst.TeamStats.Batting.Runs), strconv.Itoa(bst.TeamStats.Batting.Hits), strconv.Itoa(bst.TeamStats.Batting.Rbi), strconv.Itoa(bst.TeamStats.Batting.BaseOnBalls), strconv.Itoa(bst.TeamStats.Batting.StrikeOuts), strconv.Itoa(bst.TeamStats.Batting.LeftOnBase), "", ""})
	table.Render()
}

// PrintPitchingTable formats a team's battinng boxscore data for table printing
func (bst *BoxScoreTeam) PrintPitchingTable() {
	data := bst.PitchersTableData()
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{bst.Team.LocationName + " Pitchers", "IP", "H", "R", "ER", "BB", "SO", "HR", "ERA"})
	// table.SetBorder(false)
	table.AppendBulk(data)
	table.SetFooter([]string{"TOTALS", bst.TeamStats.Pitching.InningsPitched, strconv.Itoa(bst.TeamStats.Pitching.Hits), strconv.Itoa(bst.TeamStats.Pitching.Runs), strconv.Itoa(bst.TeamStats.Pitching.EarnedRuns), strconv.Itoa(bst.TeamStats.Pitching.BaseOnBalls), strconv.Itoa(bst.TeamStats.Pitching.StrikeOuts), strconv.Itoa(bst.TeamStats.Pitching.HomeRuns), ""})
	table.Render()
}

// Print prints a boxscore
func (bs *BoxScore) Print() {
	away := bs.Teams.Away
	home := bs.Teams.Home
	away.PrintBattingTable()
	fmt.Println()
	if len(away.Note) != 0 {
		for _, n := range away.Note {
			fmt.Println(n.Label + "-" + n.Value)
		}
		fmt.Println()
	}
	home.PrintBattingTable()
	fmt.Println()
	if len(home.Note) != 0 {
		for _, n := range home.Note {
			fmt.Println(n.Label + "-" + n.Value)
		}
		fmt.Println()
	}
	away.PrintPitchingTable()
	fmt.Println()
	home.PrintPitchingTable()
}
