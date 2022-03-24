package main

import (
	"fmt"
	"log"
	"net/http"
	"proj1/internal/config"
	"proj1/pkg/logger"

	"go.uber.org/zap"
)

func main() {
	logger.Init()

	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		rw.WriteHeader(http.StatusOK)
		_, _ = rw.Write([]byte("ok"))
	})

	logger.Info("intializing config")
	cfg := config.GetConfig()

	logger.Info("starting server", zap.String("host", cfg.App.Host), zap.String("port", cfg.App.Port))
	err := http.ListenAndServe(fmt.Sprintf("%s:%s", cfg.App.Host, cfg.App.Port), nil)
	if err != nil {
		log.Fatal("server has been crashed")
	}
}
