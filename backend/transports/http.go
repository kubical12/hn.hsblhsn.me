package transports

import (
	"net/http"

	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"github.com/hsblhsn/hn.hsblhsn.me/backend/endpoints"
	"github.com/hsblhsn/hn.hsblhsn.me/backend/types"
)

func HTTP(router *mux.Router, ep endpoints.Endpoint) {
	router.Path("/lists/{type}.json").
		Queries("page", "{page:[0-9]+}").
		Methods(http.MethodGet).
		Handler(httptransport.NewServer(
			ep.GetListByType,
			types.DecodeGetListByTypeRequest,
			types.EncodeGetListByTypeResponse,
			httptransport.ServerErrorEncoder(ErrorEncoder),
		))

	router.Path("/items/{id:[0-9]+}.json").
		Methods(http.MethodGet).
		Handler(httptransport.NewServer(
			ep.GetItemByID,
			types.DecodeGetItemByIDRequest,
			types.EncodeGetItemByIDResponse,
			httptransport.ServerErrorEncoder(ErrorEncoder),
		))

	router.Path("/image.jpeg").
		Queries("src", "{src}", "size", "{size:full|thumbnail}").
		Methods(http.MethodGet).
		Handler(httptransport.NewServer(
			ep.GetResizedImage,
			types.DecodeGetResizedImage,
			types.EncodeGetResizedImageResponse,
			httptransport.ServerErrorEncoder(ErrorEncoder),
		))
}
