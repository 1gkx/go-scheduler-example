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
	Id int `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Params string `json:"params,omitempty"`
	Interval string `json:"interval,omitempty"`
	Repeatable bool `json:"repeatable,omitempty"`
	Fn func(ctx context.Context, t *Task)
	Cancel context.CancelFunc
}

func (ts *Task) Invoke(ctx context.Context, t *Task) {
	ts.Fn(ctx, t)
}

func (ts *Task) GetInterval() (time.Duration, error) {
	return time.ParseDuration(ts.Interval)
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
