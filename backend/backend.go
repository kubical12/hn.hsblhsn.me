package backend

import (
	"github.com/gorilla/mux"
	"github.com/hsblhsn/hackernews"
	"github.com/hsblhsn/hn.hsblhsn.me/backend/endpoints"
	"github.com/hsblhsn/hn.hsblhsn.me/backend/internal/grpc/readabilityserver"
	"github.com/hsblhsn/hn.hsblhsn.me/backend/services"
	"github.com/hsblhsn/hn.hsblhsn.me/backend/transports"
)

func RegisterRoutes(router *mux.Router) {
	go readabilityserver.Initialize()
	var (
		hnCfg    = hackernews.NewConfiguration()
		hnClient = hackernews.NewAPIClient(hnCfg)
		svc      = services.NewService(hnClient)
		eps      = endpoints.NewEndpoints(svc)
	)
	transports.HTTP(router, eps)
}
