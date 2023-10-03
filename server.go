package main

import (
	"fmt"
	"net/http"
	"time"
)

func StartServer(addr int) error {
	srv := http.Server{
		Addr:              fmt.Sprintf(":%d", addr),
		Handler:           HandleRoutes(),
		IdleTimeout:       30 * time.Second,
		WriteTimeout:      30 * time.Second,
		ReadTimeout:       30 * time.Second,
		ReadHeaderTimeout: 30 * time.Second,
	}

	fmt.Printf("Backend started at %d\n", addr)

	return srv.ListenAndServe()
}
