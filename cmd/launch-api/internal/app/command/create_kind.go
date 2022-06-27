package command

import (
	"context"
	"errors"
	"strings"

	"github.com/edanko/nx/cmd/launch-api/internal/app/event"
	"github.com/edanko/nx/cmd/launch-api/internal/domain/kind"
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
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
	sanitazer  sanitazer
}

func NewCreateKindHandler(
	eventBus eventBus,
	repo KindRepository,
	s sanitazer,
) *CreateKindHandler {
	return &CreateKindHandler{
		eventBus:   eventBus,
		repository: repo,
		sanitazer:  s,
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
		return errors.New("invalid command")
	}

	c.Name = h.sanitazer.Sanitize(c.Name)
	c.Name = strings.TrimSpace(c.Name)

	if c.Description != nil {
		description := h.sanitazer.Sanitize(*c.Description)
		description = strings.TrimSpace(description)
		c.Description = &description
	}

	kind, err := kind.NewKind(
		c.ID,
		c.Name,
		c.Description,
		c.Status,
	)
	if err != nil {
		log.Warn().Err(err).Msg("error creating kind")
		return nil
	}

	err = kind.Validate()
	if err != nil {
		log.Warn().Err(err).Msg("error validating kind")
		return nil
	}

	exist, err := h.repository.Exist(ctx, c.Name)
	if err != nil {
		log.Warn().Err(err).Msg("error checking if kind with the name already exists")
		return nil
	}
	if exist {
		log.Warn().Msg("kind with the name already exists")
		return nil
	}

	err = h.repository.Create(ctx, kind)
	if err != nil {
		log.Warn().Err(err).Msg("error creating kind in repository")
		return nil
	}

	err = h.eventBus.Publish(ctx, event.KindCreated{
		ID:          kind.ID(),
		Name:        kind.Name(),
		Description: kind.Description(),
		Status:      kind.Status().String(),
	})
	if err != nil {
		log.Warn().Err(err).Msg("error publishing kind created event")
		return nil
	}

	return nil
}
