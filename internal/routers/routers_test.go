package routers

import (
	"bytes"
	"encoding/json"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/poggerr/avito/internal/app"
	"github.com/poggerr/avito/internal/config"
	"github.com/poggerr/avito/internal/logger"
	"github.com/poggerr/avito/internal/models"
	"github.com/poggerr/avito/internal/storage"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func NewDefConf() *config.Config {
	conf := config.Config{
		ServAddr: ":8080",
		DB:       "host=localhost user=avito password=password dbname=avito sslmode=disable",
	}
	return &conf
}

var sugaredLogger = logger.Initialize()
var cfg = NewDefConf()
var strg = storage.NewStorage(connectDB(), cfg)
var newApp = app.NewApp(cfg, strg, sugaredLogger)

func connectDB() *sqlx.DB {
	db, err := sqlx.Connect("postgres", cfg.DB)
	if err != nil {
		log.Fatalln(err)
	}
	return db
}

func testRequestPost(t *testing.T, ts *httptest.Server, method,
	path string, data string) (*http.Response, string) {

	req, err := http.NewRequest(method, ts.URL+path, bytes.NewBuffer([]byte(data)))
	require.NoError(t, err)

	resp, err := ts.Client().Do(req)
	require.NoError(t, err)

	respBody, err := io.ReadAll(resp.Body)
	require.NoError(t, err)

	return resp, string(respBody)
}

func testRequestJSON(t *testing.T, ts *httptest.Server, method, path string, data []byte) (*http.Response, string) {
	req, err := http.NewRequest(method, ts.URL+path, bytes.NewBuffer(data))
	require.NoError(t, err)

	resp, err := ts.Client().Do(req)
	require.NoError(t, err)

	respBody, err := io.ReadAll(resp.Body)
	require.NoError(t, err)

	return resp, string(respBody)
}

func TestHandlersPost(t *testing.T) {
	ts := httptest.NewServer(Router(newApp))
	defer ts.Close()

	var testTable = []struct {
		name   string
		api    string
		method string
		status int
	}{
		{name: "segment_create", api: "/api/segment/create", method: "POST", status: 201},
		{name: "segment_create_duplicate", api: "/api/segment/create", method: "POST", status: 409},
		{name: "segment_create_negative", api: "/api/segment/create", method: "POST", status: 400},
		{name: "user_segment_create", api: "/api/user/segment", method: "POST", status: 201},
		{name: "user_segment_delete", api: "/api/user/segment", method: "POST", status: 201},
		{name: "get_user_segment", api: "/api/segment/2376e110-e40d-41d0-85ba-22db804c4f51", method: "GET", status: 200},
	}

	for _, v := range testTable {
		switch {
		case v.name == "segment_create" || v.name == "segment_create_duplicate":
			var segment models.Segment
			segment.Slug = "AVITO_DISCOUNT_50"
			marshal, err := json.Marshal(segment)
			if err != nil {
				sugaredLogger.Info(err)
			}
			resp, _ := testRequestJSON(t, ts, v.method, v.api, marshal)
			defer resp.Body.Close()
			assert.Equal(t, v.status, resp.StatusCode)
		case v.name == "user_segment_create":
			add := []string{"AVITO_DISCOUNT_50"}
			user := uuid.Must(uuid.Parse("2376e110-e40d-41d0-85ba-22db804c4f51"))
			data := models.CRUDSegmentToUser{
				Add:    add,
				Delete: nil,
				UserID: user,
			}
			marshal, err := json.Marshal(data)
			if err != nil {
				sugaredLogger.Info(err)
			}
			resp, _ := testRequestJSON(t, ts, v.method, v.api, marshal)
			defer resp.Body.Close()
			assert.Equal(t, v.status, resp.StatusCode)
		case v.name == "user_segment_delete":
			deleteSlug := []string{"AVITO_DISCOUNT_50"}
			user := uuid.Must(uuid.Parse("2376e110-e40d-41d0-85ba-22db804c4f51"))
			data := models.CRUDSegmentToUser{
				Add:    nil,
				Delete: deleteSlug,
				UserID: user,
			}
			marshal, err := json.Marshal(data)
			if err != nil {
				sugaredLogger.Info(err)
			}
			resp, _ := testRequestJSON(t, ts, v.method, v.api, marshal)
			defer resp.Body.Close()
			assert.Equal(t, v.status, resp.StatusCode)
		case v.name == "get_user_segment":
			resp, _ := testRequestPost(t, ts, v.method, v.api, "")
			defer resp.Body.Close()
			assert.Equal(t, v.status, resp.StatusCode)
		}
	}

}
