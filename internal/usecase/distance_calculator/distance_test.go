package distance_calculator

import (
	"context"
	"testing"

	"github.com/danieeelfr/cartesian/internal/config"
	"github.com/danieeelfr/cartesian/internal/entity"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/suite"
)

const (
	AppNameTest = "CARTESIANAPI"
)

type DistanceUsecaseTestSuite struct {
	suite.Suite
	usecase *distanceUsecase
	ctx     context.Context
}

func (suite *DistanceUsecaseTestSuite) SetupSuite() {

	logrus.SetLevel(logrus.InfoLevel)
	log.Debug("Starting test usecase...")

	c := &config.CartesianApiConfig{}
	suite.usecase = newDistanceUsecase(c)
	suite.ctx = context.Background()
}

func (suite *DistanceUsecaseTestSuite) BeforeTest(suiteName, testName string) {
	log.WithField("suite", suiteName).WithField("name", testName).Debug("Test begin.")

	switch testName {
	case "TestGetPointsByDistance":
		// response := new(entity.DistanceResponse)
		// uuid := suite.ctx.Value(fxlog.TraceFieldName).(string)
		// suite.acc.On("CompleteWithdraw", suite.ctx, MyUniqueID).Return(nil)
		// suite.acc.On("RevertWithdraw", suite.ctx, MyUniqueID, entity.InternalValidationReversalError).Return(nil)

		// // INVALID
		// invalidPay := createInvalidDistanceParsed()
		// suite.prv.On("Distance", suite.ctx, invalidPay, response).Return(nil)

		// invalidPayload := createInvalidDistanceParsed()
		// invalidPayload.EnableDefaultParser()
		// invalidPayload.Payer.Ispb = 0
		// msgFail := createMessageWithError(&suite.begun, &suite.startedAt, invalidPayload, response)
		// suite.msg.On("Publish", suite.ctx, uuid, msgFail).Return(nil)

		// // VALID
		// validPay := createDistanceParsed()
		// queryResponse := new(entity.DistanceResponse)
		// suite.prv.On("Distance", suite.ctx, validPay, response).Return(nil)
		// suite.prv.On("ConsultDistance", suite.ctx, MyUniqueExternalID, queryResponse).Return(nil)

		// validPayload := createDistanceParsed()
		// validPayload.EnableDefaultParser()
		// msgSucess := createMessage(&suite.begun, &suite.startedAt, validPayload, createResponse())
		// suite.msg.On("Publish", suite.ctx, uuid, msgSucess).Return(nil)
	}
}

func (suite *DistanceUsecaseTestSuite) TestGetPointsByDistance() {

	type Output struct {
		resp *DistanceResponse
		err  error
	}

	// values := &[]struct {
	// 	description string
	// 	input       *DistanceParams
	// 	output      *Output
	// }{
	// 	{
	// 		description: "For parameters not filled in correctly, validation should fail.",
	// 		input: &DistanceParams{
	// 			Ctx:       suite.ctx,
	// 			Begun:     suite.begun,
	// 			StartedAt: &suite.startedAt,
	// 			Distance:  createInvalidDistance(),
	// 		},
	// 		output: &Output{
	// 			resp: nil,
	// 			err:  createValidationError(),
	// 		},
	// 	},
	// 	{
	// 		description: "For parameters filled in correctly, it must be successful.",
	// 		input: &DistanceParams{
	// 			Ctx:       suite.ctx,
	// 			Begun:     suite.begun,
	// 			StartedAt: &suite.startedAt,
	// 			Distance:  createDistance(),
	// 		},
	// 		output: &Output{
	// 			resp: &DistanceResponse{
	// 				Data: createResponse(),
	// 			},
	// 			err: nil,
	// 		},
	// 	},
	// }

	// for _, v := range *values {
	// 	resp, err := suite.usecase.GetPointsByDistance(v.input)
	// 	// TODO: Add assert expectations.
	// 	assert.Equal(suite.T(), v.output.resp, resp, v.description)
	// 	assert.Equal(suite.T(), v.output.err, err, v.description)
	// 	suite.usecase.wait.Wait()
	// }
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

func createDistance() entity.DistanceRequest {

	return entity.DistanceRequest{
		// ID: MyUniqueID,

	}
}

func createInvalidDistance() entity.DistanceRequest {

	return entity.DistanceRequest{
		// ID:       MyUniqueID,

	}
}

func createResponse() *entity.DistanceResponse {

	points := make([]entity.Point, 2)
	p := entity.Point{
		Name: "point A",
		PosX: 1,
		PosY: 1,
	}

	points[0] = p

	p = entity.Point{
		Name: "point B",
		PosX: 2,
		PosY: 2,
	}

	points[1] = p

	resp := &entity.DistanceResponse{
		Points: points,
	}

	return resp
}
