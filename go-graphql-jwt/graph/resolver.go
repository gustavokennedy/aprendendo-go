package graph

//go:generate go run github.com/99designs/gqlgen

import "gitlab.com/pragmaticreviews/graphql-go/graph/model"

type Resolver struct {
	videos []*model.Video
}
