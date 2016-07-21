package api

import (
	"testing"
	"time"
)

var testGameOne = Game{
	ID:           "2016/06/25/chnmlb-cinmlb-1",
	Time:         "6:10",
	Status:       "Final",
	Inning:       9,
	IsTop:        true,
	AwayTeam:     "Cubs",
	HomeTeam:     "Reds",
	AwayTeamRuns: 6,
	HomeTeamRuns: 2,
	AwayTeamHits: 10,
	HomeTeamHits: 5,
	AwayTeamErrs: 0,
	HomeTeamErrs: 1,
}

var testGameTwo = Game{
	ID:           "2016/06/25/phimlb-arimlb-1",
	Time:         "9:40",
	Status:       "In Progress",
	Inning:       7,
	IsTop:        false,
	AwayTeam:     "Phillies",
	HomeTeam:     "D-backs",
	AwayTeamRuns: 1,
	HomeTeamRuns: 4,
	AwayTeamHits: 3,
	HomeTeamHits: 6,
	AwayTeamErrs: 0,
	HomeTeamErrs: 0,
}

var exampleGames = Games{
	Date:     "20160625",
	GameList: []Game{testGameOne, testGameTwo},
}

func TestFormatURL(t *testing.T) {
	testLoc, _ := time.LoadLocation("America/Chicago")
	testTime := time.Date(2016, time.June, 30, 12, 50, 0, 0, testLoc)
	result := FormatURL(testTime)
	expected := "http://gd2.mlb.com/components/game/mlb/year_2016/month_06/day_30/miniscoreboard.json"
	if result != expected {
		t.Errorf("Got %v, expected %v.", result, expected)
	}
}

func TestIsOver(t *testing.T) {
	result := testGameOne.IsOver()
	if result != true {
		t.Errorf("Got %v, expected %v", result, true)
	}

	resultTwo := testGameTwo.IsOver()
	if resultTwo != false {
		t.Errorf("Got %v, expected %v", result, false)
	}
}

func TestHasTeam(t *testing.T) {
	result := testGameOne.HasTeam("tex")
	if result != false {
		t.Errorf("Got %v, expected %v", result, false)
	}

	resultTwo := testGameTwo.HasTeam("arizona")
	if resultTwo != true {
		t.Errorf("Got %v, expected %v", result, true)
	}

	resultFail := testGameOne.HasTeam("CH")
	if resultFail != false {
		t.Errorf("Got %v, expected %v", result, false)
	}
}

func TestFindTeam(t *testing.T) {
	result := testGameOne.FindTeam("Cubs")
	if result != true {
		t.Errorf("Got %v, expected %v", result, true)
	}
}
