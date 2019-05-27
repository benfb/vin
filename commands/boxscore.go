package commands

import (
	"strconv"
	"strings"
	"time"

	"github.com/benfb/vin/api"
	"github.com/urfave/cli"
)

// BoxScoreCmd is the command run by `vin box`
func BoxScoreCmd(date, team string) error {
	timeFmtStr := "1/_2/06"
	if date == "today" {
		date = time.Now().Format(timeFmtStr)
	}
	parsedTime, timeErr := time.Parse(timeFmtStr, date)
	if timeErr != nil {
		return cli.NewExitError("Error! \""+date+"\" is not a valid date.", 1)
	}
	list := api.FetchGames(parsedTime)
	var gameIDs []int
	for _, g := range list {
		// fmt.Println(g.GamePk)
		// fmt.Println(team)
		if g.FindTeam(strings.Title(team)) {
			gameIDs = append(gameIDs, g.GamePk)
		}
	}
	for _, gID := range gameIDs {
		boxscore := api.FetchBoxScore(strconv.Itoa(gID))
		boxscore.Print()
	}

	return nil
}
