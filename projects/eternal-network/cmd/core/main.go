package main

import (
	"fmt"
	"net/http"

	"eternal_network/middleware/request_id"
	"eternal_network/middleware/request_logger"
	"eternal_network/pkg/config"
	"eternal_network/pkg/logger"

	"github.com/gorilla/mux"
)

func main() {
	cfg := config.GetConfig()

	logger.Debug(*cfg)

	r := mux.NewRouter()

	r.Use(request_id.RequestIdMiddleware, request_logger.LoggingMiddleware)

	// Обработчик запроса
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello, world!")
	}).Methods("GET")

	http.Handle("/", r)

	logger.Infof("App started on port: %v. Environment: %s\n", cfg.Port, cfg.AppEnv)

	logger.Fatalf("App is crashed: %s", http.ListenAndServe(fmt.Sprintf(":%v", cfg.Port), nil))
}
