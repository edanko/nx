package kind

import (
	"errors"
)

var (
	ErrKindValidationUUID        = errors.New("kind uuid validation failed")
	ErrKindValidationName        = errors.New("kind name cannot be empty")
	ErrKindValidationDescription = errors.New("kind description cannot be empty")
	ErrKindValidationStatus      = errors.New("kind status validation failed")
	ErrKindAlreadyExist          = errors.New("kind is already exist")
	ErrKindAlreadyPublished      = errors.New("kind is already published")
	ErrKindAlreadyDraft          = errors.New("kind is already draft")
)
