package server

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"golang.org/x/sync/errgroup"
)

type Options struct {
	Port   string
	Router http.Handler
}

func New(options *Options) *Server {

	if options.Port == "" {
		panic("missing Port while creating server")
	}

	if options.Router == nil {
		panic("missing Router while creating server")
	}

	return &Server{
		Server: &http.Server{
			Addr:    fmt.Sprintf(":%s", options.Port),
			Handler: options.Router,
		},
	}
}

type Server struct {
	*http.Server
}

func (svr *Server) Start() error {
	mainCtx, cancel := context.WithCancel(context.Background())

	go func() {
		quit := make(chan os.Signal, 1)
		signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

		<-quit
		cancel()
	}()

	stages, stagesCtx := errgroup.WithContext(mainCtx)

	stages.Go(func() error {
		log.Println("server is ready to accept connections")
		return svr.ListenAndServe()
	})

	stages.Go(func() error {
		<-stagesCtx.Done()
		return svr.Shutdown(context.Background())
	})

	if err := stages.Wait(); errors.Is(err, http.ErrServerClosed) {
		log.Printf("server was shut down at %s", time.Now().UTC().String())
		return nil
	}

	return stages.Wait()
}
