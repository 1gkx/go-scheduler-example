package job

import (
	"context"
	"fmt"
	"time"
)

type Task struct {
	Id int
	Interval time.Duration
	Repeatable bool
	Fn func(ctx context.Context)
	Cancel context.CancelFunc
}

func ParseSubscriptionData(ctx context.Context) {
	time.Sleep(time.Second * 1)
	fmt.Printf(
		"Subscription parsed successfuly at %s\n",
		time.Now().String(),
	)
}

func SendStatistics(ctx context.Context) {
	time.Sleep(time.Second*3)
	fmt.Printf(
		"Statistics send at %s\n",
		time.Now().String(),
	)
}
