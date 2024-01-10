package request_id

import (
	"context"
	"net/http"

	"github.com/google/uuid"
)

type RequestIDKey string

var RequestID RequestIDKey = "requestID"

func RequestIdMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestId := r.Header.Get("X-Request-ID")

		if requestId == "" {
			requestId = uuid.New().String()

			r.Header.Set("X-Request-ID", requestId)
		}

		ctx := r.Context()
		ctx = context.WithValue(ctx, RequestID, requestId)
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}
