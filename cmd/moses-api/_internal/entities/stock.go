package entities

import (
	"github.com/gofrs/uuid"
)

type Stock struct {
	ID      uuid.UUID `bson:"uuid" json:"id"`
	Dim     string    `bson:"dim" json:"dim"`
	Quality string    `bson:"quality" json:"quality"`
	Length  float64   `bson:"length" json:"length"`
}

func (s *Stock) ToBar() *Bar {
	b := new(Bar)

	b.Dim = s.Dim
	b.Quality = s.Quality
	b.Length = s.Length

	return b
}

func (s *Stock) Validate() error {
	if s.Dim == "" {
		return ErrDim
	}
	if s.Quality == "" {
		return ErrQuality
	}
	if s.Length <= 0 {
		return ErrLength
	}

	return nil
}
