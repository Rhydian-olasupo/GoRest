package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	//Create App using CLI
	app := cli.NewApp()

	//add flags with three arguments
	app.Flags = []cli.Flag{
		&cli.StringFlag{
			Name:  "name",
			Value: "stranger",
			Usage: "your wonderful name",
		},

		&cli.IntFlag{
			Name:  "age",
			Value: 0,
			Usage: "Your graceful age",
		},
	}
	//This function parses and brings data in cli.Context struct
	app.Action = func(ctx *cli.Context) error {
		//c.String, C.Int looks for value of given flag
		log.Printf("Hello %s (%dyears), welcome to the command line world", ctx.String("name"), ctx.Int("age"))
		return nil
	}

	app.Run(os.Args)
}
