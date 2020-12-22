package stop

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

type Handler struct{}

func (Handler) ServeHTTP(http.ResponseWriter, *http.Request) {
	var srv http.Server

	idleConnsClosed := make(chan struct{})
	go func() {
		dur, _ := time.ParseDuration("5sec")
		ctx, ctxCancel := context.WithTimeout(context.Background(), dur)
		defer ctxCancel()

		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt)
		<-sigint

		// We received an interrupt signal, shut down.
		if err := srv.Shutdown(ctx); err != nil {
			// Error from closing listeners, or context timeout:
			log.Printf("HTTP server Shutdown: %v", err)
		}
		close(idleConnsClosed)
	}()

	if err := srv.ListenAndServe(); err != http.ErrServerClosed {
		// Error starting or closing listener:
		log.Fatalf("HTTP server ListenAndServe: %v", err)
	}

	<-idleConnsClosed
}

func New() *Handler {
	return &Handler{}
}
