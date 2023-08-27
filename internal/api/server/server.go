package server

import (
	"github.com/IlyaZayats/dynus/internal/api/routing"
)

type Server struct {
	handlers routing.HandlerInterface
}

func InitNewServer(handlers routing.HandlerInterface) Server {
	return Server{handlers: handlers}
}

func (s *Server) Run() {
	s.handlers.Init()
	s.handlers.Run()
	defer s.handlers.CloseConnection()
}
