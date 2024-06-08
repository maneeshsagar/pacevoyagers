package tests

import (
	"testing"

	"github.com/maneeshsagar/pacevoyagers/model"
	"github.com/maneeshsagar/pacevoyagers/service"
	"github.com/stretchr/testify/assert"
)

func TestCreateExoplanet(t *testing.T) {
	exoplanet := &model.Exoplanet{
		Name:              "PlanetX",
		Description:       "A new exoplanet",
		DistanceFromEarth: 50,
		Radius:            1.5,
		Type:              "GasGiant",
	}

	spaceService := service.NewExoplanetService()
	fetchedExoplanet, err := spaceService.AddExoplanet(exoplanet)
	assert.NoError(t, err)
	assert.Equal(t, "PlanetX", fetchedExoplanet.Name)
}

func TestListExoplanets(t *testing.T) {

	exoplanet1 := &model.Exoplanet{
		Name:              "PlanetA",
		Description:       "First exoplanet",
		DistanceFromEarth: 100,
		Radius:            2.0,
		Type:              "Terrestrial",
		Mass:              5.0,
	}
	exoplanet2 := &model.Exoplanet{
		Name:              "PlanetB",
		Description:       "Second exoplanet",
		DistanceFromEarth: 200,
		Radius:            3.0,
		Type:              "GasGiant",
	}

	spaceService := service.NewExoplanetService()

	spaceService.AddExoplanet(exoplanet1)
	spaceService.AddExoplanet(exoplanet2)

	exoplanetList, _, _, _ := spaceService.GetExoplanets()

	assert.Len(t, exoplanetList, 2)
}

func TestGetExoplanet(t *testing.T) {

	exoplanet := &model.Exoplanet{
		Name:              "PlanetC",
		Description:       "Third exoplanet",
		DistanceFromEarth: 150,
		Radius:            1.8,
		Type:              "Terrestrial",
		Mass:              4.0,
	}

	spaceService := service.NewExoplanetService()
	exoplanet, _ = spaceService.AddExoplanet(exoplanet)

	fetchedExoplanet, _, _, err := spaceService.GetExoplanetById(exoplanet.Id)
	assert.NoError(t, err)
	assert.Equal(t, "PlanetC", fetchedExoplanet.Name)
}

func TestUpdateExoplanet(t *testing.T) {

	exoplanet := &model.Exoplanet{
		Name:              "PlanetD",
		Description:       "Fourth exoplanet",
		DistanceFromEarth: 180,
		Radius:            2.2,
		Type:              "GasGiant",
	}

	spaceService := service.NewExoplanetService()
	addedExoplanet, _ := spaceService.AddExoplanet(exoplanet)

	updatedExoplanet := &model.Exoplanet{
		Id:                addedExoplanet.Id,
		Name:              "PlanetD-Updated",
		Description:       "Updated description",
		DistanceFromEarth: 180,
		Radius:            2.2,
		Type:              "GasGiant",
	}
	resultExoplanet, _, _, err := spaceService.UpdateExoplanet(updatedExoplanet)
	assert.NoError(t, err)
	assert.Equal(t, "PlanetD-Updated", resultExoplanet.Name)
}
func TestCalculateFuelEstimation(t *testing.T) {

	exoplanet := &model.Exoplanet{
		Name:              "PlanetF",
		Description:       "Sixth exoplanet",
		DistanceFromEarth: 300,
		Radius:            2.5,
		Type:              "GasGiant",
	}
	spaceService := service.NewExoplanetService()
	addedExoplanet, _ := spaceService.AddExoplanet(exoplanet)

	fuel, _, _, err := spaceService.CalculateFuelEstimation(addedExoplanet.Id, 10)
	assert.NoError(t, err)
	assert.NotZero(t, fuel)
}
