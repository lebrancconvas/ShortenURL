package server

import (
	"log"

	"github.com/gin-gonic/gin"
)

type Server struct {
	router *gin.Engine
	port string
}

func NewServer(r *gin.Engine, port string) *Server {
	return &Server{
		router: r,
		port: port,
	}
}

func (s *Server) Run() {
	if err := s.router.Run(s.port); err != nil {
		log.Fatalf("Failed to run server on port %s with %v", s.port, err)
	}
}
