package event

import (
	"context"

	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
)

type KindCreated struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Description *string   `json:"description"`
	Status      string    `json:"status"`
}

type KindCreatedHandler struct {
	// 	Repository persistence.KindRepository
	// CommandBus *cqrs.CommandBus
}

func (k KindCreatedHandler) HandlerName() string {
	return "kind-created"
}

func (k KindCreatedHandler) NewEvent() any {
	return &KindCreated{}
}

func (k KindCreatedHandler) Handle(ctx context.Context, event any) error {
	// ctx, cancel := context.WithTimeout(ctx, time.Second*120)
	// defer cancel()

	// e := event.(*KindCreated)
	// _ = e

	log.Info().Msg("KindCreated event received")

	return nil
}
