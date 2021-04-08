package points

import (
	"encoding/json"
	"io/ioutil"

	"github.com/danieeelfr/cartesian/internal/entity"
)

type Points struct {
}

func NewPoints() (*Points, error) {
	srv := new(Points)
	return srv, nil
}

func (ref *Points) GetPoints(filePath string) ([]entity.Point, error) {

	file, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	p := entity.DistanceResponse{}
	if err := json.Unmarshal([]byte(file), &p); err != nil {
		return nil, err
	}

	return p.Points, nil
}
