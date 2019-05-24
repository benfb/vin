package util

import (
	"fmt"
	"strings"
	"time"
)

// ContainsString checks if slice s contains string e
func ContainsString(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

// ContainsStringAny checks if any string in slice s contains string e
func ContainsStringAny(s []string, e string) bool {
	for _, a := range s {
		if strings.Contains(a, e) {
			return true
		}
	}
	return false
}

// FindInStringSlice returns the location of the first string in slice s which contains string e
func FindInStringSlice(s []string, e string) int {
	if len(e) < 1 {
		return -1
	}
	for i, a := range s {
		if strings.Contains(a, e) {
			return i
		}
	}
	return -1
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
func FormatInning(inning string, isTop bool, status string) string {
	if status == "Final" {
		return fmt.Sprintf("Final %s", "\u2714")
	}
	if isTop {
		return fmt.Sprintf("%s %s", inning, "\u25B4")
	}
	return fmt.Sprintf("%s %s", inning, "\u25BE")
}
