package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/olekukonko/tablewriter"
)

// FetchStandings gets the latest game data from the MLB API
// and returns a list of games on the day specified by `t`
func FetchStandings() Standings {
	url := "https://erikberg.com/mlb/standings.json"

	standingsStruct := &StandingsResponse{}
	GetStandingsJSON(url, standingsStruct)

	return standingsStruct.StandingList
}

// GetStandingsJSON unmarshals an XML API response into a list of games
func GetStandingsJSON(url string, target *StandingsResponse) error {
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

// RestrictLeague restricts standings to a particular league
func (standings Standings) RestrictLeague(league string) Standings {
	newStandings := []Standing{}
	for _, s := range standings {
		if s.League == strings.ToUpper(league) {
			newStandings = append(newStandings, s)
		}
	}
	return newStandings
}

// RestrictDivision restricts standings to a particular division
func (standings Standings) RestrictDivision(division string) Standings {
	newStandings := []Standing{}
	for _, s := range standings {
		if s.Division == strings.ToUpper(division) {
			newStandings = append(newStandings, s)
		}
	}
	return newStandings
}

// PrintStandingsTable prints a standings table for a particular league and
// division to Stdout
func (standings Standings) PrintStandingsTable(league, division string) {
	standings = standings.RestrictLeague(league).RestrictDivision(division)
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"#", "Team", "Pct", "Won", "Lost", "Back", "Str"})
	for i, s := range standings {
		gamesBack := strconv.FormatFloat(s.GamesBack, 'G', 100, 32)
		if gamesBack == "0" {
			gamesBack = "-"
		}
		table.Append([]string{strconv.Itoa(i + 1), s.FirstName + " " + s.LastName, s.WinPercentage, strconv.Itoa(s.Won), strconv.Itoa(s.Lost), gamesBack, s.Streak})
	}
	table.Render()
}

// PrintMasterStandingsTable prints a game-wide standings table to Stdout
func (standings Standings) PrintMasterStandingsTable() {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"#", "Team", "Pct", "Won", "Lost", "Str"})
	for i, s := range standings {
		table.Append([]string{strconv.Itoa(i + 1), s.FirstName + " " + s.LastName, s.WinPercentage, strconv.Itoa(s.Won), strconv.Itoa(s.Lost), s.Streak})
	}
	table.Render()
}

// Len is implemented to make standings sortable
func (standings Standings) Len() int {
	return len(standings)
}

// Less is implemented to make standings sortable
func (standings Standings) Less(i, j int) bool {
	return standings[i].WinPercentage > standings[j].WinPercentage
}

// Swap is implemented to make standings sortable
func (standings Standings) Swap(i, j int) {
	standings[i], standings[j] = standings[j], standings[i]
}
