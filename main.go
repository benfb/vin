package main

import (
	"fmt"
	"os"

	"github.com/benfb/vin/commands"

	"gopkg.in/urfave/cli.v1"
)

func main() {
	app := cli.NewApp()
	app.Name = "vin"
	app.Usage = "the baseball command-line companion"
	app.Version = "0.2.0"
	app.Authors = []cli.Author{
		cli.Author{
			Name:  "Ben Bailey",
			Email: "bennettbailey@gmail.com",
		},
	}
	app.Commands = []cli.Command{
		{
			Name:      "watch",
			Aliases:   []string{"w"},
			Usage:     "get notified when a blacked-out game is available",
			ArgsUsage: "[phone]",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:  "team, t",
					Value: "texas",
					Usage: "name of team to watch",
				},
				&cli.Uint64Flag{
					Name:  "interval, i",
					Value: 20,
					Usage: "how often to check if a game is over (in seconds)",
				},
			},
			Action: func(c *cli.Context) error {
				phone := c.Args().Get(0)
				if fmt.Sprintf("%T", phone) == "string" && phone != "" {
					commands.WatchClient(c.Uint64("interval"), c.String("team"), phone)
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
				division := c.Args().Get(0)
				if division == "" && c.Bool("aggregate") {
					division = "agg"
				} else if division == "" {
					division = "all"
				}
				commands.StandingsCmd(division)
				return nil
			},
		},
		{
			Name:      "server",
			Aliases:   []string{"serve", "se"},
			Usage:     "run a vin server",
			ArgsUsage: "[address]",
			Action: func(c *cli.Context) error {
				commands.ServerCmd()
				return nil
			},
		},
		{
			Name:      "results",
			Aliases:   []string{"r"},
			Usage:     "get results for all the games from a particular day, formatted as mm/dd/yy",
			ArgsUsage: "[date]",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:  "team, t",
					Value: "all",
					Usage: "name of team to get box score for",
				},
			},
			Action: func(c *cli.Context) error {
				day := c.Args().Get(0)
				if day == "" {
					day = "today"
				}
				return commands.ResultsCmd(day, c.String("team"))
			},
		},
	}

	app.Run(os.Args)
}
