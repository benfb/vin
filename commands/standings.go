package commands

import (
	"fmt"
	"sort"
	"strings"

	"github.com/benfb/vin/api"
	"github.com/daviddengcn/go-colortext"
)

// StandingsCmd gets standings from the xmlstats API
func StandingsCmd(division string) {
	standings := api.FetchStandings()
	sort.Sort(standings)

	if division == "agg" {
		standings.PrintMasterStandingsTable()
	} else {
		division = strings.ToUpper(division)[:3]
		leagueDivisionMap := map[string][]string{
			"AL": []string{"E", "C", "W"},
			"NL": []string{"E", "C", "W"},
		}

		shortDivisionMap := map[string]string{
			"ALE": "American League East",
			"ALC": "American League Central",
			"ALW": "American League West",
			"NLE": "National League East",
			"NLC": "National League Central",
			"NLW": "National League West",
		}

		for league, divisionSlice := range leagueDivisionMap {
			for _, standingsDivision := range divisionSlice {
				if league+standingsDivision == division || division == "ALL" {
					ct.Foreground(ct.Cyan, true)
					fmt.Println("\n" + shortDivisionMap[league+standingsDivision])
					ct.ResetColor()
					standings.PrintStandingsTable(league, standingsDivision)
				}
			}
		}
	}
}
