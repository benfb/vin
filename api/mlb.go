package api

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/benfb/vin/util"
	ct "github.com/daviddengcn/go-colortext"
	"github.com/olekukonko/tablewriter"
)

// FetchGames gets the latest game data from the MLB API
// and returns a list of games on the day specified by `t`
func FetchGames(t time.Time) []ScheduleGame {
	year := t.Year()
	month := util.PadDate(int(t.Month()))
	day := util.PadDate(int(t.Day()))
	scheduleStruct := &Schedule{}

	c := NewClient()
	req, _ := c.NewRequest("GET", "schedule", map[string]string{
		"sportId": "1",
		"date":    fmt.Sprint(month, "/", day, "/", year),
	})

	c.Do(req, scheduleStruct)

	return scheduleStruct.Dates[0].Games
}

// FetchLineScore gets a line score from the MLB API
func FetchLineScore(id string) *LineScore {
	ls := &LineScore{}

	c := NewClient()
	req, _ := c.NewRequest("GET", "game/"+id+"/linescore", nil)

	c.Do(req, ls)

	return ls
}

// IsOver determines whether or not a game is over
func (g ScheduleGame) IsOver() bool {
	return g.Status.AbstractGameState == "Final"
}

// ParseTime returns a game's time localized to the current time zone
func (g ScheduleGame) ParseTime() time.Time {
	return g.GameDate.Local()
}

// HasTeam determines if the team `abbrv` is playing in `game`
func (g ScheduleGame) HasTeam(abbrv string) bool {
	if len(abbrv) < 3 {
		log.Println("Team abbreviation too short!")
		return false
	}
	abbrv = strings.ToLower(abbrv[:3])
	return strings.Contains(strconv.Itoa(g.GamePk), abbrv)
}

// FindTeam determines if the team `team` is playing in `game`
func (g ScheduleGame) FindTeam(team string) bool {
	teamIndex := util.FindInStringSlice(TeamsWithLocs, team)
	if teamIndex == -1 {
		return false
	}
	team = TeamsWithLocs[util.FindInStringSlice(TeamsWithLocs, team)]
	return g.Teams.Away.Team.Name == team || g.Teams.Home.Team.Name == team
}

// PrintBoxScoreTable prints a box score to Stdout
func PrintBoxScoreTable(sg ScheduleGame, ls *LineScore) {
	ct.Foreground(ct.Cyan, true)
	gameTime := sg.ParseTime()
	fmt.Printf("%s (%d - %d) @ %s (%d - %d) [%v]\n", sg.Teams.Away.Team.Name, sg.Teams.Away.LeagueRecord.Wins, sg.Teams.Away.LeagueRecord.Losses, sg.Teams.Home.Team.Name, sg.Teams.Home.LeagueRecord.Wins, sg.Teams.Home.LeagueRecord.Losses, gameTime.Format("3:04PM"))
	ct.ResetColor()
	if sg.Status.AbstractGameState != "Preview" {
		data := [][]string{
			{
				sg.Teams.Away.Team.Name,
				strconv.Itoa(ls.Teams.Away.Runs),
				strconv.Itoa(ls.Teams.Away.Hits),
				strconv.Itoa(ls.Teams.Away.Errors),
			},
			{
				sg.Teams.Home.Team.Name,
				strconv.Itoa(ls.Teams.Home.Runs),
				strconv.Itoa(ls.Teams.Home.Hits),
				strconv.Itoa(ls.Teams.Home.Errors),
			},
		}
		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"Team", "Runs", "Hits", "Errs"})
		table.AppendBulk(data)
		table.Render()
		fmt.Println("Inning: " + util.FormatInning(ls.CurrentInningOrdinal, ls.IsTopInning, sg.Status.AbstractGameState))
		fmt.Println()
	}
}
