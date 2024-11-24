package server

import (
	"example.com/boiletplate/ent"
	otphandler "example.com/boiletplate/infrastructure/OTPHandler"
	"example.com/boiletplate/infrastructure/queue"
	"github.com/gin-gonic/gin"
)

type Server struct {
	entClient  *ent.Client
	otpHandler otphandler.OTPHandler
	Gin        *gin.Engine
	publisher  queue.IPublisher
}

func NewServer(entClient *ent.Client, otpHandler otphandler.OTPHandler, publisher queue.IPublisher) *Server {
	return &Server{
		Gin:        gin.Default(),
		entClient:  entClient,
		otpHandler: otpHandler,
		publisher:  publisher,
	}
}

func (s *Server) Bootstrap() {
	NewHandlers(s, s.publisher, s.otpHandler)
	s.Gin.Use(gin.Logger())
	s.Gin.Use(gin.Recovery())
	err := s.Gin.Run()
	if err != nil {
		panic(err)
	}
}
