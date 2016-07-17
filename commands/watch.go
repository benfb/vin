package commands

import (
	"fmt"
	"log"
	"net"
	"strconv"
	"time"

	"github.com/benfb/vin/api"
	"github.com/benfb/vin/util"
	"github.com/jasonlvhit/gocron"
)

// checkGame checks to see if the game being played
// by `team` in `list` is over. If it is, it
// notifies `phone`
func checkGame(team, phone string, sent time.Time) {
	now := time.Now()
	if now.Hour() >= 14 {
		list := api.FetchGames(now)
		for _, g := range list {
			if g.HasTeam(team) {
				if g.IsOver() {
					timeAvailable := now.Add(90 * time.Minute)
					localTimeAvailable := util.LocateTime(timeAvailable, "America/Chicago") // TODO: Don't hardcode this timezone
					textString := fmt.Sprintf("The game is over! You can watch it at %02d:%02d", localTimeAvailable.Hour(), localTimeAvailable.Minute())
					err := util.SendNotification(phone, textString)
					if err != nil {
						log.Println("The game is over, but we couldn't notify that number.")
					}
					log.Println("The game is over and you were successfully notified at \"" + phone + "\"! Sleeping...")
					time.Sleep(13 * time.Hour)
					checkGame(team, phone, now)
				} else {
					log.Println("The game is not over!")
				}
			}
		}
	} else {
		log.Println("The current time is not within the notification window. Waiting...")
		time.Sleep(10 * time.Minute)
	}
}

// WatchCmd does things
func WatchCmd(interval uint64, team, phone string) {
	go util.Spinner()

	gocron.Every(interval).Seconds().Do(checkGame, team, phone, time.Time{})
	<-gocron.Start()
}

// WatchClient does things
func WatchClient(interval uint64, team, phone string) {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	conn.Write([]byte(strconv.Itoa(int(interval)) + "," + team + "," + phone))

	buf := make([]byte, 1024)
	_, readErr := conn.Read(buf)
	if readErr != nil {
		log.Fatal(readErr)
	}
	log.Println(string(buf))
}
