package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/go-api/controllers/handlers"
)

// GenerateToken godoc
// @Summary Generate JWT
// @Description Responds with a JWT
// @Produce json
// @Success 200
// @Router /auth/generatetoken [post]
func Auth(g *gin.RouterGroup) {
	g.POST("/generatetoken", handlers.HandleAuthorize)
}
