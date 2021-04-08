package entity

import (
	"errors"
	"strings"

	"github.com/sirupsen/logrus"
)

const (
	ValidationMessageFormat = "Your request is not valid. Check the details."
)

var (
	log = logrus.WithField("package", "entity")
)

type DistanceRequest struct {
	X        string `json:"x"`
	Y        string `json:"y"`
	Distance string `json:"distance"`
}

func (ref *DistanceRequest) Validate() (err error) {

	if len(strings.TrimSpace(ref.X)) == 0 {
		return errors.New(BadRequestError)
	}

	if len(strings.TrimSpace(ref.Y)) == 0 {
		return errors.New(BadRequestError)
	}

	if len(strings.TrimSpace(ref.Distance)) == 0 {
		return errors.New(BadRequestError)
	}

	return
}

type DistanceResponse struct {
	Points []Point `json:"points"`
}

type Point struct {
	Name               string `json:"name"`
	PosX               int64  `json:"x"`
	PosY               int64  `json:"y"`
	DistanceFromOrigin int64  `json:"distance_from_origin"`
}
