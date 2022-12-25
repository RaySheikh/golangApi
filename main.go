package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-api/controllers/routes"
	_ "github.com/go-api/docs"
	"github.com/go-api/middleware"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func getUsers(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, "test")
}

// @title     Golang Open AI api
// @version         1.0
// @description     An api to communicate with AI.
// @host      localhost:8082
// @BasePath  /api/v1
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name API_TOKEN
func main() {
	router := gin.Default()
	v1 := router.Group("/api/v1")
	{
		routes.Auth(v1.Group("auth"))

		v1.Use(middleware.TokenAuthMiddleware()) //auth required for routes below this line
		routes.User(v1.Group("user"))
	}
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.Run("localhost:8082")
}
