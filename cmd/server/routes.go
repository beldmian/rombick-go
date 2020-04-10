package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) indexHandler(c *gin.Context) {
	c.JSON(http.StatusOK, struct{ Msg string }{Msg: "Hello"})
}
