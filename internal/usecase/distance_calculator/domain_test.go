package distance_calculator

import (
	"testing"
	"time"
	// "github.com/stretchr/testify/assert"
)

func TestDistanceParams_validateParams(t *testing.T) {

	now := time.Now()
	println(now)
	// lateHour := now.Add(-time.Duration(6) * time.Second)
	// earlyHour := now.Add(time.Duration(1) * time.Second)
	// withinHour := now.Add(-time.Duration(3) * time.Second)
	// maxDelay = time.Duration(5) * time.Second

	// values := &[]struct {
	// 	description string
	// 	input       *DistanceParams
	// 	output      error
	// }{
	// 	{
	// 		description: "For StartedAt equal to the current time, you must fail.",
	// 		input: &DistanceParams{
	// 			Begun:     now,
	// 			StartedAt: &now,
	// 		},
	// 		output: entity.ErrorInvalidStartedAt,
	// 	},
	// 	{
	// 		description: "For StartedAt before the acceptable delay, you must fail.",
	// 		input: &DistanceParams{
	// 			Begun:     now,
	// 			StartedAt: &lateHour,
	// 		},
	// 		output: entity.ErrorInvalidStartedAt,
	// 	},
	// 	{
	// 		description: "For StartedAt after the current time, you must fail",
	// 		input: &DistanceParams{
	// 			Begun:     now,
	// 			StartedAt: &earlyHour,
	// 		},
	// 		output: entity.ErrorInvalidStartedAt,
	// 	},
	// 	{
	// 		description: "For StartedAt within the acceptable delay, you must succeed.",
	// 		input: &DistanceParams{
	// 			Begun:     now,
	// 			StartedAt: &withinHour,
	// 		},
	// 		output: nil,
	// 	},
	// }

	// for _, v := range *values {

	// 	err := v.input.startedAtValidate()
	// 	assert.Equal(t, v.output, err, v.description)
	// }
}
