package main

import (
	"github.com/gin-gonic/gin"
	"github.com/maneeshsagar/pacevoyagers/handlers"
	"github.com/maneeshsagar/pacevoyagers/service"
)

func main() {
	router := gin.Default()

	spaceService := service.NewExoplanetService()
	router.GET("/exoplanet", handlers.GetListOfExoplanet(spaceService))
	router.GET("/exoplanet:id", handlers.GetExoplanet(spaceService))
	router.POST("/exoplanet", handlers.AddExoplanet(spaceService))
	router.PUT("/exoplanet", handlers.UpdateExoplanet(spaceService))
	router.DELETE("/exoplanet:id", handlers.DeleteExoplanet(spaceService))
	router.GET("/exoplanets/:id/fuel", handlers.GetFuelEstimation(spaceService))
	router.Run()

}
