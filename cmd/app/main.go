package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"proj1/internal/composites"
	"proj1/internal/config"
	"proj1/internal/handlers/api"
	"proj1/pkg/logger"
	"syscall"
	"time"

	"github.com/go-chi/chi"
	_ "github.com/lib/pq"
)

func main() {
	logger.Init()
	logger.Info("intializing config")
	cfg := config.GetConfig()

	pgComposite, err := composites.NewPostgresDBComposite(context.TODO(), cfg.PostgresDB.Host, cfg.PostgresDB.Port, cfg.PostgresDB.Username, cfg.PostgresDB.Password, cfg.PostgresDB.Database, cfg.PostgresDB.Timeout, cfg.PostgresDB.MaxConns)
	if err != nil {
		logger.Fatal(err.Error())
		os.Exit(1)
	}
	logger.Info("postgres composite initializing")

	containerComposite := composites.NewContainerComposite(pgComposite)
	logger.Info("container composite initializing")

	server := &http.Server{
		Addr:    ":8000",
		Handler: service(containerComposite.Handler),
	}

	serverCtx, serverStopCtx := context.WithCancel(context.Background())

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	go func() {
		<-sig

		shutdownCtx, _ := context.WithTimeout(serverCtx, 30*time.Second)
		go func() {
			<-shutdownCtx.Done()
			if shutdownCtx.Err() == context.DeadlineExceeded {
				logger.Fatal("graceful shutdown timed out.. forcing exit.")
			}
		}()

		err = server.Shutdown(shutdownCtx)
		if err != nil {
			logger.Fatal(err.Error())
		}

		serverStopCtx()
	}()

	// run
	err = server.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		logger.Fatal(err.Error())
	}

	<-serverCtx.Done()
}

func service(handlers ...api.Handler) http.Handler {
	r := chi.NewMux()
	r.Use(api.Authentication)
	for _, v := range handlers {
		v.Register(r)
	}

	r.Put("/ping", Ping)

	return r
}

func Ping(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("pong"))
}
