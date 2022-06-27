package queries

import (
	"time"

	"github.com/google/uuid"

	"github.com/edanko/nx/cmd/launch-api/internal/domain/kind"
)

type KindModel struct {
	ID          uuid.UUID
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Name        string
	Description *string
	Status      string
}

func fromDomain(d *kind.Kind) KindModel {
	return KindModel{
		ID:          d.ID(),
		CreatedAt:   d.CreatedAt(),
		UpdatedAt:   d.UpdatedAt(),
		Name:        d.Name(),
		Description: d.Description(),
		Status:      d.Status().String(),
	}
}
