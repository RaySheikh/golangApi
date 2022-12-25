package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-api/services"
)

// GetUsers godoc
// @Summary Get users array
// @Description Responds with a list of all users as JSON
// @Tags users
// @Produce json
// @Success 200 {array} models.User
// @Router /user/users [get]
// @Security ApiKeyAuth
func GetUsers(g *gin.Context) {
	users := services.GetUsers()

	g.JSON(http.StatusOK, users)
	return
}
