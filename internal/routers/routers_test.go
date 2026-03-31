package routers

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

type mockHandlers struct{}

func (mockHandlers) CreateSegment(res http.ResponseWriter, _ *http.Request) {
	res.WriteHeader(http.StatusCreated)
}

func (mockHandlers) DeleteSegment(res http.ResponseWriter, _ *http.Request) {
	res.WriteHeader(http.StatusOK)
}

func (mockHandlers) SegmentByUserID(res http.ResponseWriter, _ *http.Request) {
	res.WriteHeader(http.StatusOK)
}

func (mockHandlers) SegmentsToUser(res http.ResponseWriter, _ *http.Request) {
	res.WriteHeader(http.StatusCreated)
}

func (mockHandlers) CreateCSVSegment(res http.ResponseWriter, _ *http.Request) {
	res.WriteHeader(http.StatusOK)
}

func TestRoutes(t *testing.T) {
	ts := httptest.NewServer(Router(mockHandlers{}))
	defer ts.Close()

	tests := []struct {
		name   string
		method string
		path   string
		body   string
		status int
	}{
		{name: "create segment", method: http.MethodPost, path: "/api/segment/create", body: `{"segment":"A"}`, status: http.StatusCreated},
		{name: "delete segment", method: http.MethodPost, path: "/api/segment/delete", body: `{"segment":"A"}`, status: http.StatusOK},
		{name: "create user segment", method: http.MethodPost, path: "/api/user/segment", body: `{"add":["A"],"delete":[],"user":"2376e110-e40d-41d0-85ba-22db804c4f51"}`, status: http.StatusCreated},
		{name: "get user segments", method: http.MethodGet, path: "/api/segment/2376e110-e40d-41d0-85ba-22db804c4f51", status: http.StatusOK},
		{name: "create csv", method: http.MethodPost, path: "/api/segment/csv/2376e110-e40d-41d0-85ba-22db804c4f51", body: `{"period":"03-2026"}`, status: http.StatusOK},
		{name: "method not allowed", method: http.MethodGet, path: "/api/user/segment", status: http.StatusMethodNotAllowed},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, err := http.NewRequest(tt.method, ts.URL+tt.path, bytes.NewBuffer([]byte(tt.body)))
			assert.NoError(t, err)

			resp, err := ts.Client().Do(req)
			assert.NoError(t, err)
			defer resp.Body.Close()

			_, _ = io.ReadAll(resp.Body)
			assert.Equal(t, tt.status, resp.StatusCode)
		})
	}
}
