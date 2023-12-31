package main

import (
	"context"
	"flag"
	"github.com/IlyaZayats/dynus/internal/db"
	"github.com/IlyaZayats/dynus/internal/handlers"
	"github.com/IlyaZayats/dynus/internal/repository"
	"github.com/IlyaZayats/dynus/internal/services"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {

	var (
		dbUrl    string
		listen   string
		logLevel string
	)
	//postgres://dynus:dynus@localhost:5555/dynus
	//postgres://dynus:dynus@postgres:5432/dynus
	flag.StringVar(&dbUrl, "db", "postgres://dynus:dynus@postgres:5432/dynus", "database connection url")
	flag.StringVar(&listen, "listen", ":8080", "server listen interface")
	flag.StringVar(&logLevel, "log-level", "error", "log level: panic, fatal, error, warning, info, debug, trace")

	flag.Parse()

	ctx := context.Background()

	level, err := logrus.ParseLevel(logLevel)
	if err != nil {
		logrus.Panicf("unable to get log level: %v", err)
	}
	logrus.SetLevel(level)

	dbc, err := db.NewPostgresPool(dbUrl)
	if err != nil {
		logrus.Panicf("unable get postgres pool: %v", err)
	}

	slugRepo, err := repository.NewPostgresSlugRepository(dbc)
	if err != nil {
		logrus.Panicf("unable build slug repo: %v", err)
	}

	slugService, err := services.NewSlugService(slugRepo)
	if err != nil {
		logrus.Panicf("unable build slug service: %v", err)
	}

	g := gin.New()

	valid, err := handlers.GetValidator()
	if err != nil {
		logrus.Panicf("unable build slug validator: %v", err)
	}

	_, err = handlers.NewSlugHandlers(g, slugService, valid)
	if err != nil {
		logrus.Panicf("unable build slug handlers: %v", err)
	}

	doneC := make(chan error)

	go func() { doneC <- g.Run(listen) }()

	signalChan := make(chan os.Signal)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGABRT, syscall.SIGHUP, syscall.SIGTERM)

	childCtx, cancel := context.WithCancel(ctx)
	go func() {
		sig := <-signalChan
		logrus.Debugf("exiting with signal: %v", sig)
		cancel()
	}()

	go func(ctx context.Context) {
		ticker := time.NewTicker(1 * time.Second)
		for {
			select {
			case <-ctx.Done():
				doneC <- ctx.Err()
			case <-ticker.C:
				if err := slugRepo.CleanupByTTL(); err != nil {
					logrus.WithFields(logrus.Fields{
						"err": err,
					}).Error("unable cleanup slug by ttl")
				}
			}
		}
	}(childCtx)

	<-doneC
}
