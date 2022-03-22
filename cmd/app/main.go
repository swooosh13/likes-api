package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"proj1/internal/config"
	"proj1/pkg/logger"

	"go.uber.org/zap"
)

type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func main() {
	logger.Init()
	logger.Info("Check logger zap", zap.String("infomsg", "<value of string>"))
	// logger.Fatal("Check logger zap", zap.String("fatalmsg", "<value of string>"))
	logger.Error("Check logger zap", zap.String("errormsg", "<value of string>"))
	logger.Debug("Check logger zap", zap.String("debugmsg", "<value of string>"))
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {

		mp := map[string]struct {
			Name string `json:"name"`
			Age  int    `json:"age"`
		}{
			"1": {
				Age:  19,
				Name: "Artyom",
			},
		}

		rw.Header().Set("Content-Type", "application/json")

		jsonData, err := json.Marshal(mp)
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			resp := &ErrorResponse{http.StatusInternalServerError, "Error data"}
			jsonData, _ = json.Marshal(resp)

			_, _ = rw.Write(jsonData)
			return
		}

		rw.WriteHeader(http.StatusOK)
		_, _ = rw.Write(jsonData)
	})

	cfg := config.GetConfig()
	err := http.ListenAndServe(fmt.Sprintf("%s:%s", cfg.App.BindIP, cfg.App.Port), nil)
	if err != nil {
		log.Fatal("server has been crashed")
	}
}
