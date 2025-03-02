package main

import (
	"context"
	"log/slog"

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

	db, err = mongodb.New(
		context.Background(), mongoConf)
	if err != nil {
		slog.Error("db error", "error", err)
	}

	api = http_api_user.New(httpConf, db)

	err = api.Start(context.TODO())
	if err != nil {
		slog.Error("api error", "error", err)
	}

	slog.Info("User API stopped")
}
