package scheduler

import (
	"context"
	"sync"
	"time"

	"scheduler/internal/job"
)

type Scheduler struct {
	wg            *sync.WaitGroup
	tasks map[int]*job.Task
}

func NewScheduler() *Scheduler {
	return &Scheduler{
		wg:            new(sync.WaitGroup),
		//cancellations: make([]context.CancelFunc, 0),
		tasks: make(map[int]*job.Task, 0),
	}
}

func (s *Scheduler) Add(ctx context.Context, j job.Task) {
	c, cancel := context.WithCancel(ctx)
	j.Cancel = cancel
	s.tasks[j.Id] = &j
	s.wg.Add(1)
	go s.process(c, j)
}

func (s *Scheduler) Stop() {
	for _, task := range s.tasks {
		task.Cancel()
	}
	s.wg.Wait()
}

func (s *Scheduler) process(ctx context.Context, j job.Task) {
	ticker := time.NewTicker(j.Interval)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			j.Fn(ctx)
			if !j.Repeatable {
				j.Cancel()
				delete(s.tasks, j.Id)
			}
		case <-ctx.Done():
			s.wg.Done()
			return
		}
	}
}

//func (s *Scheduler) Delete() {}
