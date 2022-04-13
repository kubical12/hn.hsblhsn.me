package backend

import (
	"github.com/gorilla/mux"
	"github.com/hsblhsn/hn.hsblhsn.me/backend/graph"
)

func RegisterRoutes(router *mux.Router, gql *graph.GraphQLHandler, img *graph.ImageHandler) {
	router.Handle("/graphql", gql)
	router.Handle("/images.jpeg", img)
}
