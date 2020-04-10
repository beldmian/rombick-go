package server

import (
	"os"

	"github.com/beldmian/Rombick/cmd/config"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// Server ...
type Server struct {
	bindAddr string
	logger   *logrus.Logger
	router   *gin.Engine
}

// New ...
func New() *Server {
	return &Server{
		bindAddr: os.Getenv("PORT"),
		logger:   logrus.New(),
		router:   gin.New(),
	}
}

// Start ...
func (s *Server) Start() error {
	if err := s.configureLogger(); err != nil {
		return err
	}
	s.logger.Info("Logger configured successfuly")

	s.configureRouter()
	s.logger.Info("Server started successful on port :" + s.bindAddr)
	return s.router.Run(":" + s.bindAddr)
}

func (s *Server) configureLogger() error {
	level, err := logrus.ParseLevel(config.Config.LogLevel)
	if err != nil {
		return err
	}
	s.logger.SetLevel(level)
	return nil
}

func (s *Server) configureRouter() {
	s.router.Use(gin.Recovery())
	s.router.GET("/", s.indexHandler)
}
