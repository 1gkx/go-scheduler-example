package server

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"scheduler/internal/job"
	"scheduler/internal/scheduler"
)

type Srv struct {
	ctx context.Context
	s *scheduler.Scheduler
	r *mux.Router
}

var PORT = ":3333"

func NewServer(ctx context.Context) *Srv{

	sh := scheduler.NewScheduler()

	route := mux.NewRouter().StrictSlash(true)
	route.HandleFunc("/add", addTask(ctx, sh)).Methods("POST")

	return &Srv{
		ctx: ctx,
		r: route,
		s: sh,
	}
}

func addTask (ctx context.Context, s *scheduler.Scheduler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var t job.Task
		if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		t.Fn = job.Greeting

		fmt.Printf("Task: %+v\n", t)
		s.Add(ctx, t)

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Ok"))
	}
}

func (s *Srv) Serve() {
	fmt.Printf("Server start at %s port\n", PORT)
	if err := http.ListenAndServe(PORT, s.r); err != nil {
		log.Fatal(err)
	}
}

func (s *Srv) Shutdown() {
	s.s.Stop()
}
