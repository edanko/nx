package entities

type Hole struct {
	Name   string             `bson:"name" json:"name"`
	Params map[string]float64 `bson:"params" json:"params"`
}
