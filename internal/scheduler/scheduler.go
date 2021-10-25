package scheduler

import (
	"context"
	"sync"
	"time"

	"scheduler/internal/job"
)

type Scheduler struct {
	wg            *sync.WaitGroup
	cancellations []context.CancelFunc
}

func NewScheduler() *Scheduler {
	return &Scheduler{
		wg:            new(sync.WaitGroup),
		cancellations: make([]context.CancelFunc, 0),
	}
}

func (s *Scheduler) Add(
	ctx context.Context,
	j job.Job,
	interval time.Duration,
) {
	ctx, cancel := context.WithCancel(ctx)
	s.cancellations = append(s.cancellations, cancel)
	s.wg.Add(1)
	go s.process(ctx, j, interval)
}

func (s *Scheduler) Stop() {
	for _, cancel := range s.cancellations {
		cancel()
	}
	s.wg.Wait()
}

func (s *Scheduler) process(
	ctx context.Context,
	j job.Job,
	interval time.Duration,
) {
	ticker := time.NewTicker(interval)
	for {
		select {
		case <-ticker.C:
			j(ctx)
		case <-ctx.Done():
			s.wg.Done()
			return
		}

	}
}

//func (s *Scheduler) Delete() {}
