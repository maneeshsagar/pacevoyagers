package service

import (
	"errors"
	"log"
	"math"

	"github.com/maneeshsagar/pacevoyagers/model"
)

type ExoplanetService struct {
	GlobalId          int
	PlanetIdPlanetMap map[int]*model.Exoplanet
}

func NewExoplanetService() Driver {
	return &ExoplanetService{
		GlobalId:          0,
		PlanetIdPlanetMap: make(map[int]*model.Exoplanet),
	}
}

func (e *ExoplanetService) AddExoplanet(exoplanet *model.Exoplanet) (err error) {
	e.GlobalId++
	e.PlanetIdPlanetMap[e.GlobalId] = exoplanet
	return
}

func (e *ExoplanetService) UpdateExoplanet(exoplanet *model.Exoplanet) (msg string, code int, err error) {
	functionName := "service.UpdateExoplanet"

	oldExoplanet, ok := e.PlanetIdPlanetMap[exoplanet.Id]
	if !ok {
		msg = "planet doest not exist"
		code = 404
		err = errors.New("planet doesn't exists")
		log.Fatal(functionName, "planet doesn't exists planet Id ", exoplanet.Id)
		return
	}

	e.PlanetIdPlanetMap[oldExoplanet.Id] = exoplanet
	return
}

func (e *ExoplanetService) DeleteExoplanet(exoplanetId int) (msg string, code int, err error) {
	functionName := "service.DeleteExoplanet"
	exoplanet, ok := e.PlanetIdPlanetMap[exoplanetId]
	if !ok {
		msg = "planet doest not exist"
		code = 404
		err = errors.New("planet doesn't exists")
		log.Fatal(functionName, "planet doesn't exists planet Id ", exoplanet.Id)
		return
	}
	delete(e.PlanetIdPlanetMap, exoplanetId)
	return
}

func (e *ExoplanetService) GetExoplanets() (exoplanets []*model.Exoplanet, msg string, code int, err error) {

	// if there is no planet
	if len(e.PlanetIdPlanetMap) <= 0 {
		msg = "planet doest not exist"
		code = 404
		err = errors.New("planet doesn't exists")
		return
	}

	// to make the list of planet itreating over map
	for _, value := range e.PlanetIdPlanetMap {
		exoplanets = append(exoplanets, value)
	}
	return
}

func (e *ExoplanetService) GetExoplanetById(exoplanetId int) (exoplanet *model.Exoplanet, msg string, code int, err error) {
	functionName := "service.GetExoplanetById"

	exoplanet, ok := e.PlanetIdPlanetMap[exoplanetId]
	if !ok {
		msg = "planet doest not exist"
		code = 404
		err = errors.New("planet doesn't exists")
		log.Fatal(functionName, "planet doesn't exists planet Id ", exoplanet.Id)
		return
	}

	return
}

func (e *ExoplanetService) CalculateFuelEstimation(id, crewCapacity int) (fuel float64, msg string, code int, err error) {
	exoplanet, ok := e.PlanetIdPlanetMap[id]
	if !ok {
		msg = "exoplanet not found"
		code = 404
		err = errors.New(msg)
		return
	}

	var gravity float64
	switch exoplanet.Type {
	case "GasGiant":
		gravity = 0.5 / math.Pow(exoplanet.Radius, 2)
	case "Terrestrial":
		gravity = exoplanet.Mass / math.Pow(exoplanet.Radius, 2)
	}

	fuel = float64(exoplanet.DistanceFromEarth) / math.Pow(gravity, 2) * float64(crewCapacity)
	return
}
