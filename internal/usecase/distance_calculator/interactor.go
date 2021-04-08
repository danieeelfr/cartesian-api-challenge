package distance_calculator

import (
	"github.com/danieeelfr/cartesian/internal/config"
)

type (
	DistanceInteractor interface {
		GetPointsByDistance(params *DistanceParams) (*DistanceResponse, error)
	}
)

func New(c *config.CartesianApiConfig) DistanceInteractor {
	return newDistanceUsecase(c)
}
