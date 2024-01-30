package request_id

import (
	"context"
	"net/http"

	"github.com/google/uuid"
)

type RequestIDKey string

var RequestID RequestIDKey = "requestID"

const XRequestIdKey = "X-Request-ID"

func RequestIdMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestId := r.Header.Get(XRequestIdKey)

		if requestId == "" {
			requestId = uuid.New().String()

			r.Header.Set(XRequestIdKey, requestId)
		}

		ctx := r.Context()
		ctx = context.WithValue(ctx, RequestID, requestId)
		r = r.WithContext(ctx)

		w.Header().Add(XRequestIdKey, requestId)

		next.ServeHTTP(w, r)
	})
}
