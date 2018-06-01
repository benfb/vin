package api

import (
	"encoding/json"
	"fmt"
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
	c := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return err
	}

	req.Header.Set("User-Agent", "Vin/1.0 (http://github.com/benfb/vin)")

	resp, htmlErr := c.Do(req)
	fmt.Println(url)
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
func (std Standings) RestrictLeague(league string) Standings {
	newStandings := []Standing{}
	for _, s := range std {
		if s.League == strings.ToUpper(league) {
			newStandings = append(newStandings, s)
		}
	}
	return newStandings
}

// RestrictDivision restricts standings to a particular division
func (std Standings) RestrictDivision(division string) Standings {
	newStandings := []Standing{}
	for _, s := range std {
		if s.Division == strings.ToUpper(division) {
			newStandings = append(newStandings, s)
		}
	}
	return newStandings
}

// PrintStandingsTable prints a standings table for a particular league and
// division to Stdout
func (std Standings) PrintStandingsTable(league, division string) {
	std = std.RestrictLeague(league).RestrictDivision(division)
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"#", "Team", "Pct", "Won", "Lost", "Back", "Str"})
	for i, s := range std {
		gamesBack := strconv.FormatFloat(s.GamesBack, 'G', 100, 32)
		if gamesBack == "0" {
			gamesBack = "-"
		}
		table.Append([]string{strconv.Itoa(i + 1), s.FirstName + " " + s.LastName, s.WinPercentage, strconv.Itoa(s.Won), strconv.Itoa(s.Lost), gamesBack, s.Streak})
	}
	table.Render()
}

// PrintMasterStandingsTable prints a game-wide standings table to Stdout
func (std Standings) PrintMasterStandingsTable() {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"#", "Team", "Pct", "Won", "Lost", "Str"})
	for i, s := range std {
		table.Append([]string{strconv.Itoa(i + 1), s.FirstName + " " + s.LastName, s.WinPercentage, strconv.Itoa(s.Won), strconv.Itoa(s.Lost), s.Streak})
	}
	table.Render()
}

func (std Standings) Len() int {
	return len(std)
}

func (std Standings) Less(i, j int) bool {
	return std[i].WinPercentage > std[j].WinPercentage
}

func (std Standings) Swap(i, j int) {
	std[i], std[j] = std[j], std[i]
}
