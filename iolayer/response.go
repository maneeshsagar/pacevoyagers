package iolayer

import "github.com/maneeshsagar/pacevoyagers/model"

type ExoplanetResponse struct {
	Id                int     `json:"id"`
	Name              string  `json:"name"`
	Description       string  `json:"description"`
	DistanceFromEarth float32 `json:"distance_from_earth"`
	Radius            float32 `json:"radius"`
	Mass              float32 `json:"mass"`
	Type              string  `json:"type"`
}

type ExoplanetListResponse struct {
	Palnets []*ExoplanetResponse `json:"palnets"`
}

func (e *ExoplanetListResponse) DomainToResponse(exoplanets []*model.Exoplanet) {
	for _, exoplanet := range exoplanets {
		newExoplanet := new(ExoplanetResponse)
		newExoplanet.DomainToResponse(exoplanet)
		e.Palnets = append(e.Palnets, newExoplanet)
	}
}

func (e *ExoplanetResponse) DomainToResponse(exoplanet *model.Exoplanet) {
	e.Id = exoplanet.Id
	e.Name = exoplanet.Name
	e.Description = exoplanet.Description
	e.DistanceFromEarth = exoplanet.DistanceFromEarth
	e.Radius = exoplanet.Radius
	e.Mass = exoplanet.Mass
	e.Type = exoplanet.Type
}
