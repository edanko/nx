package command

import (
	"context"
	"errors"

	"github.com/edanko/nx/cmd/launch-api/internal/app/event"
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
)

type MakeKindPublished struct {
	ID uuid.UUID `json:"id"`
}

type MakeKindPublishedHandler struct {
	eventBus   eventBus
	repository KindRepository
}

func NewMakeKindPublishedHandler(
	eventBus eventBus,
	repo KindRepository,
) *MakeKindPublishedHandler {
	return &MakeKindPublishedHandler{
		eventBus:   eventBus,
		repository: repo,
	}
}

func (h MakeKindPublishedHandler) HandlerName() string {
	return "make-kind-published"
}

func (h MakeKindPublishedHandler) NewCommand() any {
	return &MakeKindPublished{}
}

func (h MakeKindPublishedHandler) Handle(ctx context.Context, command any) error {
	c, ok := command.(*MakeKindPublished)
	if !ok {
		return errors.New("invalid command")
	}

	k, err := h.repository.Get(ctx, c.ID)
	if err != nil {
		log.Warn().Err(err).Msg("error getting kind from repository")
		return nil
	}

	err = k.MakePublished()
	if err != nil {
		log.Warn().Err(err).Msg("error making kind published")
		return nil
	}

	err = h.repository.Update(ctx, k)
	if err != nil {
		log.Warn().Err(err).Msg("error updating kind in repository")
		return nil
	}

	err = h.eventBus.Publish(ctx, event.KindMadePublished{
		ID: k.ID(),
	})
	if err != nil {
		log.Warn().Err(err).Msg("error publishing kind made published event")
		return nil
	}

	return nil
}
