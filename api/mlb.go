package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/benfb/vin/util"
	"github.com/daviddengcn/go-colortext"
	"github.com/olekukonko/tablewriter"
)

// FormatURL takes a time and returns the appropriate API URL to call
func FormatURL(t time.Time) string {
	year := t.Year()
	month := util.PadDate(int(t.Month()))
	day := util.PadDate(int(t.Day()))
	return fmt.Sprint("https://statsapi.mlb.com/api/v1/schedule?sportId=1&date=", month, "%2F", day, "%2F", year)
}

// FetchGames gets the latest game data from the MLB API
// and returns a list of games on the day specified by `t`
func FetchGames(t time.Time) []ScheduleGame {
	url := FormatURL(t)

	scheduleStruct := &Schedule{}
	GetSchedule(url, scheduleStruct)

	return scheduleStruct.Dates[0].Games
}

// GetSchedule unmarshals an MLB JSON API response into an MLBRoot
func GetSchedule(url string, target *Schedule) error {
	resp, htmlErr := http.Get(url)

	if htmlErr != nil {
		return htmlErr
	}

	defer resp.Body.Close()

	r, ioErr := ioutil.ReadAll(resp.Body)

	if ioErr != nil {
		return ioErr
	}

	if jsonErr := json.Unmarshal(r, &target); jsonErr != nil {
		return jsonErr
	}

	return nil
}

func FetchLineScore(id string) *LineScore {
	ls := &LineScore{}
	GetLineScore("https://statsapi.mlb.com/api/v1/game/"+id+"/linescore", ls)

	return ls
}

func GetLineScore(url string, target *LineScore) error {
	resp, htmlErr := http.Get(url)

	if htmlErr != nil {
		return htmlErr
	}

	defer resp.Body.Close()

	r, ioErr := ioutil.ReadAll(resp.Body)

	if ioErr != nil {
		return ioErr
	}

	if jsonErr := json.Unmarshal(r, &target); jsonErr != nil {
		return jsonErr
	}

	return nil
}

// IsOver determines whether or not a game is over
func (g ScheduleGame) IsOver() bool {
	if g.Status.AbstractGameState == "Final" {
		return true
	}

	return false
}

// HasTeam determines if the team `abbrv` is playing in `game`
func (g ScheduleGame) HasTeam(abbrv string) bool {
	if len(abbrv) < 3 {
		log.Println("Team abbreviation too short!")
		return false
	}
	abbrv = strings.ToLower(abbrv[:3])
	return strings.Contains(strconv.Itoa(g.ID), abbrv)
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
	fmt.Printf("%s (%d - %d) @ %s (%d - %d)\n", sg.Teams.Away.Team.Name, sg.Teams.Away.LeagueRecord.Wins, sg.Teams.Away.LeagueRecord.Losses, sg.Teams.Home.Team.Name, sg.Teams.Home.LeagueRecord.Wins, sg.Teams.Home.LeagueRecord.Losses)
	ct.ResetColor()
	if sg.Status.AbstractGameState != "Preview" {
		data := [][]string{
			[]string{
				sg.Teams.Away.Team.Name,
				strconv.Itoa(ls.Teams.Away.Runs),
				strconv.Itoa(ls.Teams.Away.Hits),
				strconv.Itoa(ls.Teams.Away.Errors),
			},
			[]string{
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

// PrintProbablePitchers prints the probable pitchers (if any exist)
// func (g LineScore) PrintProbablePitchers() {
// 	emptyPitcher := ProbablePitcher{}
// 	if g.AwayProbablePitcher != emptyPitcher {
// 		fmt.Println("Probable pitchers:")
// 		away := g.AwayProbablePitcher
// 		home := g.HomeProbablePitcher
// 		data := [][]string{
// 			[]string{
// 				away.FirstName + " " + away.LastName,
// 				strconv.Itoa(away.Wins),
// 				strconv.Itoa(away.Losses),
// 			},
// 			[]string{
// 				home.FirstName + " " + home.LastName,
// 				strconv.Itoa(home.Wins),
// 				strconv.Itoa(home.Losses),
// 			},
// 		}
// 		table := tablewriter.NewWriter(os.Stdout)
// 		table.SetHeader([]string{"Pitcher", "Won", "Lost"})
// 		table.AppendBulk(data)
// 		table.Render()
// 		fmt.Println()
// 	}
// }
