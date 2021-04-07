package httpserver

import (
	"context"

	"github.com/danieeelfr/cartesian/internal/config"
	"github.com/danieeelfr/cartesian/pkg/wait"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type server struct {
	conf   *config.CartesianApiConfig
	wait   *wait.Wait
	e      *echo.Echo
	router *router
}

func NewHttpServer(cfg *config.Config, wg *wait.Wait) (*server, error) {

	srv := new(server)
	srv.conf = cfg.GetCartesianApiConfig()
	srv.wait = wg
	srv.e = echo.New()

	var err error
	srv.router, err = newRouter(srv.e, cfg, wg)
	if err != nil {
		return nil, err
	}
	return srv, nil
}

func (srv *server) Start() error {

	log.Infof("Starting API HTTP server. Host: [%s]...", srv.conf.HttpServerHost)
	srv.wait.Add()

	srv.e.Use(middleware.LoggerWithConfig(middleware.DefaultLoggerConfig))

	srv.setValidator()
	srv.router.build()

	go func() {
		if err := srv.e.Start(srv.conf.HttpServerHost); err != nil {
			if !srv.wait.IsBlock() {
				log.WithError(err).Fatalf("Failed to start http server. Host: [%s].", srv.conf.HttpServerHost)
			}
		}
	}()
	return nil
}

func (srv *server) Shutdown() {

	defer srv.wait.Done()
	if err := srv.e.Shutdown(context.Background()); err != nil {
		log.WithError(err).Error("Failed to shutdown server.")
	}
}

func (srv *server) setValidator() {

	srv.e.Validator = &ParameterValidator{
		validator: validator.New(),
	}
}

type ParameterValidator struct {
	validator *validator.Validate
}

func (val *ParameterValidator) Validate(i interface{}) error {
	return val.validator.Struct(i)
}
