package main

import (
	"fmt"
	"net/http"
	"time"

	"eternal_network/middleware/request_logger"
	"eternal_network/packages/config"
	"eternal_network/packages/logger"

	"github.com/gorilla/mux"
)

func initConfig() *config.Config {
	cfg := config.NewConfig()

	cfg.InitConfigByEnv()

	return cfg
}

func main() {
	cfg := initConfig()

	logger.InfoLogger.Println(*cfg)

	r := mux.NewRouter()

	r.Use(request_logger.LoggingMiddleware)

	// Обработчик запроса
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		for i := 0; i < 1000000; i++ {
			time.Sleep(1000)
		}
		fmt.Fprint(w, "Hello, world!")
	}).Methods("GET")

	http.Handle("/", r)

	logger.ErrorLogger.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", cfg.Port), nil))

	logger.InfoLogger.Printf("App started on port: %v. Environment: %s\n", cfg.Port, cfg.AppEnv)
}
