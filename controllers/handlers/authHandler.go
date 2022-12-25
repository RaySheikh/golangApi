package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-api/services"
)

func HandleAuthorize(g *gin.Context) {
	token, err := services.GenerateJWT()

	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	tokenResult := map[string]string{"token": token}

	g.JSON(http.StatusOK, tokenResult)
	return
}
