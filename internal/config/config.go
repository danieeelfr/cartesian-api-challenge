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
	}

	return cfg
}

func getEnvironmentVariables(app string) (map[string]string, error) {

	switch app {
	case CartesianApp:
		requiredEnvVars = []string{
			HttpServerHostEnvVar,
		}
	}

	missing, values := checkMissing(requiredEnvVars)
	if len(missing) > 0 {
		return nil, fmt.Errorf("Not found required environment variables. Missing: %v", missing)
	}
	return values, nil
}

// func parseUint(value string) uint {

// 	number, err := strconv.Atoi(value)
// 	if err != nil {
// 		log.WithError(err).Fatalf("Failed to uint parsing. Value[%s].", value)
// 	}
// 	return uint(number)
// }

// func parseBool(value string) bool {

// 	v, err := strconv.ParseBool(value)
// 	if err != nil {
// 		log.Fatalf("Failed to bool parsing. Value[%s].", value)
// 	}
// 	return v
// }

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
