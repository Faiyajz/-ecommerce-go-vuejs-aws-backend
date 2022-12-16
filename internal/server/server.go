package server

import (
	"backend/internal/category"
	"github.com/gin-gonic/gin"
)

type Server struct {
}

type Config struct {
}

func New(config Config) (*Server, error) {
	return &Server{}, nil
}

func (server *Server) Categories(ctx *gin.Context) {
	categories := []category.Category{
		{
			ID:          "1",
			Name:        "HardCover",
			Description: "A Hard Copy of the Handbook",
		},
		{
			ID:          "2",
			Name:        "Digital Version",
			Description: "A Soft Copy of the Handbook",
		},
	}
	ctx.Header("Access-Control-Allow-Origin", "http://localhost:8080")
	ctx.JSON(200, categories)
}
