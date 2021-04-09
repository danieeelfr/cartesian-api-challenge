package httpserver

import (
	"github.com/danieeelfr/cartesian/internal/config"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type server struct {
	conf   *config.CartesianApiConfig
	e      *echo.Echo
	router *router
}

func NewHttpServer(cfg *config.Config) (*server, error) {

	srv := new(server)
	srv.conf = cfg.GetCartesianApiConfig()
	srv.e = echo.New()

	var err error
	srv.router, err = newRouter(srv.e, cfg)
	if err != nil {
		return nil, err
	}
	return srv, nil
}

func (srv *server) Start() error {

	log.Infof("Starting API HTTP server. Host: [%s]...", srv.conf.HttpServerHost)

	srv.e.Use(middleware.LoggerWithConfig(middleware.DefaultLoggerConfig))

	srv.setValidator()
	srv.router.build()

	if err := srv.e.Start(srv.conf.HttpServerHost); err != nil {
		log.WithError(err).Fatalf("Failed to start http server. Host: [%s].", srv.conf.HttpServerHost)
	}

	return nil
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
