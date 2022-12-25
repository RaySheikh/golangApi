package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/go-api/controllers/handlers"
)

func User(g *gin.RouterGroup) {
	g.GET("/users", handlers.GetUsers)
}
