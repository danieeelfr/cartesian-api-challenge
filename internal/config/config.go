package config

import (
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
)

const (
	CartesianApp = "cartesian-api"
)

var (
	log             = logrus.WithField("package", "config")
	requiredEnvVars []string
)

func New(app string) *Config {

	values, err := getEnvironmentVariables(app)
	if err != nil {
		log.Fatal(err)
	}

	cfg := new(Config)
	switch app {
	case CartesianApp:
		cfg.CartesianApiConfig = &CartesianApiConfig{
			HttpServerHost: values[HttpServerHostEnvVar],
		}
		cfg.PointsFilePath = values[PointsFilePathEnvVar]
	}

	return cfg
}

func getEnvironmentVariables(app string) (map[string]string, error) {

	switch app {
	case CartesianApp:
		requiredEnvVars = []string{
			HttpServerHostEnvVar,
			PointsFilePathEnvVar,
		}
	}

	missing, values := checkMissing(requiredEnvVars)
	if len(missing) > 0 {
		return nil, fmt.Errorf("Not found required environment variables. Missing: %v", missing)
	}
	return values, nil
}

func checkMissing(requiredVars []string) (missing []string, values map[string]string) {
	missing = []string{}
	values = map[string]string{}

	for _, env := range requiredVars {

		if value := os.Getenv(env); value == "" {
			missing = append(missing, env)
		} else {
			values[env] = value
		}
	}

	return missing, values
}
