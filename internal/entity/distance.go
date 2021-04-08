package entity

import (
	"encoding/json"
	"errors"
	"fmt"
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

func (ref *DistanceRequest) String() string {

	b, err := json.Marshal(ref)
	if err != nil {
		return fmt.Sprintf("Fail to parsing distance. Detail: [%s]", err)
	}
	return string(b)
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

func (ref *DistanceResponse) String() string {

	b, err := json.Marshal(ref)
	if err != nil {
		return fmt.Sprintf("Fail to parsing distance response. Detail: [%s]", err)
	}
	return string(b)
}
