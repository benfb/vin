package api

// import (
// 	"testing"
// 	"time"
// )

// var testGameOne = ScheduleGame{
// 	ID:   530274,
// 	Time: "2018-06-02T23:15:00Z",
// 	Status: ScheduleGameStatus{
// 		AbstractGameState: "Final",
// 		CodedGameState:    "F",
// 		DetailedState:     "Final",
// 		StatusCode:        "F",
// 		AbstractGameCode:  "F",
// 	},
// 	Teams: ScheduleGameTeams{
// 		Away: GameTeam{
// 			LeagueRecord: GameTeamLeagueRecord{
// 				Wins:   28,
// 				Losses: 30,
// 				Pct:    ".483",
// 			},
// 			Score: 12,
// 			Team: GameTeamTeam{
// 				ID:   119,
// 				Name: "Los Angeles Dodgers",
// 				Link: "/api/v1/teams/119",
// 			},
// 			IsWinner:     true,
// 			SplitSquad:   false,
// 			SeriesNumber: 19,
// 		},
// 		Home: GameTeam{
// 			LeagueRecord: GameTeamLeagueRecord{
// 				Wins:   30,
// 				Losses: 28,
// 				Pct:    ".517",
// 			},
// 			Score: 4,
// 			Team: GameTeamTeam{
// 				ID:   115,
// 				Name: "Colorado Rockies",
// 				Link: "/api/v1/teams/115",
// 			},
// 			IsWinner:     false,
// 			SplitSquad:   false,
// 			SeriesNumber: 19,
// 		},
// 	},
// }

// var testGameTwo = ScheduleGame{
// 	ID:   530299,
// 	Time: "2018-06-03T17:35:00Z",
// 	Status: ScheduleGameStatus{
// 		AbstractGameState: "Live",
// 		CodedGameState:    "I",
// 		DetailedState:     "In Progress",
// 		StatusCode:        "I",
// 		AbstractGameCode:  "L",
// 	},
// 	Teams: ScheduleGameTeams{
// 		Away: GameTeam{
// 			LeagueRecord: GameTeamLeagueRecord{
// 				Wins:   33,
// 				Losses: 24,
// 				Pct:    ".579",
// 			},
// 			Score: 0,
// 			Team: GameTeamTeam{
// 				ID:   120,
// 				Name: "Washington Nationals",
// 				Link: "/api/v1/teams/120",
// 			},
// 			SplitSquad:   false,
// 			SeriesNumber: 19,
// 		},
// 		Home: GameTeam{
// 			LeagueRecord: GameTeamLeagueRecord{
// 				Wins:   34,
// 				Losses: 24,
// 				Pct:    ".586",
// 			},
// 			Score: 1,
// 			Team: GameTeamTeam{
// 				ID:   144,
// 				Name: "Atlanta Braves",
// 				Link: "/api/v1/teams/144",
// 			},
// 			SplitSquad:   false,
// 			SeriesNumber: 19,
// 		},
// 	},
// }

// var exampleGames = ScheduleDate{
// 	Date:  "2018-06-02",
// 	Games: []ScheduleGame{testGameOne, testGameTwo},
// }

// func TestFormatURL(t *testing.T) {
// 	testLoc, _ := time.LoadLocation("America/Chicago")
// 	testTime := time.Date(2016, time.June, 30, 12, 50, 0, 0, testLoc)
// 	result := FormatURL(testTime)
// 	expected := "https://statsapi.mlb.com/api/v1/schedule?sportId=1&date=06%2F30%2F2016"
// 	if result != expected {
// 		t.Errorf("Got %v, expected %v.", result, expected)
// 	}
// }

// func TestIsOver(t *testing.T) {
// 	result := testGameOne.IsOver()
// 	if result != true {
// 		t.Errorf("Got %v, expected %v", result, true)
// 	}

// 	resultTwo := testGameTwo.IsOver()
// 	if resultTwo != false {
// 		t.Errorf("Got %v, expected %v", result, false)
// 	}
// }

// // func TestHasTeam(t *testing.T) {
// // 	result := testGameOne.HasTeam("tex")
// // 	if result != false {
// // 		t.Errorf("Got %v, expected %v", result, false)
// // 	}

// // 	resultTwo := testGameTwo.HasTeam("WAS")
// // 	if resultTwo != true {
// // 		t.Errorf("Got %v, expected %v", result, true)
// // 	}

// // 	resultFail := testGameOne.HasTeam("CH")
// // 	if resultFail != false {
// // 		t.Errorf("Got %v, expected %v", result, false)
// // 	}
// // }

// func TestFindTeam(t *testing.T) {
// 	resultOne := testGameOne.FindTeam("Rockies")
// 	if resultOne != true {
// 		t.Errorf("Got %v, expected %v", resultOne, true)
// 	}

// 	resultTwo := testGameOne.FindTeam("Cubs")
// 	if resultTwo != false {
// 		t.Errorf("Got %v, expected %v", resultTwo, false)
// 	}

// 	resultThree := testGameOne.FindTeam("")
// 	if resultThree != false {
// 		t.Errorf("Got %v, expected %v", resultThree, false)
// 	}
// }
