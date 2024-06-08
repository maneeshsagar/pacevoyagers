package service

import "github.com/maneeshsagar/pacevoyagers/model"

type Driver interface {
	AddExoplanet(exoplanet *model.Exoplanet) (savedExoplanet *model.Exoplanet, err error)
	UpdateExoplanet(exoplanet *model.Exoplanet) (savedExoplanet *model.Exoplanet, msg string, code int, err error)
	DeleteExoplanet(exoplanetId int) (msg string, code int, err error)
	GetExoplanets() (exoplanets []*model.Exoplanet, msg string, code int, err error)
	GetExoplanetById(exoplanetId int) (exoplanet *model.Exoplanet, msg string, code int, err error)
	CalculateFuelEstimation(id, crewCapacity int) (fuel float64, msg string, code int, err error)
}
