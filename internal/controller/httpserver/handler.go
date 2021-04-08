package httpserver

import (
	"net/http"

	"github.com/danieeelfr/cartesian/internal/config"
	usecase "github.com/danieeelfr/cartesian/internal/usecase/distance_calculator"
	"github.com/danieeelfr/cartesian/pkg/wait"
	"github.com/sirupsen/logrus"

	"github.com/labstack/echo/v4"
)

var (
	log = logrus.WithField("package", "controller.httpserver")
)

type handler struct {
	wait *wait.Wait
	uc   usecase.DistanceInteractor
}

func newHandler(cfg *config.Config, wg *wait.Wait) (*handler, error) {

	hdl := new(handler)
	hdl.wait = wg

	CartesianApiConfig := cfg.GetCartesianApiConfig()

	hdl.uc = usecase.New(CartesianApiConfig, wg)
	return hdl, nil
}

func (hdl *handler) GetPointsByDistance(c echo.Context) error {

	if !hdl.wait.Add() {
		return http.ErrServerClosed
	}
	defer hdl.wait.Done()

	params := new(usecase.DistanceParams)

	if err := c.Bind(params); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	if err := c.Validate(params); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	log.Infof("Processing the request. Parameters: %v.", params.Distance)
	response, err := hdl.uc.GetPointsByDistance(params)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	log.Infof("Responding to successful request. Status code: [%d].", http.StatusOK)
	c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)

	return c.JSON(http.StatusOK, response.Data)
}
