package server

import (
	rmsMiddleware "github.com/RacoonMediaServer/rms-packages/pkg/middleware"
	"net/http"
)

func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return rmsMiddleware.PanicHandler(rmsMiddleware.RequestsCountHandler(rmsMiddleware.UnauthorizedRequestsCountHandler(handler)))
}
