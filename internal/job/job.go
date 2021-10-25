package job

import (
	"context"
	"fmt"
	"time"
)

type Job func(ctx context.Context)

func ParseSubscriptionData(ctx context.Context) {
	time.Sleep(time.Second * 1)
	fmt.Printf(
		"Subscription parsed successfuly at %s\n",
		time.Now().String(),
	)
}

func SendStatistics(ctx context.Context) {
	time.Sleep(time.Second*5)
	fmt.Printf(
		"Statistics send at %s\n",
		time.Now().String(),
	)
}
