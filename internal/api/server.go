package api

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/jdan1217/Ell_AI/internal/traffic"
)

type Server struct {
	history *traffic.History
}

func NewServer(history *traffic.History) *Server {
	return &Server{
		history: history,
	}
}

func (s *Server) Handler() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /api/traffic", s.handleGetTraffic)
	mux.HandleFunc("GET /api/traffic/{id}", s.handleGetTrafficByID)
	return mux
}

func (s *Server) handleGetTraffic(w http.ResponseWriter, r *http.Request) {
	limitString := r.URL.Query().Get("limit")
	offsetString := r.URL.Query().Get("offset")
	limit := 50
	offset := 0
	if limitString != "" {
		parsedlimit, err := strconv.Atoi(limitString)
		if err != nil {
			http.Error(w, "invalid limit", http.StatusBadRequest)
			return
		}
		limit = parsedlimit
		if limit < 1 {
			http.Error(w, "limit must be greater than 0", http.StatusBadRequest)
			return
		}
		if limit > 500 {
			http.Error(w, "limit cannot exceed 500", http.StatusBadRequest)
			return
		}
	}
	if offsetString != "" {
		parsedoffset, err := strconv.Atoi(offsetString)
		if err != nil {
			http.Error(w, "invalid offset", http.StatusBadRequest)
			return
		}
		offset = parsedoffset
		if offset < 0 {
			http.Error(w, "offset cannot be less than 0", http.StatusBadRequest)
			return
		}
	}
	entries := s.history.Paginated(limit, offset)
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(entries); err != nil {
		log.Printf("Failed to encode traffic history: %v", err)
		return
	}
}

func (s *Server) handleGetTrafficByID(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	entry, found := s.history.GetByID(id)
	if !found {
		http.Error(w, "Failed to get Traffic by ID", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(entry); err != nil {
		log.Printf("Failed to encode traffic entry: %v", err)
	}
}
