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
	return fmt.Sprint("http://gd2.mlb.com/components/game/mlb/year_", year, "/month_", month, "/day_", day, "/miniscoreboard.json")
}

// FetchGames gets the latest game data from the MLB API
// and returns a list of games on the day specified by `t`
func FetchGames(t time.Time) []Game {
	url := FormatURL(t)
	// fmt.Println("Getting game data...")

	gamesStruct := &MLBRoot{}
	GetMLB(url, gamesStruct)

	return gamesStruct.Data.Games.GameList
}

// GetMLB unmarshals an MLB JSON API response into an MLBRoot
func GetMLB(url string, target *MLBRoot) error {
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
func (game Game) IsOver() bool {
	if game.Status == "Final" {
		return true
	}

	return false
}

// HasTeam determines if the team `abbrv` is playing in `game`
func (game Game) HasTeam(abbrv string) bool {
	if len(abbrv) < 3 {
		log.Println("Team abbreviation too short!")
		return false
	}
	abbrv = strings.ToLower(abbrv[:3])
	return strings.Contains(game.ID, abbrv)
}

// FindTeam determines if the team `team` is playing in `game`
func (game Game) FindTeam(team string) bool {
	return game.AwayTeam == team || game.HomeTeam == team
}

// PrintBoxScoreTable prints a box score to Stdout
func (game Game) PrintBoxScoreTable() {
	ct.Foreground(ct.Cyan, true)
	fmt.Printf("%s (%d - %d) @ %s (%d - %d)\n", game.AwayTeam, game.AwayTeamWins, game.AwayTeamLosses, game.HomeTeam, game.HomeTeamWins, game.HomeTeamLosses)
	ct.ResetColor()
	if game.Status != "Preview" {
		data := [][]string{
			[]string{
				game.AwayTeam,
				strconv.Itoa(game.AwayTeamRuns),
				strconv.Itoa(game.AwayTeamHits),
				strconv.Itoa(game.AwayTeamErrs),
			},
			[]string{
				game.HomeTeam,
				strconv.Itoa(game.HomeTeamRuns),
				strconv.Itoa(game.HomeTeamHits),
				strconv.Itoa(game.HomeTeamErrs),
			},
		}
		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"Team", "Runs", "Hits", "Errs"})
		table.AppendBulk(data)
		table.Render()
		fmt.Println("Inning: " + util.FormatInning(game.Inning, game.IsTop, game.Status))
		fmt.Println()
	}
}

// PrintProbablePitchers prints the probable pitchers (if any exist)
func (game Game) PrintProbablePitchers() {
	emptyPitcher := ProbablePitcher{}
	if game.AwayProbablePitcher != emptyPitcher {
		fmt.Println("Probable pitchers:")
		away := game.AwayProbablePitcher
		home := game.HomeProbablePitcher
		data := [][]string{
			[]string{
				away.FirstName + " " + away.LastName,
				strconv.Itoa(away.Wins),
				strconv.Itoa(away.Losses),
			},
			[]string{
				home.FirstName + " " + home.LastName,
				strconv.Itoa(home.Wins),
				strconv.Itoa(home.Losses),
			},
		}
		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"Pitcher", "Won", "Lost"})
		table.AppendBulk(data)
		table.Render()
		fmt.Println()
	}
}
