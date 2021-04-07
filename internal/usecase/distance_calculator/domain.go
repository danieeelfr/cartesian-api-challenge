package distance_calculator

import (
	"encoding/json"
	"time"

	"github.com/danieeelfr/cartesian/internal/entity"

	"github.com/sirupsen/logrus"
)

const (
	DistanceResource = "distance"
)

var (
	log      = logrus.WithField("package", "usecase.spi")
	maxDelay time.Duration
)

type DistanceParams struct {
	entity.Distance
}

func (p *DistanceParams) UnmarshalJSON(data []byte) error {

	parser := struct {
		entity.Distance
	}{}
	if err := json.Unmarshal(data, &parser); err != nil {
		return err
	}
	p.Distance = parser.Distance
	return nil
}

func (p *DistanceParams) validate() error {
	return p.Distance.Validate()
}

type DistanceResponse struct {
	Data *entity.DistanceResponse
}
