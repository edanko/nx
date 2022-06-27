package commands

import (
	"context"
	"errors"
	"strings"

	"github.com/google/uuid"
	"github.com/rs/zerolog/log"

	"github.com/edanko/nx/cmd/launch-api/internal/app/events"
	"github.com/edanko/nx/cmd/launch-api/internal/domain/kind"
)

type CreateKind struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Description *string   `json:"description"`
	Status      string    `json:"status"`
}

type CreateKindHandler struct {
	eventBus   eventBus
	repository KindRepository
	sanitizer  sanitizer
}

func NewCreateKindHandler(
	eventBus eventBus,
	repo KindRepository,
	s sanitizer,
) *CreateKindHandler {
	return &CreateKindHandler{
		eventBus:   eventBus,
		repository: repo,
		sanitizer:  s,
	}
}

func (h CreateKindHandler) HandlerName() string {
	return "create-kind"
}

func (h CreateKindHandler) NewCommand() any {
	return &CreateKind{}
}

func (h CreateKindHandler) Handle(ctx context.Context, command any) error {
	c, ok := command.(*CreateKind)
	if !ok {
		return errors.New("invalid commands")
	}

	c.Name = h.sanitizer.Sanitize(c.Name)
	c.Name = strings.TrimSpace(c.Name)

	if c.Description != nil {
		description := h.sanitizer.Sanitize(*c.Description)
		description = strings.TrimSpace(description)
		c.Description = &description
	}

	k, err := kind.NewKind(
		c.ID,
		c.Name,
		c.Description,
		c.Status,
	)
	if err != nil {
		log.Warn().Err(err).Msg("error creating kind")
		return nil
	}

	exist, err := h.repository.Exist(ctx, c.Name)
	if err != nil {
		log.Warn().Err(err).Msg("error checking if kind with the name already exists")
		return nil
	}
	if exist {
		log.Warn().Err(kind.ErrKindAlreadyExist).Msg("kind with the name already exists")
		return nil
	}

	err = h.repository.Create(ctx, k)
	if err != nil {
		log.Warn().Err(err).Msg("error creating kind in repository")
		return nil
	}

	err = h.eventBus.Publish(ctx, events.KindCreated{
		ID:          k.ID(),
		Name:        k.Name(),
		Description: k.Description(),
		Status:      k.Status().String(),
	})
	if err != nil {
		log.Warn().Err(err).Msg("error publishing kind created events")
		return nil
	}

	return nil
}
