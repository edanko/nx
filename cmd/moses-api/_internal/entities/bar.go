package entities

import (
	"github.com/gofrs/uuid"
)

type Bar struct {
	Dim        string     `bson:"dim" json:"dim"`
	Quality    string     `bson:"quality" json:"quality"`
	Length     float64    `bson:"length" json:"length"`
	UsedLength float64    `bson:"used_length" json:"used_length"`
	Remnant    *Remnant   `bson:"-" json:"-"`
	RemnantID  *uuid.UUID `bson:"remnant,omitempty" json:"remnant,omitempty"`
}

func (b *Bar) Validate() error {
	if b.Dim == "" {
		return ErrDim
	}
	if b.Quality == "" {
		return ErrQuality
	}
	if b.Length <= 0 {
		return ErrLength
	}
	if b.UsedLength > b.Length {
		return ErrUsedLength
	}
	return nil
}
