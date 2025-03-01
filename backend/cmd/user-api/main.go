package main

import (
	"context"

	"github.com/devflex-pro/tma-starter-kit/backend/config"
	"github.com/devflex-pro/tma-starter-kit/backend/domain"
	http_api_user "github.com/devflex-pro/tma-starter-kit/backend/http-api/user"
)

func main() {

	var (
		db         domain.UserDB
		api        domain.UserAPI
		httpConfig config.HTTPServerConfig
	)

	httpConfig.Load()

	api = http_api_user.New(httpConfig, db)

	err := api.Start(context.TODO())
	if err != nil {

	}

}
