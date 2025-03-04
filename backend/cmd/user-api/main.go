package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/devflex-pro/tma-starter-kit/backend/config"
	"github.com/devflex-pro/tma-starter-kit/backend/db/mongodb"
	"github.com/devflex-pro/tma-starter-kit/backend/domain"
	http_api_user "github.com/devflex-pro/tma-starter-kit/backend/http-api/user"
)

func main() {
	slog.Info("User API started")

	var (
		db        domain.UserDB
		api       domain.UserAPI
		httpConf  config.HTTPServerConfig
		mongoConf config.MongoUserDBConfig
		err       error
	)

	httpConf.Load()
	mongoConf.Load()

	ctx, cancel := context.WithCancel(
		context.Background(),
	)
	defer cancel()

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-sigCh
		slog.Info("context canceled")
		cancel()
	}()

	db, err = mongodb.New(
		ctx, mongoConf)
	if err != nil {
		slog.Error("db error", "error", err)
		os.Exit(1)
	}

	api = http_api_user.New(httpConf, db)

	err = api.Start(ctx)
	if err != nil {
		slog.Error("api error", "error", err)
		os.Exit(1)
	}

	slog.Info("User API stopped")
}
