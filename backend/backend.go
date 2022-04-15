package backend

import (
	"net/http"

	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gorilla/mux"
	"github.com/hsblhsn/hn.hsblhsn.me/backend/graphql"
)

func RegisterRoutes(router *mux.Router, gql *graphql.GQLHandler, img *graphql.ImageHandler) {
	router.Path("/graphql").
		Methods(http.MethodGet, http.MethodPost).
		Handler(gql)
	router.Path("/explorer").
		Methods(http.MethodGet).
		Handler(playground.Handler("GraphQL Explorer", "/graphql"))
	router.Path("/images.jpeg").
		Methods(http.MethodGet).
		Queries("size", "{size}", "src", "{src}").
		Handler(img)
}
