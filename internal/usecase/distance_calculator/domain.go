package distance_calculator

import (
	"github.com/danieeelfr/cartesian/internal/entity"
)

type DistanceParams struct {
	entity.DistanceRequest
}

func (p *DistanceParams) validate() error {
	return p.DistanceRequest.Validate()
}

type DistanceResponse struct {
	Data *entity.DistanceResponse
}
