package entities

import (
	"os"
	"strconv"

	"github.com/gofrs/uuid"
)

type Nest struct {
	ID          uuid.UUID   `bson:"uuid" json:"id"`
	Project     string      `bson:"project" json:"project"`
	Launch      string      `bson:"launch" json:"launch"`
	NestName    string      `bson:"nest_name" json:"nest_name"`
	Bar         *Bar        `bson:"bar" json:"bar"`
	Profiles    []*Profile  `bson:"-" json:"-"`
	ProfilesIds []uuid.UUID `bson:"profiles" json:"profiles"`
	Spacings    []float64   `bson:"spacings" json:"spacings"`
	Machine     MachineType `bson:"machine" json:"machine"`
}

type NestSlice []*Nest

func (bs NestSlice) Len() int { return len(bs) }
func (bs NestSlice) Less(i, j int) bool {
	return bs[i].Bar.Length < bs[j].Bar.Length
}
func (bs NestSlice) Swap(i, j int) {
	bs[i], bs[j] = bs[j], bs[i]
}

func (n *Nest) GetRemnant() *Remnant {
	if !n.HasUsefulRemnant() {
		return nil
	}

	r := new(Remnant)

	r.Project = n.Project
	r.Dim = n.Bar.Dim
	r.Quality = n.Bar.Quality
	r.Length = n.Bar.Length - n.Bar.UsedLength
	r.Marking = n.NestName + "R01"
	r.From = n.NestName

	return r
}

func (n *Nest) HasUsefulRemnant() bool {
	minRemnant, _ := strconv.ParseFloat(os.Getenv("MIN_REMNANT"), 64)
	return n.Bar.Length-n.Bar.UsedLength >= minRemnant
}

func (n *Nest) RemnantLength() float64 {
	if n.HasUsefulRemnant() {
		return n.Bar.Length - n.Bar.UsedLength
	}
	return 0
}

func (n *Nest) RemnantLengthString() string {
	return strconv.FormatFloat(n.RemnantLength(), 'g', -1, 64)
}

func (n *Nest) PartsLength() float64 {
	var res float64
	for _, p := range n.Profiles {
		res += p.Length
	}
	return res
}

func (n *Nest) PartsLengthString() string {
	return strconv.FormatFloat(n.PartsLength(), 'g', -1, 64)
}

func (n *Nest) Scrap() float64 {
	var res float64
	for _, s := range n.Spacings {
		res += s
	}
	if !n.HasUsefulRemnant() {
		res += n.Bar.Length - n.Bar.UsedLength
	}
	return res
}

func (n *Nest) NestingPercent() float64 {
	return n.PartsLength() / (n.Bar.Length - n.RemnantLength()) * 100
}

func (n *Nest) Validate() error {
	if n.NestName == "" {
		return ErrNestName
	}
	if n.Project == "" {
		return ErrProject
	}
	if n.Machine == 0 {
		return ErrMachine
	}
	if len(n.Profiles) == 0 {
		return ErrProfiles
	}
	if n.Launch == "" {
		return ErrLaunch
	}
	if len(n.Spacings) != len(n.Profiles)*2 {
		return ErrSpacings
	}

	return nil
}
