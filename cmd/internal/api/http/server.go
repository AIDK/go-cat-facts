package server

import (
	"context"
	"net/http"

	"github.com/aidk/gocatfacts/cmd/internal/api/service"
	"github.com/aidk/gocatfacts/cmd/internal/api/util"
)

// Server is the interface that provides the HTTP server.
type Server struct {
	svc service.Service
}

// NewServer returns a new Server. It requires a Service. It returns a pointer to a Server.
func NewServer(svc service.Service) *Server {
	return &Server{
		svc: svc,
	}
}

// Start starts the HTTP server on the specified port.
func (s *Server) Start(port string) error {

	// we register our handler on the default router
	http.HandleFunc("/fact", s.handleGetFact)

	// we start the HTTP server
	return http.ListenAndServe(port, nil)
}

// handleGetFact is the handler for the GET /fact endpoint.
func (s *Server) handleGetFact(w http.ResponseWriter, r *http.Request) {

	// we get a fact from the service
	fact, err := s.svc.GetFact(context.Background())
	if err != nil {
		// if there was an error, we write the error to the response body
		util.WriteToJSON(w, http.StatusInternalServerError, err.Error())
		return
	}

	// we write the fact to the response body
	util.WriteToJSON(w, http.StatusOK, fact)

}
