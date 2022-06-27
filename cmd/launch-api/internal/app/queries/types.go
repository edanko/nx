package queries

import (
	"github.com/google/uuid"

	"github.com/edanko/nx/cmd/launch-api/internal/domain/kind"
)

type KindModel struct {
	ID          uuid.UUID
	Name        string
	Description *string
	Status      string
}

func fromDomain(d *kind.Kind) KindModel {
	return KindModel{
		ID:          d.ID(),
		Name:        d.Name(),
		Description: d.Description(),
		Status:      d.Status().String(),
	}
}
