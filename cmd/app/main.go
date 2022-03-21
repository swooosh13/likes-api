package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func main() {
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

	err := http.ListenAndServe("127.0.0.1:9000", nil)
	if err != nil {
		log.Fatal("server has been crashed")
	}
}
