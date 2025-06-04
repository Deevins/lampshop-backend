package server

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type HTTPServer struct{}

func NewHTTPServer() *HTTPServer {
	return &HTTPServer{}
}

func (s *HTTPServer) Run(addr string, router *gin.Engine) error {
	return http.ListenAndServe(addr, router)
}
