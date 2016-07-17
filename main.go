package main

import (
	"os"

	"github.com/benfb/vin/commands"

	"gopkg.in/urfave/cli.v1"
)

func main() {
	app := cli.NewApp()
	app.Name = "vin"
	app.Usage = "the baseball command-line companion"
	app.Version = "0.2.0"
	app.Commands = []cli.Command{
		{
			Name:    "watch",
			Aliases: []string{"w"},
			Usage:   "get notified when a blacked-out game is available",
			Flags: []cli.Flag{
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
			},
			Action: func(c *cli.Context) error {
				if c.String("phone") != "" {
					commands.WatchCmd(c.Uint64("interval"), c.String("team"), c.String("phone"))
				} else {
					return cli.NewExitError("Error! You must supply a phone number", 1)
				}
				return nil
			},
		},
		{
			Name:    "standings",
			Aliases: []string{"s"},
			Usage:   "get the current standings",
			Flags: []cli.Flag{
				&cli.BoolFlag{
					Name:  "aggregate, a",
					Usage: "get all standings in one table",
				},
			},
			Action: func(c *cli.Context) error {
				var division string
				if c.Args().Get(0) != "" {
					division = c.Args().Get(0)
				} else if c.Bool("aggregate") {
					division = "agg"
				} else {
					division = "all"
				}
				commands.StandingsCmd(division)
				return nil
			},
		},
		{
			Name:    "results",
			Aliases: []string{"r"},
			Usage:   "get results for all the games from a particular day, formatted as m/d/yy",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:  "team, t",
					Value: "all",
					Usage: "name of team to watch",
				},
			},
			Action: func(c *cli.Context) error {
				var day string
				if c.Args().Get(0) != "" {
					day = c.Args().Get(0)
				}
				commands.ResultsCmd(day, c.String("team"))
				return nil
			},
		},
	}

	app.Run(os.Args)
}
