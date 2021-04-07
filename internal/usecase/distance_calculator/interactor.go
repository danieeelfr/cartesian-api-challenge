package distance_calculator

import (
	"github.com/danieeelfr/cartesian/internal/config"
	"github.com/danieeelfr/cartesian/pkg/wait"
)

type (
	DistanceInteractor interface {
		GetDistance(params *DistanceParams) (*DistanceResponse, error)
	}
)

func New(c *config.CartesianApiConfig, w *wait.Wait) DistanceInteractor {
	return newDistanceUsecase(c, w)
}
