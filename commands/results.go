package commands

import (
	"log"
	"os"
	"strconv"
	"time"

	"github.com/benfb/vin/api"
	"github.com/olekukonko/tablewriter"
)

// ResultsCmd does things
func ResultsCmd(date string) {
	parsedTime, timeErr := time.Parse("1/_2/06", date)
	if timeErr != nil {
		log.Fatalln("That is not a valid date!")
	}
	list := api.FetchGames(parsedTime)
	for _, g := range list {
		data := [][]string{
			[]string{g.AwayTeam, strconv.Itoa(g.AwayTeamRuns), strconv.Itoa(g.AwayTeamHits), strconv.Itoa(g.AwayTeamErrs)},
			[]string{g.HomeTeam, strconv.Itoa(g.HomeTeamRuns), strconv.Itoa(g.HomeTeamHits), strconv.Itoa(g.HomeTeamErrs)},
		}
		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"Team", "Runs", "Hits", "Errs"})
		table.AppendBulk(data)
		table.Render()
	}

}
