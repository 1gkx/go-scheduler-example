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

func (s *Scheduler) addTask(j *job.Task) {
	s.tasks[j.Id] = j
	s.wg.Add(1)
}

func (s *Scheduler) deleteTask(j *job.Task) {
	j.Cancel()
	delete(s.tasks, j.Id)
}

func (s *Scheduler) Add(ctx context.Context, j job.Task) {
	ctx, j.Cancel = context.WithCancel(ctx)
	s.addTask(&j)
	go s.process(ctx, &j)
}

func (s *Scheduler) Stop() {
	for _, task := range s.tasks {
		task.Cancel()
	}
	s.wg.Wait()
}

func (s *Scheduler) process(ctx context.Context, j *job.Task) {
	interval, err := j.GetInterval()
	if err != nil {
		panic(err)
	}
	ticker := time.NewTicker(interval)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			j.Invoke(ctx, j)
			if !j.Repeatable {
				s.deleteTask(j)
			}
		case <-ctx.Done():
			s.wg.Done()
			return
		}
	}
}

//func (s *Scheduler) Delete() {}
