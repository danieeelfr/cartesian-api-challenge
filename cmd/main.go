package main

import (
	"os"

	"github.com/danieeelfr/cartesian/internal/config"
	httpServerCtrl "github.com/danieeelfr/cartesian/internal/controller/httpserver"
	pointsRepositoryCtrl "github.com/danieeelfr/cartesian/internal/controller/repository/points"

	"github.com/sirupsen/logrus"
)

var (
	log = logrus.WithField("package", "main.distance")
)

func init() {
	logrus.SetLevel(logrus.InfoLevel)
}

func main() {
	log.Info("Starting application...")

	ctrlPoints, err := pointsRepositoryCtrl.NewPoints()
	if err != nil {
		log.WithError(err).Fatal("Could not run the application.")
	}

	points, err := ctrlPoints.GetPoints(os.Getenv(config.PointsFilePathEnvVar))
	if err != nil {
		log.WithError(err).Fatal("Could not run the application.")
	} else {
		log.Info(points)
	}

	cfg := config.New(config.CartesianApp)
	cfg.CartesianApiConfig.Points = points

	ctrlHttp, err := httpServerCtrl.New(cfg)

	if err != nil {
		log.WithError(err).Fatal("Could not run the application.")
	}

	if err := ctrlHttp.Start(); err != nil {
		log.WithError(err).Fatal("Fail starting the http server.")
	}

}
