package main

import (
	"GoAppOne/app/config"
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := &cli.App{
		Name:  "Migration",
		Usage: "Membuat migration dan seeders",
		Action: func(ctx *cli.Context) error {
			var command string = ctx.Args().Get(0)
			if command == "migrate" {
				fmt.Println("Run with migration")
				config.Run(true)
				return nil
			}
			config.Run(false)
			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
