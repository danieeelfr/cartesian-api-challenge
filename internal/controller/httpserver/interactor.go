package httpserver

import (
	"github.com/danieeelfr/cartesian/internal/config"

	"github.com/danieeelfr/cartesian/pkg/wait"
)

type Interactor interface {
	Start() error
	Shutdown()
}

func New(cfg *config.Config, wg *wait.Wait) (Interactor, error) {

	httpSrv, err := NewHttpServer(cfg, wg)
	if err != nil {
		return nil, err
	}

	return httpSrv, nil
}
