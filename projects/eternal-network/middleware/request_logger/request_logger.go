package request_logger

import (
	"eternal_network/middleware/request_id"
	"eternal_network/pkg/logger"
	"net/http"
	"time"
)

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()

		// Вызываем следующий обработчик в цепочке
		next.ServeHTTP(w, r)

		endTime := time.Now()
		duration := endTime.Sub(startTime)

		requestId, ok := r.Context().Value(request_id.RequestID).(string)

		if !ok {
			logger.Warn("RequestID is missed in context")
		}

		logger.Infof("[%s] [%v] %s %s - %s", r.Method, requestId, r.URL.Path, r.RemoteAddr, duration)
	})
}
