package backend

import (
	"github.com/gorilla/mux"
	"github.com/hsblhsn/hn.hsblhsn.me/backend/graphql"
)

func RegisterRoutes(router *mux.Router, gql *graphql.GQLHandler, img *graphql.ImageHandler) {
	router.Handle("/graphql", gql)
	router.Handle("/images.jpeg", img)
}
