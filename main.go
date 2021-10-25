package main

import (
	"log"
	"os"

	"github.com/urfave/cli"
	"scheduler/internal/cmd"
)

func main() {
	app := cli.NewApp()
	app.Name = "scheduler"
	app.Usage = "Example project"
	app.Version = "0.0.1"
	app.Commands = []cli.Command{cmd.Run}
	if err := app.Run(os.Args); err != nil {
		log.Fatalf("Failed to start application: %v", err)
	}
}
