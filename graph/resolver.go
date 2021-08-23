package graph

import (
	"github.com/megajon/gqlgen-todos/graph/model"
	"github.com/megajon/gqlgen-todos/postgres"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	users     []*model.User
	UsersRepo postgres.UsersRepo
}
