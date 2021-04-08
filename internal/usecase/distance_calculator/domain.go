package distance_calculator

import (
	"encoding/json"

	"github.com/danieeelfr/cartesian/internal/entity"

	"github.com/sirupsen/logrus"
)

var (
	log = logrus.WithField("package", "usecase.distance_calculator")
)

type DistanceParams struct {
	entity.DistanceRequest
}

func (p *DistanceParams) UnmarshalJSON(data []byte) error {

	parser := struct {
		entity.DistanceRequest
	}{}
	if err := json.Unmarshal(data, &parser); err != nil {
		return err
	}
	p.Distance = parser.Distance
	return nil
}

func (p *DistanceParams) validate() error {
	return p.DistanceRequest.Validate()
}

type DistanceResponse struct {
	Data *entity.DistanceResponse
}
