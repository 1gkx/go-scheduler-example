package cmd

import (
	"context"
	"github.com/urfave/cli"
	"os"
	"os/signal"
	"scheduler/internal/job"
	"scheduler/internal/scheduler"
	"time"
)

var Run = cli.Command{
	Name:        "run",
	Usage:       "Start web server",
	Description: `Description`,
	Action:      run,
}

func run(c *cli.Context) {

	ctx := context.Background()

	worker := scheduler.NewScheduler()
	worker.Add(ctx, job.ParseSubscriptionData, time.Second*5)
	worker.Add(ctx, job.SendStatistics, time.Second*10)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, os.Interrupt)

	<-quit
	worker.Stop()
}
