package main

import (
	"context"
	"fmt"
	"os"
	"proj1/internal/composites"
	"proj1/internal/config"
	"proj1/pkg/logger"

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

	rows, err := pgComposite.Client.Query(context.TODO(), "SELECT datname FROM pg_database;")
	if err != nil {
		fmt.Println(err)
	}
	for rows.Next() {
		var output string

		err = rows.Scan(&output)
		if err != nil {
			logger.Fatal(err.Error())
		}
		fmt.Println(output)
	}

	// http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
	// 	rw.WriteHeader(http.StatusOK)
	// 	_, _ = rw.Write([]byte("ok"))
	// })

	// logger.Info("starting server", zap.String("host", cfg.Listen.Host), zap.String("port", cfg.Listen.Port))
	// err := http.ListenAndServe(fmt.Sprintf("%s:%s", cfg.Listen.Host, cfg.Listen.Port), nil)
	// if err != nil {
	// 	log.Fatal("server has been crashed")
	// }
}
