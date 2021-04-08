package httpserver

import (
	"github.com/danieeelfr/cartesian/internal/config"

)

type Interactor interface {
	Start() error
	Shutdown()
}

func New(cfg *config.Config) (Interactor, error) {

	httpSrv, err := NewHttpServer(cfg)
	if err != nil {
		return nil, err
	}

	return httpSrv, nil
}
