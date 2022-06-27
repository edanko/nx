package kind

import (
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

type Kind struct {
	id          uuid.UUID `validate:"required"`
	name        string    `validate:"required"`
	description *string
	status      Status `validate:"required"`
}

func (k *Kind) ID() uuid.UUID {
	return k.id
}

func (k *Kind) Name() string {
	return k.name
}

func (k *Kind) Description() *string {
	return k.description
}

func (k *Kind) Status() Status {
	return k.status
}

func NewKind(
	id uuid.UUID,
	name string,
	description *string,
	status string,
) (*Kind, error) {
	if id == uuid.Nil {
		return nil, ErrKindValidationUUID
	}
	if name == "" {
		return nil, ErrKindValidationName
	}
	if description != nil && *description != "" {
		return nil, ErrKindValidationDescription
	}

	s, err := NewStatusFromString(status)
	if err != nil {
		return nil, errors.Wrap(err, ErrKindValidationStatus.Error())
	}

	kind := &Kind{
		id:          id,
		name:        name,
		description: description,
		status:      s,
	}

	return kind, nil
}

func (k *Kind) ChangeName(newName string) {
	k.name = newName
}

func (k *Kind) ChangeDescription(newDescription string) {
	k.description = &newDescription
}

func (k *Kind) IsPublished() bool {
	return k.status == Published
}

func (k *Kind) IsDraft() bool {
	return k.status == Draft
}

func (k *Kind) MakePublished() error {
	if k.IsPublished() {
		return ErrKindAlreadyPublished
	}

	k.status = Published
	return nil
}

func (k *Kind) MakeDraft() error {
	if k.IsDraft() {
		return ErrKindAlreadyDraft
	}

	k.status = Draft
	return nil
}
