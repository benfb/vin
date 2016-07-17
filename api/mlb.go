package api

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/benfb/vin/util"
	"github.com/olekukonko/tablewriter"
)

// FormatURL takes a time and returns the appropriate API URL to call
func FormatURL(t time.Time) string {
	year := t.Year()
	month := util.PadDate(int(t.Month()))
	day := util.PadDate(int(t.Day()))
	return fmt.Sprint("http://gd2.mlb.com/components/game/mlb/year_", year, "/month_", month, "/day_", day, "/miniscoreboard.xml")
}

// FetchGames gets the latest game data from the MLB API
// and returns a list of games on the day specified by `t`
func FetchGames(t time.Time) []Game {
	url := FormatURL(t)
	// fmt.Println("Getting game data...")

	gamesStruct := &Games{}
	GetXML(url, gamesStruct)

	return gamesStruct.GameList
}

// GetXML unmarshals an XML API response into a list of games
func GetXML(url string, target *Games) error {
	resp, htmlErr := http.Get(url)

	if htmlErr != nil {
		return htmlErr
	}

	defer resp.Body.Close()

	r, ioErr := ioutil.ReadAll(resp.Body)

	if ioErr != nil {
		return ioErr
	}

	if xmlErr := xml.Unmarshal(r, &target); xmlErr != nil {
		return xmlErr
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
		log.Fatalln("Team abbreviation too short!")
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
}
