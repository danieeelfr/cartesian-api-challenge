package config

import "github.com/danieeelfr/cartesian/internal/entity"

const (
	HttpServerHostEnvVar = "HTTP_SERVER_HOST"
	PointsFilePathEnvVar = "POINTS_FILE_PATH"
)

type Config struct {
	CartesianApiConfig *CartesianApiConfig
	PointsFilePath     string
}

type CartesianApiConfig struct {
	HttpServerHost string
	Points         []entity.Point
}

func (cfg *Config) GetCartesianApiConfig() *CartesianApiConfig { return cfg.CartesianApiConfig }
