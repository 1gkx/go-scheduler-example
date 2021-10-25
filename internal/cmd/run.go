package cmd

import (
	"context"
	"github.com/urfave/cli"
	"os"
	"os/signal"
	"scheduler/internal/server"
)

var Run = cli.Command{
	Name:        "run",
	Usage:       "Start web server",
	Description: `Description`,
	Action:      run,
}

func run(c *cli.Context) {

	ctx := context.Background()

	srv := server.NewServer(ctx)
	go srv.Serve()
	defer srv.Shutdown()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, os.Interrupt)

	<-quit
}
