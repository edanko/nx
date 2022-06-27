package command

import (
	"context"
	"errors"

	"github.com/edanko/nx/cmd/launch-api/internal/app/event"
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
)

type DeleteKind struct {
	ID uuid.UUID `json:"id"`
}

type DeleteKindHandler struct {
	eventBus   eventBus
	repository KindRepository
}

func NewDeleteKindHandler(
	eventBus eventBus,
	repository KindRepository,
) *DeleteKindHandler {
	return &DeleteKindHandler{eventBus, repository}
}

func (h DeleteKindHandler) HandlerName() string {
	return "delete-kind"
}

func (h DeleteKindHandler) NewCommand() any {
	return &DeleteKind{}
}

func (h DeleteKindHandler) Handle(ctx context.Context, command any) error {
	c, ok := command.(*DeleteKind)
	if !ok {
		return errors.New("invalid command")
	}

	err := h.repository.Delete(ctx, c.ID)
	if err != nil {
		log.Warn().Err(err).Msg("failed to delete kind")
		return nil
	}

	err = h.eventBus.Publish(ctx, event.KindDeleted{
		ID: c.ID,
	})
	if err != nil {
		log.Warn().Err(err).Msg("error publishing kind deleted event")
		return nil
	}

	return nil
}
