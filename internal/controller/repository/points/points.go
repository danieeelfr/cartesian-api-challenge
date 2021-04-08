package points

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path"

	"github.com/danieeelfr/cartesian/internal/entity"
)

type Points struct {
	// Points *entity.DistanceResponse
}

func NewPoints() (*Points, error) {
	srv := new(Points)
	return srv, nil
}

func (ref *Points) GetPoints() ([]entity.Point, error) {
	root := root_dir()

	file, err := ioutil.ReadFile(root + "/data/points.json")
	if err != nil {
		return nil, err
	}

	p := entity.DistanceResponse{}
	if err := json.Unmarshal([]byte(file), &p); err != nil {
		return nil, err
	}

	return p.Points, nil
}

func root_dir() string {
	dirname, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Current directory: %v\n", dirname)
	dir, err := os.Open(path.Join(dirname, "../"))
	if err != nil {
		panic(err)
	}

	return dir.Name()
	//fmt.Printf("Name of ../../: %v\n", dir.Name())
}
