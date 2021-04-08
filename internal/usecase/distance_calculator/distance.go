package distance_calculator

import (
	"errors"
	"sort"
	"strconv"

	"github.com/danieeelfr/cartesian/internal/config"
	"github.com/danieeelfr/cartesian/internal/entity"
	"github.com/sirupsen/logrus"
)

var (
	log = logrus.WithField("package", "entity")
)

type distanceUsecase struct {
	points []entity.Point
}

func newDistanceUsecase(c *config.CartesianApiConfig) *distanceUsecase {
	uc := &distanceUsecase{
		points: c.Points,
	}
	return uc
}

func (uc *distanceUsecase) GetPointsByDistance(params *DistanceParams) (*DistanceResponse, error) {
	var err error
	response := new(entity.DistanceResponse)

	if err = params.validate(); err != nil {
		return nil, errors.New(entity.BusinessError)
	}

	origin := entity.Point{
		Name:               "origin",
		PosX:               stringToInt64(params.DistanceRequest.X),
		PosY:               stringToInt64(params.DistanceRequest.Y),
		DistanceFromOrigin: stringToInt64(params.DistanceRequest.Distance),
	}

	points := make([]entity.Point, 0)

	for i := 0; i < len(uc.points); i++ {
		p := uc.points[i]

		manhattanDistance := abs(origin.PosX-p.PosX) + abs(origin.PosY-p.PosY)
		p.DistanceFromOrigin = manhattanDistance

		if manhattanDistance <= origin.DistanceFromOrigin {
			points = append(points, p)
		}
	}

	sort.SliceStable(points, func(i, j int) bool {
		return points[i].DistanceFromOrigin < points[j].DistanceFromOrigin
	})

	response.Points = points

	return &DistanceResponse{
		Data: response,
	}, nil
}

func abs(x int64) int64 {
	if x < 0 {
		return -x
	}
	return x
}

func stringToInt64(value string) int64 {
	number, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		log.WithError(err).Errorf("Failed to converting to int64. Value: [%s].", value)
		return 0
	}
	return number
}
