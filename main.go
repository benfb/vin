package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/jasonlvhit/gocron"

	"gopkg.in/urfave/cli.v1"
)

// Games represents a root
type Games struct {
	Date     string `xml:"date"`
	GameList []Game `xml:"game"`
}

// Game is an individual game
type Game struct {
	ID     string `xml:"id,attr"`
	Time   string `xml:"time,attr"`
	Status string `xml:"status,attr"`
}

// fetchGames gets the latest game data from the MLB API
// and returns a list of games on the day specified by `t`
func fetchGames(t time.Time) []Game {
	url := fmt.Sprint("http://gd2.mlb.com/components/game/mlb/year_", t.Year(), "/month_", padDate(int(t.Month())), "/day_", padDate(int(t.Day())), "/miniscoreboard.xml")
	fmt.Println("Getting game data...")

	gamesStruct := &Games{}
	getXML(url, gamesStruct)

	return gamesStruct.GameList
}

// getXML unmarshals an XML API response into a list of games
func getXML(url string, target *Games) error {
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

func padDate(toPad int) string {
	return fmt.Sprintf("%02d", toPad)
}

// sendNotification sends `message` to `phonenumber`
func sendNotification(phoneNumber, message string) error {
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

// hasTeam determines if the team `abbrv` is playing in `game`
func (game Game) hasTeam(abbrv string) bool {
	if len(abbrv) < 3 {
		log.Fatalln("Team abbreviation too short!")
	}
	abbrv = strings.ToLower(abbrv[:3])
	return strings.Contains(game.ID, abbrv)
}

// isOver determines whether or not a game is over
func (game Game) isOver() bool {
	if game.Status == "Final" {
		return true
	}

	return false
}

// checkGame checks to see if the game being played
// by `team` in `list` is over. If it is, it
// notifies `phone`
func checkGame(team, phone string, list []Game) {
	for _, g := range list {
		if g.hasTeam(team) {
			if g.isOver() {
				timeAvailable := time.Now().Add(90 * time.Minute)
				textString := fmt.Sprintf("The game is over! You can watch it at %02d:%02d", timeAvailable.Hour(), timeAvailable.Minute())
				err := sendNotification(phone, textString)
				if err != nil {
					log.Println("The game is over, but we couldn't notify that number!")
				}
				log.Println(textString)
				log.Println("The game is over and you were successfully notified!")
				os.Exit(0)
			} else {
				log.Println("The game is not over!")
			}
		}
	}
}

func main() {
	app := cli.NewApp()
	app.Name = "ball"
	app.Usage = "get notified when you can watch a blacked out mlb.tv game"
	app.Version = "0.1.0"
	app.Flags = []cli.Flag{
		&cli.StringFlag{
			Name:  "team, t",
			Value: "texas",
			Usage: "name of team to watch",
		},
		&cli.StringFlag{
			Name:  "phone, p",
			Usage: "phone number to notify when game is available",
		},
		&cli.Uint64Flag{
			Name:  "interval, i",
			Value: 20,
			Usage: "how often to check if a game is over (in seconds)",
		},
	}
	app.Action = func(c *cli.Context) error {
		if c.String("phone") != "" {
			gocron.Every(c.Uint64("interval")).Seconds().Do(checkGame, c.String("team"), c.String("phone"), fetchGames(time.Now()))
			<-gocron.Start()
		} else {
			return cli.NewExitError("Error! You must supply a phone number", 1)
		}
		return nil
	}

	app.Run(os.Args)
}
