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

	//worker := scheduler.NewScheduler()
	//task1 := job.Task{ Id: 1, Name: "task-1", Interval: time.Second*2, Repeatable: true, Fn: job.Greeting}
	//task2 := job.Task{ Id: 2, Name: "task-2", Interval: time.Second*5, Repeatable: false, Fn: job.Greeting}
	//
	//worker.Add(ctx, task1)
	//worker.Add(ctx, task2)

	go srv.Serve()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, os.Interrupt)

	<-quit
	srv.Shutdown()
}
