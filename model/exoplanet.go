package model

type Exoplanet struct {
	Id                int
	Name              string
	Description       string
	DistanceFromEarth float64
	Radius            float64
	Mass              float64
	Type              string
}

func NewExoplanet() Exoplanet {
	return Exoplanet{}
}
