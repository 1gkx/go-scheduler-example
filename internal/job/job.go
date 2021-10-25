package job

import (
	"context"
	"fmt"
	"time"
)

type TaskInterface interface {
	Invoke(ctx context.Context)
}

type Task struct {
	Id int
	Name string
	Params string
	Interval time.Duration
	Repeatable bool
	Fn func(ctx context.Context, t *Task)
	Cancel context.CancelFunc
}

func (ts *Task) Invoke(ctx context.Context, t *Task) {
	ts.Fn(ctx, t)
}

func Greeting(ctx context.Context, t *Task) {
	fmt.Printf("I`m task %s with id %d\n", t.Name, t.Id)
}

func ParseSubscriptionData(ctx context.Context) {
	fmt.Printf(
		"Subscription parsed successfuly at %s\n",
		time.Now().String(),
	)
}

func SendStatistics(ctx context.Context) {
	fmt.Printf(
		"Statistics send at %s\n",
		time.Now().String(),
	)
}
