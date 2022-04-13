package backend

import (
	graphql "github.com/99designs/gqlgen/graphql/handler"
	"github.com/gorilla/mux"
	"github.com/hsblhsn/hn.hsblhsn.me/backend/internal/images"
)

func RegisterRoutes(router *mux.Router, gql *graphql.Server, img *images.ImageResizeHandler) {
	router.Handle("/graphql", gql)
	router.Handle("/images.jpeg", img)
}
