package commands

import (
	"strings"

	"github.com/benfb/vin/api"
	"github.com/fatih/color"
)

// StandingsCmd gets standings from the xmlstats API
func StandingsCmd(division string) {
	standings := api.FetchStandings()

	if division == "agg" {
		standings.PrintMasterStandingsTable()
	} else {
		division = strings.ToUpper(division)[:3]

		shortDivisionMap := map[string]string{
			"ALE": "American League East",
			"ALC": "American League Central",
			"ALW": "American League West",
			"NLE": "National League East",
			"NLC": "National League Central",
			"NLW": "National League West",
		}

		for key := range shortDivisionMap {
			if key == division || division == "ALL" {
				color.New(color.FgHiCyan).Println("\n" + shortDivisionMap[key])
				standings.PrintStandingsTable(key)
			}
		}
	}
}
