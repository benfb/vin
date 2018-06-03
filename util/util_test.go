package util

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestContainsString(t *testing.T) {
	if result := ContainsString([]string{"hello", "world"}, "world"); result != true {
		t.Errorf("Got %v, expected true.", result)
	}

	if result := ContainsString([]string{"hello", "world"}, "there"); result != false {
		t.Errorf("Got %v, expected false.", result)
	}
}

func TestContainsStringAny(t *testing.T) {
	if result := ContainsStringAny([]string{"hello", "world"}, "world"); result != true {
		t.Errorf("Got %v, expected true.", result)
	}

	if result := ContainsStringAny([]string{"hello", "world"}, "hel"); result != true {
		t.Errorf("Got %v, expected true.", result)
	}

	if result := ContainsStringAny([]string{"hello", "world"}, "there"); result != false {
		t.Errorf("Got %v, expected false.", result)
	}
}

func TestFindInStringSlice(t *testing.T) {
	if result := FindInStringSlice([]string{"hello", "world"}, "world"); result != 1 {
		t.Errorf("Got %v, expected 1.", result)
	}

	if result := FindInStringSlice([]string{"hello", "world"}, "hel"); result != 0 {
		t.Errorf("Got %v, expected 0.", result)
	}

	if result := FindInStringSlice([]string{"hello", "world"}, ""); result != -1 {
		t.Errorf("Got %v, expected -1.", result)
	}

	if result := FindInStringSlice([]string{"hello", "world"}, "there"); result != -1 {
		t.Errorf("Got %v, expected -1.", result)
	}
}

func TestLocateTime(t *testing.T) {
	result, _ := LocateTime(time.Now(), "America/Chicago")
	central, _ := time.LoadLocation("America/Chicago")
	if result.Location().String() != central.String() {
		t.Errorf("Got %v, expected %v.", result, central)
	}

	_, err := LocateTime(time.Now(), "Some bad location")
	if err == nil {
		t.Errorf("Got nil, expected an error.")
	}
}

func TestPadDate(t *testing.T) {
	result := PadDate(3)
	expected := "03"
	if result != expected {
		t.Errorf("Got %v, expected %v.", result, expected)
	}
}

func TestFormatInning(t *testing.T) {
	result := FormatInning("1st", true, "")
	expected := "1st \u25B4"
	if result != expected {
		t.Errorf("Got %v, expected %v.", result, expected)
	}

	resultBottom := FormatInning("1st", false, "")
	expectedBottom := "1st \u25BE"
	if resultBottom != expectedBottom {
		t.Errorf("Got %v, expected %v.", resultBottom, expectedBottom)
	}

	resultFinal := FormatInning("9th", false, "Final")
	expectedFinal := "Final \u2714"
	if resultFinal != expectedFinal {
		t.Errorf("Got %v, expected %v.", resultFinal, expectedFinal)
	}
}

func TestSendNotification(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "{\"success\":true}")
	}))
	defer ts.Close()

	res, err := SendNotification(ts.URL, "5555555555", "test")
	if err != nil {
		log.Fatal(err)
	}
	result := string(res)
	expected := "{\"success\":true}\n"
	if result != expected {
		t.Errorf("Got %v, expected %v.", result, expected)
	}

	_, reqErr := SendNotification("badurl", "5555555555", "badtext")
	if reqErr == nil {
		t.Errorf("Got nil, expected an error.")
	}
}
