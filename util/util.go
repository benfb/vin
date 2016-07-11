package util

import (
	"fmt"
	"net/http"
	"strings"
)

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

// SendNotification sends `message` to `phonenumber`
func SendNotification(phoneNumber, message string) error {
	body := strings.NewReader("number=" + phoneNumber + "&message=" + message)
	req, err := http.NewRequest("POST", "http://textbelt.com/text", body)
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	resp.Body.Close()
	return nil
}
