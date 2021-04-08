package httpserver

import (
	"net/http"

	"github.com/danieeelfr/cartesian/internal/config"
	"github.com/danieeelfr/cartesian/pkg/wait"

	"github.com/labstack/echo/v4"
)

const (
	DistanceEndpoint = "api/points"
)

type router struct {
	e      *echo.Echo
	routes []*route
}

type route struct {
	method   string
	endpoint string
	handler  func(c echo.Context) error
}

func newRouter(e *echo.Echo, cfg *config.Config, wg *wait.Wait) (*router, error) {

	handler, err := newHandler(cfg, wg)
	if err != nil {
		return nil, err
	}

	return &router{
		e: e,
		routes: []*route{
			&route{
				method:   http.MethodGet,
				endpoint: DistanceEndpoint,
				handler:  handler.GetPointsByDistance,
			},
		},
	}, nil
}

func (rtr *router) build() {

	for _, route := range rtr.routes {
		rtr.setRoute(route)
	}
}

func (rtr *router) setRoute(r *route) {

	switch r.method {
	case http.MethodGet:
		rtr.e.GET(r.endpoint, r.handler)

	default:
		log.Errorf("Method not implemented. Received method: [%s].", r.method)
	}
}
