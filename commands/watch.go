package commands

import (
	"fmt"
	"log"
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
	if now.Day() != sent.Day() {
		list := api.FetchGames(now)
		for _, g := range list {
			if g.HasTeam(team) {
				if g.IsOver() {
					timeAvailable := time.Now().Add(90 * time.Minute)
					localTimeAvailable := util.LocateTime(timeAvailable, "America/Chicago") // TODO: Don't hardcode this timezone
					textString := fmt.Sprintf("The game is over! You can watch it at %02d:%02d", localTimeAvailable.Hour(), localTimeAvailable.Minute())
					err := util.SendNotification(phone, textString)
					if err != nil {
						log.Println("The game is over, but we couldn't notify that number!")
					}
					log.Println("The game is over and you were successfully notified!")
					checkGame(team, phone, now)
				} else {
					log.Println("The game is not over!")
				}
			}
		}
	} else {
		log.Println("You were already notified today. Sleeping until tomorrow...")
		time.Sleep(12 * time.Hour)
	}
}

// WatchCmd does things
func WatchCmd(interval uint64, team, phone string) {
	gocron.Every(interval).Seconds().Do(checkGame, team, phone, time.Time{})
	<-gocron.Start()
}
