package commands

import (
	"fmt"
	"log"
	"time"

	"github.com/benfb/vin/api"
	"github.com/benfb/vin/util"
)

// ResultsCmd is the command run by `vin results`
func ResultsCmd(date, team string) {
	if date == "" {
		date = time.Now().Format("1/_2/06")
	}
	parsedTime, timeErr := time.Parse("1/_2/06", date)
	if timeErr != nil {
		log.Fatalln("That is not a valid date!")
	}
	list := api.FetchGames(parsedTime)
	for _, g := range list {
		if g.FindTeam(team) || team == "all" {
			g.PrintBoxScoreTable()
			fmt.Println("Inning: " + util.FormatInning(g.Inning, g.IsTop, g.Status))
		}
	}
}
