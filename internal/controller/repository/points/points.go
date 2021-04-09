package points

import (
	"encoding/json"
	"fmt"
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
		e := fmt.Errorf("Error reading the points file. Defails: %w", err)
		return nil, e
	}

	p := entity.DistanceResponse{}
	if err := json.Unmarshal([]byte(file), &p); err != nil {
		e := fmt.Errorf("Error parsing the points file. Defails: %w", err)
		return nil, e
	}

	return p.Points, nil
}
