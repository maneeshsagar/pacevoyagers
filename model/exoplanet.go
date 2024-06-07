package model

type Exoplanet struct {
	Id                int
	Name              string
	Description       string
	DistanceFromEarth float32
	Radius            float32
	Mass              float32
	Type              string
}

func NewExoplanet() Exoplanet {
	return Exoplanet{}
}
