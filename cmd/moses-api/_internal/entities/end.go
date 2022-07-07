package entities

type End struct {
	Name        string             `bson:"name" json:"name"`
	Params      map[string]float64 `bson:"params" json:"params"`
	Uncut       bool               `bson:"uncut,omitempty" json:"uncut,omitempty"`
	Excess      float64            `bson:"excess,omitempty" json:"excess,omitempty"`
	WebBevel    *Bevel             `bson:"web_bevel,omitempty" json:"web_bevel,omitempty"`
	FlangeBevel *Bevel             `bson:"flange_bevel,omitempty" json:"flange_bevel,omitempty"`
	Scallop     *Scallop           `bson:"scallop,omitempty" json:"scallop,omitempty"`
}
