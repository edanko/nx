package command

import (
	"context"
	"errors"

	"github.com/edanko/nx/cmd/launch-api/internal/app/event"
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
)

type MakeKindDraft struct {
	ID uuid.UUID `json:"id"`
}

type MakeKindDraftHandler struct {
	eventBus   eventBus
	repository KindRepository
}

func NewMakeKindDraftHandler(
	eventBus eventBus,
	repo KindRepository,
) *MakeKindDraftHandler {
	return &MakeKindDraftHandler{
		eventBus:   eventBus,
		repository: repo,
	}
}

func (h MakeKindDraftHandler) HandlerName() string {
	return "make-kind-draft"
}

func (h MakeKindDraftHandler) NewCommand() any {
	return &MakeKindDraft{}
}

func (h MakeKindDraftHandler) Handle(ctx context.Context, command any) error {
	c, ok := command.(*MakeKindDraft)
	if !ok {
		return errors.New("invalid command")
	}

	k, err := h.repository.Get(ctx, c.ID)
	if err != nil {
		log.Warn().Err(err).Msg("error getting kind from repository")
		return nil
	}

	err = k.MakeDraft()
	if err != nil {
		log.Warn().Err(err).Msg("error making kind draft")
		return nil
	}

	err = h.repository.Update(ctx, k)
	if err != nil {
		log.Warn().Err(err).Msg("error updating kind in repository")
		return nil
	}

	err = h.eventBus.Publish(ctx, event.KindMadeDraft{
		ID: k.ID(),
	})
	if err != nil {
		log.Warn().Err(err).Msg("error publishing kind made draft event")
		return nil
	}

	return nil
}
