package distance_calculator

import (
	"fmt"
	"testing"

	"github.com/danieeelfr/cartesian/internal/config"
	"github.com/danieeelfr/cartesian/internal/entity"
	"github.com/go-playground/assert/v2"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/suite"
)

const (
	AppNameTest = "CARTESIAN_API"
)

type DistanceUsecaseTestSuite struct {
	suite.Suite
	usecase *distanceUsecase
}

func (suite *DistanceUsecaseTestSuite) SetupSuite() {

	logrus.SetLevel(logrus.InfoLevel)
	log.Debug("Starting test usecase...")

	c := &config.CartesianApiConfig{}
	suite.usecase = newDistanceUsecase(c)
	suite.usecase.points = get_fake_points()
}

func (suite *DistanceUsecaseTestSuite) TestGetPointsByDistance() {

	type Output struct {
		resp *DistanceResponse
		err  error
	}

	values := &[]struct {
		description string
		input       *DistanceParams
		output      *Output
	}{
		{
			description: "Distance calculator with valid distance params should return only the within points",
			input: &DistanceParams{
				entity.DistanceRequest{
					X:        "1",
					Y:        "1",
					Distance: "5",
				},
			},
			output: &Output{
				resp: &DistanceResponse{
					Data: &entity.DistanceResponse{
						Points: suite.usecase.points,
					},
				},
				err: nil,
			},
		},
		{
			description: "Distance calculator with invalid distance params should return a business error",
			input: &DistanceParams{
				entity.DistanceRequest{
					Y:        "1",
					Distance: "5",
				},
			},
			output: &Output{
				resp: nil,
				err:  fmt.Errorf("Bad Request. Details: The field X is mandatory"),
			},
		},
		{
			description: "Distance calculator with invalid distance params should return a business error",
			input: &DistanceParams{
				entity.DistanceRequest{
					X:        "1",
					Distance: "2",
				},
			},
			output: &Output{
				resp: nil,
				err:  fmt.Errorf("Bad Request. Details: The field Y is mandatory"),
			},
		},
		{
			description: "Distance calculator with invalid distance params should return a business error",
			input: &DistanceParams{
				entity.DistanceRequest{
					X: "1",
					Y: "2",
				},
			},
			output: &Output{
				resp: nil,
				err:  fmt.Errorf("Bad Request. Details: The field Distance is mandatory"),
			},
		},
		{
			description: "Distance calculator with invalid distance params should return a business error",
			input: &DistanceParams{
				entity.DistanceRequest{},
			},
			output: &Output{
				resp: nil,
				err:  fmt.Errorf("Bad Request. Details: The field X is mandatory"),
			},
		},
	}

	for _, v := range *values {
		resp, err := suite.usecase.GetPointsByDistance(v.input)

		if err == nil {
			assert.Equal(suite.T(), len(resp.Data.Points), 2)
		} else {
			assert.Equal(suite.T(), fmt.Sprint(v.output.err), fmt.Sprint(err))
		}
	}
}

func (suite *DistanceUsecaseTestSuite) AfterTest(suiteName, testName string) {
	log.WithField("suite", suiteName).WithField("name", testName).Debug("Test done.")
}

func (suite *DistanceUsecaseTestSuite) TearDownSuite() {
	log.Debug("Finishing test usecase...")
}

func TestUsecase(t *testing.T) {
	suite.Run(t, new(DistanceUsecaseTestSuite))
}

func get_fake_points() []entity.Point {

	points := make([]entity.Point, 10)
	for i := 0; i < 10; i++ {
		p := entity.Point{
			Name: fmt.Sprintf("point %v", i),
			PosX: int64(i + 1),
			PosY: int64(i*-2 - 1),
		}
		points[i] = p
	}

	return points
}
