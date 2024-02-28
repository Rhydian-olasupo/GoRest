package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := cli.NewApp()

	app.Flags = []cli.Flag{
		&cli.StringFlag{
			Name:  "save",
			Value: "no",
			Usage: "Should save to database (yes/no)",
		},
	}

	app.Version = "1.0"
	//define Action
	app.Action = func(ctx *cli.Context) error {
		var args []string
		if ctx.NArg() > 0 {
			//Fetch arguments in an array
			args := ctx.Args().Get(0)
			personName := args[0]
			marks := args[1:len(args)]
			log.Println("Person: ", personName)
			log.Println("marks", marks)
		}
		//Check the flag value
		if ctx.String("save") == "no" {
			log.Println("Skipping saving to the database")
		} else {
			//Add database logiv here
			log.Println("Saving to the database", args)
		}

		return nil
	}
	app.Run(os.Args)
}
