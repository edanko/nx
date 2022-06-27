package commands

import (
	"context"
	"errors"
	"strings"

	"github.com/google/uuid"
	"github.com/rs/zerolog/log"

	"github.com/edanko/nx/cmd/launch-api/internal/app/events"
)

type ChangeKindName struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

type ChangeKindNameHandler struct {
	eventBus   eventBus
	repository KindRepository
	sanitizer  sanitizer
}

func NewChangeKindNameHandler(
	eventBus eventBus,
	repo KindRepository,
	s sanitizer,
) *ChangeKindNameHandler {
	return &ChangeKindNameHandler{
		eventBus:   eventBus,
		repository: repo,
		sanitizer:  s,
	}
}

func (h ChangeKindNameHandler) HandlerName() string {
	return "change-kind-name"
}

func (h ChangeKindNameHandler) NewCommand() any {
	return &ChangeKindName{}
}

func (h ChangeKindNameHandler) Handle(ctx context.Context, command any) error {
	c, ok := command.(*ChangeKindName)
	if !ok {
		return errors.New("invalid commands")
	}

	c.Name = h.sanitizer.Sanitize(c.Name)
	c.Name = strings.TrimSpace(c.Name)
	c.Name = strings.ToUpper(c.Name)

	k, err := h.repository.Get(ctx, c.ID)
	if err != nil {
		log.Warn().Err(err).Msg("error getting kind from repository")
		return nil
	}

	k.ChangeName(c.Name)

	err = h.repository.Update(ctx, k)
	if err != nil {
		log.Warn().Err(err).Msg("error updating kind in repository")
		return nil
	}

	err = h.eventBus.Publish(ctx, events.KindNameChanged{
		ID:   k.ID(),
		Name: k.Name(),
	})
	if err != nil {
		log.Warn().Err(err).Msg("error publishing kind name changed events")
		return nil
	}

	return nil
}
