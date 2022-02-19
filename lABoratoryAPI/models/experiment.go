package models

type Experiment struct {
	Id          string       `bson:"_id,omitempty" json:"id,omitempty"`
	Name        string       `json:"name"`
	Assignments []Assignment `json:"assignments"`
	Owner       User         `json:"owner"`
}
