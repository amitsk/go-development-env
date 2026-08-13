package api

import (
	"encoding/json"
	"io"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/amitsk/go-development-env/heroes-service/internal/store"
)

func newTestHandler() http.Handler {
	logger := slog.New(slog.NewTextHandler(io.Discard, nil))
	return NewServer(store.NewMemory(), logger, false).Handler()
}

func TestHealthz(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/healthz", nil)
	rec := httptest.NewRecorder()
	newTestHandler().ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("status %d", rec.Code)
	}
}

func TestCreateAndGetHero(t *testing.T) {
	h := newTestHandler()

	req := httptest.NewRequest(http.MethodPost, "/heroes", strings.NewReader(`{"name":"Batman"}`))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Request-ID", "test-req-1")
	rec := httptest.NewRecorder()
	h.ServeHTTP(rec, req)
	if rec.Code != http.StatusCreated {
		t.Fatalf("create status %d body %s", rec.Code, rec.Body.String())
	}
	if rec.Header().Get("X-Request-ID") != "test-req-1" {
		t.Fatalf("missing request id echo")
	}

	var created struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	}
	if err := json.Unmarshal(rec.Body.Bytes(), &created); err != nil {
		t.Fatal(err)
	}
	if created.ID == 0 || created.Name != "Batman" {
		t.Fatalf("created %+v", created)
	}

	get := httptest.NewRequest(http.MethodGet, "/heroes/1", nil)
	grec := httptest.NewRecorder()
	h.ServeHTTP(grec, get)
	if grec.Code != http.StatusOK {
		t.Fatalf("get status %d", grec.Code)
	}
}

func TestCreateHeroRejectsEmptyName(t *testing.T) {
	req := httptest.NewRequest(http.MethodPost, "/heroes", strings.NewReader(`{"name":""}`))
	rec := httptest.NewRecorder()
	newTestHandler().ServeHTTP(rec, req)
	if rec.Code != http.StatusBadRequest {
		t.Fatalf("status %d", rec.Code)
	}
}

func TestGetHeroNotFound(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/heroes/99", nil)
	rec := httptest.NewRecorder()
	newTestHandler().ServeHTTP(rec, req)
	if rec.Code != http.StatusNotFound {
		t.Fatalf("status %d", rec.Code)
	}
}
