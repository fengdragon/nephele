package app

import (
	"github.com/nephele/logger"
	"github.com/nephele/service"
)

// Define server initialization function type.
type ServerInitializeFunc func(*Server) error

// Server represents holder for services and
// is the entry for all components to initialize, open or close.
type Server struct {
	init    ServerInitializeFunc
	logger  *logger.Logger
	service *service.Service
}

// Call to make external initialization.
func (s *Server) Init(init ServerInitializeFunc) error {
	return nil
}

// Open services and other components.
func (s *Server) Open() error {
	s.init(s)
	return nil
}

// Close services and other components elegantly.
func (s *Server) Close() error {
	return nil
}

// Return service to register router or modify image handler path.
func (s *Server) Service() *service.Service {
	return s.service
}