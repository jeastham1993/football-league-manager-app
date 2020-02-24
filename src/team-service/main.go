package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/rs/cors"

	"team-service/infrastructure"
	"team-service/usecases"
)

func main() {
	var (
		httpAddr = flag.String("http.addr", ":8080", "HTTP listen address")
	)

	flag.Parse()

	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.NewSyncLogger(logger)
		logger = level.NewFilter(logger, level.AllowDebug())
		logger = log.With(logger,
			"svc", "order",
			"ts", log.DefaultTimestampUTC,
			"caller", log.DefaultCaller,
		)
	}

	level.Info(logger).Log("msg", "service started")
	defer level.Info(logger).Log("msg", "service ended")

	teamInteractor := new(usecases.TeamInteractor)
	// teamInteractor.TeamRepository = infrastructure.NewInMemTeamRepo()
	teamInteractor.TeamRepository = infrastructure.NewDynamoDbRepo()
	teamInteractor.Logger = new(infrastructure.Logger)
	// teamInteractor.EventHandler = new(infrastructure.MockEventBus)
	teamInteractor.EventHandler = infrastructure.NewAmazonSnsEventBus()

	var h http.Handler
	{
		endpoints := MakeEndpoints(teamInteractor)
		h = MakeHandler(endpoints)
	}

	errs := make(chan error)
	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errs <- fmt.Errorf("%s", <-c)
	}()

	go func() {
		level.Info(logger).Log("transport", "HTTP", "addr", *httpAddr)
		server := &http.Server{
			Addr:    *httpAddr,
			Handler: cors.Default().Handler(h),
		}

		errs <- server.ListenAndServe()
	}()

	level.Error(logger).Log("exit", <-errs)
}
