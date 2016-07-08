package api

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/benfb/vin/util"
)

// Games represents a root
type Games struct {
	Date     string `xml:"date"`
	GameList []Game `xml:"game"`
}

// Game is an individual game
type Game struct {
	ID           string `xml:"id,attr"`
	Time         string `xml:"time,attr"`
	Status       string `xml:"status,attr"`
	AwayTeam     string `xml:"away_team_name,attr"`
	HomeTeam     string `xml:"home_team_name,attr"`
	AwayTeamRuns int    `xml:"away_team_runs,attr"`
	HomeTeamRuns int    `xml:"home_team_runs,attr"`
	AwayTeamHits int    `xml:"away_team_hits,attr"`
	HomeTeamHits int    `xml:"home_team_hits,attr"`
	AwayTeamErrs int    `xml:"away_team_errors,attr"`
	HomeTeamErrs int    `xml:"home_team_errors,attr"`
}

// FetchGames gets the latest game data from the MLB API
// and returns a list of games on the day specified by `t`
func FetchGames(t time.Time) []Game {
	url := fmt.Sprint("http://gd2.mlb.com/components/game/mlb/year_", t.Year(), "/month_", util.PadDate(int(t.Month())), "/day_", util.PadDate(int(t.Day())), "/miniscoreboard.xml")
	fmt.Println("Getting game data...")

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
