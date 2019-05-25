package commands

import (
	"fmt"
	"strings"

	"github.com/benfb/vin/api"
	ct "github.com/daviddengcn/go-colortext"
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
				ct.Foreground(ct.Cyan, true)
				fmt.Println("\n" + shortDivisionMap[key])
				ct.ResetColor()
				standings.PrintStandingsTable(key)
			}
		}
	}
}
