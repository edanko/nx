package kind

import (
	"errors"
)

var (
	ErrKindAlreadyExist     = errors.New("kind is already exist")
	ErrKindAlreadyPublished = errors.New("kind is already published")
	ErrKindAlreadyDraft     = errors.New("kind is already draft")
)
