package http_api_user

import (
	"context"
	"net/http"

	"github.com/devflex-pro/tma-starter-kit/backend/config"
	"github.com/devflex-pro/tma-starter-kit/backend/domain"
	"github.com/valyala/fasthttp/reuseport"
)

type API struct {
	conf       config.HTTPServerConfig
	httpServer *http.Server
}

func New(
	conf config.HTTPServerConfig,
	db domain.UserDB,
) *API {

	mux := http.NewServeMux()
	RegisterRoutes(mux, db)

	return &API{
		httpServer: &http.Server{
			Handler: mux,
		},
	}
}

func (api *API) Start(ctx context.Context) error {
	ln, err := reuseport.Listen(
		*api.conf.Network,
		*api.conf.Addr)
	if err != nil {
		return err
	}
	defer ln.Close()

	errChan := make(chan error, 1)

	go func() {
		defer close(errChan)
		errServe := api.httpServer.Serve(ln)
		if errServe != nil && errServe != http.ErrServerClosed {
			select {
			case errChan <- errServe:
			default:
			}
		}
	}()

	select {
	case <-ctx.Done():
		return api.httpServer.Shutdown(ctx)
	case err := <-errChan:
		return err
	}
}
