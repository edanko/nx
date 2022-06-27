package command

import (
	"context"
	"errors"
	"strings"

	"github.com/edanko/nx/cmd/launch-api/internal/app/event"
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
)

type ChangeKindDescription struct {
	ID          uuid.UUID `json:"id"`
	Description string    `json:"description"`
}

type ChangeKindDescriptionHandler struct {
	eventBus   eventBus
	repository KindRepository
	sanitazer  sanitazer
}

func NewChangeKindDescriptionHandler(
	eventBus eventBus,
	repo KindRepository,
	s sanitazer,
) *ChangeKindDescriptionHandler {
	return &ChangeKindDescriptionHandler{
		eventBus:   eventBus,
		repository: repo,
		sanitazer:  s,
	}
}

func (h ChangeKindDescriptionHandler) HandlerName() string {
	return "change-kind-description"
}

func (h ChangeKindDescriptionHandler) NewCommand() any {
	return &ChangeKindDescription{}
}

func (h ChangeKindDescriptionHandler) Handle(ctx context.Context, command any) error {
	c, ok := command.(*ChangeKindDescription)
	if !ok {
		return errors.New("invalid command")
	}

	c.Description = h.sanitazer.Sanitize(c.Description)
	c.Description = strings.TrimSpace(c.Description)

	k, err := h.repository.Get(ctx, c.ID)
	if err != nil {
		log.Warn().Err(err).Msg("error getting kind from repository")
		return nil
	}

	k.ChangeDescription(c.Description)

	err = h.repository.Update(ctx, k)
	if err != nil {
		log.Warn().Err(err).Msg("error updating kind in repository")
		return nil
	}

	err = h.eventBus.Publish(ctx, event.KindDescriptionChanged{
		ID:          k.ID(),
		Description: *k.Description(),
	})
	if err != nil {
		log.Warn().Err(err).Msg("error publishing kind name changed event")
		return nil
	}

	return nil
}
