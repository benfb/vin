package util

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/benfb/vin/api"
)

// ContainsString checks is slice s contains string e
func ContainsString(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

// Spinner prints out a cool spinner to prove that we're doing something
func Spinner() {
	defer fmt.Println()
	for {
		for _, r := range `-\|/` {
			fmt.Printf("\r%c", r)
			time.Sleep(50 * time.Millisecond)
		}
	}
}

// LocateTime locates a time in a place
func LocateTime(t time.Time, l string) (time.Time, error) {
	loc, err := time.LoadLocation(l)
	if err != nil {
		return time.Time{}, err
	}
	return t.In(loc), nil
}

// PadDate formats an integer as a two-digit string
func PadDate(toPad int) string {
	return fmt.Sprintf("%02d", toPad)
}

// FormatInning takes an inning and a half and returns a pretty-formatted string
func FormatInning(inning int, isTop bool, status string) string {
	if status == "Final" {
		return fmt.Sprintf("Final %s", "\u2713")
	}
	if isTop {
		return fmt.Sprintf("%d %s", inning, "\u23F6")
	}
	return fmt.Sprintf("%d %s", inning, "\u25BE")
}

// SendNotification sends `message` to `phonenumber` using the API at `apiURL`
func SendNotification(apiURL, phoneNumber, message string) ([]byte, error) {
	body := strings.NewReader("number=" + phoneNumber + "&message=" + message)
	req, err := http.NewRequest("POST", apiURL, body)
	if err != nil {
		return []byte{}, err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return []byte{}, err
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	return respBody, nil
}

// SeparateGames separates games
func SeparateGames(games []api.Game, with, without []string) []api.Game {
	validGames := []Game{}
	for _, g := range games {
		if g.FindTeam(with) && !g.FindTeam(without) {
			validGames = append(validGames, g)
		}
	}
	return validGames
}
