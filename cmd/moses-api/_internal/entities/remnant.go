package entities

import (
	"github.com/gofrs/uuid"
)

type Remnant struct {
	ID      uuid.UUID `bson:"uuid" json:"id"`
	Used    bool      `bson:"used" json:"used"`
	Length  float64   `bson:"length" json:"length"`
	Project string    `bson:"project" json:"project"`
	From    string    `bson:"from" json:"from"`
	Dim     string    `bson:"dim" json:"dim"`
	Quality string    `bson:"quality" json:"quality"`
	Marking string    `bson:"marking" json:"marking"`
	UsedIn  string    `bson:"used_in,omitempty" json:"used_in,omitempty"`
}

func (r *Remnant) ToBar() *Bar {
	if r.Used {
		return nil
	}

	b := new(Bar)
	b.Dim = r.Dim
	b.Quality = r.Quality
	b.Length = r.Length
	// b.IsRemnant = true
	b.RemnantID = &r.ID

	return b
}

func (r *Remnant) Validate() error {
	if r.Project == "" {
		return ErrProject
	}
	if r.From == "" {
		return ErrFrom
	}
	if r.Dim == "" {
		return ErrDim
	}
	if r.Quality == "" {
		return ErrQuality
	}
	if r.Length <= 0 {
		return ErrLength
	}
	return nil
}
