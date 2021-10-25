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
	"time"
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
		s.Add(ctx, t)
		data, err := json.Marshal(t)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write(data)
	}
}

func (s *Srv) Serve() {
	fmt.Printf("Server start at %s port\n", PORT)
	if err := http.ListenAndServe(PORT, s.r); err != nil {
		log.Fatal(err)
	}
}

func (s *Srv) Shutdown() {
	fmt.Printf("Server stop at %s\n", time.Now().String())
	s.s.Stop()
}
