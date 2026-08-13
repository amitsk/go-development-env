package api

import (
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"
	"net/http/pprof"
	"strconv"
	"strings"

	"github.com/amitsk/go-development-env/heroes-service/internal/store"
)

// Server wires HTTP routes to a Store.
type Server struct {
	store  store.Store
	logger *slog.Logger
	pprof  bool
}

func NewServer(s store.Store, logger *slog.Logger, enablePprof bool) *Server {
	if logger == nil {
		logger = slog.Default()
	}
	return &Server{store: s, logger: logger, pprof: enablePprof}
}

// Handler returns the fully wrapped mux.
func (s *Server) Handler() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /healthz", s.healthz)
	mux.HandleFunc("GET /heroes", s.listHeroes)
	mux.HandleFunc("GET /heroes/{id}", s.getHero)
	mux.HandleFunc("POST /heroes", s.createHero)
	if s.pprof {
		mux.HandleFunc("GET /debug/pprof/", pprof.Index)
		mux.HandleFunc("GET /debug/pprof/cmdline", pprof.Cmdline)
		mux.HandleFunc("GET /debug/pprof/profile", pprof.Profile)
		mux.HandleFunc("GET /debug/pprof/symbol", pprof.Symbol)
		mux.HandleFunc("GET /debug/pprof/trace", pprof.Trace)
	}
	return withRequestID(withLogging(s.logger, mux))
}

func (s *Server) healthz(w http.ResponseWriter, _ *http.Request) {
	writeJSON(w, http.StatusOK, map[string]string{"status": "ok"})
}

func (s *Server) listHeroes(w http.ResponseWriter, r *http.Request) {
	heroes, err := s.store.List(r.Context())
	if err != nil {
		s.fail(w, r, http.StatusInternalServerError, "list heroes", err)
		return
	}
	writeJSON(w, http.StatusOK, map[string]any{"heroes": heroes})
}

func (s *Server) getHero(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "id must be an integer"})
		return
	}
	hero, err := s.store.Get(r.Context(), id)
	if errors.Is(err, store.ErrNotFound) {
		writeJSON(w, http.StatusNotFound, map[string]string{"error": "not found"})
		return
	}
	if err != nil {
		s.fail(w, r, http.StatusInternalServerError, "get hero", err)
		return
	}
	writeJSON(w, http.StatusOK, hero)
}

type createHeroRequest struct {
	Name string `json:"name"`
}

func (s *Server) createHero(w http.ResponseWriter, r *http.Request) {
	var req createHeroRequest
	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()
	if err := dec.Decode(&req); err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "invalid json"})
		return
	}
	hero, err := s.store.Create(r.Context(), strings.TrimSpace(req.Name))
	if errors.Is(err, store.ErrInvalidName) {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}
	if err != nil {
		s.fail(w, r, http.StatusInternalServerError, "create hero", err)
		return
	}
	writeJSON(w, http.StatusCreated, hero)
}

func (s *Server) fail(w http.ResponseWriter, r *http.Request, status int, msg string, err error) {
	s.logger.Error(msg,
		slog.String("request_id", RequestIDFrom(r.Context())),
		slog.Any("err", err),
	)
	writeJSON(w, status, map[string]string{"error": "internal error"})
}

func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(v); err != nil {
		return
	}
}
