package request_logger

import (
	"net/http"
	"time"

	"eternal_network/packages/logger"
)

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()

		// Вызываем следующий обработчик в цепочке
		next.ServeHTTP(w, r)

		endTime := time.Now()
		duration := endTime.Sub(startTime)

		logger.InfoLogger.Printf("[%s] %s %s - %s", r.Method, r.URL.Path, r.RemoteAddr, duration)
	})
}
