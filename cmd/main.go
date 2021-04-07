package main

import (
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/danieeelfr/cartesian/internal/config"
	httpServerCtrl "github.com/danieeelfr/cartesian/internal/controller/httpserver"
	pointsCtrl "github.com/danieeelfr/cartesian/internal/controller/points"

	"github.com/danieeelfr/cartesian/pkg/wait"

	"github.com/sirupsen/logrus"
)

var (
	log            = logrus.WithField("package", "main.distance")
	w              = wait.New()
	waitToShutdown time.Duration
)

func init() {
	logrus.SetLevel(logrus.InfoLevel)
}

func main() {
	log.Info("Starting application...")

	ctrlPoints, err := pointsCtrl.NewPoints()
	if err != nil {
		log.WithError(err).Fatal("Could not run the application.")
	}

	points, err := ctrlPoints.GetPoints()
	if err != nil {
		log.WithError(err).Fatal("Could not run the application.")
	} else {
		log.Info(points)
	}

	cfg := config.New(config.CartesianApp)
	cfg.CartesianApiConfig.Points = points

	waitToShutdown = time.Duration(5) * time.Second

	ctrlHttp, err := httpServerCtrl.New(cfg, w)

	if err != nil {
		log.WithError(err).Fatal("Could not run the application.")
	}

	w.Add()

	if err := ctrlHttp.Start(); err != nil {
		log.WithError(err).Fatal("Fail starting the http server.")
	}

	w.Wait()
	log.Infof("Finishing %s!", config.CartesianApp)
}

func shutdownSignal(ctrlHttp httpServerCtrl.Interactor) {

	signalChannel := make(chan os.Signal, 2)
	signal.Notify(signalChannel, syscall.SIGTERM, syscall.SIGINT)

	go func() {
		sig := <-signalChannel
		switch sig {
		case syscall.SIGTERM, syscall.SIGINT:
			log.Infof("Interruption request. Signal: [%v].", sig)
			w.Block()
			ctrlHttp.Shutdown()

			log.Infof("Waiting [%v] for open processes.", waitToShutdown)
			time.Sleep(waitToShutdown)

			log.Infof("Finishing...")
			for w.Done() {
				log.Infof("Ignoring open process to kill...")
			}
		}
	}()
}
