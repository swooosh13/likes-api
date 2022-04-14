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
	
	logger.Info("intializing config")
	cfg := config.GetConfig()
	fmt.Println(cfg.PostgresDB.Host)

	// poolConfig, err := postgresdb.NewPoolConfig()

	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		rw.WriteHeader(http.StatusOK)
		_, _ = rw.Write([]byte("ok"))
	})

	logger.Info("starting server", zap.String("host", cfg.Listen.Host), zap.String("port", cfg.Listen.Port))
	err := http.ListenAndServe(fmt.Sprintf("%s:%s", cfg.Listen.Host, cfg.Listen.Port), nil)
	if err != nil {
		log.Fatal("server has been crashed")
	}
}
