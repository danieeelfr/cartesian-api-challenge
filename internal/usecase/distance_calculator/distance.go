package distance_calculator

import (
	"strconv"

	"github.com/danieeelfr/cartesian/internal/config"
	"github.com/danieeelfr/cartesian/internal/entity"
	"github.com/danieeelfr/cartesian/pkg/wait"
)

type distanceUsecase struct {
	wait   *wait.Wait
	points []entity.Point
}

func newDistanceUsecase(c *config.CartesianApiConfig, w *wait.Wait) *distanceUsecase {
	uc := &distanceUsecase{
		wait:   w,
		points: c.Points,
	}
	return uc
}

func (uc *distanceUsecase) GetDistance(params *DistanceParams) (*DistanceResponse, error) {
	var err error
	response := new(entity.DistanceResponse)

	if err = params.validate(); err != nil {
		return nil, err
	}

	origin := entity.Point{
		Name:               "origin",
		PosX:               stringToInt64(params.Distance.X),
		PosY:               stringToInt64(params.Distance.Y),
		DistanceFromOrigin: stringToInt64(params.Distance.Distance),
	}

	log.Info(origin)

	log.Info(uc.points)

	return &DistanceResponse{
		Data: response,
	}, nil
}

func stringToInt64(value string) int64 {
	number, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		log.WithError(err).Errorf("Failed to convert document number. Value: [%s].", value)
		return 0
	}
	return number
}

