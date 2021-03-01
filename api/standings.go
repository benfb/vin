package api

import (
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/olekukonko/tablewriter"
)

// FetchStandings gets the latest game data from the MLB API
// and returns a list of games on the day specified by `t`
func FetchStandings() *Standings {
	c := NewClient()
	req, _ := c.NewRequest("GET", "standings", map[string]string{
		"leagueId": "103,104",
		"season":   "2019",
	})

	standingsStruct := &Standings{}
	c.Do(req, standingsStruct)

	return standingsStruct
}

// RestrictLeague restricts standings to a particular league
func RestrictLeague(records []StandingsRecord, league string) []StandingsRecord {
	var newStandings []StandingsRecord
	for _, s := range records {
		if s.League.ID == leagueIDMap[strings.ToUpper(league)] {
			newStandings = append(newStandings, s)
		}
	}
	return newStandings
}

// RestrictDivision restricts standings to a particular division
func RestrictDivision(records []StandingsRecord, division string) []StandingsRecord {
	var newStandings []StandingsRecord
	for _, s := range records {
		if s.Division.ID == divisionIDMap[strings.ToUpper(division)] {
			newStandings = append(newStandings, s)
		}
	}
	return newStandings
}

// PrintStandingsTable prints a standings table for a particular league and
// division to Stdout
func (std *Standings) PrintStandingsTable(division string) {
	narrowed := RestrictDivision(std.Records, division)
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"#", "Team", "Pct", "Won", "Lost", "Back", "Str"})
	for i, s := range narrowed[0].TeamRecords {
		gamesBack := s.GamesBack
		if gamesBack == "0" {
			gamesBack = "-"
		}
		table.Append([]string{strconv.Itoa(i + 1), s.Team.Name, s.WinningPercentage, strconv.Itoa(s.Wins), strconv.Itoa(s.Losses), gamesBack, s.Streak.StreakCode})
	}
	table.Render()
}

// PrintMasterStandingsTable prints a game-wide standings table to Stdout
func (std *Standings) PrintMasterStandingsTable() {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"#", "Team", "Pct", "Won", "Lost", "Str"})
	trs := std.FlattenToTeamRecordsSlice()
	sort.SliceStable(trs, func(i, j int) bool {
		return trs[i].WinningPercentage > trs[j].WinningPercentage
	})
	for i, s := range trs {
		table.Append([]string{strconv.Itoa(i + 1), s.Team.Name, s.WinningPercentage, strconv.Itoa(s.Wins), strconv.Itoa(s.Losses), s.Streak.StreakCode})
	}
	table.Render()
}

// FlattenToTeamRecordsSlice takes a Standings pointer and returns a flattened slice of TeamRecords
func (std *Standings) FlattenToTeamRecordsSlice() []StandingsTeamRecord {
	var trs []StandingsTeamRecord
	for _, r := range std.Records {
		trs = append(trs, r.TeamRecords...)
	}
	return trs
}
