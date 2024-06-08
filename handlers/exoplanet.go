package handlers

import (
	"encoding/json"
	"io"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/maneeshsagar/pacevoyagers/iolayer"
	"github.com/maneeshsagar/pacevoyagers/service"
	"github.com/maneeshsagar/pacevoyagers/util"
	"github.com/spf13/cast"
)

func AddExoplanet(spaceService service.Driver) gin.HandlerFunc {

	funtionName := "handlers.AddExoplanet"

	return func(ctx *gin.Context) {

		body, err := io.ReadAll(ctx.Request.Body)
		if err != nil {
			log.Println("invalid request")
			response := util.NewResponse(400, "invalid request", nil)
			ctx.JSON(200, response)
			return
		}

		var originalRequest iolayer.ExoplanetRequest
		err = json.Unmarshal(body, &originalRequest)
		if err != nil {
			log.Println(funtionName, "error : failed to unmarshal ", err)
			response := util.NewResponse(400, "invalid request", nil)
			ctx.JSON(200, response)
			return
		}

		request := originalRequest.RequestToServiceDomain()

		err = spaceService.AddExoplanet(request)
		if err != nil {
			log.Println(funtionName, "error : failed to add ", err)
			response := util.NewResponse(400, "invalid request", nil)
			ctx.JSON(200, response)
			return
		}

		response := util.NewResponse(200, "success", nil)
		ctx.JSON(200, response)
	}
}

func GetListOfExoplanet(spaceService service.Driver) gin.HandlerFunc {
	funtionName := "handlers.GetListOfExoplanet"
	return func(ctx *gin.Context) {

		serviceExoplanet, msg, code, err := spaceService.GetExoplanets()
		if err != nil {
			log.Println(funtionName, "error : failed to get all exoplanet  ", err)
			response := util.NewResponse(code, msg, nil)
			ctx.JSON(200, response)
			return
		}

		response := new(iolayer.ExoplanetListResponse)
		response.DomainToResponse(serviceExoplanet)

		finalResponse := util.NewResponse(200, "success", response)
		ctx.JSON(200, finalResponse)
	}
}

func GetExoplanet(spaceService service.Driver) gin.HandlerFunc {
	funtionName := "handlers.GetExoplanet"

	return func(ctx *gin.Context) {

		exoplanetId := ctx.Param("id")

		serviceExoplanet, msg, code, err := spaceService.GetExoplanetById(cast.ToInt(exoplanetId))
		if err != nil {
			log.Println(funtionName, "error : failed to get exoplanet ", err)
			response := util.NewResponse(code, msg, nil)
			ctx.JSON(200, response)
			return
		}

		response := new(iolayer.ExoplanetResponse)
		response.DomainToResponse(serviceExoplanet)
		finalResponse := util.NewResponse(200, "success", nil)
		ctx.JSON(200, finalResponse)
	}
}

func UpdateExoplanet(spaceService service.Driver) gin.HandlerFunc {
	funtionName := "handlers.UpdateExoplanet"

	return func(ctx *gin.Context) {
		body, err := io.ReadAll(ctx.Request.Body)
		if err != nil {
			log.Println("invalid request")
			response := util.NewResponse(400, "invalid request", nil)
			ctx.JSON(200, response)
			return
		}

		var originalRequest iolayer.ExoplanetRequest
		err = json.Unmarshal(body, &originalRequest)
		if err != nil {
			log.Println(funtionName, "error : failed to unmarshal ", err)
			response := util.NewResponse(400, "invalid request", nil)
			ctx.JSON(200, response)
			return
		}

		request := originalRequest.RequestToServiceDomain()

		msg, code, err := spaceService.UpdateExoplanet(request)
		if err != nil {
			log.Println(funtionName, msg, err)
			response := util.NewResponse(code, msg, nil)
			ctx.JSON(200, response)
			return
		}

		response := util.NewResponse(200, "success", nil)
		ctx.JSON(200, response)
	}
}

func DeleteExoplanet(spaceService service.Driver) gin.HandlerFunc {
	funtionName := "handlers.DeleteExoplanet"

	return func(ctx *gin.Context) {

		exoplanetId := ctx.Param("id")

		msg, code, err := spaceService.DeleteExoplanet(cast.ToInt(exoplanetId))
		if err != nil {
			log.Println(funtionName, "error : failed to delete ", err)
			response := util.NewResponse(code, msg, nil)
			ctx.JSON(200, response)
			return
		}

		response := util.NewResponse(200, "success", nil)
		ctx.JSON(200, response)
	}
}

func GetFuelEstimation(spaceService service.Driver) gin.HandlerFunc {
	funtionName := "handlers.GetFuelEstimation"

	return func(ctx *gin.Context) {

		exoplanetId := ctx.Param("id")
		crewCapacityStr := ctx.Query("crew")
		if crewCapacityStr == "" {
			response := util.NewResponse(400, "crew capacity is required", nil)
			ctx.JSON(200, response)
			return
		}

		crewCapacity := cast.ToInt(crewCapacityStr)
		if crewCapacity <= 0 {
			response := util.NewResponse(400, "crew capacity is required", nil)
			ctx.JSON(200, response)
			return
		}

		fuel, msg, code, err := spaceService.CalculateFuelEstimation(cast.ToInt(exoplanetId), crewCapacity)
		if err != nil {
			log.Println(funtionName, "error : failed to estimate ", err)
			response := util.NewResponse(code, msg, nil)
			ctx.JSON(200, response)
			return
		}

		actualResponse := iolayer.FuelResposne{
			Fuel: fuel,
		}

		response := util.NewResponse(200, "success", actualResponse)
		ctx.JSON(200, response)
	}
}
