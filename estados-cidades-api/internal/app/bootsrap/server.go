package bootstrap

import (
	"fmt"

	"github.com/gin-gonic/gin"
	location "github.com/limadlp/estados-cidades-api/internal/app/handlers/locations"
	repositories "github.com/limadlp/estados-cidades-api/internal/infrastructure/repositories/location"
)

func StartServer() {
	var e *gin.Engine = gin.Default()
	configureRoutes(e)
	var err error = e.Run(":8080")

	if err != nil {
		panic(err)
	}

	fmt.Println("Server Started")
}

func configureRoutes(e *gin.Engine) {

	locationRepository := repositories.NewLocationRepository()
	locationHandler := location.NewLocationHandler(locationRepository)

	var g *gin.RouterGroup = e.Group("/api/v1")
	{
		g.GET("/states", locationHandler.GetAllStates)

	}
}
