package entities

import (
	"github.com/gofrs/uuid"
)

type Spacing struct {
	ID         uuid.UUID   `bson:"uuid" json:"id"`
	HasBevel   bool        `bson:"has_bevel" json:"has_bevel"`
	HasScallop bool        `bson:"has_scallop" json:"has_scallop"`
	Length     float64     `bson:"length" json:"length"`
	Machine    MachineType `bson:"machine" json:"machine"`
	Dim        string      `bson:"dim" json:"dim"`
	Name       string      `bson:"name" json:"name"`
}

func (s *Spacing) Validate() error {
	if s.Machine == 0 {
		return ErrMachine
	}
	if s.Dim == "" {
		return ErrDim
	}
	if s.Length <= 0 {
		return ErrLength
	}

	return nil
}
