package points

import "github.com/danieeelfr/cartesian/internal/entity"

type Interactor interface {
	GetPoints(filePath string) ([]entity.Point, error)
}

func New() (Interactor, error) {

	pointsSrv, err := NewPoints()
	if err != nil {
		return nil, err
	}

	return pointsSrv, nil
}
