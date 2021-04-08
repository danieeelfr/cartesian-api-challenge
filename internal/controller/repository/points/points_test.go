package points

import (
	"os"
	"path"
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestGetPoints(t *testing.T) {
	p, _ := NewPoints()
	result, err := p.GetPoints(get_points_file_path())
	assert.Equal(t, nil, err)
	assert.Equal(t, len(result) > 0, true)
}

func get_points_file_path() string {
	dirname, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	dir, err := os.Open(path.Join(dirname, "../../../../"))
	if err != nil {
		panic(err)
	}

	return dir.Name() + "/data/points.json"
}
