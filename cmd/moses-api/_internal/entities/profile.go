package entities

import (
	"github.com/gofrs/uuid"
)

type Profile struct {
	ID           uuid.UUID `bson:"uuid" json:"id"`
	Project      string    `bson:"project" json:"project"`
	Launch       string    `bson:"launch" json:"launch"`
	Section      string    `bson:"section" json:"section"`
	PosNo        string    `bson:"pos_no" json:"pos_no"`
	Quality      string    `bson:"quality" json:"quality"`
	Dim          string    `bson:"dim" json:"dim"`
	Length       float64   `bson:"length" json:"length"`
	FullLength   float64   `bson:"-" json:"-"`
	Quantity     int       `bson:"quantity" json:"quantity"`
	L            *End      `bson:"l" json:"l"`
	R            *End      `bson:"r" json:"r"`
	Spacing      []float64 `bson:"-" json:"-"`
	TraceBevel   *Bevel    `bson:"trace_bevel,omitempty" json:"trace_bevel,omitempty"`
	Holes        []*Hole   `bson:"holes,omitempty" json:"holes,omitempty"`
	BendingCurve []float64 `bson:"bending_curve,omitempty" json:"bending_curve,omitempty"`
}

type ProfileSliceAsc []*Profile

func (ps ProfileSliceAsc) Len() int { return len(ps) }
func (ps ProfileSliceAsc) Less(i, j int) bool {
	return ps[i].Length < ps[j].Length
}
func (ps ProfileSliceAsc) Swap(i, j int) {
	ps[i], ps[j] = ps[j], ps[i]
}

func (p *Profile) Validate() error {
	if p.Launch == "" {
		return ErrLaunch
	}
	if p.Project == "" {
		return ErrProject
	}
	if p.Section == "" {
		return ErrSection
	}
	if p.PosNo == "" {
		return ErrPosNo
	}
	if p.Quality == "" {
		return ErrQuality
	}
	if p.Dim == "" {
		return ErrDim
	}
	if p.Length <= 0 {
		return ErrLength
	}
	if p.Quantity <= 0 {
		return ErrQuantity
	}
	if p.L == nil || p.R == nil {
		return ErrEnd
	}

	return nil
}

func (p *Profile) InvertHolesX() {
	for _, h := range p.Holes {
		h.Params["X"] = p.Length - h.Params["X"]
	}
}
