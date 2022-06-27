package app

import (
	"context"

	"github.com/edanko/nx/cmd/launch-api/internal/app/queries"
)

type mediator interface {
	Send(ctx context.Context, cmd any) error
}

type Application struct {
	Commands mediator
	Queries  Queries
}

type Queries struct {
	ListKinds     queries.ListKindsHandler
	GetKind       queries.GetKindHandler
	GetKindByName queries.GetKindByNameHandler
	// CountKinds    queries.CountKindsHandler
}
