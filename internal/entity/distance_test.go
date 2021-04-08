package entity

import (
	"errors"
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestDistanceBadRequestValidate(t *testing.T) {

	values := &[]struct {
		description string
		input       *DistanceRequest
		output      error
	}{
		{
			description: "Empty request should return bad request error",
			input:       &DistanceRequest{},
			output:      errors.New(BadRequestError),
		},
		{
			description: "Request without X value should return bad request error",
			input: &DistanceRequest{
				Y:        "1",
				Distance: "5",
			},
			output: errors.New(BadRequestError),
		},
		{
			description: "Request without Y value should return bad request error",
			input: &DistanceRequest{
				X:        "1",
				Distance: "10",
			},
			output: errors.New(BadRequestError),
		},
		{
			description: "Request without Distance value should return bad request error",
			input: &DistanceRequest{
				X: "6",
				Y: "7",
			},
			output: errors.New(BadRequestError),
		},
		{
			description: "Request with valid values should pass in validations",
			input: &DistanceRequest{
				Y:        "5",
				X:        "1",
				Distance: "10",
			},
			output: nil,
		},
	}

	for _, v := range *values {
		err := v.input.Validate()
		assert.Equal(t, v.output, err)
	}
}
