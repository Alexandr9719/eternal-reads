package request_id

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestWithNewReqId(t *testing.T) {
	nextHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		reqIdCtx, ok := r.Context().Value(RequestID).(string)
		reqIdW := w.Header().Get(XReqIdKey)

		if !ok || reqIdCtx == "" {
			t.Errorf("Context request id must be set")
		}

		if reqIdW == "" {
			t.Errorf("Response X-Request-Id must be set")
		}

		if ok && reqIdCtx != reqIdW {
			t.Errorf("Context and response request id must be equal")
		}
	})

	handlerToTest := RequestIdMiddleware(nextHandler)

	req := httptest.NewRequest(http.MethodGet, "/", nil)

	handlerToTest.ServeHTTP(httptest.NewRecorder(), req)
}
