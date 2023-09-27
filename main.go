package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"log"
	"net/http"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("ok"))
	})

	srv := &http.Server{
		Addr:         SrvBindAddr,
		WriteTimeout: TimeoutWrite,
		ReadTimeout:  TimeoutRead,
		IdleTimeout:  TimeoutIdle,
		Handler:      r,
	}

	// Run our server in a goroutine so that it doesn't block.
	if err := srv.ListenAndServe(); err != nil {
		log.Println(err)
	}
}
