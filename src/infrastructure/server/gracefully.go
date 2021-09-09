package server

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var loggerHandler LoggerHandler

type GracefullyShutDown struct {
	httpserver *http.Server
}

func NewGracefullyShutDown(h http.Handler, address string) *GracefullyShutDown {
	return &GracefullyShutDown {
		httpserver: &http.Server{
			Handler: h,
			Addr: address,
		},
	}
}

func (r *GracefullyShutDown) RunGracefully() {
	var err error
	go func() {
		err = r.httpserver.ListenAndServe()

		if err != nil &&err != http.ErrServerClosed {
			loggerHandler.LogError("ERROR WHILE RUNNING GRACEFULLY %v", err)
		}
	}()
	loggerHandler.LogAccess("Running gracefully")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	loggerHandler.LogAccess("Stopping ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := r.httpserver.Shutdown(ctx); err != nil{
		loggerHandler.LogError("Server forced to stop ... %v", err)
	}
	loggerHandler.LogAccess("Gracefully Stopped")
	os.Exit(0)
}