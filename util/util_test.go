package util

import (
	"fmt"
	"testing"
	"time"
)

func TestContainsString(t *testing.T) {
	if result := ContainsString([]string{"hello", "world"}, "world"); result != true {
		t.Errorf("Got %v, expected true.", result)
	}
}

func TestLocateTime(t *testing.T) {
	result := LocateTime(time.Now(), "America/Chicago")
	central, _ := time.LoadLocation("America/Chicago")
	if result.Location().String() != central.String() {
		t.Errorf("Got %v, expected %v.", result, central)
	}
}

func TestFormatInning(t *testing.T) {
	result := FormatInning(1, true, "")
	expected := fmt.Sprintf("%d %s", 1, "\u23F6")
	if result != expected {
		t.Errorf("Got %v, expected %v.", result, expected)
	}
}
