package main

import (
	"goGINLearn/router"
	"golang.org/x/sync/errgroup"
	"net/http"
	"time"
)

func main() {
	var g errgroup.Group

	server := &http.Server{
		Addr:         "0.0.0.0:8888",
		Handler:      router.Router,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	g.Go(func() error {
		return server.ListenAndServe()
	})

	if err := g.Wait(); err != nil {
		panic(err)
	}

}
