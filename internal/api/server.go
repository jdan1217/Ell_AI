package api

import (
	"encoding/json"
	"log"
	"net/http"

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
	entries := s.history.All()
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
