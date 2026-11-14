package api

import (
    "net/http"

    "github.com/go-chi/chi/v5"
)

type Server struct {
	Router *chi.mux
}

func NewServer() *Server {
	server := Server{
		Router: chi.NewRouter()
	}

	return server
}

func (s *Server) InitHandlers {
	s.Router.Post("v1/events", ProcessEvent)
}

func ProcessEvent(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusOK)
    w.Write([]byte("Event processed!!"))
}