package iolayer

import "github.com/maneeshsagar/pacevoyagers/model"

type ExoplanetRequest struct {
	Name              string  `json:"name"`
	Description       string  `json:"description"`
	DistanceFromEarth float32 `json:"distance_from_earth"`
	Radius            float32 `json:"radius"`
	Mass              float32 `json:"mass"`
	Type              string  `json:"type"`
}

func (e *ExoplanetRequest) RequestToServiceDomain() *model.Exoplanet {
	return &model.Exoplanet{
		Name:              e.Name,
		Description:       e.Description,
		DistanceFromEarth: e.DistanceFromEarth,
		Radius:            e.Radius,
		Mass:              e.Mass,
		Type:              e.Type,
	}
}
